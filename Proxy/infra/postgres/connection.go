package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/pkg/config"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/pkg/log"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func Connection() (*DB, error) {
	var dsn string = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.AppConfig.DatabaseConfig.Host,
		config.AppConfig.DatabaseConfig.Port,
		config.AppConfig.DatabaseConfig.Username,
		config.AppConfig.DatabaseConfig.Password,
		config.AppConfig.DatabaseConfig.Database,
		config.AppConfig.DatabaseConfig.SSLMode,
	)

	con, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Logger.Error("database connection failed")
		return nil, err
	}

	var status bool = false
	for key := range 3 {
		if err := con.Ping(); err != nil {
			log.Logger.Info(fmt.Sprintf("%d. deneme başarısız", key+1))
			time.Sleep(1 * time.Second)
			continue
		}

		status = true
		break
	}

	if !status {
		log.Logger.Error("database connection failed")
		return nil, err
	}

	Db, err := gorm.Open(
		postgres.New(postgres.Config{
			Conn: con,
		}), &gorm.Config{
			// NamingStrategy: schema.NamingStrategy{
			// 	TablePrefix:   config.AppConfig.DatabaseConfig.Schema,
			// 	SingularTable: false,
			// },
		},
	)

	if err != nil {
		log.Logger.Error("database connection failed")
		return nil, err
	}

	return &DB{
		db: Db,
	}, nil
}
