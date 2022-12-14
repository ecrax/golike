package main

type GameData struct {
	MapWidth   int
	MapHeight  int
	TileWidth  int
	TileHeight int
}

func NewGameData() GameData {
	g := GameData{
		MapWidth:   50,
		MapHeight:  30,
		TileWidth:  16,
		TileHeight: 16,
	}
	return g
}
