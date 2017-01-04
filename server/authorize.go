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

	if r.Method == "POST" {
		r.ParseForm()
		if r.Form.Get("response_type") == "" || r.Form.Get("client_id") == "" || r.Form.Get("redirect_uri") == "" {
			fmt.Printf("%v", r.PostForm.Encode())
			resp.Output["code"] = 400
			resp.Output["msg"] = "invalid params given"
			osin.OutputJSON(resp, w, r)
			return
		}
	} else {
		http.Redirect(w, r, "/auth?"+r.URL.RawQuery, 301)
		return
	}

	api := &APIResponse{}

	if ar := server.HandleAuthorizeRequest(resp, r); ar != nil {

		// 校验用户数据
		resposne, _ := PostAPI(cfg.API+"login", map[string]interface{}{
			"email":  r.PostForm.Get("login"),
			"passwd": r.PostForm.Get("password"),
		})

		api = DecodeAPIResponse(resposne)

		ar.Authorized = false // 默认

		if api.Code == 1200 {
			ar.UserData = api.Result["id"].(float64)

			resp.Output["uid"] = api.Result["id"].(float64)
			ar.Authorized = true
		}
		server.FinishAuthorizeRequest(resp, r, ar)
	}

	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}
	if api.Code != 1200 {
		resp.Output["api_msg"] = api.Msg
		resp.Output["api_code"] = api.Code
	}

	osin.OutputJSON(resp, w, r)
}
