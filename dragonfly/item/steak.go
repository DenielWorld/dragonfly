package item

import (
	"github.com/df-mc/dragonfly/dragonfly/world"
	"time"
)

// Steak is a food item that can be obtained from cooking raw beef and can be eaten by the player.
type Steak struct{}

// AlwaysConsumable ...
func (s Steak) AlwaysConsumable() bool {
	return false
}

// ConsumeDuration ...
func (s Steak) ConsumeDuration() time.Duration {
	return DefaultConsumeDuration
}

// Consume ...
func (s Steak) Consume(_ *world.World, c Consumer) Stack {
	c.Saturate(8, 12.8)
	return Stack{}
}

// EncodeItem ...
func (s Steak) EncodeItem() (id int32, meta int16) {
	return 364, 0
}
