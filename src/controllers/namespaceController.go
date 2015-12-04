package controllers

import (
    "github.com/emicklei/go-restful"
    bm "models/baseModel"
    // "services"
    // "services/baseService"
)
import (
    "fmt"
)

type NamespaceController struct {
    // bm.BaseModel
}

func (this *NamespaceController) Get(request *restful.Request, response *restful.Response) {
	fmt.Println("[Get] Namespace")
	
    response.WriteEntity(bm.Namespace)
    return
}


