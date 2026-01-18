package registry

import "github.com/heyjasonn/design-pattern/factory"

var registry = map[string]func() factory.Payment{}

func Register(key string, value func() factory.Payment) {
	registry[key] = value
}

func Create(key string) factory.Payment {
	return registry[key]()
}
