package Config

import (
	"fmt"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	DBPORT, err := strconv.Atoi(Environment.GetEnv("DBPORT"))
	if err != nil {
		panic(err)
	}
	dbConfig := DBConfig{
		Host:     Environment.GetEnv("HOST"),
		Port:     DBPORT,
		User:     Environment.GetEnv("USER"),
		Password: Environment.GetEnv("PASSWORD"),
		DBName:   Environment.GetEnv("DBNAME"),
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func ConnectDB() (*gorm.DB, error) {
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       DbURL(BuildDBConfig()), // data source name
		DefaultStringSize:         256,                    // default size for string fields
		DisableDatetimePrecision:  true,                   // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                   // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                   // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                  // auto configure based on currently MySQL version
	}), &gorm.Config{})
	return DB, err
}
