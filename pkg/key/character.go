package key

type Character string

const (
	Arlan                 Character = "arlan"
	Blade                 Character = "blade"
	Bronya                Character = "bronya"
	Clara                 Character = "clara"
	DanHeng               Character = "danheng"
	DanHengImbibitorLunae Character = "danhengimbibitorlunae"
	DummyCharacter        Character = "dummy_character"
	Gepard                Character = "gepard"
	March7th              Character = "march7th"
	Natasha               Character = "natasha"
	Pela                  Character = "pela"
	Qingque               Character = "qingque"
	Sampo                 Character = "sampo"
	Seele                 Character = "seele"
	SilverWolf            Character = "silverwolf"
	Sushang               Character = "sushang"
	Welt                  Character = "welt"
)

func (c Character) String() string {
	return string(c)
}
