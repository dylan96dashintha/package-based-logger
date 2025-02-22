package main

import (
	"context"
	"github.com/package-based-logger/config"
	"github.com/package-based-logger/log"
	"github.com/package-based-logger/product"
)

func main() {
	config.InitLogConfig()
	logger := log.NewLogger()

	// initialize the global logger
	gLog, err := logger.GetLogger(config.LogConf)
	if err != nil {
		return
	}

	// create a product object
	prd, err := product.NewProduct(logger, config.LogConf)
	if err != nil {
		gLog.Err(err).Msg("unable in initializing product")
		return
	}

	ctx := context.Background()
	prd.SetProduct(ctx, 231, "soap")

	// get the updated product
	prd.GetProduct(ctx, 231)
	prd.GetProduct(ctx, 345)
}
