[BACK](../README.md)

```bash
curl -X GET \
    'localhost:8080/api/v1/files?name=photo.jpeg' \
    -o photo_1.jpeg
```

```bash
curl -i -v -X DELETE \
    'localhost:8080/api/v1/files?name=photo.jpeg'
```

```bash
curl -i -v -X PUT \
    'localhost:8080/api/v1/files' \
    -H 'Content-Type: application/octet-stream' \
    -H 'Content-Disposition: attachment; filename=photo.jpeg' \
    --data-binary '@docs/files/photo.jpeg'
```
