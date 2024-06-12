package main

import (
	"github.com/sageil/consumer-interfaces/internals/english"
	multi "github.com/sageil/consumer-interfaces/internals/multinational"
	"github.com/sageil/consumer-interfaces/internals/spanish"
)

func main() {
	// construct the concrete struct
	multi := &multi.Multinational{}
	// construct the english consumer using the multi as source
	english := english.NewEnglish(multi)
	spanish := spanish.NewSpanish(multi)

	english.Greet()
	spanish.Greet()
}
