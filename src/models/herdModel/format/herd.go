package format

// https://github.com/gosharplite/herd/blob/master/api_spec
type Data struct {
	Services []string `json:"services,omitempty"`
	Clusters []string `json:"clusters,omitempty"`
	Containers []string `json:"containers,omitempty"`
}

type HerdGet struct {
	Services []service `json:"services,omitempty"`
	Clusters []cluster `json:"clusters,omitempty"`
	Containers []container `json:"containers,omitempty"`
}

type service struct {
	Service_name string `json:"service_name,omitempty"` 
	Clusters []service_cluster `json:"clusters,omitempty"`
}

type service_cluster struct {
	Cluster_name string `json:"cluster_name,omitempty"`
	Containers []cluster_container `json:"containers,omitempty"`
}

type cluster_container struct {
	Container_name string `json:"container_name"`
	Cpu int `json:"cpu"`
	Mem int `json:"mem"`
}

type cluster struct {
	Containers []container `json:"containers"`
	Enable_auto_scale int `json:"enable_auto_scale"`
	Cpu_min int `json:"cpu_min"`
	Cpu_max int `json:"cpu_max"`
	Pod_min int `json:"pod_min"`
	Pod_max int `json:"pod_max"`
}

type container struct {
	Container_name string `json:"container_name"`
	Cpu int `json:"cpu"`
	Mem int `json:"mem"`
}

//---------------

type HerdGetScale struct {
	Clusters []scale_Cluster `json:"clusters,omitempty"`
}

type scale_Cluster struct {
	Cluster_id string `json:"cluster_id"`
	Enable_auto_scale int `json:"enable_auto_scale"`
	Cpu_min int `json:"cpu_min"`
	Cpu_max int `json:"cpu_max"`
	Pod_min int `json:"pod_min"`
	Pod_max int `json:"pod_max"`
}