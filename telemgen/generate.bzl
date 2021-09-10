"""Build macros for gNMI API generation."""

def ondatra_go_path_api(
        name,
        config = False,
        structs_file_split = 10,
        path_structs_file_split = 5,
        helper_func_file_split = 10,
        telem_type_file_split = 10,
        extra_yang_files = None,
        extra_yang_deps = None,
        extra_exclude_modules = []):
    """Generates a Go path API with helper methods for telemetry or config.

    The library provides a set of generated Go structs and path structs for the
    OpenConfig library based on a set of common schema definition files, in
    addition to extra YANG files and dependencies. The path structs are
    augmented with telemetry/config API methods.

    Args:
      name: Name; required
      config: Generate config API instead of telemetry API; optional
      structs_file_split: Non-negative integer of output GoStruct files; required
      path_structs_file_split: Non-negative integer of output PathStruct files; required
      helper_func_file_split: Non-negative integer of output helper function files; required
      telem_type_file_split: Non-negative integer of output telemetry type files; required
      extra_yang_files: List of strings of google3-relative paths; optional
      extra_yang_deps: List of labels; optional
      extra_exclude_modules: List of strings, yang modules to exclude from generation; optional
    """

    common_yang_deps = [
        "//ops/netops/openconfig/yang:all_modules",
        "//third_party/openconfig/public/release/models:openconfig_modules",
        "//third_party/openconfig/public/third_party/ietf:yang_modules",
        "//third_party/openconfig/ondatra/yang:all_modules",
        "//third_party/openconfig/vendor/arista/yang/eos4261f:arista_modules",
        "//platforms/networking/gpins/yang:all_modules",
    ]

    if extra_yang_deps:
        common_yang_deps += extra_yang_deps

    extra_yang_files_str = ""
    if extra_yang_files:
        extra_yang_files_str = " ".join(extra_yang_files) + " "

    native.vardef(
        "YANG_FILES",
        extra_yang_files_str +
        "ops/netops/openconfig/yang/platform/openconfig-platform-pipeline-counters.yang " +
        "platforms/networking/gpins/yang/google/google-pins-interfaces.yang " +
        "platforms/networking/gpins/yang/google/google-pins-platform.yang " +
        "platforms/networking/gpins/yang/google/google-pins-system.yang " +
        "platforms/networking/gpins/yang/openconfig-pins-interfaces.yang " +
        "platforms/networking/gpins/yang/openconfig-pins-platform-chassis.yang " +
        "platforms/networking/gpins/yang/openconfig-pins-platform-node.yang " +
        "platforms/networking/gpins/yang/openconfig-pins-platform-port.yang " +
        "platforms/networking/gpins/yang/openconfig-pins-platform.yang " +
        "third_party/openconfig/ondatra/yang/openconfig-ate-flow.yang " +
        "third_party/openconfig/ondatra/yang/openconfig-ate-intf.yang " +
        "third_party/openconfig/public/release/models/interfaces/openconfig-if-aggregate.yang " +
        "third_party/openconfig/public/release/models/interfaces/openconfig-if-ethernet.yang " +
        "third_party/openconfig/public/release/models/interfaces/openconfig-if-ip.yang " +
        "third_party/openconfig/public/release/models/interfaces/openconfig-if-ip-ext.yang " +
        "third_party/openconfig/public/release/models/interfaces/openconfig-interfaces.yang " +
        "third_party/openconfig/public/release/models/lacp/openconfig-lacp.yang " +
        "third_party/openconfig/public/release/models/lldp/openconfig-lldp.yang " +
        "third_party/openconfig/public/release/models/lldp/openconfig-lldp-types.yang " +
        "third_party/openconfig/public/release/models/network-instance/openconfig-network-instance.yang " +
        "third_party/openconfig/public/release/models/platform/openconfig-platform-cpu.yang " +
        "third_party/openconfig/public/release/models/platform/openconfig-platform-software.yang " +
        "third_party/openconfig/public/release/models/platform/openconfig-platform-transceiver.yang " +
        "third_party/openconfig/public/release/models/platform/openconfig-platform.yang " +
        "third_party/openconfig/public/release/models/qos/openconfig-qos.yang " +
        "third_party/openconfig/public/release/models/qos/openconfig-qos-elements.yang " +
        "third_party/openconfig/public/release/models/qos/openconfig-qos-interfaces.yang " +
        "third_party/openconfig/public/release/models/qos/openconfig-qos-types.yang " +
        "third_party/openconfig/public/release/models/system/openconfig-system.yang " +
        "third_party/openconfig/vendor/arista/yang/eos4261f/release/openconfig/models/interfaces/arista-intf-augments.yang",
    )

    structs_srcs = []
    if not config:
        structs_srcs = [
            "enum.go",
            "enum_map.go",
            "schema.go",
            "union.go",
        ]
        for i in range(structs_file_split):
            structs_srcs.append("structs-{}.go".format(i))
    for i in range(path_structs_file_split):
        structs_srcs.append("path_structs-{}.go".format(i))

    helpers_srcs = [
        "telem_helpers.go",
    ]
    for i in range(helper_func_file_split):
        helpers_srcs.append("telemetry-{}.go".format(i))
    for i in range(telem_type_file_split):
        helpers_srcs.append("telem_types-{}.go".format(i))

    go_library_deps = [
        "//third_party/golang/goyang/pkg/yang",
        "//third_party/golang/protobuf/v1/ptypes",
        "//third_party/golang/protobuf/v2/proto",
        "//third_party/golang/ygot/ygot",
        "//third_party/golang/ygot/ytypes",
        "//third_party/openconfig/gnmi/proto/gnmi:gnmi_go_proto",
        "//third_party/openconfig/gnmi/value",
        "//third_party/openconfig/ondatra/proto:ondatra_go_proto",
        "//third_party/openconfig/ondatra/telemgen/telemgo",
    ]
    if config:
        go_library_deps.append("//third_party/openconfig/ondatra/telemetry")
        helper_method_opts = [
            "-generate_config_func ",
            "-prefer_shadow_path ",
            "-schema_struct_path=google3/third_party/openconfig/ondatra/telemetry/telemetry ",
        ]
    else:
        helper_method_opts = []

    native.go_library(
        name = name,
        srcs = structs_srcs + helpers_srcs,
        deps = go_library_deps,
    )

    exclude_modules = ["ietf-interfaces"] + extra_exclude_modules

    native.genrule(
        name = "generate_helper_methods",
        srcs = common_yang_deps,
        outs = helpers_srcs,
        cmd = "$(locations //third_party/openconfig/ondatra/telemgen/main:main) " +
              "-path=third_party/openconfig/public/release/models,third_party/openconfig/public/third_party " +
              "".join(helper_method_opts) +
              "-package_name={} ".format(name) +
              "-exclude_modules={} ".format(",".join(exclude_modules)) +
              "-output_dir=$(RULEDIR) " +
              "-telemetry_funcs_file_split={} ".format(helper_func_file_split) +
              "-telemetry_types_file_split={} ".format(telem_type_file_split) +
              native.varref("YANG_FILES") +
              " && $(locations //third_party/go:gofmt) -w $(OUTS)" +
              " && //third_party/golang/go_tools/cmd/goimports:goimports -w $(OUTS)",
        tools = [
            "//third_party/go:gofmt",
            "//third_party/golang/go_tools/cmd/goimports",
            "//third_party/openconfig/ondatra/telemgen/main",
        ],
    )

    if config:
        generator_opts = [
            "-generate_structs=false ",
            "-generate_path_structs ",
            "-exclude_state ",
            "-schema_struct_path=google3/third_party/openconfig/ondatra/telemetry/telemetry ",
        ]
    else:  # telemetry
        generator_opts = [
            "-generate_structs ",
            "-generate_path_structs ",
            "-prefer_operational_state ",
            "-list_builder_key_threshold=4 ",
        ]

    native.genrule(
        name = "generate_structs",
        srcs = common_yang_deps,
        outs = structs_srcs,
        cmd = "$(locations //third_party/golang/ygot/generator:generator) " +
              "".join(generator_opts) +
              "-output_dir=$(RULEDIR) -package_name={} ".format(name) +
              "-compress_paths=true -exclude_modules={} ".format(",".join(exclude_modules)) +
              "-generate_fakeroot -fakeroot_name=device " +
              # shadow-path values are ignored during unmarshalling.
              "-ignore_shadow_schema_paths " +
              "-generate_simple_unions " +
              "-shorten_enum_leaf_names " +
              "-typedef_enum_with_defmod " +
              "-enum_suffix_for_simple_union_enums " +
              "-trim_enum_openconfig_prefix " +
              "-include_schema " +
              "-generate_append -generate_getters -generate_rename " +
              "-generate_delete " +
              "-generate_leaf_getters " +
              "-structs_split_files_count={} ".format(structs_file_split) +
              "-path_structs_split_files_count={} ".format(path_structs_file_split) +
              "-path=third_party/openconfig/public/release/models,third_party/openconfig/public/third_party " +
              "-ygot_path=google3/third_party/golang/ygot/ygot/ygot " +
              "-ytypes_path=google3/third_party/golang/ygot/ytypes/ytypes " +
              "-goyang_path=google3/third_party/golang/goyang/pkg/yang/yang " +
              native.varref("YANG_FILES") +
              " && $(locations //third_party/go:gofmt) -w $(OUTS)" +
              " && //third_party/golang/go_tools/cmd/goimports:goimports -w $(OUTS)",
        tools = [
            "//third_party/go:gofmt",
            "//third_party/golang/go_tools/cmd/goimports",
            "//third_party/golang/ygot/generator",
        ],
    )
