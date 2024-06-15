package grid

import (
	"strings"
	"testing"
)

func TestCreateGrid(t *testing.T) {
	grid := CreateGrid("test")
	if grid.data.name != "test" {
		t.Errorf("Expected %q, but got %q", "test", grid.data.name)
	}
	if len(grid.Tiles) != 3 {
		t.Errorf("Expected %v, but got %v", 3, len(grid.Tiles))
	}
}

func TestAddUser(t *testing.T) {
	grid := CreateGrid("test")
	err := grid.AddUser("username1", 0)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
	got := grid.data.users[0]
	if got != "username1" {
		t.Errorf("Expected %q, but got %q", "username1", got)
	}
	gotTile := grid.Tiles[0].GetTile()
	if gotTile.Occupant != "username1" {
		t.Errorf("Expected %q, but got %q", "username1", gotTile.Occupant)
	}
	gotDisplay := gotTile.Display()
	if strings.Contains(gotDisplay, "O") == false {
		t.Errorf("Expected %q, but got %q", "O", gotDisplay)
	}

	err = grid.AddUser("username2", 0)
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}
}
