package bdm

// Commands
const (
	GetVersion byte = 0xA2

	GetPowerState byte = 0x19
	SetPowerState byte = 0x18

	GetUserInputControl      byte = 0x1D
	SetUserInputControl      byte = 0x1C
	GetUserInputControlState byte = 0x1B
	SetUserInputControlState byte = 0x1A

	GetPowerAtColdStart byte = 0xA4
	SetPowerAtColdStart byte = 0xA3

	// TODO: argument and result
	SetInputSource byte = 0xAC

	GetCurrentSource byte = 0xAD

	GetAutoSignalDetecting byte = 0xAF
	SetAutoSignalDetecting byte = 0xAE

	GetVideoParameters  byte = 0x33
	SetVideoParameters  byte = 0x32
	GetColorTemperature byte = 0x35
	SetColorTemperature byte = 0x34
	GetColorParameters  byte = 0x37
	SetColorParameters  byte = 0x36

	GetPictureFormat      byte = 0x3B
	SetPictureFormat      byte = 0x3A
	GetVGAVideoParameters byte = 0x39
	SetVGAVideoParameters byte = 0x38

	GetPictureInPicture byte = 0x3D
	SetPictureInPicture byte = 0x3C

	GetPIPSource byte = 0x85
	SetPIPSource byte = 0x84

	GetVolume byte = 0x45
	SetVolume byte = 0x44

	SetVolumeLimits byte = 0xB8

	GetAudioParameters byte = 0x43
	SetAudioParameters byte = 0x42

	GetMiscInfo byte = 0x0F

	GetSmartPower byte = 0xDE
	SetSmartPower byte = 0xDD

	SetVideoAlignment byte = 0x70

	GetTemperature byte = 0x2F

	GetSerialCode byte = 0x15

	GetTiling byte = 0x23
	SetTiling byte = 0x22

	GetlightSensor byte = 0x25
	SetlightSensor byte = 0x24

	GetOSDRotating byte = 0x27
	SetOSDRotating byte = 0x26

	GetInformationOSDFeature byte = 0x2D
	SetInformationOSDFeature byte = 0x2C

	GetMEMCEffect byte = 0x29
	SetMEMCEffect byte = 0x28

	GetTouchFeature byte = 0x1F
	SetTouchFeature byte = 0x1E

	GetNoiseReductionFeature byte = 0x2B
	SetNoiseReductionFeature byte = 0x2A

	GetScanModeFeature byte = 0x51
	SetScanModeFeature byte = 0x50

	GetScanConversionFeature byte = 0x53
	SetScanConversionFeature byte = 0x52

	GetSwitchOnDelayFeature byte = 0x55
	SetSwitchOnDelayFeature byte = 0x54

	SetFactoryReset byte = 0x56
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

	UserInputLockAll          byte = 0x01
	UserInputLockAllButVolume byte = 0x02
	UserInputLockAllButPower  byte = 0x03
)

// Argumetns and result for power state at cold start.
const (
	PowerAtColdStartPowerOff   byte = 0x00
	PowerAtColdStartForcedOn   byte = 0x01
	PowerAtColdStartLastStatus byte = 0x02
)
