package contextimpl

import (
	"fmt"
	"testing"
)

func TestBackgroundNotTODO(t *testing.T) {
	todo := fmt.Sprint(TODO())
	bg := fmt.Sprint(Background())
	if todo == bg {
		t.Errorf("TODO and Background are equal: %q vs %q", todo, bg)
	}
}

func TestWithCancel(t *testing.T) {

}
