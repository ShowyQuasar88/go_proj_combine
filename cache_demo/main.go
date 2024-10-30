package main

import "github.com/showyquasar88/proj-combine/cache_demo/cache"

func main() {
	c := cache.NewMemCache()
	c.SetMaxMemory("100MB")
}
