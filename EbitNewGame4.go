// EbitNewGame4 project EbitNewGame4.go
package main

import (
	E "EbitNew4"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	//ebiten.SetVsyncEnabled(false)
	//ebiten.SetTPS(ebiten.SyncWithFPS)
	//ebiten.SetFPSMode(ebiten.)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

type entity1 struct {
	id int
	Position
	Image
}

type entity2 struct {
	id int
}

func (s entity1) Getid() int {
	return s.id
}

func (s entity2) Getid() int {
	return s.id
}

type Position struct {
	X int
	Y int
}

func (p Position) getX() int {
	return p.X
}

func (p Position) getY() int {
	return p.Y
}

type Image struct {
	image string
}

func (s Image) getImage() string {
	return s.image
}

type image interface {
	getImage() string
}

type position interface {
	getX() int
	getY() int
}

type imagepos interface {
	image
	position
}

var ImagePosSys []imagepos

func init() {

	ent := E.NewEnt[entity1](entity1{id: E.NewId()})
	ImagePosSys = append(ImagePosSys, ent)
	print(ent.getX())

}

func printImage(entity E.Entity) {
	if entity, ok := entity.(imagepos); ok {
		println(entity.getImage())
	}
}

type Game struct{}

func (g *Game) Update() error {

	for _, s := range E.Entitys {

		if s, ok := s.(imagepos); ok {
			println(s.getImage())
		}

		if s, ok := s.(position); ok {
			print(s.getX())
		}
		if s, ok := s.(position); ok {
			if z, ok := s.(image); ok {
				println(z.getImage())
				println("hej")
				println(s.getY())
			}
		}
		/*
			if s, ok := s.(image); ok {
				printHej := true
			}

			if s, ok := s.(position); ok {
				printHej := true
			}
		*/
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
