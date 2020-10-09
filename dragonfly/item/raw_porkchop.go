package item

import (
	"github.com/df-mc/dragonfly/dragonfly/world"
	"time"
)

// RawPorkchop is the uncooked version of a cooked porkchop which can be obtained from adult pigs. It can
// be eaten by the player.
type RawPorkchop struct{}

// AlwaysConsumable ...
func (r RawPorkchop) AlwaysConsumable() bool {
	return false
}

// ConsumeDuration ...
func (r RawPorkchop) ConsumeDuration() time.Duration {
	return DefaultConsumeDuration
}

// Consume ...
func (r RawPorkchop) Consume(_ *world.World, c Consumer) Stack {
	c.Saturate(3, 1.8)
	return Stack{}
}

// EncodeItem ...
func (r RawPorkchop) EncodeItem() (id int32, meta int16) {
	return 319, 0
}
