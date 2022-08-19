package main

import (
	"fmt"

	"github.com/yeisonLucio/shopping-cart/src/config"
	"github.com/yeisonLucio/shopping-cart/src/database/connections"
	"github.com/yeisonLucio/shopping-cart/src/database/migrations"
	"github.com/yeisonLucio/shopping-cart/src/helpers"
	"github.com/yeisonLucio/shopping-cart/src/providers"
)

func main() {

	err := helpers.LoadConfig(".env", ".", "env", &config.App)
	if err != nil {
		fmt.Println(err)
	}
	app := providers.Initialize()
	connections.CreateConnection()
	migrations.Execute()
	app.Start()
}
