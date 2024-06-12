package spanish

type speaker interface {
	SayHola()
}
type Spanish struct {
	spanish speaker
}

func NewSpanish(s speaker) *Spanish {
	return &Spanish{s}
}

func (s *Spanish) Greet() {
	s.spanish.SayHola()
}
