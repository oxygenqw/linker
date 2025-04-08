# linker


## Docker

``` bash
docker build -t linker .
docker run -d -p 8080:8080 --name linker_app linker
```


## Tunnel
lt --port 8080 --subdomain linker

## Postgres
psql -h localhost -p 5432 -U linker -d linker_db

## TODO:

context

graceful shutdown +

docker-compose +

yaml

Makefile

Logger
