# Rust drop project on Go

## To set up .env file with connections to db's:
##### MongoDB Connection
``` env
MONGODB_HOST=go-rust-drop-mdb
MONGODB_PORT=27017
MONGODB_DBNAME=api
MONGODB_USER=rust-drop
MONGODB_PASSWORD="<H;wFO&:L:ym;9"
MONGODB_AUTH_MECHANISM=SCRAM-SHA-1
MONGODB_AUTH_DATABASE=admin
```
##### MySQL Connection 
``` env
MYSQL_HOST=go-rust-drop-mysql
MYSQL_PORT=3306
MYSQL_DBNAME=api
MYSQL_USER=root
MYSQL_PASSWORD="6i6Eo0v812"
```

## Makefile commands
- `make build` - using by air plugin of live reload
- `make docker` or `make` - build and up project stack

## Documentation for creating newest data in project
- [Creating migrations for MySQL](https://github.com/popcornrus/go-rust-drop/issues/1)
- [Creating Enum](https://github.com/popcornrus/go-rust-drop/issues/2)
