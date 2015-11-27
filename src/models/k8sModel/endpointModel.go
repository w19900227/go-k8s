package k8sModel

import (
	"fmt"

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
	body := this.req.Post("http://192.168.12.8:8080/api/v1/namespaces/default/endpoints", data)
	
	var _endpoint_data format.Endpoints
	json.Unmarshal(body, &_endpoint_data)

	return _endpoint_data
}

func (this *EndpointModel) Test22() {
	fmt.Println("ssss")
}





