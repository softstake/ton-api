package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type TonAPIConfig struct {
	TonlibCfgPath string `env:"TONLIB_CFG_PATH" envDefault:"/usr/local/bin/app/tonlib.config.json.example"`
	ListenPort    int32  `env:"LISTEN_PORT" envDefault:"5400"`
	ContractAddr  string `env:"CONTRACT_ADDR,required"`
}

func GetConfig() TonAPIConfig {
	cfg := &TonAPIConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal("Cannot parse initial ENV vars: ", err)
	}
	return *cfg
}
