# goliath

```bash
go run server.go
```

```bash
curl -i -v -X GET \
    'localhost:8080/api/v1/users'
```

```bash
curl -i -v -X DELETE \
    'localhost:8080/api/v1/users?id=1'
```

```bash
curl -i -v -X POST \
    localhost:8080/api/v1/users \
    -H 'Content-Type: application/json' \
    -d '{"name":"Eldarian"}'
```

```bash
curl -i -v -X POST \
    localhost:8080/api/v1/users \
    -H 'Content-Type: application/json' \
    -d '{"id":1,"name":"Eldarian"}'
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
