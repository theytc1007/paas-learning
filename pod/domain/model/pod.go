package model

type Pod struct {
	ID           int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	PodName      string `gorm:"unique_index;not_null" json:"pod_name"`
	PodNamespace string `json:"pod_namespace"`

	// Pod 所属团队 ?
	PodTeamID int64 `json:"pod_team_id"`

	PodReplicas  int32   `json:"pod_replicas"`
	PodCpuMin    float32 `json:"pod_cpu_min"`
	PodCpuMax    float32 `json:"pod_cpu_max"`
	PodMemoryMin float32 `json:"pod_memory_min"`
	PodMemoryMax float32 `json:"pod_memory_max"`

	// Pod 开放端口
	PodPort []PodPort `gorm:"ForeignKey:PodID" json:"pod_port"`

	// Pod 使用的环境变量
	PodEnv []PodEnv `gorm:"ForeignKey:PodID" json:"pod_env"`

	// 镜像拉取策略
	// Always: 总是拉取 pull
	// IfNotPresent: 默认，本地有则使用本地镜像，否则才拉取
	// Never: 只使用本地镜像，从不拉取
	PodPullPolicy string `json:"pod_pull_policy"`

	// 重启策略
	// Always: 容器失效时，由 kubelet 自动重启该容器
	// OnFailure: 容器终止运行且错误码不为 0 时，由 kubelet 自动重启该容器
	// Never: 不重启
	// 控制方式为 replica 和 DaemonSet 时，必须保证容器的运行
	// Job: OnFailure 或者 Never，确保容器执行完成后不再重启
	PodRestart string `json:"pod_restart"`

	// 发布策略
	// recreate rolling-update blue/green canary a/b_testing
	PodType string `json:"pod_type"`

	PodImageWithTag string `json:"pod_image_with_tag"`

	// TODO 挂盘/域名设置
}
