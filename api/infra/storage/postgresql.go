package storage

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	DB   *gorm.DB
	PG   *PostgreSQL
	once = sync.Once{}
)

type PostgreSQL struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (pg PostgreSQL) DSN() string {
	return fmt.
		Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			pg.Host, pg.Username, pg.Password, pg.Database, pg.Port,
		)
}

func NewPostgres() *gorm.DB {
	once.Do(func() {
		PG = new(PostgreSQL)
		err := viper.UnmarshalKey("postgresql", PG)
		if err != nil {
			logrus.Panic("unable to unmarshal postgres config: ", err)
		}
		DB, err = gorm.Open(postgres.Open(PG.DSN()), &gorm.Config{})
		if err != nil {
			logrus.Panic("failed to connect to PostgreSQL: ", err)
		}
	})
	return DB
}
