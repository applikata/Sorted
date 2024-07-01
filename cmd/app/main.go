package main

import (
	app "arr-sorted/internal"
	conf "arr-sorted/internal/config"
	"fmt"
)

func main() {

	config := conf.ReadConfig()

	fmt.Printf("Listen Port: %s\nListen Addr: %s\n", config.SERVER_PORT, config.LISTEN_ADDR)
	app.StartServer(config)

}
