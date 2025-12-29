[BACK](../README.md)

```bash
curl -i -v -X POST \
    localhost:8080/api/v1/log \
    -H 'Content-Type: application/json' \
    -d '{"level":"INFO","message":"Hello world!"}'
```
