package stack

import (
	"fmt"
)

type Transaction struct {
	store  map[string]string
	parent *Transaction
}

type Stack struct {
	global map[string]string
	top    *Transaction
}

func New() *Stack {
	return &Stack{global: make(map[string]string)}
}

func (s Stack) Set(k, v string) {

	if s.top == nil {
		s.global[k] = v
	} else {
		s.top.store[k] = v
	}

}

func (s Stack) Get(k string) {

	var ok bool
	var v string

	if s.top == nil {
		v, ok = s.global[k]
	} else {
		v, ok = s.top.store[k]
	}

	if !ok {
		fmt.Printf("%s not set\n", k)
		return
	}

	fmt.Println(v)
}

func (s Stack) Delete(k string) {

	if s.top == nil {
		delete(s.global, k)
	} else {
		delete(s.top.store, k)
	}

}

func (s *Stack) Push() {

	s.top = &Transaction{
		store:  make(map[string]string),
		parent: s.top,
	}

}

func (s *Stack) End() {

	if s.top != nil {
		s.top = s.top.parent
	}

}

func (s *Stack) Roll() {

	if s.top != nil {
		s.top.store = make(map[string]string)
	}

}

func (s *Stack) Commit() {

	if s.top != nil {

		if s.top.parent != nil {
			for k, v := range s.top.store {
				s.top.parent.store[k] = v
			}
		} else {
			for k, v := range s.top.store {
				s.global[k] = v
			}
		}

		s.End()
	}

}
