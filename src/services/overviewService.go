package services

// // import (
// // 	"fmt"
// // 	"runtime"
// // 	"time"
// // )
import (
	bs "services/baseService"
)

type OverviewService struct {
	bs.BaseService
}

// // func (this *BaseService) init() {

// 	// this.baseurl = "http://192.168.6.13:8080/api/v1"
// 	// this.namespace = "/namespaces/default/"
// 	// this.hred_ip = "http://192.168.6.14:8090"
// // }

func (this *OverviewService) GetTest() string {
	// return "sssss"
	return this.GetBaseUrl()
}

// func (this *OverviewService) GetService() string {
// 	return this.GetBaseUrl()
// }

// func (this *OverviewService) GetCluster() string {
// 	return this.GetBaseUrl()
// }

// func (this *OverviewService) GetContainer() string {
// 	return this.GetBaseUrl()
// }

// func (this *OverviewService) GetMachine() string {
// 	return this.GetBaseUrl()
// }
