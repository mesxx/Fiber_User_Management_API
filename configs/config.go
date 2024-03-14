package configs

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConfig() (*gorm.DB, error) {
	var (
		db_host     = os.Getenv("DB_HOST")
		db_username = os.Getenv("DB_USERNAME")
		db_password = os.Getenv("DB_PASSWORD")
		db_name     = os.Getenv("DB_NAME")
		db_port     = os.Getenv("DB_PORT")
	)

	dsn := "host=" + db_host + " user=" + db_username + " password=" + db_password + " dbname=" + db_name + " port=" + db_port + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
