package services

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )
import (
	bs "./baseService"
)

type EditService struct {
	bs.BaseService
}

// func (this *BaseService) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

func (this *EditService) GetTest() string {
	return this.GetBaseUrl()
}

func (this *EditService) GetService() string {
	return this.GetBaseUrl()
}

func (this *EditService) GetCluster() string {
	return this.GetBaseUrl()
}

func (this *EditService) GetContainer() string {
	return this.GetBaseUrl()
}

func (this *EditService) GetMachine() string {
	return this.GetBaseUrl()
}
