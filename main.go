package main

import (
	"github.com/facundocenteno/bookstore_users-api/app"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app.StartApplication()

}
