package product

import (
	"context"
	"errors"
	"fmt"
	"github.com/package-based-logger/config"
	"github.com/package-based-logger/log"
	"github.com/package-based-logger/util"
	logger "github.com/rs/zerolog"
)

type Product interface {
	GetProduct(ctx context.Context, id int64) (err error)
	SetProduct(ctx context.Context, id int64, value string)
}

type prd struct {
	logger logger.Logger
	list   map[int64]string
}

func NewProduct(log log.Logger, logConf config.LogConfig) (productObj Product, err error) {
	name, exist := util.GetPackageName()
	if !exist {
		return productObj, errors.New("package name could not found")
	}
	logObj, err := log.GetPackageBasedLogger(name, logConf)
	if err != nil {
		return productObj, err
	}
	logObj = logObj.With().Str("prefix", "Product").Logger()
	prd := prd{
		logger: logObj,
		list:   make(map[int64]string),
	}

	return &prd, nil
}

func (p prd) GetProduct(ctx context.Context, id int64) (err error) {
	name, ok := p.list[id]
	if !ok {
		er := errors.New("error in finding the product")
		p.logger.Error().Err(er).Msg("product is missing")
		return er
	}

	p.logger.Info().Msg(fmt.Sprintf("product is found, name: %v", name))
	return nil
}

func (p prd) SetProduct(ctx context.Context, id int64, value string) {
	p.list[id] = value
}
