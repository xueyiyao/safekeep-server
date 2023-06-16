package main

import (
	"github.com/xueyiyao/safekeep/initializers"
	"github.com/xueyiyao/safekeep/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Container{})
	initializers.DB.AutoMigrate(&models.Item{})
}
