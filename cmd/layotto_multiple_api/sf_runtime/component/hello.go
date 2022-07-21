package component_sf

import "mosn.io/layotto/components/custom"

type Hello interface {
	custom.Component
	SayHello1(name string) (string, error)
}
