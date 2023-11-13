package client

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/fesInformatics/jiro-stamp-rally/infrastructure/config"
	database "github.com/fesInformatics/jiro-stamp-rally/infrastructure/gateway/database"
)

type DbClient struct {
	config *dbconfig.DBconfig
}

func (dc *DbClient) InsertSQL(tablename string, valueMap map[string]any) error {
	var (
		columns string = ""
		here    string = ""
		values         = make([]any, 0, len(valueMap))
	)
	db, err := dc.connectDB()
	if err != nil {
		return err
	}

	for k, v := range valueMap {
		if columns != "" {
			columns += ","
		}
		columns += k
		if here == "" {
			here += "?"
		} else {
			here += ",?"
		}
		values = append(values, v)
	}

	ins, err := db.Prepare(fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", tablename, columns, here))
	if err != nil {
		return err
	}
	_, err = ins.Exec(values...)
	return err
}

func (dc *DbClient) connectDB() (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", dc.config.User, dc.config.Password, dc.config.Host, dc.config.Port, dc.config.DBName))
}

func NewDbClient(config *dbconfig.DBconfig) database.DbClient {
	return &DbClient{
		config: config,
	}
}
