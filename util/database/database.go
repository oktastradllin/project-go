package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/vwa/util"
)

func Connect() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", util.Cfg.Sqlhost, util.Cfg.Sqlport, util.Cfg.User, util.Cfg.Password, util.Cfg.Dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
