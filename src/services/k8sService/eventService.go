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
	// k8s_format "models/k8sModel/format"
	// "strings"
	// "strconv"

	// herd_format "models/herdModel/format"
)

type EventService struct {
	bs.BaseService
	// k8s.GetK8s
	// bm.BaseModel
}

type data struct {
	Object_type string `json:"object_type,omitempty"`
	Lbject_name string `json:"object_name,omitempty"`
	First_time string `json:"first_time,omitempty"`
	Last_time string `json:"last_time,omitempty"`
	Frequency int `json:"frequency,omitempty"`
	Reason string `json:"reason,omitempty"`
	Detail string `json:"detail,omitempty"`
	Host_ip string `json:"host_ip,omitempty"`
}


type Events struct {
	Status string `json:"status,omitempty"`
	Errno string `json:"errno,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
	Data []data `json:"data,omitempty"`
}

func (this *EventService) GetTest() string {
	fmt.Println("EventService")
	return "EventService"
}

func (this *EventService) GetEventList() Events {
	var _event_list Events
	_eventModel := k8sModel.EventModel{}
	_event_data := _eventModel.GetEventList()

	for _, event := range _event_data.Items {
		var _data data
		_data.Object_type = event.InvolvedObject.Kind
		_data.Lbject_name = event.Metadata.Name
		_data.First_time = event.FirstTimestamp
		_data.Last_time = event.LastTimestamp
		_data.Frequency = event.Count
		_data.Reason = event.Reason
		_data.Detail = event.Message
		if event.Source.Host != "" {
			_data.Host_ip = event.Source.Host
		} else {
			_data.Host_ip = "N/A"
		}
		_event_list.Data = append(_event_list.Data, _data)
	}

	_event_list.Status = "ok"
	return _event_list
}

