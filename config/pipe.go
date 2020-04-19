package config

type PipeConf struct {
	Name    string
	Inputs  []InputConf
	Filters []FilterConf
	Outputs []OutputConf
}
