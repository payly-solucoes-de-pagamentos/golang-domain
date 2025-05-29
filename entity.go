package domain

import events "github.com/payly-solucoes-de-pagamentos/golang-events"

type Entity struct {
	id              string
	eventDispatcher events.IEventDispatcher
}

func (entity *Entity) AddDomainEvent(domainEvent events.IDomainEvent) {
	entity.eventDispatcher.AddDomainEvent(domainEvent)
}

func (entity *Entity) SetDispatcher(dispatcher events.IEventDispatcher) {
	entity.eventDispatcher = dispatcher
}
