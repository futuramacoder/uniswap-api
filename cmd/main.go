package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"

	"github.com/futuramacoder/uniswap-api/app/api"
	"github.com/futuramacoder/uniswap-api/app/config"
	"github.com/futuramacoder/uniswap-api/app/pkg/client/eth"
	"github.com/futuramacoder/uniswap-api/app/service/uniswap"
)

func init() {
	time.Local = time.UTC

	_ = godotenv.Load()
}

func main() {
	envConfig, err := config.LoadEnvConfig()
	if err != nil {
		log.WithError(err).Fatal("error processing application env config")
	}

	apiConfig := api.Config{
		LogLevel:    envConfig.LogLevel,
		Port:        envConfig.Port,
		CorsOrigins: envConfig.CorsOrigins,
		CorsMethods: envConfig.CorsMethods,
	}

	ethClient, err := ethclient.Dial(envConfig.NodeUrl)
	if err != nil {
		log.WithError(err).Fatal("error connecting to eth client")
	}

	client := eth.NewClient(ethClient)

	uniswapSvc := uniswap.NewService(client)

	server, err := api.NewServer(apiConfig, uniswapSvc)
	if err != nil {
		log.WithError(err).Fatal("error to configure api server")
	}

	errCh := make(chan error, 1)

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
		errCh <- fmt.Errorf("%s", <-sigs)
	}()

	go func() {
		err = server.Start()
		if err != nil {
			log.WithError(err).Errorf("api server stopped")

			errCh <- err
		}
	}()

	check(<-errCh)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
