// initializers/database.go

package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// db := os.Getenv("DB_DRIVER")
	dsn := os.Getenv("DB_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})
	// DB, err := gorm.Open(postgres.New(postgres.Config{
	// 	DSN:                  dsn,
	// 	PreferSimpleProtocol: true, // Use simple protocol (disables prepared statements)
	// }), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	// sqlDB, err := DB.DB()
	// if err != nil {
	// 	log.Fatal("Failed to connect to DB:", err)
	// }
	// // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// sqlDB.SetMaxIdleConns(10)

	// // SetMaxOpenConns sets the maximum number of open connections to the database.
	// sqlDB.SetMaxOpenConns(100)

	// // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// sqlDB.SetConnMaxLifetime(time.Hour)

}
