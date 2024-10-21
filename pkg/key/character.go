package key

type Character string

const (
	Arlan                 Character = "arlan"
	Asta                  Character = "asta"
	Blade                 Character = "blade"
	Bronya                Character = "bronya"
	Clara                 Character = "clara"
	DanHeng               Character = "danheng"
	DanHengImbibitorLunae Character = "danhengimbibitorlunae"
	Gepard                Character = "gepard"
	Jingliu               Character = "jingliu"
	Guinaifen             Character = "guinaifen"
	Herta                 Character = "herta"
	Himeko                Character = "himeko"
	Hook                  Character = "hook"
	Luka                  Character = "luka"
	Kafka                 Character = "kafka"
	Pela                  Character = "pela"
	Qingque               Character = "qingque"
	SilverWolf            Character = "silverwolf"
	TrailblazerImaginary  Character = "trailblazerimaginary"
	DummyCharacter        Character = "dummy_character"
	Sampo                 Character = "sampo"
	Serval                Character = "serval"
	Sushang               Character = "sushang"
	Natasha               Character = "natasha"
	March7th              Character = "march7th"
	Seele                 Character = "seele"
	Huohuo                Character = "huohuo"
)

func (c Character) String() string {
	return string(c)
}
