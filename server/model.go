package server

import (
	"encoding/json"
)

type APIResponse struct {
	Code   int                    `json:"code"`
	Result map[string]interface{} `json:"result"`
	Msg    string                 `json:"msg"`
}

func DecodeAPIResponse(response []byte) *APIResponse {
	ar := APIResponse{}
	json.Unmarshal(response, &ar)
	return &ar
}
