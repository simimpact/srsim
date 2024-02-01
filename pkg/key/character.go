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
	Himeko                Character = "himeko"
	Hook                  Character = "hook"
	Kafka                 Character = "kafka"
	Pela                  Character = "pela"
	Qingque               Character = "qingque"
	SilverWolf            Character = "silverwolf"
	DummyCharacter        Character = "dummy_character"
	Sampo                 Character = "sampo"
	Serval                Character = "serval"
	Sushang               Character = "sushang"
	Natasha               Character = "natasha"
	March7th              Character = "march7th"
	Seele                 Character = "seele"
)

func (c Character) String() string {
	return string(c)
}
