package services

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )
import (
	bs "./baseService"
)

type DeleteService struct {
	bs.BaseService
}

// func (this *BaseService) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

func (this *DeleteService) GetTest() string {
	return this.GetBaseUrl()
}

func (this *DeleteService) GetService() string {
	return this.GetBaseUrl()
}

func (this *DeleteService) GetCluster() string {
	return this.GetBaseUrl()
}

func (this *DeleteService) GetContainer() string {
	return this.GetBaseUrl()
}

func (this *DeleteService) GetMachine() string {
	return this.GetBaseUrl()
}
