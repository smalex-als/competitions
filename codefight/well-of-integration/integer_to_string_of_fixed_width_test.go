package switchlights

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestIntegerToStringOfFixedWidth(t *testing.T) {
	res := integerToStringOfFixedWidth(1234, 2)
	if !reflect.DeepEqual("34", res) {
		_, file, line, _ := runtime.Caller(0)
		fmt.Printf("%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, "34", res)
		t.FailNow()
	}
}
