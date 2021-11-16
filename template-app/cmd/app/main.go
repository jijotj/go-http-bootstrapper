package main

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"{{.ServiceName}}/pkg/config"
	"{{.ServiceName}}/pkg/router"
)

func main() {
	appConfig, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("Init config")
	}

	r := router.New()

	serverTimeout := time.Duration(appConfig.ServerTimeOut) * time.Millisecond
	writeTimeout := serverTimeout + 2*time.Second
	srv := http.Server{
		Addr:         ":" + appConfig.AppPort,
		WriteTimeout: writeTimeout,
		Handler:      http.TimeoutHandler(r, serverTimeout, ""),
	}

	log.Print("Starting {{.ServiceName}} service...")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Start http server")
	}
}
