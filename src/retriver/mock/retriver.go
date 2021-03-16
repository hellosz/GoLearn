package mock

import "fmt"

type Retriver interface {
	Get(url string) string
}

func Download(r Retriver) string {
	return r.Get("www.imooc.com")
}

type Container struct {
	Content string
}

func (c Container) Get(url string) string {
	return fmt.Sprintf("%s get %s", c.Content, url)
}

// 实现String接口
func (c Container) String() string {
	return fmt.Sprintf("Container {Content=%s}", c.Content)
}
