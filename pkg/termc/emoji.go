package termc

type emoji struct {
	GreenTick string
	RedCross  string
	Person    string
	Rocket    string
}

var Emoji *emoji

func init() {
	Emoji = &emoji{
		GreenTick: greenTick,
		RedCross:  redCross,
		Person:    person,
		Rocket:    rocket,
	}
}

const (
	greenTick = "✅"
	redCross  = "❌"
	person    = "🧑‍💼"
	rocket    = "🚀"
)
