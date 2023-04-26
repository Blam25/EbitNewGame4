package main

import (
	E "EbitNew4"
	//"log"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {

	for _, s := range E.Entitys {

		Move(s)
		/*
			if s, ok := s.(imagepos); ok {
				println(s.getImage())
			}

			if s, ok := s.(position); ok {
				print(s)
			}
			if s, ok := s.(position); ok {
				if z, ok := s.(image); ok {
					println(z.getImage())
					println("hej")
					//println(s.getY())
				}
			}*/
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

	for _, s := range E.Entitys {
		Render(screen, s)
	}

}

func Render(screen *ebiten.Image, s any) {
	if s, ok := s.(imagepos); ok {
		x, y := s.getPosition()
		println(x)
		println(y)
		s.getOp().GeoM.Reset()
		s.getOp().GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(s.getImage(), s.getOp())
		//println("hej")
	}
}

func Move(s any) {
	if s, ok := s.(wasd); ok {
		//print("lol")
		x, y := s.getPosition()
		if ebiten.IsKeyPressed(ebiten.KeyW) {
			y = y - s.getSpeed()
			println("hmmmmmmmmmmmmm")
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			y += s.getSpeed()

		}
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			x += s.getSpeed()

		}
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			x -= s.getSpeed()
		}
		s.setPosition(x, y)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
