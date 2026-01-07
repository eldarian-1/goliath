# goliath

[0 Auth](docs/00_auth.md)

[1 Log](docs/01_log.md)

[2 Users](docs/02_users.md)

[3 Files](docs/03_files.md)

[4 Cache](docs/04_cache.md)

```bash
docker-compose \
  -f docker/backend.yml \
  -f docker/repositories.yml \
  -f docker/brokers.yml \
   up
```

```bash
docker-compose \
  -f docker/backend.yml \
  -f docker/repositories.yml \
  -f docker/brokers.yml \
  -f docker/infrastructure.yml \
   up
```

```bash
docker-compose \
  -f docker/backend.yml \
  -f docker/repositories.yml \
  -f docker/brokers.yml \
  -f docker/infrastructure.yml \
   down
```

```bash
docker-compose -f docker/frontend.yml up --build
```
