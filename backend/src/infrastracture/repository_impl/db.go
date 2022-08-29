package repository_impl

import (
	"database/sql"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"

	"hirotoohira.link/time-bucket/infrastracture/model"
)

const (
	driverName          = "mysql"
	dockerComposeDBHost = "db"
	dsn                 = "app:app@tcp(" + dockerComposeDBHost + ":3306)/db"
)

func InitDB() (*gorp.DbMap, error) {
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{
		Db:      db,
		Dialect: gorp.MySQLDialect{},
	}

	addTableWithName(dbmap)

	return dbmap, nil
}

func addTableWithName(dbmap *gorp.DbMap) {
	dbmap.AddTableWithName(model.BucketItem{}, "bucket_items").SetKeys(false, "id")
}
