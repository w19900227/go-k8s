package format

// https://github.com/gosharplite/herd/blob/master/api_spec
type Data struct {
	Services []string `json:"services"`
	Clusters []string `json:"clusters"`
	Containers []string `json:"containers"`
}

type HerdList struct {
	Services []service `json:"services"`
	Clusters []cluster `json:"clusters"`
	Containers []container `json:"containers"`
}

type service struct {
	Service_name string `json:"service_name"` 
	Clusters []service_cluster `json:"clusters"`
}

type service_cluster struct {
	Cluster_name string `json:"cluster_name"`
	Containers []cluster_container `json:"containers"`
}

type cluster_container struct {
	Container_name string `json:"container_name"`
	Cpu int `json:"cpu"`
	Mem int `json:"mem"`
}

type cluster struct {
	Enable_auto_scale int `json:"enable_auto_scale"`
	Cpu_min int `json:"cpu_min"`
	Cpu_max int `json:"cpu_max"`
	Pod_min int `json:"pod_min"`
	Pod_max int `json:"pod_max"`
}

type container struct {}

