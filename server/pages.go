package server

import (
	"fmt"
	"net/http"
)

const OAUTH_TEMPLATE = `
<html><body>
<form id="loginForm">Login: <input type="text" name="login" /><br/>Password: <input type="password" name="password" /><br/><input type="button" id="login" value="登陆"/></form>

<script src="//cdn.bootcss.com/jquery/3.1.1/jquery.min.js"></script>
<script type="text/javascript">
$('#login').on('click', function(){
$.ajax({
    type: 'POST',
    url: '%s/authorize?%s',
    data: $('#loginForm').serialize(),
    success: function(data) {
	console.log(data);
	if (data.error != undefined) {
	    alert(data.error_description);
	} else if (data.redirect_uri != undefined) {
	    top.location.href = data.redirect_uri;
	}
    },
});
});
</script>
</body></html>
`

func OAuthIframe(w http.ResponseWriter, r *http.Request) {

	// authorize?response_type=code&client_id=tiki&redirect_uri=http%3A%2F%2Flocalhost%3A14000%2Fappauth%2Fcode
	// ------------------
	// response_type
	// client_id
	// redirect_uri

	// GET 方式
	if r.Method != "POST" {
		/*		w.Write([]byte("<html><body>"))
				w.Write([]byte(fmt.Sprintf("<form action=\"%s/authorize?%s\" method=\"POST\">", cfg.Domain, r.URL.RawQuery)))
				w.Write([]byte("Login: <input type=\"text\" name=\"login\" /><br/>"))
				w.Write([]byte("Password: <input type=\"password\" name=\"password\" /><br/>"))
				w.Write([]byte("<input type=\"submit\"/>"))
				w.Write([]byte("</form>"))
				w.Write([]byte("</body></html>"))
		*/
		w.Write([]byte(fmt.Sprintf(OAUTH_TEMPLATE, cfg.Domain, r.URL.RawQuery)))
		return
	}
}
