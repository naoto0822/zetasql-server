load("@rules_cc//cc:defs.bzl", "cc_library")
load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "zetasql-ast-server",
    srcs = [
        "main.go",
        "zetasql.go",
    ],
    cdeps = [
        ":zetasql-parser",
    ],
    cgo = True,
    pure = "off",
)

cc_library(
    name = "zetasql-parser",
    srcs = [
        "parse_query.h",
        "parse_query.cc",
    ],
	  deps=[
        "@com_google_zetasql//zetasql/public:sql_formatter",
        "@com_google_zetasql//zetasql/public:parse_helpers",
        "@com_google_zetasql//zetasql/public:analyzer",
        "@com_google_zetasql//zetasql/public:analyzer_options",
        "@com_google_zetasql//zetasql/public:simple_catalog",
    ],
)

