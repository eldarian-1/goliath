[BACK](../README.md)

```bash
curl -i -v -X GET \
    'localhost:8080/api/v1/cache?key=super'
```

```bash
curl -i -v -X DELETE \
    'localhost:8080/api/v1/cache?key=super'
```

```bash
curl -i -v -X POST \
    'localhost:8080/api/v1/cache?key=super' \
    -H 'Content-Type: application/json' \
    -d '{"name":"Patriot"}'
```
