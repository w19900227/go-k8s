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


//curl -H "Content-Type: application/json" -X GET  http://192.168.12.8:8080/api/v1/namespaces/default/nodes
func (this *EventModel) GetEventList() (format.EventList) {
	url := this.GetK8SUrl("events", "")
	body := this.req.Get(url)

    var data format.EventList
	json.Unmarshal(body, &data)
	return data
}





