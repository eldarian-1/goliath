# goliath

```bash
go run server.go
```

```bash
curl -i -v -X POST \
    localhost:8080/api/v1/echo \
    -H 'Content-Type: application/json' \
    -d '{"name":"Eldarian"}'
```

```bash
docker-compose up -d
```

```bash
curl -i -v -X POST \
    localhost:8080/api/v1/log \
    -H 'Content-Type: application/json' \
    -d '{"level":"INFO","message":"Hello world!"}'
```
