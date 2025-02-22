package cart

import (
	"context"
	"errors"
	"fmt"
	"github.com/package-based-logger/config"
	"github.com/package-based-logger/log"
	"github.com/package-based-logger/util"
	logger "github.com/rs/zerolog"
)

type Cart interface {
	AddProduct(ctx context.Context, id int64)
	RemoveProduct(ctx context.Context, id int64) (err error)
}

type crt struct {
	logger logger.Logger
	list   map[int64]interface{}
}

func NewCart(log log.Logger, logConf config.LogConfig) (cartObj Cart, err error) {
	name, exist := util.GetPackageName()
	if !exist {
		return cartObj, errors.New("package name could not found")
	}
	logObj, err := log.GetPackageBasedLogger(name, logConf)
	if err != nil {
		return cartObj, err
	}
	logObj = logObj.With().Str("prefix", "cart").Logger()
	crt := crt{
		logger: logObj,
		list:   make(map[int64]interface{}),
	}

	return &crt, nil
}

func (c *crt) AddProduct(ctx context.Context, product int64) {
	c.list[product] = new(interface{})
	c.logger.Info().Msg(fmt.Sprintf("succesfully added the product, id:%d", product))
}

func (c *crt) RemoveProduct(ctx context.Context, productId int64) (err error) {
	_, exist := c.list[productId]
	if !exist {
		err := errors.New("product is not added to the cart")
		c.logger.Err(err).Msg("product does not exist")
		return err
	}

	delete(c.list, productId)
	c.logger.Info().Msg("successfully removed the product from the cart")
	return nil
}
