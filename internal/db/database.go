package db

import (
	"fmt"
	"time"

	"github.com/asadbek21coder/demoproject/config"
	"github.com/jmoiron/sqlx"
)

// type Config struct {
// 	Host     string
// 	Port     string
// 	Username string
// 	Password string
// 	DBName   string
// 	SSLMode  string
// }

func ConnectToDb(cfg config.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)
	for cfg.PostgresConnectionTry > 0 {
		c, err := sqlx.Connect("postgres", psqlString)
		if err == nil {
			return c, nil
		}
		fmt.Println(err)
		time.Sleep(time.Second * time.Duration(cfg.PostgresConnectionTimeOut))
		cfg.PostgresConnectionTry--
	}
	return sqlx.Connect("postgres", psqlString)
}
