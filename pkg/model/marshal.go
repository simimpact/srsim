package model

import "google.golang.org/protobuf/encoding/protojson"

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
