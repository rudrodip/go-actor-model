package actor

import (
	"errors"
	"log"

	"sync"

	"github.com/rudrodip/go-actor-model/entities"
	"github.com/rudrodip/go-actor-model/tracker"
)

const taskQueueSize = 10

type TaskActor struct {
	id       int
	closeSig chan bool
	wg       *sync.WaitGroup
	tasks    chan entities.Task
	tracker  *tracker.Tracker
}

func (a *TaskActor) AddTask(task entities.Task) error {
	if len(a.tasks) >= taskQueueSize {
		return errors.New("filled queue")
	}
	a.tasks <- task
	return nil
}

func (a *TaskActor) Start() {
	defer a.wg.Done()
	a.wg.Add(1)
	log.Printf("starting actor :%d", a.id)
	for task := range a.tasks {
		task.Execute()
		a.tracker.GetTrackerChan() <- tracker.CreateCounterTrack(tracker.Task, tracker.Completed)
	}
	log.Printf("stopped actor :%d", a.id)
	a.closeSig <- true
}

func (a *TaskActor) Stop() {
	close(a.tasks)
	<-a.closeSig
}

func CreateActor(wg *sync.WaitGroup, id int, tracker *tracker.Tracker) entities.Actor {
	actor := &TaskActor{
		id:       id,
		wg:       wg,
		closeSig: make(chan bool),
		tasks:    make(chan entities.Task, taskQueueSize),
		tracker:  tracker,
	}
	go actor.Start()
	return actor
}
