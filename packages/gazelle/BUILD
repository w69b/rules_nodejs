load("@bazel_gazelle//:def.bzl", "DEFAULT_LANGUAGES", "gazelle", "gazelle_binary")
load("@io_bazel_rules_go//go:def.bzl", "go_library")

# ts_library:load //ev/tooling/bazel:defaults.bzl
gazelle(
    name = "gazelle",
    gazelle = ":gazelle_ts",
    args = [
        "-build_file_name",
        "BUILD",
    ],
)

gazelle_binary(
    name = "gazelle_ts",
    languages = [
        ":go_default_library",
    ],
    visibility = [
        "//visibility:public",
    ],
)

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "ts.go",
        "resolve.go",
    ],
    importpath = "github.com/evertz/bzlgen/gazelle",
    visibility = ["//visibility:public"],
    deps = [
        "@bazel_gazelle//config:go_default_library",
        "@bazel_gazelle//label:go_default_library",
        "@bazel_gazelle//language:go_default_library",
        "@bazel_gazelle//repo:go_default_library",
        "@bazel_gazelle//resolve:go_default_library",
        "@bazel_gazelle//rule:go_default_library",
    ],
)
