package key

type Character string

const (
	Arlan          Character = "arlan"
	Bronya         Character = "bronya"
	Clara          Character = "clara"
	DanHeng        Character = "danheng"
	Gepard         Character = "gepard"
	Pela           Character = "pela"
	Qingque        Character = "qingque"
	SilverWolf     Character = "silverwolf"
	DummyCharacter Character = "dummy_character"
	Sampo          Character = "sampo"
	Sushang        Character = "sushang"
	Natasha        Character = "natasha"
	Seele          Character = "seele"
)

func (c Character) String() string {
	return string(c)
}
