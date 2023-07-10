package key

type LightCone string

// Destruction
const (
	TheMolesWelcomeYou     LightCone = "the_moles_welcome_you"
	NowheretoRun           LightCone = "nowhere_to_run"
	MutualDemise           LightCone = "mutual_demise"
	ASecretVow             LightCone = "a_secret_vow"
	ShatteredHome          LightCone = "shattered_home"
	SomethingIrreplaceable LightCone = "something_irreplaceable"
	UndertheBlueSky        LightCone = "under_the_blue_sky"
	CollapsingSky          LightCone = "collapsing_sky"
	OntheFallofanAeon      LightCone = "on_the_fall_of_an_aeon"
)

// Hunt
const (
	DartingArrow            LightCone = "darting_arrow"
	CruisingintheStellarSea LightCone = "cruising_in_the_stellar_sea"
	Arrows                  LightCone = "arrows"
	Adversarial             LightCone = "adversarial"
	OnlySilenceRemains      LightCone = "only_silence_remains"
	RiverFlowsinSpring      LightCone = "river_flows_in_spring"
	SubscribeforMore        LightCone = "subscribe_for_more"
	Swordplay               LightCone = "swordplay"
	ReturntoDarkness        LightCone = "return_to_darkness"
	SleepLiketheDead        LightCone = "sleep_like_the_dead"
	IntheNight              LightCone = "in_the_night"
)

// Nihility
const (
	Fermata               LightCone = "fermata"
	EyesofthePrey         LightCone = "eyes_of_the_prey"
	GoodNightandSleepWell LightCone = "good_night_and_sleep_well"
	IncessantRain         LightCone = "incessant_rain"
)

// Erudition
const (
	Passkey                   LightCone = "passkey"
	DataBank                  LightCone = "data_bank"
	BeforeDawn                LightCone = "before_dawn"
	TodayIsAnotherPeacefulDay LightCone = "today_is_another_peaceful_day"
	TheSeriousnessofBreakfast LightCone = "the_seriousness_of_breakfast"
)

// Harmony
const (
	Chorus      LightCone = "chorus"
	MeshingCogs LightCone = "meshing_cogs"
)

// Preservation
const (
	MomentOfVictory           LightCone = "moment_of_victory"
	Amber                     LightCone = "amber"
	DayOneofMyNewLife         LightCone = "day_one_of_my_new_life"
	Defense                   LightCone = "defense"
	TextureofMemories         LightCone = "texture_of_memories"
	ThisIsMe                  LightCone = "this_is_me"
	TrendoftheUniversalMarket LightCone = "trend_of_the_universal_market"
	Pioneering                LightCone = "pioneering"
	WeAreWildfire             LightCone = "we_are_wildfire"
	LandausChoice             LightCone = "landaus_choice"
)

// Abundance
const (
	FineFruit                LightCone = "fine_fruit"
	Multiplication           LightCone = "multiplication"
	Cornucopia               LightCone = "cornucopia"
	WarmthShortensColdNights LightCone = "warmth_shortens_cold_nights"
	PostOpConversation       LightCone = "post_op_conversation"
	SharedFeeling            LightCone = "shared_feeling"
	QuidProQuo               LightCone = "quid_pro_quo"
	PerfectTiming            LightCone = "perfect_timing"
)

func (l LightCone) String() string {
	return string(l)
}
