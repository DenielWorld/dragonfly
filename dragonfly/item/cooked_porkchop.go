package item

import (
	"github.com/df-mc/dragonfly/dragonfly/world"
	"time"
)

// CookedPorkchop is a food item that can be obtained from cooking raw porkchop and can be eaten by the player.
type CookedPorkchop struct{}

// AlwaysConsumable ...
func (c CookedPorkchop) AlwaysConsumable() bool {
	return false
}

// ConsumeDuration ...
func (c CookedPorkchop) ConsumeDuration() time.Duration {
	return DefaultConsumeDuration
}

// Consume ...
func (c CookedPorkchop) Consume(_ *world.World, co Consumer) Stack {
	co.Saturate(8, 12.8)
	return Stack{}
}

// EncodeItem ...
func (c CookedPorkchop) EncodeItem() (id int32, meta int16) {
	return 320, 0
}
