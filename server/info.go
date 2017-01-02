package server

import (
	"fmt"
	"github.com/RangelReale/osin"
	"net/http"
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

	if !resp.IsError && r.Form.Get("uid") != "" { // Load account data from *INNER* API
		resp = server.NewResponse()
		resp.Output["uid"] = r.Form.Get("uid")
		resp.Output["name"] = "Neo"
		resp.Output["profile"] = "https://nex.tw"
	}
	osin.OutputJSON(resp, w, r)
}
