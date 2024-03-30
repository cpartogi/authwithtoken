package init

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/spf13/viper"
)

func ConnectToPgServer() (*pg.DB, error) {

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString(`database.postgres.host`), viper.GetString(`database.postgres.port`)),
		User:     viper.GetString(`database.postgres.user`),
		Password: viper.GetString(`database.postgres.password`),
		Database: viper.GetString(`database.postgres.dbname`),
	})

	_, err := db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}
	return db, err
}
