load("@rules_cc//cc:defs.bzl", "cc_library")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/naoto0822/zetasql-ast-server
# gazelle:exclude parse_query.cc
# gazelle:exclude parse_query.h
gazelle(name = "gazelle")

go_binary(
    name = "zetasql-ast-server",
    cdeps = [
        ":zetasql-parser",
    ],
    embed = [":zetasql-ast-server_lib"],
    pure = "off",
    visibility = ["//visibility:public"],
)

cc_library(
    name = "zetasql-parser",
    srcs = [
        "parse_query.cc",
        "parse_query.h",
    ],
    deps = [
        "@com_google_zetasql//zetasql/public:parse_helpers",
        "@com_google_zetasql//zetasql/public:sql_formatter",
    ],
)

go_library(
    name = "zetasql-ast-server_lib",
    srcs = [
        "main.go",
        "zetasql.go",
    ],
    cgo = True,
    importpath = "github.com/naoto0822/zetasql-ast-server",
    visibility = ["//visibility:private"],
)
