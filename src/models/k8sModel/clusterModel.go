package k8sModel

import (
	bm "models/baseModel"
	"encoding/json"
	"libs/request"
	"models/k8sModel/format"
)

type ReplicationControllerModel struct {
	bm.BaseModel
	req request.Request
}


// func (this *ClusterModel) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

// func (this *ClusterModel) GetBaseUrl() string {
// 	return "http://192.168.6.13:8080/api/v1" + "/namespaces/default/"
// }


// func (this *ClusterModel) GetCluster() (format.ReplicationController) {
// }

// curl -H "Content-Type: application/json" -X GET  http://192.168.12.8:8080/api/v1/namespaces/ReplicationController/replicationcontrollers
func (this *ReplicationControllerModel) GetReplicationControllerList() (format.ReplicationControllerList) {
	body := this.req.Get("http://192.168.12.8:8080/api/v1/namespaces/default/replicationcontrollers")

    var data format.ReplicationControllerList
	json.Unmarshal(body, &data)
	// fmt.Println(err)
	// fmt.Println(b)
	return data
}

func (this *ReplicationControllerModel) GetReplicationController(name string) (format.ReplicationController) {
	body := this.req.Get("http://192.168.12.8:8080/api/v1/namespaces/default/replicationcontrollers/"+name)

    var data format.ReplicationController
	json.Unmarshal(body, &data)
	// fmt.Println(err)
	// fmt.Println(b)
	return data
}

func (this *ReplicationControllerModel) CreateReplicationController(data format.ReplicationController) (format.ReplicationController) {
	body := this.req.Post("http://192.168.12.8:8080/api/v1/namespaces/default/replicationcontrollers", data)
	
	var _replication_controller_data format.ReplicationController
	json.Unmarshal(body, &_replication_controller_data)

	return _replication_controller_data
}

func (this *ReplicationControllerModel) UpdateReplicationController(name string, data format.ReplicationController) (format.ReplicationController) {
	body := this.req.Put("http://192.168.12.8:8080/api/v1/namespaces/default/replicationcontrollers/"+name, data)
	
	var _replication_controller_data format.ReplicationController
	json.Unmarshal(body, &_replication_controller_data)

	return _replication_controller_data
}

func (this *ReplicationControllerModel) DeleteReplicationController(name string) (format.ReplicationController) {
	body := this.req.Delete("http://192.168.12.8:8080/api/v1/namespaces/default/replicationcontrollers/"+name, nil)
	
	var _replication_controller_data format.ReplicationController
	json.Unmarshal(body, &_replication_controller_data)
	
	return _replication_controller_data
}

