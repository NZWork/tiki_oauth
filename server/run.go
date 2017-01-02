package server

import (
	"encoding/json"
	"github.com/RangelReale/osin"
	"github.com/ShaleApps/osinredis"
	"github.com/garyburd/redigo/redis"

	"net/http"
)

var server *osin.Server
var storage *osinredis.Storage

func Run(sconfig *osin.ServerConfig, pool *redis.Pool) {
	storage = osinredis.New(pool, "tiki")

	server = osin.NewServer(sconfig, storage)

	// Authorization code endpoint
	http.HandleFunc("/authorize", AuthorizeHandler)
	// Access Token
	http.HandleFunc("/token", TokenHandler)
	// Information endpoint
	http.HandleFunc("/info", InfoHandler)

	// Manage API
	http.HandleFunc("/manage/add", addOAuthClient)
	http.HandleFunc("/manage/update", updateOAuthClient)

	http.HandleFunc("/appauth/code", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		json.NewEncoder(w).Encode(map[string]string{
			"code": r.Form.Get("code"),
			"uid":  r.Form.Get("uid"),
		})
	})

	http.ListenAndServe(":7000", nil)
}
