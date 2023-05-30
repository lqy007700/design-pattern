package simple_factory

import "testing"

func TestSimple(t *testing.T) {
	api := NewApi(2)
	say := api.Say()
	t.Log(say)
}
