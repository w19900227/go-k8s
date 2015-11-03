package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_volume
type Volume struct {
	Name string `json:"name"`
	// HostPath HostPathVolumeSource `json:"hostPath"`
	// EmptyDir EmptyDirVolumeSource `json:"emptyDir"`
	// GcePersistentDisk GCEPersistentDiskVolumeSource `json:"gcePersistentDisk"`
	// AwsElasticBlockStore AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore"`
	// GitRepo GitRepoVolumeSource `json:"gitRepo"`
	// Secret SecretVolumeSource `json:"secret"`
	// Nfs NFSVolumeSource `json:"nfs"`
	// Iscsi ISCSIVolumeSource `json:"iscsi"`
	// Glusterfs GlusterfsVolumeSource `json:"glusterfs"`
	// PersistentVolumeClaim PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim"`
	// Rbd RBDVolumeSource `json:"rbd"`
}