func (this *ReplicationControllerModel) getPagInfo_data() []byte {
	body := []byte(`{"kind": "ReplicationControllerList","apiVersion": "v1","metadata": {"selfLink": "/api/v1/namespaces/default/replicationcontrollers","resourceVersion": "1000717"},"items": [{"metadata": {"name": "app","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/app","uid": "4af54540-5603-11e5-9b89-fa163e78675c","resourceVersion": "943709","creationTimestamp": "2015-09-08T08:26:24Z","labels": {"name": "app"}},"spec": {"replicas": 3,"selector": {"name": "app"},"template": {"metadata": {"creationTimestamp": null,"labels": {"name": "app"}},"spec": {"containers": [{"name": "app","image": "192.168.12.7:5000/app","ports": [{"containerPort": 8000,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 3}},{"metadata": {"name": "dev","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/dev","uid": "47e1772f-5603-11e5-9b89-fa163e78675c","resourceVersion": "645593","creationTimestamp": "2015-09-08T08:26:19Z","labels": {"name": "dev"}},"spec": {"replicas": 1,"selector": {"name": "dev"},"template": {"metadata": {"creationTimestamp": null,"labels": {"name": "dev"}},"spec": {"containers": [{"name": "dev","image": "192.168.12.7:5000/dev:p31000","ports": [{"containerPort": 80,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 1}},{"metadata": {"name": "frontend","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/frontend","uid": "15d2e7ac-550c-11e5-820f-fa163e78675c","resourceVersion": "234954","creationTimestamp": "2015-09-07T02:56:49Z","labels": {"name": "frontend"}},"spec": {"replicas": 3,"selector": {"name": "frontend"},"template": {"metadata": {"creationTimestamp": null,"labels": {"name": "frontend"}},"spec": {"containers": [{"name": "php-redis","image": "192.168.12.7:5000/kubernetes/example-guestbook-php-redis:v2","ports": [{"containerPort": 80,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 3}},{"metadata": {"name": "kube-dns-v3","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/kube-dns-v3","uid": "9db5efa5-5507-11e5-820f-fa163e78675c","resourceVersion": "234971","creationTimestamp": "2015-09-07T02:24:50Z","labels": {"k8s-app": "kube-dns-v3","kubernetes.io/cluster-service": "true"}},"spec": {"replicas": 1,"selector": {"k8s-app": "kube-dns","version": "v3"},"template": {"metadata": {"creationTimestamp": null,"labels": {"k8s-app": "kube-dns","kubernetes.io/cluster-service": "true","version": "v3"}},"spec": {"volumes": [{"name": "dns-token","secret": {"secretName": "token-system-dns"}}],"containers": [{"name": "etcd","image": "192.168.12.7:5000/google_containers/etcd:2.0.9","command": ["/usr/local/bin/etcd","-listen-client-urls","http://127.0.0.1:2379,http://127.0.0.1:4001","-advertise-client-urls","http://127.0.0.1:2379,http://127.0.0.1:4001","-initial-cluster-token","skydns-etcd"],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"},{"name": "kube2sky","image": "192.168.12.7:5000/google_containers/kube2sky:1.9","args": ["-domain=kubernetes.local","-kubecfg_file=/etc/dns_token/kubeconfig","-v=2","-kube_master_url=https://${KUBERNETES_SERVICE_HOST}:${KUBERNETES_SERVICE_PORT}"],"resources": {},"volumeMounts": [{"name": "dns-token","readOnly": true,"mountPath": "/etc/dns_token"}],"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"},{"name": "skydns","image": "192.168.12.7:5000/google_containers/skydns:2015-03-11-001","args": ["-machines=http://localhost:4001","-addr=0.0.0.0:53","-domain=kubernetes.local.","-verbose=true"],"ports": [{"name": "dns","containerPort": 53,"protocol": "UDP"},{"name": "dns-tcp","containerPort": 53,"protocol": "TCP"}],"resources": {},"livenessProbe": {"exec": {"command": ["/bin/sh","-c","nslookup kubernetes.default.kubernetes.local localhost \u003e/dev/null"]},"initialDelaySeconds": 30,"timeoutSeconds": 5},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "Default"}}},"status": {"replicas": 1}},{"metadata": {"name": "redis-master","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/redis-master","uid": "1c244c76-550c-11e5-820f-fa163e78675c","resourceVersion": "233648","creationTimestamp": "2015-09-07T02:57:00Z","labels": {"name": "redis-master"}},"spec": {"replicas": 1,"selector": {"name": "redis-master"},"template": {"metadata": {"creationTimestamp": null,"labels": {"name": "redis-master"}},"spec": {"containers": [{"name": "master","image": "192.168.12.7:5000/redis","ports": [{"containerPort": 6379,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 1}},{"metadata": {"name": "redis-slave","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/redis-slave","uid": "3895c897-550c-11e5-820f-fa163e78675c","resourceVersion": "234978","creationTimestamp": "2015-09-07T02:57:48Z","labels": {"name": "redis-slave"}},"spec": {"replicas": 2,"selector": {"name": "redis-slave"},"template": {"metadata": {"creationTimestamp": null,"labels": {"name": "redis-slave"}},"spec": {"containers": [{"name": "slave","image": "192.168.12.7:5000/kubernetes/redis-slave:v2","ports": [{"containerPort": 6379,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 2}},{"metadata": {"name": "web-app","namespace": "default","selfLink": "/api/v1/namespaces/default/replicationcontrollers/web-app","uid": "ccf3572a-5603-11e5-9b89-fa163e78675c","resourceVersion": "646576","creationTimestamp": "2015-09-08T08:30:02Z","labels": {"Type": "test","Version": "1.0","app": "web-app","name": "web-app"}},"spec": {"replicas": 1,"selector": {"name": "web-app"},"template": {"metadata": {"creationTimestamp": null,"labels": {"name": "web-app"}},"spec": {"containers": [{"name": "web-app","image": "training/webapp","ports": [{"containerPort": 5000,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst"}}},"status": {"replicas": 1}}]}`)
	return body	
}
