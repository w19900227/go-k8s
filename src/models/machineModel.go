package models

import (
	bm "models/baseModel"
	"encoding/json"
)
import (
	"fmt"
)

type MachineModel struct {
	bm.BaseModel
}

// func (this *MachineModel) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

// func (this *MachineModel) GetBaseUrl() string {
// 	return "http://192.168.6.13:8080/api/v1" + "/namespaces/default/"
// }



// type promotion_block struct {
// 	Status string `json:"status"`
// 	Data []block_array `json:"data"`
// }

// type block_array struct {
// 	Id string `json:"id"` 
// 	Country_code string `json:"country_code"` 
// 	Title string `json:"title"` 
// 	Duration_from string `json:"duration_from"` 
// 	Duration_to string `json:"duration_to"` 
// 	Types string `json:"type"` 
// 	Channel_id string `json:"channel_id"` 
// 	Url string `json:"url"` 
// 	Image_hash string `json:"image_hash"` 
// 	Position string `json:"position"` 
// 	Status string `json:"status"` 
// 	Created_at string `json:"created_at"` 
// 	Updated_at string `json:"updated_at"` 
// 	Package_id string `json:"package_id"` 
// }

type list_self struct {
	Status string `json:"status"`
	Data []list_data `json:"data"`
}

type list_data struct {
	Official []official_array `json:"official"`
}

type official_array struct {
	Ch_id string `json:"ch_id"`
	Ch_name string `json:"ch_name"`
	Ch_description string `json:"ch_description"`
	Owner_id string `json:"owner_id"`
	Owner_name string `json:"owner_name"`
	Owner_photo string `json:"owner_photo"`
	Ch_number string `json:"ch_number"`
	Ch_type string `json:"ch_type"`
	Ch_rating_system string `json:"ch_rating_system"`
	Ch_views string `json:"ch_views"`
	Ch_likes string `json:"ch_likes"`
	Ch_subscribes string `json:"ch_subscribes"`
	Ch_has_cover string `json:"ch_has_cover"`
	Ch_cover_hash string `json:"ch_cover_hash"`
}

func (this *MachineModel) GetPagInfo() string {
	url := "http://www.miii.tv/channels/list/self"
	body := this.GetRequest(url)

    var data list_self
	json.Unmarshal(body, &data)
	fmt.Println(data)
	fmt.Println("==============================")
	fmt.Println(data.Data)
	// fmt.Println(data.Data[0].Id)
	// fmt.Println(url)
	return url

}