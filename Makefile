# VAR
MYSQL_CMD = mysql --host=127.0.0.1 --port=3306 --user=root --password=root
DSN = mysql://root:root@tcp(127.0.0.1:3306)
DATABASE_NAME = db
DB_OPTIONS = ?parseTime=true&loc=Asia%2FTokyo&interpolateParams=true
DB_CONNECTION = $(DSN)/$(DATABASE_NAME)$(DB_OPTIONS)

MIGRATION_DIR = db/migrations

# RECIPE
db_diff:
	schemadiff db/schema.sql '$(DB_CONNECTION)'

db_migrate:
	@$(MYSQL_CMD) -e "CREATE DATABASE IF NOT EXISTS $(DATABASE_NAME)" && \
	cat db/schema.sql | schemalex '$(DB_CONNECTION)' - > tmp/schema.sql && \
	$(MYSQL_CMD) $(DATABASE_NAME) < tmp/schema.sql

create_migration:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq create_table

migrate_db:
	migrate -path $(MIGRATION_DIR) -database '$(DB_CONNECTION)' up

force_to_migration_version:
	migrate -path $(MIGRATION_DIR) -database '$(DB_CONNECTION)' force 1
