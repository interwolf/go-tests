package syncv1

import "sync"

// Counter safely counts concurrently
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc increases Counter's private number by 1
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns Counter's current value
func (c *Counter) Value() int {
	return c.value
}

// NewCounter constructs a new Counter
func NewCounter() *Counter {
	return &Counter{}
}
