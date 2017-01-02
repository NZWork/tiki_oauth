package server

import (
	"fmt"
	"github.com/RangelReale/osin"
	"net/http"
)

// /token?code=ymu84PwfTua06vKqzSkhDw&grant_type=authorization_code&client_id=tiki&client_secret=aabbccdd&redirect_uri=http%3A%2F%2Flocalhost%3A14000%2Fappauth%2Fcode
// code
// grant_type
// client_id
// client_secret
// redirect_uri

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	resp := server.NewResponse()
	defer resp.Close()

	if ar := server.HandleAccessRequest(resp, r); ar != nil {
		switch ar.Type {
		case osin.AUTHORIZATION_CODE:
			ar.Authorized = true
		case osin.REFRESH_TOKEN:
			ar.Authorized = true
		}
		server.FinishAccessRequest(resp, r, ar)
	}
	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}
	if !resp.IsError {
		//		resp.Output["custom_parameter"] = 19923
	}
	osin.OutputJSON(resp, w, r)
}
