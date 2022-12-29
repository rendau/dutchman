# Dutchman

### DB dump:

```
pg_dump --no-owner -Fc -U postgres dutchman -f ./dutchman.custom
```

### DB restore:

```
dropdb -U postgres dutchman
createdb -U postgres dutchman
pg_restore --no-owner -d dutchman -U postgres ./dutchman.custom
```

### Install `migrate` command-tool:

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

### Create new migration:

```
migrate create -ext sql -dir migrations mg_name
```

### Apply migration:

```
migrate -path migrations -database "postgres://localhost:5432/db_name?sslmode=disable" up
```
