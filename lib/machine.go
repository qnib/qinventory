package qinv

// Machine holds informations about a container
type Machine struct {
    Hostname string `json:"hostname"`
    Kernel string `json:"kernel"`
}

// NewMachine returns a container instance
func NewMachine(hostname, kernel string) (Machine, error) {
    return Machine{
        Hostname: hostname,
        Kernel: kernel,
    }, nil
}
