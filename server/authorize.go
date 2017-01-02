package server

import (
	"fmt"
	"github.com/RangelReale/osin"
	"net/http"
)

// @params
// authorize?response_type=code&client_id=tiki&redirect_uri=http%3A%2F%2Flocalhost%3A14000%2Fappauth%2Fcode
// ------------------
// response_type
// client_id
// redirect_uri

func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	resp := server.NewResponse()
	defer resp.Close()

	if ar := server.HandleAuthorizeRequest(resp, r); ar != nil {
		r.ParseForm()
		if r.Method != "POST" || r.Form.Get("login") != "test" || r.Form.Get("password") != "test" {
			w.Write([]byte("<html><body>"))
			w.Write([]byte(fmt.Sprintf("LOGIN %s (use test/test)<br/>", ar.Client.GetId())))
			w.Write([]byte(fmt.Sprintf("<form action=\"/authorize?%s\" method=\"POST\">", r.URL.RawQuery)))
			w.Write([]byte("Login: <input type=\"text\" name=\"login\" /><br/>"))
			w.Write([]byte("Password: <input type=\"password\" name=\"password\" /><br/>"))
			w.Write([]byte("<input type=\"submit\"/>"))
			w.Write([]byte("</form>"))
			w.Write([]byte("</body></html>"))
			return
		}
		ar.UserData = struct{ ID uint64 }{ID: 7}
		ar.Authorized = true
		server.FinishAuthorizeRequest(resp, r, ar)
	}
	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}
	if !resp.IsError {
		resp.Output["uid"] = 7 // Get user id for further use
	}
	osin.OutputJSON(resp, w, r)
}
