package model

// Omifuda はおみくじの札です
type Omifuda int

const (
	Daikichi Omifuda = iota
	Chukichi
	Kichi
	Kyo
)

func (o Omifuda) String() string {
	switch o {
	case Daikichi:
		return "大吉"
	case Chukichi:
		return "中吉"
	case Kichi:
		return "吉"
	case Kyo:
		return "凶"
	default:
		return "はずれ"
	}
}
