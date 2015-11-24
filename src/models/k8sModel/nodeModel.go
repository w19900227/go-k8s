package k8sModel

import (
	"fmt"

	bm "models/baseModel"
	"encoding/json"
	"libs/request"
	"models/k8sModel/format"
)

type NodeModel struct {
	bm.BaseModel
	req request.Request
}

// func (this *NodeModel) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

// func (this *NodeModel) GetContainer() (format.Pod) {}

//curl -H "Content-Type: application/json" -X GET  http://192.168.12.8:8080/api/v1/namespaces/default/nodes
func (this *NodeModel) GetNodeList() (format.NodeList) {
	body := this.req.Get("http://192.168.12.8:8080/api/v1/nodes")
	// body := this.getPagInfo_data()//在伺服器run必須移除此行

    var data format.NodeList
	json.Unmarshal(body, &data)
	// fmt.Println(err)
	// fmt.Println(b)
	return data
}
func (this *NodeModel) NodeByName(name string) (format.Node) {
	body := this.req.Get("http://192.168.12.8:8080/api/v1/nodes/"+name)
	// body := this.getPagInfo_data()//在伺服器run必須移除此行
    var data format.Node
	json.Unmarshal(body, &data)
	// fmt.Println(err)
	// fmt.Println(b)
	return data
}
func (this *NodeModel) DeleteNode(name string) (format.Node) {
	body := this.req.Delete("http://192.168.12.8:8080/api/v1/nodes/"+name, nil)
	
	var _node_data format.Node
	json.Unmarshal(body, &_node_data)
	
	return _node_data
}

func (this *NodeModel) Test22() {
	fmt.Println("ssss")
}





