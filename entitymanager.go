package ecs

import "errors"

// EntityManager is a list of all entities in the game
type EntityManager struct {
	entities []*Entity
	components map[string][]*Entity
}

var _em = new(EntityManager)
_em.components = make(map[string][]*Entity)

func (manager EntityManager) addEntity(entity *Entity) {
	manager.entities = append(manager.entities, entity)
}

func (manager EntityManager) addComponent(componentName string, entity *Entity) {
	manager.components[componentName] = append(manager.components[componentName], entity)
}

func (manager EntityManager) findEntityByID(entityID uint32) int {
	for i, entity := range manager.entities {
		if entity.ID == entityID {
			return i
		}
	}
	return -1
}

func (manager EntityManager) findEntitiesByComponent(componentName string) []*Entity {
	var entities []*Entity
	for _, entity := range manager.entities {
		if _, found := entity.Components[componentName]; found {
			entities = append(entities, entity)
		}
	}

	return entities
}

// remove will remove the entity from each map of entities
func (manager EntityManager) remove(entityID uint32) error {
	i := manager.findEntityByID(entityID)
	if i != -1 {
		manager.entities = append(manager.entities[:i], manager.entities[i+1:]...)
		return nil
	}
	return errors.New("Entity not found")
}
