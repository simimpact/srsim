package lightcone

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/tests/testcfg/testchar"
	"github.com/simimpact/srsim/tests/testcfg/testcone"
	"github.com/simimpact/srsim/tests/teststub"
	"github.com/stretchr/testify/suite"
)

type PTBTest struct {
	teststub.Stub
}

func TestPTBTest(t *testing.T) {
	suite.Run(t, new(PTBTest))
}

// Testing that PTB indeed does add the correct ATK amount
func (t *PTBTest) Test_AtkAdd() {
	bronyaModel := testchar.Bronya()
	bronyaModel.LightCone = testcone.PoisedToBloom()
	t.Characters.ResetCharacters()
	bronya := t.Characters.AddCharacter(bronyaModel)
	t.StartSimulation()
	// She should have no other atk buffs from anywhere
	bronya.Equal(prop.ATKPercent, 0.16)
}

// Testing that PTB adds crit damage to only characters who share a path with another character in the party
func (t *PTBTest) Test_CritDMG() {
	bronyaModel := testchar.Bronya()
	bronyaModel.LightCone = testcone.PoisedToBloom()
	t.Characters.ResetCharacters()
	bronya := t.Characters.AddCharacter(bronyaModel)
	dan1 := t.Characters.AddCharacter(testchar.DanHung())
	dan2 := t.Characters.AddCharacter(testchar.DanHung())
	t.StartSimulation()
	// 0.50 from base crit damage + 0.16 from poised buff
	dan1.Equal(prop.CritDMG, 0.66)
	dan2.Equal(prop.CritDMG, 0.66)
	// Just the 0.50 base crit damage
	bronya.Equal(prop.CritDMG, 0.50)
}

// Testing that PTB adds crit damage to only characters who share a path with another character in the party
func (t *PTBTest) Test_No_Stacking() {
	// I'm a bit concerned about using the same character twice, but hopefully all should be good?
	bronyaModel1 := testchar.Bronya()
	bronyaModel1.LightCone = testcone.PoisedToBloom()
	bronyaModel2 := testchar.Bronya()
	bronyaModel2.LightCone = testcone.PoisedToBloom()
	t.Characters.ResetCharacters()
	bronya1 := t.Characters.AddCharacter(bronyaModel1)
	bronya2 := t.Characters.AddCharacter(bronyaModel2)

	t.StartSimulation()
	// 0.50 from base crit damage + 0.16 from poised buff
	bronya1.Equal(prop.CritDMG, 0.66)
	bronya2.Equal(prop.CritDMG, 0.66)
}
