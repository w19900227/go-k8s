package services

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )
import (
	bs "./baseService"
)

type ServiceRelationService struct {
	bs.BaseService
}

// func (this *BaseService) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

func (this *ServiceRelationService) GetTest() string {
	return this.GetBaseUrl()
}

func (this *ServiceRelationService) GetService() string {
	return this.GetBaseUrl()
}

func (this *ServiceRelationService) GetCluster() string {
	return this.GetBaseUrl()
}

func (this *ServiceRelationService) GetContainer() string {
	return this.GetBaseUrl()
}

func (this *ServiceRelationService) GetMachine() string {
	return this.GetBaseUrl()
}
