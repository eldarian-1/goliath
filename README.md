# goliath

```bash
go run server.go
```

```bash
curl -i -v -X POST \
    localhost:8080/echo-user \
    -H 'Content-Type: application/json' \
    -d '{"name":"Eldarian"}'
```
