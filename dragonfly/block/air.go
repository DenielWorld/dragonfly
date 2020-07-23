package block

import (
	"github.com/df-mc/dragonfly/dragonfly/entity/physics"
	"github.com/df-mc/dragonfly/dragonfly/world"
)

// Air is the block present in otherwise empty space.
type Air struct{ noNBT }

// CanDisplace ...
func (Air) CanDisplace(world.Liquid) bool {
	return true
}

// HasLiquidDrops ...
func (Air) HasLiquidDrops() bool {
	return false
}

// LightDiffusionLevel ...
func (Air) LightDiffusionLevel() uint8 {
	return 0
}

// EncodeItem ...
func (Air) EncodeItem() (id int32, meta int16) {
	return 0, 0
}

// EncodeBlock ...
func (Air) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:air", nil
}

// Hash ...
func (Air) Hash() uint64 {
	return hashAir
}

// AABB returns an empty Axis Aligned Bounding Box (as nothing can collide with air).
func (Air) AABB(world.BlockPos, *world.World) []physics.AABB {
	return nil
}

// ReplaceableBy always returns true.
func (Air) ReplaceableBy(world.Block) bool {
	return true
}
