package adapter

import "fmt"

type Target interface {
	Request()
}

type Adapter interface {
	ShowMe()
}

type adapterImpl struct {
}

func NewAdapterImpl() Adapter {
	return &adapterImpl{}
}

func (a adapterImpl) ShowMe() {
	fmt.Println("showme")
}

type targetImpl struct {
	Adapter
}

func NewTargetImpl(adapter Adapter) Target {
	return &targetImpl{Adapter: adapter}
}

func (t *targetImpl) Request() {
	t.ShowMe()
}
