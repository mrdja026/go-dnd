package tile

import (
	"testing"
)

func TestFieldBuilder(t *testing.T) {
	fieldBuilder := NewTileBuilder()
	if fieldBuilder == nil {
		t.Error("Expected field to not be nil")
	}
}

func TestFullBuild(t *testing.T) {
	fieldBuilder := NewTileBuilder()
	if fieldBuilder == nil {
		t.Error("Expected field to not be nil")
	}

	field := fieldBuilder.WithSize(3).WithDisplay("x").Build()
	expected := `xxx
xxx
xxx`
	got := field.Display()
	if expected != got {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

}
