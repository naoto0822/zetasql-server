# zetasql-ast-server

[![Go](https://github.com/naoto0822/zetasql-ast-server/actions/workflows/go.yml/badge.svg)](https://github.com/naoto0822/zetasql-ast-server/actions/workflows/go.yml)
[![Docker Automated build](https://img.shields.io/docker/automated/naoto0822/zetasql-ast-server.svg?style=flat-square)](https://hub.docker.com/r/naoto0822/zetasql-ast-server/)
[![Docker Pulls](https://img.shields.io/docker/pulls/naoto0822/zetasql-ast-server.svg?style=flat-square)](https://hub.docker.com/r/naoto0822/zetasql-ast-server/)

zetasql-ast-server is analyzing SQL using [ZetaSQL](https://github.com/google/zetasql).

## Usage

```sh
$ docker pull naoto0822/zetasql-ast-server:latest
$ docker run -p 8080:8080 -t naoto0822/zetasql-ast-server
```

## API

### `/valid`

- example
```sh
$ curl -X POST -H 'Content-type: application/text' --data 'SELECT 1' http://localhost:8080/valid
OK
```

### `/parse`

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

## Features

- [ ] Analyze SQL
- [ ] AST to JSON (DebugString to struct)
- [ ] Hosting VPS or Web assembly

## References

- [gogle/zetasql](https://github.com/google/zetasql)
- [ebendutoit/zetasql-analyzer-server](https://github.com/ebendutoit/zetasql-analyzer-server)
- [apstndb/zetasql-format-server](https://github.com/apstndb/zetasql-format-server)
