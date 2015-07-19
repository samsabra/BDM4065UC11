package bdm

// Commands
const (
	GetVersion byte = 0xA2

	GetPowerState byte = 0x19
	SetPowerState byte = 0x18

	GetUserInputControl byte = 0x1D
	SetUserInputControl byte = 0x1C
)

// Arguments for get version command.
const (
	SICPVersion     byte = 0x00
	PlatformVersion byte = 0x01
)

// Arguments and result for power state.
const (
	PowerStateOff byte = 0x01
	PowerStateOn  byte = 0x02
)

// Arguments and result for user input control
const (
	UserInputMaskLocked   byte = 0x00 // 0b00000000
	UserInputMaskUnlocked byte = 0xff // 0b11111111

	UserInputMaskLocalKeyboard byte = 0x02 // 0b00000010
	UserInputMaskRemoteControl byte = 0x01 // 0b00000001
)
