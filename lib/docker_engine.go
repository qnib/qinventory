package qinv

import ()

// DockerEngine holds the information of an Machine
type DockerEngine struct {
    DockerHost string
    Machine Machine
}

// NewDockerEngine returns a DockerEngine instance
func NewDockerEngine(dhost string, m Machine) (DockerEngine, error) {
    return DockerEngine{
        DockerHost: dhost,
        Machine: m,
    }, nil
}
