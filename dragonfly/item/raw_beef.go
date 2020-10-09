package item

import (
	"github.com/df-mc/dragonfly/dragonfly/world"
	"time"
)

// RawBeef is the uncooked version of a steak which can be obtained from adult cows or mooshrooms.
// It can be eaten by the player.
type RawBeef struct{}

// AlwaysConsumable ...
func (r RawBeef) AlwaysConsumable() bool {
	return false
}

// ConsumeDuration ...
func (r RawBeef) ConsumeDuration() time.Duration {
	return DefaultConsumeDuration
}

// Consume ...
func (r RawBeef) Consume(_ *world.World, c Consumer) Stack {
	c.Saturate(3, 1.8)
	return Stack{}
}

// EncodeItem ...
func (r RawBeef) EncodeItem() (id int32, meta int16) {
	return 363, 0
}
