package config

import (
	"gopkg.in/yaml.v3"
)

type Value struct {
	node *yaml.Node
}

func (v *Value) UnmarshalYAML(value *yaml.Node) error {
	v.node = value
	return nil
}

func (v *Value) IsMap() bool {
	return v.node.Kind == yaml.MappingNode
}
func (v *Value) IsArray() bool {
	return v.node.Kind == yaml.SequenceNode
}
func (v *Value) IsScalar() bool {
	return v.node.Kind == yaml.ScalarNode
}

func (v *Value) Get(key string) *Value {
	if v.node.Kind == yaml.MappingNode {
		for i := 0; i < len(v.node.Content); i = i + 2 {
			k := v.node.Content[i].Value
			if k == key {
				return &Value{node: v.node.Content[i+1]}
			}
		}
	}
	return nil
}

func (v *Value) Map() map[string]*Value {
	if v.node.Kind == yaml.MappingNode {
		values := make(map[string]*Value, len(v.node.Content)/2)
		for i := 0; i < len(v.node.Content); i = i + 2 {
			key := v.node.Content[i].Value
			value := v.node.Content[i+1]
			values[key] = &Value{node: value}
		}
		return values
	}
	return nil
}

func (v *Value) Array() []*Value {
	if v.node.Kind == yaml.SequenceNode {
		values := make([]*Value, len(v.node.Content))
		for i, node := range v.node.Content {
			values[i] = &Value{node: node}
		}
		return values
	}
	return nil
}

func (v *Value) GetString(key string) string {
	var str string
	_ = v.Get(key).Parse(&str)
	return str
}

func (v *Value) GetInt(key string) int {
	var val int
	_ = v.Get(key).Parse(&val)
	return val
}

func (v *Value) GetBool(key string) bool {
	var val bool
	_ = v.Get(key).Parse(&val)
	return val
}

func (v Value) Parse(target interface{}) error {
	return v.node.Decode(target)
}
