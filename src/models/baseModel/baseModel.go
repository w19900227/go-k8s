package baseModel

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

import (
	"fmt"
	// "encoding/json"
)

const (
	k8s_path = "http://192.168.12.8:8080"
	// k8s_path = "http://master:8080"
	herd_path = "http://192.168.12.7:8090"
	// herd_path = "http://herd:8090"
)

var (
	Namespace = "default"
)	

type BaseModel struct {
	baseurl string 
	namespace string 
	hred_ip string 
}

// func (this *BaseService) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

func (this *BaseModel) GetK8SUrl(uri string, namespace string) string {
	if namespace != "" {
		return k8s_path + "/api/v1" + "/namespaces/" + namespace + "/" + uri
	}
	return k8s_path + "/api/v1/" + uri
}

func (this *BaseModel) GetHerdUrl(uri string) string {
	return herd_path + "/" + uri
}

func (this *BaseModel) GetRequest(url string) []byte {
    res, err := http.Get(url)
    if err != nil {
        panic(err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        panic(err.Error())
    }

    return body
}


func (this *BaseModel) Test(result interface{}) {
	js, _ := json.Marshal(result)
	fmt.Println(string(js))
}
