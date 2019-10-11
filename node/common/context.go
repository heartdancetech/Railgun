package common

import (
	"sync"
)

type Context struct {
	services map[string]map[string]string
	lock sync.RWMutex
}


func (c *Context) AddServices(key,node,value string)  {
	defer c.lock.Unlock()
	c.lock.Lock()
	c.services[key][node] = value
}

func (c *Context) DelServices(key,node string)  {
	defer c.lock.Unlock()
	c.lock.Lock()
	delete(c.services[key], node)

}

func (c *Context) GetServices() map[string]map[string]string{
	return c.services
}

func (c *Context) UpdateServices(key,node,value string) {
	defer c.lock.Unlock()
	c.lock.Lock()
	if c.services[key] == nil {
		c.services[key] = make(map[string]string)
	}
	c.services[key][node] = value
}

func InitContext(data map[string]map[string]string) *Context {
	return &Context{services:data}
}