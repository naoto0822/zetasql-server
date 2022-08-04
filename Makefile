gazelle:
	bazel run //:gazelle

build: gazelle
	bazel build //:zetasql-server

clean:
	bazel clean --expunge

run: build
	bazel run //:zetasql-server

.PHONY: gazelle build clean run
