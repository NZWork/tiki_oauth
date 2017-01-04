package server

import (
	"github.com/RangelReale/osin"
	"net/http"
)

func addOAuthClient(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	client := &osin.DefaultClient{
		Id:          r.PostForm.Get("id"),
		Secret:      r.PostForm.Get("secret"),
		RedirectUri: r.PostForm.Get("redirect_uri"),
	}

	err := storage.CreateClient(client)

	resp := server.NewResponse()
	resp.Output["stat"] = 1
	if err != nil {
		resp.Output["stat"] = 0
		resp.Output["error"] = err.Error()
	}
	osin.OutputJSON(resp, w, r)
}

func updateOAuthClient(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	client, err := storage.GetClient(r.PostForm.Get("id"))

	resp := server.NewResponse()
	resp.Output["stat"] = 1

	if err != nil {
		resp.Output["stat"] = 0
		resp.Output["error"] = err.Error()
		osin.OutputJSON(resp, w, r)
		return
	}

	updatedClient := osin.DefaultClient{}
	updatedClient.CopyFrom(client)
	if r.PostForm.Get("secret") != "" {
		updatedClient.Secret = r.PostForm.Get("secret")
	}

	if r.PostForm.Get("redirect_uri") != "" {
		updatedClient.RedirectUri = r.PostForm.Get("redirect_uri")
	}

	err = storage.UpdateClient(&updatedClient)
	if err != nil {
		resp.Output["stat"] = 0
		resp.Output["error"] = err.Error()
	}
	osin.OutputJSON(resp, w, r)
}

func deleteOAuthClient() {}
