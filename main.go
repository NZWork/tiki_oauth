package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/RangelReale/osin"
	"os"
	"tiki_oauth/server"
)

func main() {

	config := server.RedisConfig{}
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	pool, err := server.RedisStorage(&config)

	if err != nil {
		fmt.Println(fmt.Errorf("Fail to initialize redis service: %s", err.Error()))
		os.Exit(1)
	}

	sconfig := osin.NewServerConfig()
	sconfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE}
	sconfig.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE, osin.REFRESH_TOKEN}
	sconfig.AllowGetAccessRequest = true
	sconfig.AllowClientSecretInParams = true

	server.Run(sconfig, pool)
}
