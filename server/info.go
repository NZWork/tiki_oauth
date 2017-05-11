package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RangelReale/osin"
)

// code
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	resp := server.NewResponse()
	defer resp.Close()

	if ir := server.HandleInfoRequest(resp, r); ir != nil {
		server.FinishInfoRequest(resp, r, ir)
	}
	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}

	if !resp.IsError {
		resp = server.NewResponse()
		if r.Form.Get("uid") != "" { // Load account data from *INNER* API
			response, _ := PostAPI(cfg.API+"user", map[string]interface{}{
				"uid":   r.Form.Get("uid"),
				"xauth": cfg.Secret,
			})

			data := DecodeAPIResponse(response)
			resp.Output["uid"] = data.Result["uid"]
			resp.Output["name"] = data.Result["name"]
			resp.Output["nickname"] = data.Result["nickname"]
			resp.Output["email"] = data.Result["email"]
			resp.Output["avatar"] = data.Result["avatar"]
			resp.Output["status"] = data.Result["status"]
			resp.Output["created_at"] = data.Result["created_at"]
			log.Println(data)
		} else {
			resp.Output["error"] = "invalid uid given"
		}
	}
	osin.OutputJSON(resp, w, r)
}
