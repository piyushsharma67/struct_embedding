package main

import "fmt"

type SpecialPosition struct{}


func (sp *SpecialPosition)TravelSpecialPosition(){
	fmt.Println("i have travelled to special position")
}

type Position struct{
	x float64
	y float64
}

func (p *Position)movePlayer(x,y float64){
	p.x+=x
	p.y+=y
}

func (p *Position)teleportPlayer(x,y float64){
	p.x=x
	p.y=y
}

type Player struct{
	Position
	SpecialPosition // embedded two structs inside the Player struct hence getting their attached properties and methods which are accessible inside the main
}


func NewPlayer()*Player{
	return &Player{
		Position:Position{},
	}
}

func main(){
	player:=NewPlayer()
	player.movePlayer(12,13)
	fmt.Println(player.Position)
	player.TravelSpecialPosition()
}