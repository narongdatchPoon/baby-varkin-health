// migrate/migrate.go

package main

import (
	"baby-varkin-health/initializers"
	"baby-varkin-health/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Activities{})
}
