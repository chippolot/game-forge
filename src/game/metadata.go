package game

type IMetadata interface {
	GetName() string
	GetDescription() string
}

type Metadata struct {
	name       string
	decription string
}

func NewMetadata(name string, description string) *Metadata {
	return &Metadata{
		name:        name,
		description: description
	}
}

func (m *Metadata) GetName() string {
	return m.name
}

func (m *Metadata) GetDescription() string {
	return m.decription
}