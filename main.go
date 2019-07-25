package main

import (
	"github.com/wmarbut/go-epdfuse"
	"github.com/stianeikeland/go-rpio"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	fuse := epdfuse.NewCustomEpdFuse("/dev/epd",264,176)
	err := rpio.Open()
	defer rpio.Close()
	Check(err)
	btn1 := rpio.Pin(21)
	btn2 := rpio.Pin(20)
	btn3 := rpio.Pin(26)
	btn4 := rpio.Pin(16)
	btn1.Input()
	btn2.Input()
	btn3.Input()
	btn4.Input()

	for {	
		resultat := "Button push test\n \n"
		if btn1.Read() == 0{
			resultat += "Push btn1\n"
		}
	
		if btn2.Read() == 0{
			resultat += "Push btn2\n"
		}
	
		if btn3.Read() == 0{
			resultat += "Push btn3\n"
		}
	
		if btn4.Read() == 0{
			resultat += "Push btn4\n"
		}
		err := fuse.WriteText(resultat)
		Check(err)
	}
}

