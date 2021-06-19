docker-pull:
	docker pull naoto0822/zetasql-ast-server:latest

docker-run:
	docker run -p 8080:8080 -t naoto0822/zetasql-ast-server

curl-valid:
	curl -X POST -H 'Content-type: application/text' --data 'SELECT 10, 10 from hoge where id = 1' http://localhost:8080/valid
