package model

import (
	"google.golang.org/protobuf/encoding/protojson"
)

func marshalOptions() protojson.MarshalOptions {
	//nolint:exhaustruct // unused options ok left uninitialized
	return protojson.MarshalOptions{
		AllowPartial:    true,
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}
}

func unmarshalOptions() protojson.UnmarshalOptions {
	//nolint:exhaustruct // unused options ok left uninitialized
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
