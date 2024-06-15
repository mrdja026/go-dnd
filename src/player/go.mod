module example.com/player

go 1.22.3

replace example.com/grid => ../grid

replace example.com/tile => ../tile

require example.com/grid v0.0.0-00010101000000-000000000000

require example.com/tile v0.0.0-00010101000000-000000000000 // indirect
