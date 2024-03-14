# keploy-example

Run MySQL

```
docker compose up -d
```

Record test

```sh
$ keploy record -c "go run ." --debug
```

Send request

```sh
$ curl -X POST localhost:8080/todos -d '{"name": "John", "title": "foo", "status": "todo"}'
$ curl localhost:8080/todos
```

Run recorded tests

```sh
$ keploy test -c "go run ." --delay 10
```
