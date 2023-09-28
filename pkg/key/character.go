package key

type Character string

const (
	Arlan                 Character = "arlan"
	Blade                 Character = "blade"
	Bronya                Character = "bronya"
	Clara                 Character = "clara"
	DanHeng               Character = "danheng"
	DanHengImbibitorLunae Character = "danhengimbibitorlunae"
	Gepard                Character = "gepard"
	Pela                  Character = "pela"
	Qingque               Character = "qingque"
	SilverWolf            Character = "silverwolf"
	DummyCharacter        Character = "dummy_character"
	Sampo                 Character = "sampo"
	Sushang               Character = "sushang"
	Natasha               Character = "natasha"
	March7th              Character = "march7th"
	Yanqing               Character = "yanqing"
)

func (c Character) String() string {
	return string(c)
}
