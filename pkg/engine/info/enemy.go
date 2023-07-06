package info

type Enemy struct {
	Level     int          `json:"level"`
	MaxStance float64      `json:"max_stance"`
	Weakness  WeaknessMap  `json:"weakness"`
	DebuffRES DebuffRESMap `json:"debuff_res"`
}
