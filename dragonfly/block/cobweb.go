package block

import (
	"github.com/df-mc/dragonfly/dragonfly/entity/physics"
	"github.com/df-mc/dragonfly/dragonfly/item"
	"github.com/df-mc/dragonfly/dragonfly/item/tool"
	"github.com/df-mc/dragonfly/dragonfly/world"
)

// Cobweb is a block which slows down any entity that collides with it. It can be efficiently broken by a sword or shears.
type Cobweb struct {
}

// LightDiffusionLevel ...
func (Cobweb) LightDiffusionLevel() uint8 {
	return 1
}

// EncodeItem ...
func (Cobweb) EncodeItem() (id int32, meta int16) {
	return 30, 0
}

// EncodeBlock ...
func (Cobweb) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:web", nil
}

// AABB ...
func (Cobweb) AABB(world.BlockPos, *world.World) []physics.AABB {
	return nil
	// Cobweb does not have a real AABB, but only slows down the entities that collide with it.
	// TODO: Item entity appears to be the only non-living entity which can collide and stick to a Cobweb block. In the future, item entities should stick to cobwebs.
}

// BreakInfo ...
func (c Cobweb) BreakInfo() BreakInfo {
	return BreakInfo{
		//TODO: Fix the client-side animation of cobweb breaking by various tools. (By default, the break time is 20s, but with shears or a sword it is only 0.4s)
		Hardness:    4.0,
		Harvestable: alwaysHarvestable,
		// Cobweb can be efficiently broken by sword and shears.
		Effective: func(t tool.Tool) bool {
			return t.ToolType() == tool.TypeShears || t.ToolType() == tool.TypeSword
		},
		Drops: func(t tool.Tool) []item.Stack {
			// TODO: When enchantments are implemented, if the ToolType is TypeSword and the tool has silk touch, then the drop should be a web block as well.
			if t.ToolType() == tool.TypeShears {
				return []item.Stack{item.NewStack(c, 1)}
			}
			// When an unsuitable tool is used to break cobweb, a string is dropped.
			return []item.Stack{item.NewStack(item.String{}, 1)}
		},
	}
}
