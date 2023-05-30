package adapter

import "testing"

func TestAdapter(t *testing.T) {
	impl := NewAdapterImpl()
	newTargetImpl := NewTargetImpl(impl)
	newTargetImpl.Request()

}
