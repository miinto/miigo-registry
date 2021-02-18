package registry

import "errors"

var instance *Registry

type Registry struct {
	dependencies map[string]*interface{}
}

func (r *Registry) AddDependency(name string, instance *interface{}) error {
	if r.dependencies == nil {
		r.dependencies = make(map[string]*interface{}, 0)
	}

	if _, isOk := r.dependencies[name]; isOk {
		return errors.New("registry dependencies cannot be overwritten")
	}

	r.dependencies[name] = instance
	return nil
}

func (r *Registry) GetDependency(name string) (*interface{},error) {
	if dep,isOk := r.dependencies[name]; isOk {
		return dep, nil
	}

	return nil, errors.New("requested dependency has not been registered")
}

func GetRegistry() *Registry {
	if instance == nil {
		instance = &Registry{}
	}

	return instance
}