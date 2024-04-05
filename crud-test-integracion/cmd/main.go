package main

import (
	"app/internal/application"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// env
	// ...

	// app
	// - config
	// app := application.NewApplicationDefault("", "./docs/db/json/products.json")
	cfg := &application.ConfigDefault{
		Database: mysql.Config{
			User:      "user1",
			Passwd:    "P4ssw0rd!",
			Net:       "tcp",
			Addr:      "127.0.0.1:3306",
			DBName:    "my_db",
			ParseTime: true,
		},
		Address: "127.0.0.1:8080",
	}
	app := application.NewApplicationDefault(cfg)
	// - tear down
	defer app.TearDown()
	// - set up
	/* if err := app.SetUp(); err != nil {
		fmt.Println(err)
		return
	} */
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
