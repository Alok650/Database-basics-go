# Postgres Basics

1. Created a simple bank schema using dbdiagram.io
2. Exported it as both `.sql` and `.pdf` 
3. Created a postgres12 container using cmd

```
docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
```

4. Use PSQL in the container

```
docker exec -it postgres12 psql -U root
```

## DB migrations
5. Install go-migrate cli tool using brew.

```
brew install go-migrate
```

6. Migration scripts are used for making changes to the database schema being used by an application. It allows the ability to alter/revert DB schema changes based on versioning.

7. Creating initial db schema using migration scripts:

```
migrate create -ext sql -dir migrations -seq init_schema
```

> The above creates two migration files i.e. up and down. Migration scripts follow versioning and provide functionality to update to/revert to a particular schema version.

8. Add the Create table cmds in the up migration script and in the down, add drop table cmds for the tables created.

9. Makefile to help setup the postgres container, DB resources inside it.