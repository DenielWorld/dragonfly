package item

import (
	"github.com/df-mc/dragonfly/dragonfly/item/tool"
	"github.com/df-mc/dragonfly/dragonfly/world"
)

// Shears is a tool used to obtain some blocks which would otherwise only be obtainable through silk touch enchanted tools.
type Shears struct {
}

// EncodeItem ...
func (Shears) EncodeItem() (id int32, meta int16) {
	return 359, 0
}

// ToolType...
func (Shears) ToolType() tool.Type {
	return tool.TypeShears
}

// HarvestLevel ...
func (Shears) HarvestLevel() int {
	return 1
}

// BaseMiningEfficiency ...
func (Shears) BaseMiningEfficiency(b world.Block) float64 {
	return 15
}

// DurabilityInfo ...
func (Shears) DurabilityInfo() DurabilityInfo {
	return DurabilityInfo{
		MaxDurability:    238,
		BrokenItem:       simpleItem(Stack{}),
		AttackDurability: 0,
		BreakDurability:  1,
	}
}

// MaxCount always returns 1.
func (Shears) MaxCount() int {
	return 1
}