package simple_factory

type Api interface {
	Say() string
}

func NewApi(i int) Api {
	if i == 1 {
		return &hi{}
	} else {
		return &hello{}
	}
}

type hi struct {
}

func (h *hi) Say() string {
	return "Hi"
}

type hello struct {
}

func (h *hello) Say() string {
	return "Hello"
}
