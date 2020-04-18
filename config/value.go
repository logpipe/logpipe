package config

import (
	"go.uber.org/config"
	"strconv"
)

type Value struct {
	value config.Value
}

func (v *Value) GetValue(key string) *Value {
	value := v.value.Get(key)
	return &Value{value: value}
}
func (v *Value) GetArray(key string) []*Value {
	var values []*Value
	target := v.value.Get(key)
	err := target.Populate(&values)
	if err == nil {
		for i, val := range values {
			val.value = target.Get(strconv.Itoa(i))
		}
	}
	return values
}

func (v *Value) Parse(target interface{}) error {
	return v.value.Populate(target)
}
