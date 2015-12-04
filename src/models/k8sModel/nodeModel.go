package k8sModel

import (
	bm "models/baseModel"
	"encoding/json"
	"libs/request"
	"models/k8sModel/format"
)

type NodeModel struct {
	bm.BaseModel
	req request.Request
}


//curl -H "Content-Type: application/json" -X GET  http://192.168.12.8:8080/api/v1/namespaces/default/nodes
func (this *NodeModel) GetNodeList() (format.NodeList) {
	url := this.GetK8SUrl("nodes", "")
	body := this.req.Get(url)

    var data format.NodeList
	json.Unmarshal(body, &data)
	return data
}
func (this *NodeModel) NodeByName(name string) (format.Node) {
	url := this.GetK8SUrl("nodes/"+name, "")
	body := this.req.Get(url)

    var data format.Node
	json.Unmarshal(body, &data)
	return data
}
func (this *NodeModel) DeleteNode(name string) (format.Node) {
	url := this.GetK8SUrl("nodes/"+name, "")
	body := this.req.Delete(url, nil)
	
	var _node_data format.Node
	json.Unmarshal(body, &_node_data)
	
	return _node_data
}






