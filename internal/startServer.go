package app

import (
	"arr-sorted/api"
	config "arr-sorted/internal/config"
	"log"
	"net/http"
	"time"
)

func StartServer(conf config.AppConfig) {
	r := api.Router()

	srv := &http.Server{
		Handler: r,
		Addr:    conf.LISTEN_ADDR + ":" + conf.SERVER_PORT,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
