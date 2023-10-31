package pgsql

import (
	"context"
	"fmt"
	"project/xihe-statistics/config"
	"project/xihe-statistics/utils"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var cli *client

func Initialize(cfg *config.PGSQL) (err error) {
	defer utils.ClearStringMemory(cfg.Password)
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=verify-ca TimeZone=Asia/Shanghai sslrootcert=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.DBCert)
	defer utils.ClearStringMemory(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	cli = &client{
		db: db,
	}

	return
}

type client struct {
	db *gorm.DB
}

func withContext(f func(context.Context) error) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	return f(ctx)
}
