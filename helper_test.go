package t

import "testing"

func TestTypeContext_InArray(t *testing.T) {
	var a = []interface{}{1, 2, "a"}
	t.Log(New("a").InArray(a))
}
