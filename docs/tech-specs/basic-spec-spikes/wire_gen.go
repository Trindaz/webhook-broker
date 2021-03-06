// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func InitializeDBEngineName() DBEngineName {
	dbEngineName := NewDBEngineName()
	return dbEngineName
}

func InitializeJobQueue() chan Job {
	maxQueuesConfig := NewMaxQueuesConfig()
	v := NewJobQueue(maxQueuesConfig)
	return v
}

func InitializeDispatcher() *Dispatcher {
	maxWorkersConfig := NewMaxWorkersConfig()
	maxQueuesConfig := NewMaxQueuesConfig()
	v := NewJobQueue(maxQueuesConfig)
	priorityQueue := NewJobPriorityQueue()
	priorityDispatcherSwitch := NewPriorityDispatcherSwitch()
	dispatcher := NewDispatcher(maxWorkersConfig, v, priorityQueue, priorityDispatcherSwitch)
	return dispatcher
}
