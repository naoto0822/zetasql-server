# zetasql-server

[![Go](https://github.com/naoto0822/zetasql-server/actions/workflows/go.yml/badge.svg)](https://github.com/naoto0822/zetasql-server/actions/workflows/go.yml)
[![Docker Automated build](https://img.shields.io/docker/automated/naoto0822/zetasql-server.svg?style=flat-square)](https://hub.docker.com/r/naoto0822/zetasql-server/)
[![Docker Pulls](https://img.shields.io/docker/pulls/naoto0822/zetasql-server.svg?style=flat-square)](https://hub.docker.com/r/naoto0822/zetasql-server/)

[WIP] zetasql-server is parsing/analyzing ZetaSQL using [google/zetasql](https://github.com/google/zetasql) framework.

## Usage

```sh
$ docker pull naoto0822/zetasql-server:latest
$ docker run -p 8080:8080 -t naoto0822/zetasql-server
```

## API

### `/parse`

> TODO: DebugString to JSON format...

Returns a list of ZetaSQL AST DebugString or error message.

- example
```sh
$ curl -X POST -H 'Content-type: application/text' --data 'SELECT 1' http://localhost:8080/parse
QueryStatement [0-8]
  Query [0-8]
    Select [0-8]
      SelectList [7-8]
        SelectColumn [7-8]
          IntLiteral(1) [7-8]

$ curl -X POST -H 'Content-type: application/text' --data 'SEL 1' http://localhost:8080/parse
generic::invalid_argument: Syntax error: Unexpected identifier "SEL" [zetasql.ErrorLocation] { line: 1 column: 1 }
```

### `/format`

Format SQL.

- example
```sh
$ curl -X POST -H 'Content-type: application/text' --data 'SELECT * FROM hoge' http://localhost:8080/format
SELECT
  *
FROM
  hoge;

$ curl -X POST -H 'Content-type: application/text' --data 'SEL 1' http://localhost:8080/format
generic::invalid_argument: Syntax error: Unexpected identifier "SEL" [at 1:1]
SEL 1
^
```

### `/valid`

Returns whether it is correct.

- example
```sh
$ curl -X POST -H 'Content-type: application/text' --data 'SELECT 1' http://localhost:8080/valid
OK

$ curl -X POST -H 'Content-type: application/text' --data 'SEL 1' http://localhost:8080/valid
generic::invalid_argument: Syntax error: Unexpected identifier "SEL" [zetasql.ErrorLocation] { line: 1 column: 1 }
```

## Development Run

```
$ make run
```

## Release

If creating Github Release of new tag, CI push new docker image.

## Features

- [x] Format sql
- [ ] Multi Statements
- [ ] Analyze SQL
- [ ] AST to JSON (DebugString to struct)
- [x] cli tool -> https://github.com/naoto0822/zetasql-cli (still private...)
- [ ] ~Hosting VPS or Web assembly~
- [x] gazelle
- [ ] zetasql base image

## Dependencies

- bazel

## References

- [google/zetasql](https://github.com/google/zetasql)
- [ebendutoit/zetasql-analyzer-server](https://github.com/ebendutoit/zetasql-analyzer-server)
- [apstndb/zetasql-format-server](https://github.com/apstndb/zetasql-format-server)
