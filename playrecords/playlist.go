package main

import (
	"fmt"
	"github.com/romzasandysman/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

type Toggleable interface {
	toggle()
}

type Switch string

func (s *Switch) toggle()  {
	if *s == "on"{
		*s = "off"
	}else{
		*s = "on"
	}
	fmt.Println(*s)
}

func TryOut(player Player){
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok{
		recorder.Record()
	}
}

func playList(device Player, songs []string)  {
	for _, song := range songs{
		device.Play(song)
	}
	device.Stop()
}
func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)

	recorder := gadget.TapeRecorder{}
	playList(recorder, mixtape)
	TryOut(recorder)
	TryOut(player)

	s := Switch("off")
	var t Toggleable = &s
	t.toggle()
	t.toggle()
}
