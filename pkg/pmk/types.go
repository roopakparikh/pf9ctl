package pmk

// CloudProviderType specifies the infrastructure where the cluster runs
type CloudProviderType string

// CNIBackend specifies the networking solution used for the k8s cluster
type CNIBackend string

const (
	// ContextFile is used to store context information
	ContextFile string = "pf9_context.json"
	// DefaultContextDir is the default path for context file
	DefaultContextDir string = "~/"
	// LogFile specifies the filename to which CLI logs o/p and errors
	LogFile string = "pf9ctl.log"

	// AWS cloud provider
	AWS CloudProviderType = "aws"
	// GCP cloud provider
	GCP CloudProviderType = "gcp"
	// BareOS on-prem cloud provider
	BareOS CloudProviderType = "local"
	// Openstack cloud provider
	Openstack CloudProviderType = "openstack"
	// Flannel network plugin for k8s networking
	Flannel CNIBackend = "flannel"
	// Calico network plugin for k8s networking
	Calico CNIBackend = "calico"
	// Weave network plugin for k8s networking
	Weave CNIBackend = "weave"
)

// Context specifies information required to connect to the PF9 Controller
