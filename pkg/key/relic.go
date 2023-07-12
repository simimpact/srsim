package key

type Relic string

// Cavern
const (
	MusketeerOfWildWheat     Relic = "musketeer_of_wild_wheat"
	HunterOfGlacialForest    Relic = "hunter_of_glacial_forest"
	PasserbyOfWanderingCloud Relic = "passerby_of_wandering_cloud"
	KnightOfPurityPalace     Relic = "knight_of_purity_palace"
)

// Planar
const (
	BelobogOfTheArchitects Relic = "belobog-of-the-architects"
	SpaceSealingStation    Relic = "space_sealing_station"
	InertSalsotto          Relic = "inert_salsotto"
	TaliaKingdomOfBanditry Relic = "talia_kingdom_of_banditry"
	SprightlyVonwacq       Relic = "sprightly_vonwacq"
	PanGalactic            Relic = "pan_galactic"
)

func (r Relic) String() string {
	return string(r)
}
