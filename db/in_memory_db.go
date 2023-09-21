package db

import "sync"

type InMemDB[V any] struct {
	m   map[string]V
	mux sync.Mutex
}

func NewInMemDB[V any]() *InMemDB[V] {
	return &InMemDB[V]{
		m: make(map[string]V),
	}
}

func (db *InMemDB[V]) Get(key string) (V, bool) {
	defer db.mux.Unlock()
	db.mux.Lock()

	val, ok := db.m[key]
	return val, ok
}

func (db *InMemDB[V]) GetAll() ([]V, error) {
	defer db.mux.Unlock()
	db.mux.Lock()

	var vals []V
	for _, val := range db.m {
		vals = append(vals, val)
	}
	return vals, nil
}

func (db *InMemDB[V]) Save(key string, val V) {
	defer db.mux.Unlock()
	db.mux.Lock()
	db.m[key] = val
}
