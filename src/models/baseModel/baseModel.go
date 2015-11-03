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

func (this *BaseModel) GetBaseUrl(uri string) string {
	return "http://192.168.15.77:8080/api/v1" + "/namespaces/default/" + uri
}

func (this *BaseModel) GetHredIP(uri string) string {
	return "http://192.168.15.76:8090/" + uri
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
    // return body

    // body = []byte(`{"status": "ok","data": [{"official": [{"ch_id": "57","ch_name": "UDN直播新聞台","ch_description": "","owner_id": "14","owner_name": "Ayukawa","owner_photo": "38/bc/14.jpg?1415009271","ch_number": "10","ch_type": "2","ch_rating_system": "0","ch_views": "10360","ch_likes": "0","ch_subscribes": "211","ch_has_cover": "1","ch_cover_hash": "1415107242"},{"ch_id": "1698","ch_name": "大愛一臺HD Live 直播","ch_description": "","owner_id": "618","owner_name": "Kelly","owner_photo": "c3/81/618.jpg?1423460588","ch_number": "11","ch_type": "2","ch_rating_system": "0","ch_views": "5639","ch_likes": "0","ch_subscribes": "72","ch_has_cover": "1","ch_cover_hash": "1431919773"}],"create": [ ],"subscribe": [ ]}]}`)
    // fmt.Println(body)
    // var data list_self
	// json.Unmarshal(body, &data)
}


func (this *BaseModel) Test(result interface{}) {
	js, _ := json.Marshal(result)
	fmt.Println(string(js))
}
