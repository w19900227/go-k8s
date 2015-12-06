package k8sService

import (
)

import (
	// "fmt"
	// bs "./baseService"
	// "../models"
	// k8s "./k8s"
)

import (
	bs "services/baseService"
	"models/k8sModel"
	// "models/herdModel"
	// herd_format "models/herdModel/format"
	k8s_format "models/k8sModel/format"
	// "strings"
	// "strconv"

	// herd_format "models/herdModel/format"
)

type NodeService struct {
	bs.BaseService
	// k8s.GetK8s
	// bm.BaseModel
}


type Nodes struct {
	Status string `json:"status,inline"`
	Errno string `json:"errno,inline"`
	Errmsg string `json:"errmsg,inline"`
	Data []k8s_format.Node `json:"data,inline"`
}


func (this *NodeService) GetPodList() Nodes {
	var _node_list Nodes
	_nodeModel := k8sModel.NodeModel{}
	_node_data := _nodeModel.GetNodeList()

	_node_list.Status = "ok"
	_node_list.Data = _node_data.Items
	return _node_list
}

