gazelle:
	bazel run //:gazelle

build:
	bazel build //:zetasql-server

clean:
	bazel clean --expunge

run:
	bazel run //:zetasql-server

.PHONY: gazelle build clean run
