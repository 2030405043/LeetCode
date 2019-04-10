package simple

import (
	"fmt"
	"testing"
)

func TestToLowerCase(t *testing.T) {
	lower := toLowerCase("AsDFGhjKLOp")
	fmt.Println(lower)
}
