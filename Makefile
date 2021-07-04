gazelle:
	bazel run //:gazelle

build:
	bazel build //:zetasql-ast-server

clean:
	bazel clean --expunge

run:
	bazel run //:zetasql-ast-server
