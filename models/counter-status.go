package models

import "errors"

type CounterStatus struct {
	Count int `form:"count"`
}

func (cs *CounterStatus) Increment() {
	cs.Count = cs.Count + 1
}

func (cs *CounterStatus) Decrement() error {
	if cs.Count > 0 {
		cs.Count = cs.Count - 1
		return nil
	}
	return errors.New("can not decrement under zero")
}
