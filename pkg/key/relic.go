package key

type Relic string

// Cavern
const (
	MusketeerOfWildWheat       Relic = "musketeer_of_wild_wheat"
	HunterOfGlacialForest      Relic = "hunter_of_glacial_forest"
	PasserbyOfWanderingCloud   Relic = "passerby_of_wandering_cloud"
	KnightOfPurityPalace       Relic = "knight_of_purity_palace"
	GeniusOfBrilliantStars     Relic = "genius_of_brilliant_stars"
	ChampionOfStreetwiseBoxing Relic = "champion_of_streetwise_boxing"
)

// Planar
const (
	BelobogOfTheArchitects Relic = "belobog_of_the_architects"
	SpaceSealingStation    Relic = "space_sealing_station"
	InertSalsotto          Relic = "inert_salsotto"
	TaliaKingdomOfBanditry Relic = "talia_kingdom_of_banditry"
	SprightlyVonwacq       Relic = "sprightly_vonwacq"
	FleetOfTheAgeless      Relic = "fleet_of_the_ageless"
	PanGalactic            Relic = "pan_galactic"
	RutilantArena          Relic = "rutilant_arena"
)

func (r Relic) String() string {
	return string(r)
}
