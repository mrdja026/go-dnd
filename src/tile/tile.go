package tile

import "errors"

type Tile struct {
	Size                   int
	Display_repsesentation [][]string
	Condition              string
	Occupant               string
	Id                     string
}

type ITile interface {
	Display() string
	SetOccupant(occupant string)
	GetTile() Tile
}

type TileBuilder interface {
	WithSize(size int) TileBuilder
	WithDisplay(display string) TileBuilder
	WithCondition(condition string) TileBuilder
	Build() Tile
}

type tileBuilder struct {
	field Tile
}

func (b *tileBuilder) WithCondition(condition string) TileBuilder {
	b.field.Condition = condition
	return b
}

func NewTileBuilder() TileBuilder {
	return &tileBuilder{}
}

func (b *tileBuilder) WithSize(size int) TileBuilder {
	b.field.Size = size
	return b
}

func (b *tileBuilder) WithDisplay(display string) TileBuilder {
	display_string := make([][]string, b.field.Size)
	for i := 0; i < len(display_string); i++ {
		display_string[i] = make([]string, b.field.Size)
		for j := 0; j < len(display_string[i]); j++ {
			display_string[i][j] = display
		}

	}

	b.field.Display_repsesentation = display_string
	return b
}

func (b *tileBuilder) Build() Tile {
	return b.field
}

func (f Tile) Display() string {
	result := ``
	for i := 0; i < len(f.Display_repsesentation); i++ {
		for j := 0; j < len(f.Display_repsesentation[i]); j++ {
			result += f.Display_repsesentation[i][j]
		}
		if i != len(f.Display_repsesentation)-1 {
			result += "\n"
		}
	}
	return result
}

func (t *Tile) SetOccupant(occupant string) error {
	if t.Condition == "occupied" {
		return errors.New("Tile is already occupied with" + t.Occupant)
	}
	if occupant == "" {
		return errors.New("Occupant is invalid")
	}
	t.Occupant = occupant
	t.Display_repsesentation[1][1] = "O"
	t.Condition = "occupied"
	return nil
}

func (t Tile) GetTile() Tile {
	return t
}

// TODO: Implement remove Occupant
