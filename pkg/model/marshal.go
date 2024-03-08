package model

import (
	"google.golang.org/protobuf/encoding/protojson"
)

func marshalOptions() protojson.MarshalOptions {
	return protojson.MarshalOptions{
		AllowPartial:    true,
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   true,
	}
}

func unmarshalOptions() protojson.UnmarshalOptions {
	return protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}
}

func (c *SimConfig) UnmarshalJSON(b []byte) error {
	return unmarshalOptions().Unmarshal(b, c)
}

func (c *SimConfig) MarshalJSON() ([]byte, error) {
	return marshalOptions().Marshal(c)
}

func (r *SimResult) MarshalJSON() ([]byte, error) {
	return marshalOptions().Marshal(r)
}

func (p Path) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (d DamageType) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (t AttackType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (df DamageFormula) MarshalText() ([]byte, error) {
	return []byte(df.String()), nil
}

func (hf HealFormula) MarshalText() ([]byte, error) {
	return []byte(hf.String()), nil
}

func (sf ShieldFormula) MarshalText() ([]byte, error) {
	return []byte(sf.String()), nil
}

func (st StatusType) MarshalText() ([]byte, error) {
	return []byte(st.String()), nil
}

func (t BehaviorFlag) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t TerminationReason) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}
