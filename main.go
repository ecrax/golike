package main

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
)

type Game struct {
	Map       GameMap
	World     *ecs.Manager
	WorldTags map[string]ecs.Tag
}

func (g *Game) Update() error {
	TryMovePlayer(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Map.Dungeons[0].Levels[0]
	level.DrawLevel(screen)
	ProcessRenderables(g, level, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	gd := NewGameData()
	return gd.TileWidth * gd.MapWidth, gd.TileHeight * gd.MapHeight
}

func NewGame() *Game {
	g := &Game{}
	world, tags := InitWorld()
	g.Map = NewGameMap()
	g.WorldTags = tags
	g.World = world
	return g
}

func main() {
	g := NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetWindowTitle("Super fun game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
