# SQL database

SQL_User_name := admin
SQL_Password := admin1234
SQL_Host := 127.0.0.1
SQL_Port := 5432
SQL_Database := flareDB
SQL_SSL_mode := disable
Migration_file_path := internal/db/migration
PostgreSQL_address := postgres://$(SQL_User_name):$(SQL_Password)@$(SQL_Host):$(SQL_Port)/$(SQL_Database)?sslmode=$(SQL_SSL_mode)


migrate-up:	
	migrate -path $(Migration_file_path) -database $(PostgreSQL_address) -verbose up

migrate-down:
	migrate -path $(Migration_file_path) -database $(PostgreSQL_address) -verbose down


TABLE := table

migrate-create:
	migrate create -ext sql -dir $(Migration_file_path) -seq create_$(TABLE)_table