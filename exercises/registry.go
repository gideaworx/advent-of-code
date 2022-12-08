package exercises

import (
	"errors"
	"fmt"
	"io"
	"log"
	"sync"
)

// ExerciseImplementation is the function signature that all exercises must implement
type ExerciseImplementation func(io.Reader, *log.Logger) (any, error)

var ErrAlreadyRegistered = errors.New("already registered")
var ErrNotFound = errors.New("not found")

// Exercise is a part's explanation and its implementation
type Exercise struct {
	Notes          string                 `json:"notes"`
	Implementation ExerciseImplementation `json:"-"`
}

// Day is an input plus a collection of exercises that use that input
type Day struct {
	Input     string              `json:"input"`
	Exercises map[string]Exercise `json:",inline"`
}

// Registry is a synchronized collection of days
type Registry struct {
	days map[string]Day
	m    *sync.RWMutex
}

func NewRegistry() *Registry {
	return &Registry{
		days: map[string]Day{},
		m:    &sync.RWMutex{},
	}
}

func (r *Registry) Register(name string, day Day) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.days[name]; ok {
		return fmt.Errorf("%s is %w", name, ErrAlreadyRegistered)
	}

	r.days[name] = day
	return nil
}

func (r *Registry) GetDay(name string) (Day, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var (
		d  Day
		ok bool
	)
	if d, ok = r.days[name]; !ok {
		return Day{}, fmt.Errorf("%s %w", name, ErrNotFound)
	}

	return d, nil
}
