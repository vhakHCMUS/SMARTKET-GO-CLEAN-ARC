package main

import (
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/bootstrap"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	bootstrap.RootApp.Execute()
}
