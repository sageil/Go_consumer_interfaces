package english

type speaker interface {
	SayHello()
}
type English struct {
	speaker speaker
}

func NewEnglish(s speaker) *English {
	return &English{
		speaker: s,
	}
}

func (e *English) Greet() {
	e.speaker.SayHello()
}
