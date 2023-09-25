# fiber-sqlboiler-exam
[![Go](https://img.shields.io/badge/go-1.20-blue.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/dl/)

> #### 🎯 docker-compose, fiber, sqlboiler (postgresql) 를 사용한 project example 입니다.
> #### 🎯 [Sqlboiler](https://github.com/volatiletech/sqlboiler) + [Fiber](https://github.com/gofiber/fiber) Project Example

## ⚙️ Installation
```shell
$ git clone https://github.com/LimJiAn/fiber-sqlboiler-exam
```
## 👀 Usage
#### 1. Run Postgres
```bash
$ docker compose build
```
```bash
$ docker compose up
```
#### 2. Wait 1-2 minutes
```console
[+] Running 2/0
 ✔ Network fiber-sqlboiler-exam_default  Created                                                                             0.1s
 ✔ Container postgres                    Created                                                                             0.0s
Attaching to postgres
postgres  |
postgres  | PostgreSQL Database directory appears to contain a database; Skipping initialization
postgres  |
postgres  |
postgres  | 2023-09-21 09:17:50.737 UTC [1] LOG:  starting PostgreSQL 16.0 (Debian 16.0-1.pgdg120+1) on aarch64-unknown-linux-gnu, compiled by gcc (Debian 12.2.0-14) 12.2.0, 64-bit
postgres  | 2023-09-21 09:17:50.737 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
postgres  | 2023-09-21 09:17:50.737 UTC [1] LOG:  listening on IPv6 address "::", port 5432
postgres  | 2023-09-21 09:17:50.740 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
postgres  | 2023-09-21 09:17:50.751 UTC [30] LOG:  database system was shut down at 2023-09-21 08:50:35 UTC
postgres  | 2023-09-21 09:17:50.770 UTC [1] LOG:  database system is ready to accept connections
```
#### 3. You have to migrate the database.
> ##### 🎯 It is a "database-first" ORM as opposed to "code-first" (like gorm/gorp). That means you must first create your database schema.
> ##### 🎯 I used [golang-migrate](https://github.com/golang-migrate/migrate) to proceed with the migrate.
###### 1. Make Migration files
```bash
$ migrate create -ext sql -dir ./migrations -seq create_initial_table
```
```console
fiber-sqlboiler-exam/migrations/000001_create_initial_table.up.sql
fiber-sqlboiler-exam/migrations/000001_create_initial_table.up.sql
```
###### 2. Migrate
```bash
$ migrate -path migrations -database "postgresql://testuser:testpass@localhost:5432/testdb?sslmode=disable" -verbose up
```
```console
2023/09/21 20:00:00 Start buffering 1/u create_initial_table
2023/09/21 20:00:00 Read and execute 1/u create_initial_table
2023/09/21 20:00:00 Finished 1/u create_initial_table (read 24.693541ms, ran 68.30925ms)
2023/09/21 20:00:00 Finished after 100.661625ms
2023/09/21 20:00:00 Closing source and database
```
###### 3. Rollback Migrate
```bash
$ migrate -path migrations -database "postgresql://testuser:testpass@localhost:5432/testdb?sslmode=disable" -verbose down
```
```console
2023/09/21 20:00:00 Are you sure you want to apply all down migrations? [y/N]
y
2023/09/21 20:00:00 Applying all down migrations
2023/09/21 20:00:00 Start buffering 1/d create_initial_table
2023/09/21 20:00:00 Read and execute 1/d create_initial_table
2023/09/21 20:00:00 Finished 1/d create_initial_table (read 39.681125ms, ran 66.220125ms)
2023/09/21 20:00:00 Finished after 1.83152475s
```
#### 4. Use sqlboiler
###### 1. Install
```bash
# Go 1.16 and above:
$ go install github.com/volatiletech/sqlboiler/v4@latest
$ go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
```
###### 2. Create a configuration file
> ###### 🎯 The configuration file should be named sqlboiler.toml
###### Example
```toml
output   = "models"
wipe     = true
no-tests = true
add-enum-types = true

[psql]
  dbname = "testdb"
  host   = "localhost"
  port   = 5432
  user   = "testuser"
  pass   = "testpass"
  schema = "schema"
  blacklist = ["migrations", "other"]
```
###### 3. Create models
> ###### 🎯 After creating a configuration file that points at the database we want to generate models for, we can invoke the sqlboiler command line utility.
```bash
$ sqlboiler psql
```
```text
models/
├── author.go
├── boil_queries.go
├── boil_table_names.go
├── boil_types.go
├── boil_view_names.go
├── post.go
├── schema_migrations.go
```

