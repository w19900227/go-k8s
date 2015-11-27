package k8sService

import (
)

import (
	"fmt"
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

type EventService struct {
	bs.BaseService
	// k8s.GetK8s
	// bm.BaseModel
}


type Events struct {
	Status string `json:"status,omitempty"`
	Errno string `json:"errno,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
	Data []k8s_format.Event `json:"data,omitempty"`
}

func (this *EventService) GetTest() string {
	fmt.Println("EventService")
	return "EventService"
}

func (this *EventService) GetEventList() Events {
	var _event_list Events
	_eventModel := k8sModel.EventModel{}
	_event_data := _eventModel.GetEventList()

	_event_list.Status = "ok"
	_event_list.Data = _event_data.Items
	return _event_list
}

