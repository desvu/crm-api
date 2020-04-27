package main

import (
	"log"

	"github.com/qilin/crm-api/internal/app"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal()
	}
	zap.ReplaceGlobals(logger)

	a, err := app.New()
	if err != nil {
		zap.S().Fatal(err)
	}

	err = a.Init()
	if err != nil {
		zap.S().Fatal(err)
	}

	err = a.Run()
	if err != nil {
		zap.S().Fatal(err)
	}
}
