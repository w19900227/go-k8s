package k8sModel

import (
	bm "models/baseModel"
	"encoding/json"
	"libs/request"
	"models/k8sModel/format"
)

type EndpointModel struct {
	bm.BaseModel
	req request.Request
}

func (this *EndpointModel) CreateEndpoint(data format.Endpoints) (format.Endpoints) {
	url := this.GetK8SUrl("endpoints", bm.Namespace)
	body := this.req.Post(url, data)
	
	var _endpoint_data format.Endpoints
	json.Unmarshal(body, &_endpoint_data)

	return _endpoint_data
}





