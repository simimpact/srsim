package mock

//go:generate mockgen -destination=mock_engine.go -package=mock github.com/simimpact/srsim/pkg/engine Engine
//go:generate mockgen -destination=mock_attribute.go -package=mock -mock_names Manager=MockAttribute github.com/simimpact/srsim/pkg/engine/attribute Manager
//go:generate mockgen -destination=mock_modifier.go -package=mock -mock_names Eval=MockModifier github.com/simimpact/srsim/pkg/engine/modifier Eval
//go:generate mockgen -destination=mock_shield.go -package=mock -mock_names Absorb=MockShield github.com/simimpact/srsim/pkg/engine/shield Absorb
