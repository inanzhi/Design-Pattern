package adapter_pattern

import "fmt"

type Windows struct{}

func (w *Windows) InsertInCirclePort() {
	fmt.Println("Insert circle port into windows machine")
}
