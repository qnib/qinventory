package qinv

// DockerEngine holds the information of an Machine
type DockerEngine struct {
	ID         string
	DockerHost string
}

// NewDockerEngine returns a DockerEngine instance
func NewDockerEngine(id, dhost string) (DockerEngine, error) {
	return DockerEngine{
		ID:         id,
		DockerHost: dhost,
	}, nil
}

// GetID returns the engine's ID
func (e *DockerEngine) GetID() string {
	return e.ID
}
