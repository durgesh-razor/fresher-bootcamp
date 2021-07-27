## Setup Database

```bash
docker run --platform linux/x86_64 --name MySQL-RetailerRecord -v $(pwd)/databases:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=durgesh -e MYSQL_PASSWORD=durgesh -e MYSQL_DATABASE=RetailerRecord -p 3306:3306 mysql
```

## Run Development Server

```bash
docker run --env-file=./.env.dev --link MySQL-RetailerRecord --name StudentRecord-Service -v $(pwd):/app -p 8080:8080 durgeshrazor/retailerservice
```

## Run Production Server

```bash
docker run --env-file=./.env.dev --link MySQL-RetailerRecord --name StudentRecord-Service -v $(pwd):/app -p 8080:8080 -d durgeshrazor/retailerservice bash -c "go run ."
```