package key

type Character string

const (
	Arlan          Character = "arlan"
	Bailu          Character = "bailu"
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
)

func (c Character) String() string {
	return string(c)
}
