package state_pattern

type LightState interface {
	ChangeToRedLight()
	ChangeToYellowLight()
	ChangeToGreenLight()
}
