package config

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Configs struct {
	ApiGateway ServiceConfig
	Order      ServiceConfig
	Payment    ServiceConfig
	Product    ServiceConfig
	User       ServiceConfig

	OrderDB   DBConfig
	PaymentDB DBConfig
	ProductDB DBConfig
	UserDB    DBConfig
}

type ServiceConfig struct {
	Path string
	Port string
}

type DBConfig struct {
	DSN string
}

var config *Configs
var once sync.Once

func Load() (*Configs, error) {
	var loadErr error
	once.Do(func() {
		config = &Configs{}

		root, err := os.Getwd()
		if err != nil {
			loadErr = err
			return
		}

		envPath := filepath.Join(root, ".env")
		if _, err := os.Stat(envPath); err == nil {
			if err := godotenv.Load(envPath); err != nil {
				loadErr = err
				return
			}
		}

		loadErr = loadConfigFromEnv(config)
	})

	return config, loadErr
}

func loadConfigFromEnv(cfg *Configs) error {
	if err := envconfig.Process("APIGATEWAY", &cfg.ApiGateway); err != nil {
		return errors.New("failed to load ApiGateway config: " + err.Error())
	}
	if err := envconfig.Process("ORDER", &cfg.Order); err != nil {
		return errors.New("failed to load Order config: " + err.Error())
	}
	if err := envconfig.Process("PAYMENT", &cfg.Payment); err != nil {
		return errors.New("failed to load Payment config: " + err.Error())
	}
	if err := envconfig.Process("PRODUCT", &cfg.Product); err != nil {
		return errors.New("failed to load Product config: " + err.Error())
	}
	if err := envconfig.Process("USER", &cfg.User); err != nil {
		return errors.New("failed to load User config: " + err.Error())
	}

	cfg.OrderDB.DSN = os.Getenv("POSTGRE_ORDER")
	cfg.PaymentDB.DSN = os.Getenv("POSTGRE_PAYMENT")
	cfg.ProductDB.DSN = os.Getenv("POSTGRE_PRODUCT")
	cfg.UserDB.DSN = os.Getenv("POSTGRE_USER")

	return nil
}
