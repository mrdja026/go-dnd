package player

import "dnd.go/pkg/grid"

type Player struct {
	name string
	Grid grid.Grid
}

type IPlayer interface {
	MoveForward() error
	MoveBackward() error
}

func (p *Player) MoveForward() error {
	//find users index from tile
	//move the user to index+1 tile
	//cleanup tile
	return nil
}

func (p *Player) MoveBackward() error {
	return nil
}
