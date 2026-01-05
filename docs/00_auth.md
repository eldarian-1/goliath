[BACK](../README.md)

```bash
curl -i -v -X POST \
    'localhost:8080/api/v1/auth/register' \
    -c cookies.txt \
    -H 'Content-Type: application/json' \
    -d '{"email":"admin@yandex.ru","password":"admin1234"}'
```

```bash
curl -i -v -X POST \
    'localhost:8080/api/v1/auth/login' \
    -c cookies.txt \
    -H 'Content-Type: application/json' \
    -d '{"email":"admin@yandex.ru","password":"admin1234"}'
```

```bash
curl -i -v -X GET \
    'localhost:8080/api/v1/auth/me' \
    -b cookies.txt
```

```bash
curl -i -v -X POST \
    'localhost:8080/api/v1/auth/refresh' \
    -b cookies.txt
```

```bash
curl -i -v -X POST \
    'localhost:8080/api/v1/auth/logout' \
    -b cookies.txt \
    -c cookies.txt
```
