package context

type UpdateConfig struct {
	Config
	UpdateModel UpdateModel
	Fields      []string
}
