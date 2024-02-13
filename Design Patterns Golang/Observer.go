package main

import (
	"fmt"
)

// Subject represents the subject
type Subject struct {
	observers []*Observer
	state     string
}

// Observer represents the observer
type Observer struct {
	id int
}

// NewSubject creates a new subject
func NewSubject() *Subject {
	return &Subject{}
}

// NewObserver creates a new observer
func NewObserver(id int) *Observer {
	return &Observer{id}
}

// Attach attaches an observer to the subject
func (s *Subject) Attach(observer *Observer) {
	s.observers = append(s.observers, observer)
}

// Detach detaches an observer from the subject
func (s *Subject) Detach(observer *Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// Notify notifies all observers of the state change
func (s *Subject) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.state)
	}
}

// SetState sets the state of the subject and notifies observers
func (s *Subject) SetState(state string) {
	s.state = state
	s.Notify()
}

// Update updates the observer with the subject's state
func (o *Observer) Update(state string) {
	fmt.Printf("Observer %d: updated with state %s\n", o.id, state)
}

func main() {
	// Create subject
	subject := NewSubject()

	// Create observers
	observer1 := NewObserver(1)
	observer2 := NewObserver(2)
	observer3 := NewObserver(3)

	// Attach observers to the subject
	subject.Attach(observer1)
	subject.Attach(observer2)
	subject.Attach(observer3)

	// Set state of the subject
	subject.SetState("State 1")
	subject.SetState("State 2")

	// Detach observer2
	subject.Detach(observer2)

	// Set state again
	subject.SetState("State 3")
}
