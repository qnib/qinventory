package qinv

// Machine holds informations about a container
type Machine struct {
	ID       string `json:"id"`
	Hostname string `json:"hostname"`
	Kernel   string `json:"kernel"`
}

// NewMachine returns a container instance
func NewMachine(hostname, kernel string) (Machine, error) {
	return Machine{
		ID:       hostname,
		Hostname: hostname,
		Kernel:   kernel,
	}, nil
}

// GetID returns the ID of the instance
func (m *Machine) GetID() string {
	return m.ID
}

// GetMe returns self as Qentity
func (m *Machine) GetMe() QEntity {
	return m
}
