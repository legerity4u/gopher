package keyboard

import "testing"

func TestGetFloat(t *testing.T) {
	want := 0.0
	if got, _ := GetFloat(); got != want {
		t.Errorf("GetFloat() = %f, want %f", got, want)
	}
}
