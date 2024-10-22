package actor

import (
	"sync"

	"github.com/rudrodip/go-actor-model/entities"
)

type TaskActorPool struct {
	pool     []entities.Actor
	poolLock *sync.Mutex
	wg       *sync.WaitGroup
}

func CreateTaskActorPool(wg *sync.WaitGroup) *TaskActorPool {
	return &TaskActorPool{
		pool:     []entities.Actor{},
		poolLock: &sync.Mutex{},
		wg:       wg,
	}
}
