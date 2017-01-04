package server

import (
	"fmt"
	"net/http"
)

func OAuthIframe(w http.ResponseWriter, r *http.Request) {

	// authorize?response_type=code&client_id=tiki&redirect_uri=http%3A%2F%2Flocalhost%3A14000%2Fappauth%2Fcode
	// ------------------
	// response_type
	// client_id
	// redirect_uri

	// GET 方式
	if r.Method != "POST" {
		w.Write([]byte("<html><body>"))
		w.Write([]byte(fmt.Sprintf("<form action=\"/authorize?%s\" method=\"POST\">", r.URL.RawQuery)))
		w.Write([]byte("Login: <input type=\"text\" name=\"login\" /><br/>"))
		w.Write([]byte("Password: <input type=\"password\" name=\"password\" /><br/>"))
		w.Write([]byte("<input type=\"submit\"/>"))
		w.Write([]byte("</form>"))
		w.Write([]byte("</body></html>"))
		return
	}
}
