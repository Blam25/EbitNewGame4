// EbitNewGame4 project EbitNewGame4.go
package main

import (
	E "EbitNew4"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	*Position
	Image
	Player
	Speed
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

func (p *Position) getPosition() (int, int) {
	return p.X, p.Y
}

func (p *Position) setPosition(X, Y int) {
	p.X = X
	p.Y = Y
}

type Image struct {
	image *ebiten.Image
	op    *ebiten.DrawImageOptions
}

func (s Image) getImage() *ebiten.Image {
	return s.image
}

func (s Image) getOp() *ebiten.DrawImageOptions {
	return s.op
}

type image interface {
	getImage() *ebiten.Image
	getOp() *ebiten.DrawImageOptions
}

type position interface {
	getPosition() (int, int)
	setPosition(int, int)
}

type imagepos interface {
	image
	position
}

type player interface {
	setPlayer()
}

type Player struct{}

func (Player) setPlayer() {

}

type speed interface {
	getSpeed() int
	setSpeed(int)
}

type Speed struct {
	speed int
}

func (s Speed) getSpeed() int {
	return s.speed
}

func (s Speed) setSpeed(speed int) {
	s.speed = speed
}

type wasd interface {
	position
	player
	speed
}

var ImagePosSys []imagepos
var wasds []wasd

func init() {

	var err error
	image1, _, err := ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	ent := E.NewEnt[entity1](entity1{
		id:       E.NewId(),
		Position: &Position{X: 500, Y: 500},
		Image:    Image{image: image1, op: &ebiten.DrawImageOptions{}},
		Speed:    Speed{5},
	})
	wasds = append(wasds, ent)

	print(ent.Getid())

}

func printImage(entity E.Entity) {
	if entity, ok := entity.(imagepos); ok {
		println(entity.getImage())
	}
}
