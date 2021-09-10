"""Generate Ondatra test definitions."""

load("//tools/build_rules:expand_template.bzl", "expand_template")
load("//tools/build_defs/label:def.bzl", "parse_label")
load("//tools/build_defs/build_test:build_test.bzl", "build_test")

def ondatra_test(
        name,
        srcs,
        testbed,
        glaze_kind,
        run_timeout = "30m",
        args = None,
        deps = None,
        data = None,
        tags = None,
        visibility = None):
    """Compiles and runs an Ondatra test written in Go.

    Args:
      name: Name; required
      srcs: List of labels; required
      testbed: Label of testbed file; required
      glaze_kind: Must be set to "go_test" in the BUILD file for glaze; required
      run_timeout: Timeout on the test run, excluding the wait time for the
        testbed to become available, specified as a duration string
        (see http://godoc/pkg/time#ParseDuration); default is 30 minutes
      args: List of args: optional
      deps: List of labels; optional
      data: List of labels; optional
      tags: List of arbitrary text tags; optional
      visibility: List of visibility labels; optional
    """
    if glaze_kind != "go_test":
        fail("glaze_kind must be set to \"go_test\"")

    data = (data or []) + [testbed]
    testbed_arg = "--testbed=%s/%s" % parse_label(testbed)
    args = (args or []) + [
        testbed_arg,
        "--run_time=%s" % run_timeout,
        # A zero lets the binding implementation choose an appropriate default.
        # This is useful for Google-specific bindings, where the wait time should
        # roughly be the Blaze/Guitar-specific test timeout, minus the run time.
        "--wait_time=0",
    ]

    # Manual targets may need the guitar tag to be found by Guitar:
    # go/guitar/encyclopedia/developer/workflow.md#specifying-blaze-targets
    # The no-testloasd tag requires the test to be run with --notest_loasd:
    # http://g3doc/devtools/blaze/g3doc/user-manual.html#flag--test_loasd
    #
    # go/functional-test-tags#supports-graceful-termination
    # Indicates that Blaze should attempt graceful termination of the test process,
    # if Blaze is interrupted during execution of this test target
    go_test_tags = (tags or []) + ["guitar", "manual", "no-testloasd", "supports-graceful-termination"]

    native.go_test(
        name = name,
        deps = deps,
        srcs = srcs,
        data = data,
        args = args,
        tags = go_test_tags,
        visibility = visibility,
        size = "enormous",  # Reservation is highly variable.
        local = True,  # Tests cannot run on Forge.
    )

    # Create the _reserve binary target
    test_lib = name + "_lib"
    native.go_library(
        name = test_lib,
        srcs = srcs,
        deps = deps,
        testonly = 1,
    )
    expand_template(
        name = name + "_expand",
        template = "//third_party/openconfig/ondatra/internal/reservemain:template",
        substitutions = {"{TEST_PACKAGE}": "%s/%s" % (native.package_name(), test_lib)},
        out = name + "_reserve.go",
    )
    native.go_binary(
        name = name + "_reserve",
        srcs = [name + "_reserve.go"],
        deps = [
            test_lib,
            "//third_party/openconfig/ondatra/internal/reservemain",
        ],
        data = [testbed],
        args = [
            testbed_arg,
            "--run_time=1h",
            "--wait_time=1h",
        ],
        testonly = True,
    )

    # Create a forge-able compilation test that can run on TAP.
    build_test(
        name = "%s_build_test" % name,
        targets = [name],
        tags = tags,
    )

def ondatra_test_suite(
        name,
        srcs,
        testbeds,
        glaze_kind,
        per_test_run_timeout = "30m",
        args = None,
        deps = None,
        data = None,
        tags = None,
        visibility = None):
    """Defines a suite of Ondatra tests written in Go.

    For every (testbed-name, testbed-file) entry in the provided testbeds map,
    this rule creates an ondatra_test with the name "[name]_[testbed-name]" and
    the testbed equal to testbed-file.

    Args:
      name: Name; required
      srcs: List of labels; required
      testbeds: Map of custom testbed name to testbed label; required
      glaze_kind: Must be set to "go_test" in the BUILD file for glaze; required
      per_test_run_timeout: Timeout on each test run in the suite, excluding the
        wait time for the testbed to become available, specified as a duration
        string (see http://godoc/pkg/time#ParseDuration); default is 30 minutes
      args: List of args: optional
      deps: List of labels; optional
      data: List of labels; optional
      tags: List of arbitrary text tags; optional
      visibility: List of visibility labels; optional
    """
    if len(testbeds) == 0:
        fail("must specify at least one testbed")

    tests = []
    for testbed_name, testbed_src in testbeds.items():
        test_name = "%s_%s" % (name, testbed_name)
        tests.append(test_name)
        ondatra_test(
            name = test_name,
            srcs = srcs,
            testbed = testbed_src,
            glaze_kind = glaze_kind,
            run_timeout = per_test_run_timeout,
            args = args,
            deps = deps,
            data = data,
            tags = tags,
            visibility = visibility,
        )

    # Unlike other tags, "manual" on a test_suite means the suite itself is manual.
    # http://g3doc/devtools/blaze/g3doc/be/build-encyclopedia.html#test_suite
    native.test_suite(name = name, tests = tests, visibility = visibility, tags = ["manual"])
