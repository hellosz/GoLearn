package method

import (
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	getCache()
// }

type Cache struct {
	sync.Mutex
	mapping map[string]string
}

// 查找键值
func (c *Cache) Lookup(key string) string {
	c.Mutex.Lock()
	value := c.mapping[key]
	c.Mutex.Unlock()

	return value
}

// 写键值
func (c *Cache) Write(key string, value string) {
	c.Mutex.Lock()
	c.mapping[key] = value
	c.Mutex.Unlock()
}

func GetCache() {
	// goroutine continue write
	var ca = Cache{mapping: make(map[string]string)}
	go func() {
		for {
			ca.Write("name", "fabulous")
			// ca.mapping["name"] = "fabulous"
		}
	}()

	time.Sleep(time.Second)

	value := ca.Lookup("name")
	fmt.Println(value)
}
