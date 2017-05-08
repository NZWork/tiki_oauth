package server

import (
	"fmt"
	"net/http"
)

const OAuthTemplate = `
<!DOCTYPE html>
<html>

<style type="text/css">
    body {
	width: 100%;
	height: 100%;
	font-family: Alright Sans LP, Avenir Next, Helvetica Neue, Helvetica, Arial, PingFang SC, Source Han Sans SC, Hiragino Sans GB, Microsoft YaHei, WenQuanYi MicroHei, sans-serif;
	background-color: #f6f6f6;
	background: url(https://data.zi.com/5ebc0030c51be36dde1eb1e87a143dc9.svg) repeat 0 0px;
	-webkit-font-smoothing: antialiased
    }

    .box {
	overflow: auto;
	margin: auto;
	position: absolute;
	top: 0;
	left: 0;
	bottom: 0;
	right: 0;
	padding: 0 32px;
	width: 320px;
	height: 501px;
	background: #fcfcfc;
	z-index: 2;
	border-radius: 4px;
	box-shadow: 0 -20px 40px rgba(32, 46, 71, .05), 0 40px 40px rgba(32, 46, 71, .05);
	text-align: center;
    }

    .logo {
	padding-top: 32px;
	background: url('https://tiki.im/assets/images/logo.png') no-repeat;
	background-size: 100px;
	background-position: center center;
	width: 100%;
	height: 110px
    }

    .slogen {
	margin-top: 90px;
	bottom: 0
    }

    .inputs {
	margin-top: 64px
    }

    .form-control {
	width: 70%;
	height: 32px;
	margin-bottom: 16px;
	padding: 7px 17px;
	font-size: 14px;
	line-height: 1.5;
	color: #404040;
	background-color: #fff;
	background-image: none;
	border: 2px solid #f7f9fa;
	border-radius: 3px;
	transition: border-color ease-in-out 0.15s, box-shadow ease-in-out 0.15s;
	-webkit-transition: border-color ease-in-out 0.15s, box-shadow ease-in-out 0.15s;
	-moz-transition: border-color ease-in-out 0.15s, box-shadow ease-in-out 0.15s;
	-ms-transition: border-color ease-in-out 0.15s, box-shadow ease-in-out 0.15s;
	-o-transition: border-color ease-in-out 0.15s, box-shadow ease-in-out 0.15s;
	-webkit-appearance: none;
    }

    .form-control:focus {
	outline: none;
	border-color: #47525d
    }

    .btn {
	width: 50%;
	letter-spacing: 2px;
	position: relative;
	display: inline-block;
	vertical-align: middle;
	overflow: hidden;
	cursor: pointer;
	-webkit-transition: all .15s ease-in-out;
	transition: all .15s ease-in-out;
	box-shadow: inset 0 1px 0 hsla(0, 0%, 100%, .05), 0 1px 2px rgba(0, 0, 0, .05);
	border-radius: 100px;
	color: #47525d;
	background-color: transparent;
	border: 1px solid #47525d;
	padding: 12px 32px 12px 34px;
	font-size: 14px;
	line-height: 14px;
	text-shadow: none;
    }

    .btn:hover {
	color: #fff;
	background-color: #47525d;
	border: 1px solid #47525d
    }

    .login {
	outline: none;
	margin-top: 32px
    }
</style>

<body>
    <div class="box" id="login-form">
	<div class="logo">
	    <p class="slogen"></p>
	</div>
	<div class="inputs">
	    <input class="form-control" type="text" name="login" id="email" placeholder="hello@tiki.im" />
	    <input class="form-control" type="password" name="password" id="password" placeholder="密码" />
	    <div class="form-group">
		<button class="btn login" type="button" id="login">登录</button>
	    </div>
	</div>
    </div>


`

const JS = `
    <script src="//cdn.bootcss.com/jquery/3.1.1/jquery.min.js"></script>
    <script type="text/javascript">
	$('#login').on('click', function() {
	    $.ajax({
		type: 'POST',
		url: '%s/authorize?%s',
		data: {
			login: $('#email').val(),
			password: $('#password').val(),
		},
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
</body>

</html>
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
		template := fmt.Sprintf(JS, cfg.Domain, r.URL.RawQuery)
		w.Write([]byte(OAuthTemplate + template))
		return
	}
}
