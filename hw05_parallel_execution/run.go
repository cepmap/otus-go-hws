package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type errCounter struct {
	mu    sync.Mutex
	value int
}

func (e *errCounter) Add() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.value++
}

func (e *errCounter) Get() int {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.value
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// n - max workers
	// m - max errors
	if m < 0 {
		m = 0
	}
	errCounter := errCounter{}

	var wg sync.WaitGroup
	tasksChannel := make(chan Task)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, t := range tasks {
			if errCounter.Get() <= m {
				tasksChannel <- t
			}
		}
		close(tasksChannel)
	}()

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasksChannel {
				if errCounter.Get() > m {
					return
				} else if task() != nil {
					errCounter.Add()
				}
			}
		}()
	}

	wg.Wait()
	if errCounter.Get() <= m {
		return nil
	}
	return ErrErrorsLimitExceeded
}
