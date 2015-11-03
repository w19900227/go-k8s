package models

// import (
// 	bm "./baseModel"
// 	"encoding/json"
// 	"libs/request"
// )
// import (
// 	// "fmt"
// )

// type ServiceModel struct {
// 	bm.BaseModel
// 	req request.Request
// }

// // func (this *ServiceModel) init() {

// 	// this.baseurl = "http://192.168.6.13:8080/api/v1"
// 	// this.namespace = "/namespaces/default/"
// 	// this.hred_ip = "http://192.168.6.14:8090"
// // }

// type replicationControllerList struct {
// 	Kind string `json:"kind"`
// 	ApiVersion string `json:"apiVersion"`
// 	Metadata struct {
// 		SelfLink string `json:"selfLink"`
// 		ResourceVersion string `json:"	resourceVersion"`
// 	} `json:"metadata"`
// 	Items []items `json:"items"`
// }


// type items struct {
// 	Metadata items_metadata`json:"metadata"`
// 	Spec spec `json:"spec"`
// 	Status struct {
// 		Replicas int `json:"replicas"`
// 	} `json:"status"`
// }

// type items_metadata struct {
// 	Name string `json:"name"`
// 	Namespace string `json:"namespace"`
// 	SelfLink string `json:"selfLink"`
// 	Uid string `json:"uid"`
// 	ResourceVersion string `json:"resourceVersion"`
// 	CreationTimestamp string `json:"creationTimestamp"`
// 	Labels map[string]string `json:"labels"`
// }

// type spec struct {
// 	Ports []struct {
// 		Protocol string `json:"protocol"`
// 		Port int `json:"port"`
// 		TargetPort int `json:"targetPort"`
// 		NodePort int `json:"nodePort"`
// 	} `json:"ports"`
// 	// Selector `json:"selector"`
// 	ClusterIP string `json:"clusterIP"`
// 	Types string `json:"types"`
// 	SessionAffinity string `json:"sessionAffinity"`
// 	// Replicas int `json:"replicas"`
// 	Selector struct {
// 		Name string `json:"name"`
// 	} `json:"selector"`
// 	// Template template `json:"template"`
// }

// type template struct {
// 	Metadata struct {
// 		CreationTimestamp string `json:"creationTimestamp"`
// 		Labels struct {
// 			Name string `json:"name"`
// 		} `json:"labels"`
// 	} `json:"metadata"`
// 	Spec struct {
// 		Containers []containers `json:"containers"`
// 		RestartPolicy string `json:"restartPolicy"`
// 		DnsPolicy string `json:"dnsPolicy"`
// 	}
// }

// type containers struct {
// 	Name string `json:"name"`
// 	Image string `json:"image"`
// 	Ports []struct {
// 		ContainerPort int `json:"containerPort"`
// 		Protocol string `json:"protocol"`
// 	}
// 	TerminationMessagePath string `json:"terminationMessagePath"`
// 	ImagePullPolicy string `json:"imagePullPolicy"`
// }


// func (this *ServiceModel) GetPagInfo() (replicationControllerList) {
// 	// body := this.req.Get("http://192.168.15.77:8080/api/v1/namespaces/default/services")
// 	body := getPagInfo_data()//在伺服器run必須移除此行

//     var data replicationControllerList
// 	json.Unmarshal(body, &data)
// 	// fmt.Println(err)
// 	// fmt.Println(b)
// 	return data
// }

// //以下可以除

// var url string = "http://www.miii.tv/channels/promotion/block"

