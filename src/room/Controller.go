package room

import "github.com/gopherjs/gopherjs/js"

type Controller struct {
	*js.Object

	IsPowerEnabled bool `js:"isPowerEnabled"`
	Level          int  `js:"level"`
	Progress       int  `js:"progress"`
	ProgressTotal  int  `js:"progressTotal"`
	// Reservation       Reservation `js:"reservation"`
	SafeMode          int `js:"safeMode"`
	SafeModeAvailable int `js:"safeModeAvailable"`
	SafeModeCooldown  int `js:"safeModeCooldown"`
	// Sign              Sign        `js:"sign"`
	TicksToDowngrade int `js:"ticksToDowngrade"`
	UpgradeBlocked   int `js:"upgradeBlocked"`
}
