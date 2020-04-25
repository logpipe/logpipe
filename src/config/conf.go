package config

type ConfLoader interface {
	Load(value *Value) error
}
