# Rust drop project on Go

### To set up .env file with connections to db's:
##### MongoDB Connection
```
MONGO_DB=api
MONGO_ADDRESS=go-rust-drop-mdb
MONGO_PORT=27017
MONGO_USERNAME=rust-drop
MONGO_PASSWORD="<H;wFO&:L:ym;9"
```
##### MySQL Connection 
```
MYSQL_DB=api
MYSQL_ADDRESS=go-rust-drop-mysql
MYSQL_PORT=3306
MYSQL_USERNAME=root
MYSQL_PASSWORD="6i6Eo0v812;:**'w"
```

### Makefile commands
- `make build` - using by air plugin of live reload
- `make docker` or `make` - build and up project stack
