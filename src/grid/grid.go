package grid

import (
	"errors"

	"example.com/tile"
)

type JsonGrid struct {
	name  string
	users []string
}

type Grid struct {
	Tiles []tile.Tile
	data  JsonGrid
}

type IGrid interface {
	AddUser(user string, tileIndex int)
}

func CreateGrid(name string) Grid {
	tiles := []tile.Tile{
		tile.NewTileBuilder().WithSize(3).WithDisplay("x").Build(),
		tile.NewTileBuilder().WithSize(3).WithDisplay("x").Build(),
		tile.NewTileBuilder().WithSize(3).WithDisplay("x").Build(),
	}
	return Grid{
		Tiles: tiles,
		data: JsonGrid{
			name:  name,
			users: []string{},
		},
	}
}

func (g *Grid) AddUser(user string, tileIndex int) error {
	if tileIndex < 0 || tileIndex >= len(g.Tiles) {
		return errors.New("Invalid tile index")
	}
	if user == "" {
		return errors.New("Invalid user")
	}
	g.data.users = append(g.data.users, user)
	if err := g.Tiles[tileIndex].SetOccupant(user); err != nil {
		return err
	}
	return nil
}
