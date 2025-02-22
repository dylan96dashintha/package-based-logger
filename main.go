package main

import (
	"context"
	"github.com/package-based-logger/cart"
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

	// create a cart object
	crt, err := cart.NewCart(logger, config.LogConf)
	if err != nil {
		gLog.Err(err).Msg("unable in initializing cart")
		return
	}
	ctx := context.Background()
	prd.SetProduct(ctx, 231, "soap")

	// get the updated product
	prd.GetProduct(ctx, 231)

	// get not added item
	prd.GetProduct(ctx, 345)

	// add the product to a cart
	crt.AddProduct(ctx, 231)

	// remove the item
	crt.RemoveProduct(ctx, 231)

	// remove not added item
	crt.RemoveProduct(ctx, 345)

}
