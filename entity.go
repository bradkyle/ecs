package ecs

import (
	"errors"
	"reflect"
)

var _ID uint32

func nextID() uint32 {
	_ID = _ID + 1
	return _ID
}

// Entity is a game object consisting of a unique ID and a list of components
type Entity struct {
	ID         uint32
	Components map[string]*interface{}
}

// NewEntity creates a new entity and returns it for use
func NewEntity() *Entity {
	var e = new(Entity)
	e.ID = nextID()
	e.Components = make(map[string]*interface{})
	_em.addEntity(e)
	return e
}

// AddComponent will add the component to the entity if not previously added
func (e *Entity) AddComponent(component *interface{}) error {
	componentName := reflect.TypeOf(component).Name()
	if _, found := e.Components[componentName]; !found {
		e.Components[componentName] = component
	} else {
		return errors.New("Component already added.")
	}
	return nil
}

// RemoveComponent will remove the component from the entity if found
func (e *Entity) RemoveComponent(componentName string) error {
	if _, found := e.Components[componentName]; found {
		delete(e.Components, componentName)
	} else {
		return errors.New("Component could not be found")
	}
	return nil
}
