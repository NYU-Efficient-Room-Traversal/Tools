package main

import "github.com/NYU-Efficient-Room-Traversal/Roomba-Control"


func main() {

	roombaControl.ModeStart()	
	spr := roombaControl.InitializeStepper(18,0.005,23,24,17,25)
	spr.Spin()
}


