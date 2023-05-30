package builder

import (
	"fmt"
	"testing"
)

func TestNewDirector(t *testing.T) {
	b1 := &Builder1{}
	director := NewDirector(b1)
	director.Construct()
	fmt.Println(b1.s)
}
