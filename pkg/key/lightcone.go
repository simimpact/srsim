package key

type LightCone string

// Destruction
const (
	TheMolesWelcomeYou LightCone = "the_moles_welcome_you"
	NowheretoRun       LightCone = "nowhere_to_run"
	MutualDemise       LightCone = "mutual_demise"
	ASecretVow         LightCone = "a_secret_vow"
	ShatteredHome      LightCone = "shattered_home"
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
)

// Erudition
const (
	Passkey                   LightCone = "passkey"
	DataBank                  LightCone = "data_bank"
	BeforeDawn                LightCone = "before_dawn"
	TodayIsAnotherPeacefulDay LightCone = "today_is_another_peaceful_day"
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
	TrendoftheUniversalMarket LightCone = "trend_of_the_universal_market"
	WeAreWildfire             LightCone = "we_are_wildfire"
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
)

func (l LightCone) String() string {
	return string(l)
}
