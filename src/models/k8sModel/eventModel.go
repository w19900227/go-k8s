package k8sModel

import (
	bm "models/baseModel"
	"encoding/json"
	"libs/request"
	"models/k8sModel/format"
)

type EventModel struct {
	bm.BaseModel
	req request.Request
}

// func (this *EventModel) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

// func (this *EventModel) GetContainer() (format.Pod) {}

//curl -H "Content-Type: application/json" -X GET  http://192.168.12.8:8080/api/v1/namespaces/default/nodes
func (this *EventModel) GetEventList() (format.EventList) {
	body := this.req.Get("http://192.168.12.8:8080/api/v1/events")
	// body := this.getPagInfo_data()//在伺服器run必須移除此行

    var data format.EventList
	json.Unmarshal(body, &data)
	// fmt.Println(err)
	// fmt.Println(b)
	return data
}