// func getPagInfo_data() []byte {
// 	// body := []byte(`{"kind": "ReplicationControllerList","apiVersion": "v1","metadata": {"selfLink": "/api/v1/namespaces/default/replicationcontrollers","resourceVersion": "46750615"},"items": [{"metadata": {"name": "frontend","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/frontend","uid": "ecfab3a0-3f3a-11e5-ae3a-fa163ea883d9","resourceVersion": "43538292","creationTimestamp": "2015-08-10T08:36:42Z","labels": {"name": "frontend"}},"spec": {"replicas": 3,"selector": {"name": "frontend"},"template": {"metadata": {"creationTimestamp": null,"labels": {"name": "frontend"}},"spec": {"containers": [{"name": "php-redis","image": "192.168.6.14:5000/kubernetes/example-guestbook-php-redis:v2","ports": [{"containerPort": 80,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 3}},{"metadata": {"name": "kube-dns-v3","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/kube-dns-v3","uid": "c5a3c27f-3f1a-11e5-ae3a-fa163ea883d9","resourceVersion": "43460409","creationTimestamp": "2015-08-10T04:46:32Z","labels": {"k8s-app": "kube-dns-v3","kubernetes.io/cluster-service": "true"}},"spec": {"replicas": 1,"selector": {"k8s-app": "kube-dns","version": "v3"},"template": {"metadata": {"creationTimestamp": null,"labels": {"k8s-app": "kube-dns","kubernetes.io/cluster-service": "true","version": "v3"}},"spec": {"volumes": [{"name": "dns-token","secret": {"secretName": "token-system-dns"}}],"containers": [{"name": "etcd","image": "192.168.6.14:5000/google_containers/etcd:2.0.9","command": ["/usr/local/bin/etcd","-listen-client-urls","http://127.0.0.1:2379,http://127.0.0.1:4001","-advertise-client-urls","http://127.0.0.1:2379,http://127.0.0.1:4001","-initial-cluster-token","skydns-etcd"],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"},{"name": "kube2sky","image": "192.168.6.14:5000/google_containers/kube2sky:1.9","args": ["-domain=kubernetes.local","-kubecfg_file=/etc/dns_token/kubeconfig","-v=2","-kube_master_url=https://${KUBERNETES_SERVICE_HOST}:${KUBERNETES_SERVICE_PORT}"],"resources": {},"volumeMounts": [{"name": "dns-token","readOnly": true,"mountPath": "/etc/dns_token"}],"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"},{"name": "skydns","image": "192.168.6.14:5000/google_containers/skydns:2015-03-11-001","args": ["-machines=http://localhost:4001","-addr=0.0.0.0:53","-domain=kubernetes.local.","-verbose=true"],"ports": [{"name": "dns","containerPort": 53,"protocol": "UDP"},{"name": "dns-tcp","containerPort": 53,"protocol": "TCP"}],"resources": {},"livenessProbe": {"exec": {"command": ["/bin/sh","-c","nslookup kubernetes.default.kubernetes.local localhost \u003e/dev/null"]},"initialDelaySeconds": 30,"timeoutSeconds": 5},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "Default"}}},"status": {"replicas": 1}},{"metadata": {"name": "michael-wordpress","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/michael-wordpress","uid": "3a263e54-40be-11e5-ae3a-fa163ea883d9","resourceVersion": "44453431","creationTimestamp": "2015-08-12T06:49:06Z","labels": {"Type": "ui","Version": "1.0","app": "web","name": "michael-wordpress"}},"spec": {"replicas": 3,"selector": {"name": "michael-wordpress"},"template": {"metadata": {"creationTimestamp": null,"labels": {"name": "michael-wordpress"}},"spec": {"containers": [{"name": "michael-wordpress","image": "cael/wordpress","ports": [{"containerPort": 80,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 3}},{"metadata": {"name": "monitoring-heapster-v3","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/monitoring-heapster-v3","uid": "abb736f1-3f1e-11e5-ae3a-fa163ea883d9","resourceVersion": "43469634","creationTimestamp": "2015-08-10T05:14:26Z","labels": {"k8s-app": "heapster","kubernetes.io/cluster-service": "true","version": "v3"}},"spec": {"replicas": 1,"selector": {"k8s-app": "heapster","version": "v3"},"template": {"metadata": {"creationTimestamp": null,"labels": {"k8s-app": "heapster","kubernetes.io/cluster-service": "true","version": "v3"}},"spec": {"volumes": [{"name": "ssl-certs","hostPath": {"path": "/etc/ssl/certs"}},{"name": "monitoring-token","secret": {"secretName": "token-system-monitoring"}}],"containers": [{"name": "heapster","image": "gcr.io/google_containers/heapster:v0.14.1","command": ["/heapster","--source=kubernetes:''","--sink=influxdb:http://monitoring-influxdb:8086"],"resources": {},"volumeMounts": [{"name": "ssl-certs","readOnly": true,"mountPath": "/etc/ssl/certs"},{"name": "monitoring-token","readOnly": true,"mountPath": "/etc/kubernetes/kubeconfig"}],"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 1}},{"metadata": {"name": "monitoring-influx-grafana-v1","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/monitoring-influx-grafana-v1","uid": "17e282ee-3f1e-11e5-ae3a-fa163ea883d9","resourceVersion": "43468243","creationTimestamp": "2015-08-10T05:10:18Z","labels": {"k8s-app": "influxGrafana","kubernetes.io/cluster-service": "true","version": "v1"}},"spec": {"replicas": 1,"selector": {"k8s-app": "influxGrafana","version": "v1"},"template": {"metadata": {"creationTimestamp": null,"labels": {"k8s-app": "influxGrafana","kubernetes.io/cluster-service": "true","version": "v1"}},"spec": {"containers": [{"name": "influxdb","image": "gcr.io/google_containers/heapster_influxdb:v0.3","ports": [{"hostPort": 8083,"containerPort": 8083,"protocol": "TCP"},{"hostPort": 8086,"containerPort": 8086,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"},{"name": "grafana","image": "gcr.io/google_containers/heapster_grafana:v0.7","env": [{"name": "INFLUXDB_HOST","value": "monitoring-influxdb"},{"name": "INFLUXDB_PORT","value": "8086"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 1}},{"metadata": {"name": "redis-slave","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/redis-slave","uid": "1678c3e9-3f1b-11e5-ae3a-fa163ea883d9","resourceVersion": "43461186","creationTimestamp": "2015-08-10T04:48:47Z","labels": {"name": "redis-slave"}},"spec": {"replicas": 2,"selector": {"name": "redis-slave"},"template": {"metadata": {"creationTimestamp": null,"labels": {"name": "redis-slave"}},"spec": {"containers": [{"name": "slave","image": "192.168.6.14:5000/kubernetes/redis-slave:v2","ports": [{"containerPort": 6379,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 2}},{"metadata": {"name": "test-1","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/test-1","uid": "2024261b-3f40-11e5-ae3a-fa163ea883d9","resourceVersion": "43578976","creationTimestamp": "2015-08-10T09:13:55Z","labels": {"app": "web","name": "test-1","type": "ui","version": "1.0"}},"spec": {"replicas": 1,"selector": {"name": "test-1"},"template": {"metadata": {"creationTimestamp": null,"labels": {"name": "test-1"}},"spec": {"containers": [{"name": "test-1","image": "192.168.6.14:5000/kubernetes/pause","ports": [{"containerPort": 88,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 1}},{"metadata": {"name": "test-2","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/test-2","uid": "6668b56f-4094-11e5-ae3a-fa163ea883d9","resourceVersion": "44353937","creationTimestamp": "2015-08-12T01:49:42Z","labels": {"Type": "ui","Version": "1.0","app": "web","name": "test-2"}},"spec": {"replicas": 3,"selector": {"name": "test-2"},"template": {"metadata": {"creationTimestamp": null,"labels": {"name": "test-2"}},"spec": {"containers": [{"name": "test-2","image": "192.168.6.14:5000/kubernetes/pause","ports": [{"containerPort": 88,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 3}},{"metadata": {"name": "wordpress-2","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/wordpress-2","uid": "7d88186f-40be-11e5-ae3a-fa163ea883d9","resourceVersion": "44454120","creationTimestamp": "2015-08-12T06:50:59Z","labels": {"app": "web","name": "wordpress-2","type": "ui","version": "1.0"}},"spec": {"replicas": 3,"selector": {"name": "wordpress-2"},"template": {"metadata": {"creationTimestamp": null,"labels": {"name": "wordpress-2"}},"spec": {"containers": [{"name": "wordpress-2","image": "cael/wordpress","ports": [{"containerPort": 80,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 3}}]}`)
// 	// body := []byte(`{"kind": "ReplicationControllerList","apiVersion": "v1","metadata": {"selfLink": "/api/v1/namespaces/default/replicationcontrollers","resourceVersion": "46750615"},"items": []}`)
// 	// body := []byte(`{"kind": "ReplicationControllerList","apiVersion": "v1","metadata": {"selfLink": "/api/v1/namespaces/default/replicationcontrollers","resourceVersion": "46750615"},"items": [{"metadata": {"name": "frontend","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/frontend","uid": "ecfab3a0-3f3a-11e5-ae3a-fa163ea883d9","resourceVersion": "43538292","creationTimestamp": "2015-08-10T08:36:42Z","labels": {"name": "frontend"}},"spec": {"replicas": 3,"selector": {"name": "frontend"},"template": {"metadata": {"creationTimestamp": null,"labels": {"name": "frontend"}},"spec": {"containers": [{"name": "php-redis","image": "192.168.6.14:5000/kubernetes/example-guestbook-php-redis:v2","ports": [{"containerPort": 80,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 3}},{"metadata": {"name": "frontend","selfLink": "/api/v1/namespaces/default/replicationcontrollers/frontend","uid": "ecfab3a0-3f3a-11e5-ae3a-fa163ea883d9","resourceVersion": "43538292"},"spec": {"replicas": 3,"selector": {"name": "frontend"},"template": {"metadata": {"creationTimestamp": null,"labels": {"name": "frontend"}},"spec": {"containers": [{"name": "php-redis","image": "192.168.6.14:5000/kubernetes/example-guestbook-php-redis:v2","ports": [{"containerPort": 80,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 3}}]}`)
// 	// body := []byte(`{"kind": "ReplicationControllerList","apiVersion": "v1","metadata": {"selfLink": "/api/v1/namespaces/default/replicationcontrollers","resourceVersion": "46750615"},"items": [{"metadata": {"name": "frontend","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/frontend","uid": "ecfab3a0-3f3a-11e5-ae3a-fa163ea883d9","resourceVersion": "43538292","creationTimestamp": "2015-08-10T08:36:42Z","labels": {"name": "frontend","name2": "frontend2"}},"spec": {"replicas": 3,"selector": {"name": "frontend"},"template": {"metadata": {"creationTimestamp": null,"labels": {"name": "frontend"}},"spec": {"containers": [{"name": "php-redis","image": "192.168.6.14:5000/kubernetes/example-guestbook-php-redis:v2","ports": [{"containerPort": 80,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 3}},{"metadata": {"name": "frontend","selfLink": "/api/v1/namespaces/default/replicationcontrollers/frontend","uid": "ecfab3a0-3f3a-11e5-ae3a-fa163ea883d9","resourceVersion": "43538292"},"spec": {"replicas": 3,"selector": {"name": "frontend"},"template": {"metadata": {"creationTimestamp": null,"labels": {"name": "frontend"}},"spec": {"containers": [{"name": "php-redis","image": "192.168.6.14:5000/kubernetes/example-guestbook-php-redis:v2","ports": [{"containerPort": 80,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 3}}]}`)
// 	body := []byte(`{"kind": "ServiceList","apiVersion": "v1","metadata": {"selfLink": "/api/v1/namespaces/default/services","resourceVersion": "4794785"},"items": [{"metadata": {"name": "frontend","namespace": "default","selfLink": "/api/v1/namespaces/default/services/frontend","uid": "eb883eca-47b2-11e5-8e88-fa163e293324","resourceVersion": "3755","creationTimestamp": "2015-08-21T03:15:48Z","labels": {"name": "frontend"}},"spec": {"ports": [{"protocol": "TCP","port": 80,"targetPort": 80,"nodePort": 30000}],"selector": {},"clusterIP": "10.10.108.113","type": "NodePort","sessionAffinity": "None"},"status": {"loadBalancer": {}}},{"metadata": {"name": "golang","namespace": "default","selfLink": "/api/v1/namespaces/default/services/golang","uid": "bd2818e9-4b94-11e5-bf97-fa163e293324","resourceVersion": "1813452","creationTimestamp": "2015-08-26T01:49:50Z","labels": {"name": "golang"}},"spec": {"ports": [{"protocol": "TCP","port": 9090,"targetPort": 9090,"nodePort": 30001}],"selector": {"name": "golang"},"clusterIP": "10.10.218.111","type": "NodePort","sessionAffinity": "None"},"status": {"loadBalancer": {}}},{"metadata": {"name": "golang-dh","namespace": "default","selfLink": "/api/v1/namespaces/default/services/golang-dh","uid": "c7b5b1d6-4b94-11e5-bf97-fa163e293324","resourceVersion": "1813572","creationTimestamp": "2015-08-26T01:50:08Z","labels": {"name": "golang-dh"}},"spec": {"ports": [{"protocol": "TCP","port": 9000,"targetPort": 9000,"nodePort": 30002}],"selector": {},"clusterIP": "10.10.67.4","type": "NodePort","sessionAffinity": "None"},"status": {"loadBalancer": {}}},{"metadata": {"name": "kube-dns","namespace": "default","selfLink": "/api/v1/namespaces/default/services/kube-dns","uid": "965f7469-47b2-11e5-8e88-fa163e293324","resourceVersion": "3117","creationTimestamp": "2015-08-21T03:13:25Z","labels": {"k8s-app": "kube-dns","kubernetes.io/cluster-service": "true","kubernetes.io/name": "KubeDNS"}},"spec": {"ports": [{"name": "dns","protocol": "UDP","port": 53,"targetPort": 53,"nodePort": 0},{"name": "dns-tcp","protocol": "TCP","port": 53,"targetPort": 53,"nodePort": 0}],"selector": {"k8s-app": "kube-dns"},"clusterIP": "10.10.0.10","type": "ClusterIP","sessionAffinity": "None"},"status": {"loadBalancer": {}}},{"metadata": {"name": "kubernetes","namespace": "default","selfLink": "/api/v1/namespaces/default/services/kubernetes","resourceVersion": "2262","creationTimestamp": null,"labels": {"component": "apiserver","provider": "kubernetes"}},"spec": {"ports": [{"protocol": "TCP","port": 443,"targetPort": 443,"nodePort": 0}],"clusterIP": "10.10.0.1","type": "ClusterIP","sessionAffinity": "None"},"status": {"loadBalancer": {}}},{"metadata": {"name": "monitoring-grafana","namespace": "default","selfLink": "/api/v1/namespaces/default/services/monitoring-grafana","uid": "8784e488-4d2c-11e5-bf97-fa163e293324","resourceVersion": "2558826","creationTimestamp": "2015-08-28T02:28:55Z","labels": {"kubernetes.io/cluster-service": "true","kubernetes.io/name": "Grafana"}},"spec": {"ports": [{"protocol": "TCP","port": 80,"targetPort": 8080,"nodePort": 30010}],"selector": {"k8s-app": "influxGrafana"},"clusterIP": "10.10.226.112","type": "NodePort","sessionAffinity": "None"},"status": {"loadBalancer": {}}},{"metadata": {"name": "monitoring-heapster","namespace": "default","selfLink": "/api/v1/namespaces/default/services/monitoring-heapster","uid": "8efd7590-4d2c-11e5-bf97-fa163e293324","resourceVersion": "2558875","creationTimestamp": "2015-08-28T02:29:07Z","labels": {"kubernetes.io/cluster-service": "true","name": "monitoring-heapster"}},"spec": {"ports": [{"protocol": "TCP","port": 80,"targetPort": 8082,"nodePort": 0}],"selector": {"k8s-app": "heapster"},"clusterIP": "10.10.53.18","type": "ClusterIP","sessionAffinity": "None"},"status": {"loadBalancer": {}}},{"metadata": {"name": "monitoring-influxdb","namespace": "default","selfLink": "/api/v1/namespaces/default/services/monitoring-influxdb","uid": "958aa96e-4d2c-11e5-bf97-fa163e293324","resourceVersion": "2558922","creationTimestamp": "2015-08-28T02:29:18Z","labels": {"k8s-app": "influxGrafana","kubernetes.io/cluster-service": "true","kubernetes.io/name": "influxGrafana"}},"spec": {"ports": [{"name": "http","protocol": "TCP","port": 8083,"targetPort": 8083,"nodePort": 0},{"name": "api","protocol": "TCP","port": 8086,"targetPort": 8086,"nodePort": 0}],"selector": {"k8s-app": "influxGrafana"},"clusterIP": "10.10.75.125","type": "ClusterIP","sessionAffinity": "None"},"status": {"loadBalancer": {}}},{"metadata": {"name": "redis-master","namespace": "default","selfLink": "/api/v1/namespaces/default/services/redis-master","uid": "f2f2355b-47b2-11e5-8e88-fa163e293324","resourceVersion": "3820","creationTimestamp": "2015-08-21T03:16:01Z","labels": {"name": "redis-master"}},"spec": {"ports": [{"protocol": "TCP","port": 6379,"targetPort": 6379,"nodePort": 0}],"selector": {"name": "redis-master"},"clusterIP": "10.10.87.168","type": "ClusterIP","sessionAffinity": "None"},"status": {"loadBalancer": {}}},{"metadata": {"name": "redis-slave","namespace": "default","selfLink": "/api/v1/namespaces/default/services/redis-slave","uid": "f6fc18df-47b2-11e5-8e88-fa163e293324","resourceVersion": "3828","creationTimestamp": "2015-08-21T03:16:07Z","labels": {"name": "redis-slave"}},"spec": {"ports": [{"protocol": "TCP","port": 6379,"targetPort": 6379,"nodePort": 0}],"selector": {"name": "redis-slave"},"clusterIP": "10.10.153.107","type": "ClusterIP","sessionAffinity": "None"},"status": {"loadBalancer": {}}}]}`)
// 	return body	
// }


