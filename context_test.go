package term

import (
	"testing"
)

func TestWith(t *testing.T) {
	var ctx Context
	ctx.Set(Red, BlueBackground)
	newCtx := ctx.With(Green, Underline)
	if newCtx == &ctx {
		t.Fatalf("With() should return a new Context")
	}
	t.Skip("TODO")
	// if !reflect.DeepEqual([]Attribute{Green, BlueBackground, Underline}, newCtx.attrs) {
	// 	t.Fatalf("With() should have updated attributes, got %v", newCtx.attrs)
	// }
}
