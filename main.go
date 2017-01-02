package main

import (
	"fmt"
	"github.com/RangelReale/osin"
	"os"
	"tiki_oauth/server"
)

func main() {
	pool, err := server.RedisStorage("127.0.0.1", "6379", "1", "3")

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
