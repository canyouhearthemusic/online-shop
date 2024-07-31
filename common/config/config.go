package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// TODO: Make a singleton

type Configs struct {
	ApiGateway app
	Order      app
	Payment    app
	Product    app
	User       app

	OrderDB   DB
	PaymentDB DB
	ProductDB DB
	UserDB    DB
}

type DB struct {
	DSN string
}

type app struct {
	Port string
	Path string
}

func New() (cfg Configs, err error) {
	root, err := os.Getwd()
	if err != nil {
		return
	}

	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		cfg.ApiGateway.Path = os.Getenv("APIGATEWAY_PATH")
		cfg.Order.Path = os.Getenv("ORDER_PATH")
		cfg.Payment.Path = os.Getenv("PAYMENT_PATH")
		cfg.Product.Path = os.Getenv("PRODUCT_PATH")
		cfg.User.Path = os.Getenv("USER_PATH")

		cfg.ApiGateway.Port = os.Getenv("APIGATEWAY_PORT")
		cfg.Order.Port = os.Getenv("ORDER_PORT")
		cfg.Payment.Port = os.Getenv("PAYMENT_PORT")
		cfg.Product.Port = os.Getenv("PRODUCT_PORT")
		cfg.User.Port = os.Getenv("USER_PORT")

		cfg.OrderDB.DSN = os.Getenv("POSTGRE_ORDER_DSN")
		cfg.PaymentDB.DSN = os.Getenv("POSTGRE_PAYMENT_DSN")
		cfg.ProductDB.DSN = os.Getenv("POSTGRE_PRODUCT_DSN")
		cfg.UserDB.DSN = os.Getenv("POSTGRE_USER_DSN")

		return cfg, nil
	}

	if err = envconfig.Process("APIGATEWAY", &cfg.ApiGateway); err != nil {
		return
	}

	if err = envconfig.Process("ORDER", &cfg.Order); err != nil {
		return
	}

	if err = envconfig.Process("PAYMENT", &cfg.Payment); err != nil {
		return
	}

	if err = envconfig.Process("PRODUCT", &cfg.Product); err != nil {
		return
	}

	return
}
