package filter

import (
	"encoding/json"
	"fmt"
)
import (
    "github.com/emicklei/go-restful"
	bm "models/baseModel"
	"libs/request"
    "log"
    "libs/auth"
)

type Filter struct {
	req request.Request
}

type response_format struct {
	Status string `"json:status"`
	Errno string `"json:errno"`
	Errmsg string `"json:errmsg"`
	Data interface{} `"json:data"`
}

func (this *Filter) AuthToken(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	token := req.Request.Header["Token"]
	auth_resp := auth.CheckToken(token[0])

    // log.Printf("[global-`AuthToken`-filter (logger)] %s,%s, data > %s\n", req.Request.Method, req.Request.URL, response_format)
    log.Printf("[global-`AuthToken`-filter (logger)] %s,%s\n", req.Request.Method, req.Request.URL)

	if auth_resp.Status != "ok" {
		resp.WriteEntity(auth_resp)
		return
	}
	set_namespace(token[0])
    chain.ProcessFilter(req, resp)
}

func set_namespace(auth_token string) {
	var user auth.User_format
	user_data, _ := auth.GetUser(auth_token)
	json.Unmarshal(user_data, &user)
	
	if user.Namespace != "" {
		bm.Namespace = user.Namespace
	}
	// bm.Namespace = "default"
}

func Test(result interface{}) {
	js, _ := json.Marshal(result)
	fmt.Println(string(js))
}
