package t7

import "sync"

type MapConcurrent[T1 ~int | ~string | ~float64, T2 any] struct {
	mu         sync.Mutex
	secret_map map[T1]T2
}

func NewMapConcurrent(m *map[any]any) *MapConcurrent {
	return &MapConcurrent{
		mu:         sync.Mutex{},
		secret_map: *m,
	}
}

func (m *MapConcurrent) SetValue(key any, value any) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.secret_map[key] = value
}

func (m *MapConcurrent) GetValue(key any) any {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.secret_map[key]
}

func (m *MapConcurrent) DeleteKey(key any) any {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.secret_map, key)
}

func main() {
	var m = make(map[int]int)
	s := NewMapConcurrent(&m)
}
