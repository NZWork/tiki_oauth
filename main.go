package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/RangelReale/osin"
	"os"
	"tiki_oauth/server"
)

func main() {

	config := server.TikiConfig{}

	var configFile = flag.String("c", "config", "config file name(current only support toml format, and without .toml)")

	if _, err := toml.DecodeFile(*configFile+".toml", &config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Loading config...")
	server.SetConfig(&config)

	pool, err := server.RedisStorage()
	fmt.Println("OAuth storage[redis] connecting...")

	if err != nil {
		fmt.Println(fmt.Errorf("Fail to initialize redis service: %s", err.Error()))
		os.Exit(1)
	}

	sconfig := osin.NewServerConfig()
	sconfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE}
	sconfig.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE, osin.REFRESH_TOKEN}
	sconfig.AllowGetAccessRequest = true
	sconfig.AllowClientSecretInParams = true
	sconfig.RedirectUriSeparator = "#" // 多回调地址支持，分隔符

	fmt.Println("Initialized Tiki OAuth Service\nHype Life!")
	server.Run(sconfig, pool)
}
