package herdModel

import (
	bm "models/baseModel"
	// "io/ioutil"
	"encoding/json"
	"libs/request"
	"models/herdModel/format"
)
import (
	// "fmt"
	// "reflect"
)
//https://github.com/parnurzeal/gorequest/blob/master/main.go
type GetScaleHerdModel struct {
	bm.BaseModel
	Data format.Data
	req request.Request
}


func (this *GetScaleHerdModel) Test() string {
	return "k8s test"
}
// curl -H "Content-Type: application/json" -X POST -d '{"clusters":["app","dev","redis-master","redis-slave","web-app"]}' http://192.168.12.7:8090/getscale
func (this *GetScaleHerdModel) GetPagInfo(d format.Data) (format.HerdGetscale) {
	// url := this.GetBaseUrl("services")

	body := this.req.Post("http://192.168.12.7:8090/getscale", d)
	// body := this.GetRequest(url)
	// body := this.getPagInfo_data()//在伺服器run必須移除此行

    var data format.HerdGetscale
	json.Unmarshal(body, &data)
	return data
}
 
func (this *GetScaleHerdModel) PostData() (format.HerdGetscale) {
	// fmt.Println(this.data)
	// body := this.JsonRequest("get", this.data)
	// fmt.Println(string(body))
	// body := this.req.Post("http://192.168.15.76:8090/get", this.data)
	body := this.getPagInfo_data()
	// fmt.Println(string(data))


    // body, err := ioutil.ReadAll(bodys)

    // if err != nil {
        // panic(err.Error())
    // }
    
    
    var data format.HerdGetscale
	json.Unmarshal(body, &data)
	return data
	// fmt.Println(data)
	// fmt.Println(string(data))
	// zv := reflect.TypeOf(this.data).Kind()
	// fmt.Println(zv)
	// url := "http://localhost:9090/ex2"
	// types := "application/json"
	// buf, _ := json.Marshal(this.data)
	// body := bytes.NewBuffer(buf)
	// r, _ := http.Post(url, types, body)
	// response, _ := ioutil.ReadAll(r.Body)
	// fmt.Println(string(response))
}

func (this *GetScaleHerdModel) GetData(service interface{}) () {
	
}


//以下可以除

// var url string = "http://www.miii.tv/channels/promotion/block"

func (this *GetScaleHerdModel) getPagInfo_data() []byte {
	body := []byte(`{"services":[{"service_name":"frontend","clusters":[{"cluster_name":"frontend","containers":[{"container_name":"frontend-24tyt","cpu":0,"mem":1},{"container_name":"frontend-8zzuh","cpu":0,"mem":1},{"container_name":"frontend-x7zy1","cpu":0,"mem":1}]}]},{"service_name":"golang","clusters":[{"cluster_name":"golang","containers":[{"container_name":"golang-f8cw1","cpu":0,"mem":0},{"container_name":"golang-m9ng5","cpu":0,"mem":0},{"container_name":"golang-zw3ic","cpu":0,"mem":0}]}]},{"service_name":"golang-dh","clusters":[{"cluster_name":"golang-dh","containers":[{"container_name":"golang-dh-8fwnu","cpu":0,"mem":0}]}]},{"service_name":"redis-master","clusters":[{"cluster_name":"redis-master","containers":[{"container_name":"redis-master-hbaz7","cpu":0,"mem":0}]}]},{"service_name":"redis-slave","clusters":[{"cluster_name":"redis-slave","containers":[{"container_name":"redis-slave-2goiq","cpu":0,"mem":0},{"container_name":"redis-slave-jj5as","cpu":0,"mem":0}]}]}],"clusters":[],"containers":[]}`)
	return body	
}


