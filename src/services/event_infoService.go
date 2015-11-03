package services

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )
import (
	bs "./baseService"
)

type EventInfoService struct {
	bs.BaseService
}

// func (this *BaseService) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

func (this *EventInfoService) GetTest() string {
	return this.GetBaseUrl()
}

func (this *EventInfoService) GetService() string {
	return this.GetBaseUrl()
}

func (this *EventInfoService) GetCluster() string {
	return this.GetBaseUrl()
}

func (this *EventInfoService) GetContainer() string {
	return this.GetBaseUrl()
}

func (this *EventInfoService) GetMachine() string {
	return this.GetBaseUrl()
}
