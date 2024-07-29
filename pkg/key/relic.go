package key

type Relic string

// Cavern
const (
	MusketeerOfWildWheat           Relic = "musketeer_of_wild_wheat"
	HunterOfGlacialForest          Relic = "hunter_of_glacial_forest"
	PasserbyOfWanderingCloud       Relic = "passerby_of_wandering_cloud"
	KnightOfPurityPalace           Relic = "knight_of_purity_palace"
	GeniusOfBrilliantStars         Relic = "genius_of_brilliant_stars"
	ChampionOfStreetwiseBoxing     Relic = "champion_of_streetwise_boxing"
	GuardOfWutheringSnow           Relic = "guard_of_wuthering_snow"
	FiresmithOfLavaForging         Relic = "firesmith_of_lava-forging"
	BandOfSizzlingThunder          Relic = "band_of_sizzling_thunder"
	EagleOfTwilightLine            Relic = "eagle_of_twilight_line"
	ThiefOfShootingMeteor          Relic = "thief_of_shooting_meteor"
	WastelanderOfBanditryDesert    Relic = "wastelander_of_banditry_desert"
	LongevousDisciple              Relic = "longevous_disciple"
	MessengerTraversingHackerspace Relic = "messenger_traversing_hackerspace"
)

// Planar
const (
	BelobogOfTheArchitects  Relic = "belobog_of_the_architects"
	SpaceSealingStation     Relic = "space_sealing_station"
	InertSalsotto           Relic = "inert_salsotto"
	TaliaKingdomOfBanditry  Relic = "talia_kingdom_of_banditry"
	SprightlyVonwacq        Relic = "sprightly_vonwacq"
	PanGalactic             Relic = "pan_galactic"
	RutilantArena           Relic = "rutilant_arena"
	CelestialDifferentiator Relic = "celestial_differentiator"
	BrokenKeel              Relic = "broken_keel"
)

func (r Relic) String() string {
	return string(r)
}
