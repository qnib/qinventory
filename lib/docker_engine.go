package qinv

// DockerEngine holds the information of an Machine
type DockerEngine struct {
	DockerHost string
}

// NewDockerEngine returns a DockerEngine instance
func NewDockerEngine(dhost string) (DockerEngine, error) {
	return DockerEngine{
		DockerHost: dhost,
	}, nil
}
