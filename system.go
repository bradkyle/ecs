package ecs

type System struct {
	componentName string
	update        func([]*Entity)
}

var _systems []*System

// AddSystem will add the system logic functions
func AddSystem(componentName string, update func([]*Entity)) {
	var system = new(System)
	system.componentName = componentName
	system.update = update
	_systems = append(_systems, system)
}

// SystemLoop will loop over each system
func SystemLoop() {
	for i := 0; i < len(_systems); i++ {
		go _systems[i].update(_em.components[_system[i].componentName])
	}
}
