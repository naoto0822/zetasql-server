# zetasql-ast-server

[![Go](https://github.com/naoto0822/zetasql-ast-server/actions/workflows/go.yml/badge.svg)](https://github.com/naoto0822/zetasql-ast-server/actions/workflows/go.yml)
[![Docker Automated build](https://img.shields.io/docker/automated/naoto0822/zetasql-ast-server.svg?style=flat-square)](https://hub.docker.com/r/naoto0822/zetasql-ast-server/)
[![Docker Pulls](https://img.shields.io/docker/pulls/naoto0822/zetasql-ast-server.svg?style=flat-square)](https://hub.docker.com/r/naoto0822/zetasql-ast-server/)

zetasql-ast-server is parsing/analyzing ZetaSQL using [google/zetasql](https://github.com/google/zetasql) framework.

## Usage

```sh
$ docker pull naoto0822/zetasql-ast-server:latest
$ docker run -p 8080:8080 -t naoto0822/zetasql-ast-server
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
```

### `/valid`

Returns whether it is correct.

- example
```sh
$ curl -X POST -H 'Content-type: application/text' --data 'SELECT 1' http://localhost:8080/valid
OK
```

## Features

- [x] Format sql
- [ ] Analyze SQL
- [ ] AST to JSON (DebugString to struct)
- [ ] ~cli tool~ (-> new project)
- [ ] ~Hosting VPS or Web assembly~
- [ ] gazelle
- [ ] zetasql base image

## References

- [google/zetasql](https://github.com/google/zetasql)
- [ebendutoit/zetasql-analyzer-server](https://github.com/ebendutoit/zetasql-analyzer-server)
- [apstndb/zetasql-format-server](https://github.com/apstndb/zetasql-format-server)
