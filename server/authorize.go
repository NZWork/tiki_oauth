package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/RangelReale/osin"
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
			resp.Output["error"] = "invalid_params"
			resp.Output["error_description"] = "some of required params missing"
			osin.OutputJSON(resp, w, r)
			return
		}
	} else {
		http.Redirect(w, r, "/auth?"+r.URL.RawQuery, 301)
		return
	}

	api := &APIResponse{}

	fmt.Printf("[%v]AC:%d Client[%v] User[%v]\n", time.Now(), redisPool.ActiveCount(), r.Form.Get("client_id"), r.PostForm.Get("login"))

	if ar := server.HandleAuthorizeRequest(resp, r); ar != nil {

		// 校验用户数据
		resposne, _ := PostAPI(cfg.API+"login", map[string]interface{}{
			"email":  r.PostForm.Get("login"),
			"passwd": r.PostForm.Get("password"),
			"xauth":  cfg.Secret,
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

	if url, err := resp.GetRedirectUrl(); err == nil { // 重写回调
		resp.Type = osin.DATA // 取消跳转
		resp.Output["redirect_uri"] = url
	}

	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
		goto RESPONSE
	}

	if resp.InternalError == nil && api.Code != 1200 { // 内部 API 返回错误
		resp.Output["error"] = api.Code
		resp.Output["error_description"] = api.Msg
		delete(resp.Output, "redirect_uri")
	}

RESPONSE:
	osin.OutputJSON(resp, w, r)
}
