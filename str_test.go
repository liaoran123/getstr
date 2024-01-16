package getstr

import (
	"fmt"
	"testing"
)

func TestSub(t *testing.T) {
	var rstr string
	rstr = LStr("ddddddddadgrfre", "d")
	fmt.Println(rstr)
	rstr = RStr("ddddddddadgrfre", "d")
	fmt.Println(rstr)
}
