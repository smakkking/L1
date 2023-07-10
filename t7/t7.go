package t7

import "sync"

type MapConcurrent struct {
	mu         sync.Mutex
	secret_map map[int]int
}

func NewMapConcurrent(m *map[int]int) *MapConcurrent {
	return &MapConcurrent{
		mu:         sync.Mutex{},
		secret_map: *m,
	}
}

func (m *MapConcurrent) SetValue(key int, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.secret_map[key] = value
}

func (m *MapConcurrent) GetValue(key int) int {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.secret_map[key]
}

func (m *MapConcurrent) DeleteKey(key int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.secret_map, key)
}

func DoTask() {
	var m = make(map[int]int)
	NewMapConcurrent(&m)
}
