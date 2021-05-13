package main

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

/* Other example:

type Counter struct {
    sync.Mutex
    value int
}

func (c *Counter) Inc() {
    c.Lock()
    defer c.Unlock()
    c.value++
}

Very harmful as it`s exposing Lock() and Unlock()! Embedding types means the methods of that type becomes part of the public interface.
*/
