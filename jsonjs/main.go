package main

import (
	"encoding/json"
	"fmt"
)

type HttpCode struct {
	Code   int    `json:"Code"`
	Decrip string `json:"Decrip"`
}

func main() {

	var httpCodes []HttpCode

	rcvd := `[{"Code":200,"Decrip":"StatusOK"},{"Code":301,"Decrip":"StatusMovedPermanently"},
{"Code":302,"Decrip":"StatusFound"},{"Code":303,"Decrip":"StatusSeeOther"},
{"Code":307,"Decrip":"StatusTemporaryRedirect"},{"Code":400,"Decrip":"StatusBadRequest"},
{"Code":401,"Decrip":"StatusUnauthorized"},{"Code":402,"Decrip":"StatusPaymentRequired"},
{"Code":403,"Decrip":"StatusForbidden"},{"Code":404,"Decrip":"StatusNotFound"},
{"Code":405,"Decrip":"StatusMethodNotAllowed"},{"Code":418,"Decrip":"StatusTeapot"},
{"Code":500,"Decrip":"StatusInternalServerError"}]`

	json.Unmarshal([]byte(rcvd), &httpCodes)

	for _, v := range httpCodes {
		fmt.Printf("%s code: %d \n", v.Decrip, v.Code)
	}

}
