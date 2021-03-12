import re
import sys

from invoke import task

LINTER_VERSION = "1.38.0"
LINTER = "$GOPATH/bin/golangci-lint"


def fail(message):
    print(message)
    sys.exit(1)


@task(
    help={
        "test": "<PACKAGE_PATH>::<TEST_NAME>. If empty, it will run all tests.",
        "verbose": "Show stdout of tests.",
        "show": "Show coverprofile page.",
    }
)
def test(c, test="", verbose=False, show=False):
    """Run tests."""
    test_regex = "./..."

    test = test.split("::")
    if len(test) == 2:
        test_regex = f"-run {test[1]} {test[0]}"

    r = c.run(
        f"go test {'-v' if verbose else ''} -race -count=1 -cover {'-coverprofile=coverage.out' if show else ''} {test_regex}"
    )

    packages = 0
    coverage = 0.0

    for cover in re.findall(r"[0-9]+\.[0-9]+(?=%)", r.stdout):
        packages += 1
        coverage += float(cover)

    if packages:
        coverage = round(coverage / packages, 1)

    title = "=" * (len(str(packages) + str(coverage)) + 34)
    print(title, f"    Total Coverage ({packages} pkg) : {coverage}%", title, sep="\n")

    if show:
        c.run("go tool cover -html=coverage.out")
        c.run("rm -f coverage.out")


@task(
    help={
        "yes": "Automatically say yes to the following questions.",
    }
)
def devtools(c, yes=False):
    """Check and install devtools."""

    def installed():
        r = c.run(f"{LINTER} --version", warn=True, hide="both")
        installed = not r.failed and LINTER_VERSION in r.stdout
        return installed

    if not installed():
        if not yes and input("Devtools not installed, install? y/n: ").lower() != "y":
            fail("Aborting as devtools not installed!")

        c.run(
            f"curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sudo sh -s -- -b $GOPATH/bin v{LINTER_VERSION}"
        )

        if not installed():
            fail("Aborting as devtools could not be installed!")


@task(
    help={
        "yes": "Automatically say yes to the following questions.",
    }
)
def lint(c, yes=False):
    """Run linter."""
    devtools(c, yes=yes)
    r = c.run(f"{LINTER} run ./... -c .golangci.yaml")
