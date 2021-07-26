### Setup Database pod

```bash
docker run --platform linux/x86_64 --name MySQL-StudentRecord -v $(pwd)/databases:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=durgesh -e MYSQL_PASSWORD=durgesh -e MYSQL_DATABASE=StudentRecord -p 3306:3306 mysql
```


### Spin Out the Services

```bash
docker run --env-file=./.env.dev --link MySQL-StudentRecord --name StudentRecord-Service -v $(pwd):/app -p 8080:8080 durgeshrazor/studentrecord
```