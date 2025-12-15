package terminal

import "fmt"

// limpar o terminal
func Clear() {
	fmt.Print("\033[H\033[2J")
}

func Pause(){
	var temp string
	fmt.Scan(&temp)
}