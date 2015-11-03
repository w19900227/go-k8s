package services

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )
import (
	bs "./baseService"
)

type AutoscaleSetService struct {
	bs.BaseService
}

// func (this *BaseService) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

func (this *AutoscaleSetService) GetTest() string {
	return this.GetBaseUrl()
}

func (this *AutoscaleSetService) GetService() string {
	return this.GetBaseUrl()
}

func (this *AutoscaleSetService) GetCluster() string {
	return this.GetBaseUrl()
}

func (this *AutoscaleSetService) GetContainer() string {
	return this.GetBaseUrl()
}

func (this *AutoscaleSetService) GetMachine() string {
	return this.GetBaseUrl()
}
