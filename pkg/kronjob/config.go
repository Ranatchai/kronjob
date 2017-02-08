package kronjob

// Config contains all the information about the k8s cluster and local configuration
type Config struct {
	AllowParallel bool
	ContainerName string
	Deadline      int
	Namespace     string
	Schedule      string
	Template      string
	Verbose       bool
}