package qinv

// QEntity covers all entities within the inventory (machine, container, docker-engine)
type QEntity interface {
	GetID() string
	GetMe() QEntity
}
