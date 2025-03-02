package sdl

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/purego-sdl3/internal/convert"
)

type HintPriority uint32

const (
	HintDefault HintPriority = iota
	HintNormal
	HintOverride
)

const (
	HintAllowAltTabWhileGrabbed            = "SDL_ALLOW_ALT_TAB_WHILE_GRABBED"
	HintAndroidAllowRecreateActivity       = "SDL_ANDROID_ALLOW_RECREATE_ACTIVITY"
	HintAndroidBlockOnPause                = "SDL_ANDROID_BLOCK_ON_PAUSE"
	HintAndroidLowLatencyAudio             = "SDL_ANDROID_LOW_LATENCY_AUDIO"
	HintAndroidTrapBackButton              = "SDL_ANDROID_TRAP_BACK_BUTTON"
	HintAppId                              = "SDL_APP_ID"
	HintAppName                            = "SDL_APP_NAME"
	HintAppleTvControllerUiEvents          = "SDL_APPLE_TV_CONTROLLER_UI_EVENTS"
	HintAppleTvRemoteAllowRotation         = "SDL_APPLE_TV_REMOTE_ALLOW_ROTATION"
	HintAudioAlsaDefaultDevice             = "SDL_AUDIO_ALSA_DEFAULT_DEVICE"
	HintAudioAlsaDefaultPlaybackDevice     = "SDL_AUDIO_ALSA_DEFAULT_PLAYBACK_DEVICE"
	HintAudioAlsaDefaultRecordingDevice    = "SDL_AUDIO_ALSA_DEFAULT_RECORDING_DEVICE"
	HintAudioCategory                      = "SDL_AUDIO_CATEGORY"
	HintAudioChannels                      = "SDL_AUDIO_CHANNELS"
	HintAudioDeviceAppIconName             = "SDL_AUDIO_DEVICE_APP_ICON_NAME"
	HintAudioDeviceSampleFrames            = "SDL_AUDIO_DEVICE_SAMPLE_FRAMES"
	HintAudioDeviceStreamName              = "SDL_AUDIO_DEVICE_STREAM_NAME"
	HintAudioDeviceStreamRole              = "SDL_AUDIO_DEVICE_STREAM_ROLE"
	HintAudioDiskInputFile                 = "SDL_AUDIO_DISK_INPUT_FILE"
	HintAudioDiskOutputFile                = "SDL_AUDIO_DISK_OUTPUT_FILE"
	HintAudioDiskTimescale                 = "SDL_AUDIO_DISK_TIMESCALE"
	HintAudioDriver                        = "SDL_AUDIO_DRIVER"
	HintAudioDummyTimescale                = "SDL_AUDIO_DUMMY_TIMESCALE"
	HintAudioFormat                        = "SDL_AUDIO_FORMAT"
	HintAudioFrequency                     = "SDL_AUDIO_FREQUENCY"
	HintAudioIncludeMonitors               = "SDL_AUDIO_INCLUDE_MONITORS"
	HintAutoUpdateJoysticks                = "SDL_AUTO_UPDATE_JOYSTICKS"
	HintAutoUpdateSensors                  = "SDL_AUTO_UPDATE_SENSORS"
	HintBmpSaveLegacyFormat                = "SDL_BMP_SAVE_LEGACY_FORMAT"
	HintCameraDriver                       = "SDL_CAMERA_DRIVER"
	HintCpuFeatureMask                     = "SDL_CPU_FEATURE_MASK"
	HintJoystickDirectinput                = "SDL_JOYSTICK_DIRECTINPUT"
	HintFileDialogDriver                   = "SDL_FILE_DIALOG_DRIVER"
	HintDisplayUsableBounds                = "SDL_DISPLAY_USABLE_BOUNDS"
	HintEmscriptenAsyncify                 = "SDL_EMSCRIPTEN_ASYNCIFY"
	HintEmscriptenCanvasSelector           = "SDL_EMSCRIPTEN_CANVAS_SELECTOR"
	HintEmscriptenKeyboardElement          = "SDL_EMSCRIPTEN_KEYBOARD_ELEMENT"
	HintEnableScreenKeyboard               = "SDL_ENABLE_SCREEN_KEYBOARD"
	HintEvdevDevices                       = "SDL_EVDEV_DEVICES"
	HintEventLogging                       = "SDL_EVENT_LOGGING"
	HintForceRaisewindow                   = "SDL_FORCE_RAISEWINDOW"
	HintFramebufferAcceleration            = "SDL_FRAMEBUFFER_ACCELERATION"
	HintGamecontrollerconfig               = "SDL_GAMECONTROLLERCONFIG"
	HintGamecontrollerconfigFile           = "SDL_GAMECONTROLLERCONFIG_FILE"
	HintGamecontrollertype                 = "SDL_GAMECONTROLLERTYPE"
	HintGamecontrollerIgnoreDevices        = "SDL_GAMECONTROLLER_IGNORE_DEVICES"
	HintGamecontrollerIgnoreDevicesExcept  = "SDL_GAMECONTROLLER_IGNORE_DEVICES_EXCEPT"
	HintGamecontrollerSensorFusion         = "SDL_GAMECONTROLLER_SENSOR_FUSION"
	HintGdkTextinputDefaultText            = "SDL_GDK_TEXTINPUT_DEFAULT_TEXT"
	HintGdkTextinputDescription            = "SDL_GDK_TEXTINPUT_DESCRIPTION"
	HintGdkTextinputMaxLength              = "SDL_GDK_TEXTINPUT_MAX_LENGTH"
	HintGdkTextinputScope                  = "SDL_GDK_TEXTINPUT_SCOPE"
	HintGdkTextinputTitle                  = "SDL_GDK_TEXTINPUT_TITLE"
	HintHidapiLibusb                       = "SDL_HIDAPI_LIBUSB"
	HintHidapiLibusbWhitelist              = "SDL_HIDAPI_LIBUSB_WHITELIST"
	HintHidapiUdev                         = "SDL_HIDAPI_UDEV"
	HintGPUDriver                          = "SDL_GPU_DRIVER"
	HintHidapiEnumerateOnlyControllers     = "SDL_HIDAPI_ENUMERATE_ONLY_CONTROLLERS"
	HintHidapiIgnoreDevices                = "SDL_HIDAPI_IGNORE_DEVICES"
	HintImeImplementedUi                   = "SDL_IME_IMPLEMENTED_UI"
	HintIOSHideHomeIndicator               = "SDL_IOS_HIDE_HOME_INDICATOR"
	HintJoystickAllowBackgroundEvents      = "SDL_JOYSTICK_ALLOW_BACKGROUND_EVENTS"
	HintJoystickArcadestickDevices         = "SDL_JOYSTICK_ARCADESTICK_DEVICES"
	HintJoystickArcadestickDevicesExcluded = "SDL_JOYSTICK_ARCADESTICK_DEVICES_EXCLUDED"
	HintJoystickBlacklistDevices           = "SDL_JOYSTICK_BLACKLIST_DEVICES"
	HintJoystickBlacklistDevicesExcluded   = "SDL_JOYSTICK_BLACKLIST_DEVICES_EXCLUDED"
	HintJoystickDevice                     = "SDL_JOYSTICK_DEVICE"
	HintJoystickEnhancedReports            = "SDL_JOYSTICK_ENHANCED_REPORTS"
	HintJoystickFlightstickDevices         = "SDL_JOYSTICK_FLIGHTSTICK_DEVICES"
	HintJoystickFlightstickDevicesExcluded = "SDL_JOYSTICK_FLIGHTSTICK_DEVICES_EXCLUDED"
	HintJoystickGameinput                  = "SDL_JOYSTICK_GAMEINPUT"
	HintJoystickGamecubeDevices            = "SDL_JOYSTICK_GAMECUBE_DEVICES"
	HintJoystickGamecubeDevicesExcluded    = "SDL_JOYSTICK_GAMECUBE_DEVICES_EXCLUDED"
	HintJoystickHidapi                     = "SDL_JOYSTICK_HIDAPI"
	HintJoystickHidapiCombineJoyCons       = "SDL_JOYSTICK_HIDAPI_COMBINE_JOY_CONS"
	HintJoystickHidapiGamecube             = "SDL_JOYSTICK_HIDAPI_GAMECUBE"
	HintJoystickHidapiGamecubeRumbleBrake  = "SDL_JOYSTICK_HIDAPI_GAMECUBE_RUMBLE_BRAKE"
	HintJoystickHidapiJoyCons              = "SDL_JOYSTICK_HIDAPI_JOY_CONS"
	HintJoystickHidapiJoyconHomeLed        = "SDL_JOYSTICK_HIDAPI_JOYCON_HOME_LED"
	HintJoystickHidapiLuna                 = "SDL_JOYSTICK_HIDAPI_LUNA"
	HintJoystickHidapiNintendoClassic      = "SDL_JOYSTICK_HIDAPI_NINTENDO_CLASSIC"
	HintJoystickHidapiPS3                  = "SDL_JOYSTICK_HIDAPI_PS3"
	HintJoystickHidapiPS3SixaxisDriver     = "SDL_JOYSTICK_HIDAPI_PS3_SIXAXIS_DRIVER"
	HintJoystickHidapiPS4                  = "SDL_JOYSTICK_HIDAPI_PS4"
	HintJoystickHidapiPS4ReportInterval    = "SDL_JOYSTICK_HIDAPI_PS4_REPORT_INTERVAL"
	HintJoystickHidapiPS5                  = "SDL_JOYSTICK_HIDAPI_PS5"
	HintJoystickHidapiPS5PlayerLed         = "SDL_JOYSTICK_HIDAPI_PS5_PLAYER_LED"
	HintJoystickHidapiShield               = "SDL_JOYSTICK_HIDAPI_SHIELD"
	HintJoystickHidapiStadia               = "SDL_JOYSTICK_HIDAPI_STADIA"
	HintJoystickHidapiSteam                = "SDL_JOYSTICK_HIDAPI_STEAM"
	HintJoystickHidapiSteamHomeLed         = "SDL_JOYSTICK_HIDAPI_STEAM_HOME_LED"
	HintJoystickHidapiSteamdeck            = "SDL_JOYSTICK_HIDAPI_STEAMDECK"
	HintJoystickHidapiSteamHori            = "SDL_JOYSTICK_HIDAPI_STEAM_HORI"
	HintJoystickHidapiSwitch               = "SDL_JOYSTICK_HIDAPI_SWITCH"
	HintJoystickHidapiSwitchHomeLed        = "SDL_JOYSTICK_HIDAPI_SWITCH_HOME_LED"
	HintJoystickHidapiSwitchPlayerLed      = "SDL_JOYSTICK_HIDAPI_SWITCH_PLAYER_LED"
	HintJoystickHidapiVerticalJoyCons      = "SDL_JOYSTICK_HIDAPI_VERTICAL_JOY_CONS"
	HintJoystickHidapiWii                  = "SDL_JOYSTICK_HIDAPI_WII"
	HintJoystickHidapiWiiPlayerLed         = "SDL_JOYSTICK_HIDAPI_WII_PLAYER_LED"
	HintJoystickHidapiXbox                 = "SDL_JOYSTICK_HIDAPI_XBOX"
	HintJoystickHidapiXbox360              = "SDL_JOYSTICK_HIDAPI_XBOX_360"
	HintJoystickHidapiXbox360PlayerLed     = "SDL_JOYSTICK_HIDAPI_XBOX_360_PLAYER_LED"
	HintJoystickHidapiXbox360Wireless      = "SDL_JOYSTICK_HIDAPI_XBOX_360_WIRELESS"
	HintJoystickHidapiXboxOne              = "SDL_JOYSTICK_HIDAPI_XBOX_ONE"
	HintJoystickHidapiXboxOneHomeLed       = "SDL_JOYSTICK_HIDAPI_XBOX_ONE_HOME_LED"
	HintJoystickIOKit                      = "SDL_JOYSTICK_IOKIT"
	HintJoystickLinuxClassic               = "SDL_JOYSTICK_LINUX_CLASSIC"
	HintJoystickLinuxDeadzones             = "SDL_JOYSTICK_LINUX_DEADZONES"
	HintJoystickLinuxDigitalHats           = "SDL_JOYSTICK_LINUX_DIGITAL_HATS"
	HintJoystickLinuxHatDeadzones          = "SDL_JOYSTICK_LINUX_HAT_DEADZONES"
	HintJoystickMfi                        = "SDL_JOYSTICK_MFI"
	HintJoystickRawinput                   = "SDL_JOYSTICK_RAWINPUT"
	HintJoystickRawinputCorrelateXinput    = "SDL_JOYSTICK_RAWINPUT_CORRELATE_XINPUT"
	HintJoystickRogChakram                 = "SDL_JOYSTICK_ROG_CHAKRAM"
	HintJoystickThread                     = "SDL_JOYSTICK_THREAD"
	HintJoystickThrottleDevices            = "SDL_JOYSTICK_THROTTLE_DEVICES"
	HintJoystickThrottleDevicesExcluded    = "SDL_JOYSTICK_THROTTLE_DEVICES_EXCLUDED"
	HintJoystickWgi                        = "SDL_JOYSTICK_WGI"
	HintJoystickWheelDevices               = "SDL_JOYSTICK_WHEEL_DEVICES"
	HintJoystickWheelDevicesExcluded       = "SDL_JOYSTICK_WHEEL_DEVICES_EXCLUDED"
	HintJoystickZeroCenteredDevices        = "SDL_JOYSTICK_ZERO_CENTERED_DEVICES"
	HintKeycodeOptions                     = "SDL_KEYCODE_OPTIONS"
	HintKmsdrmDeviceIndex                  = "SDL_KMSDRM_DEVICE_INDEX"
	HintKmsdrmRequireDrmMaster             = "SDL_KMSDRM_REQUIRE_DRM_MASTER"
	HintLogging                            = "SDL_LOGGING"
	HintMacBackgroundApp                   = "SDL_MAC_BACKGROUND_APP"
	HintMacCtrlClickEmulateRightClick      = "SDL_MAC_CTRL_CLICK_EMULATE_RIGHT_CLICK"
	HintMacOpenglAsyncDispatch             = "SDL_MAC_OPENGL_ASYNC_DISPATCH"
	HintMacOptionAsAlt                     = "SDL_MAC_OPTION_AS_ALT"
	HintMacScrollMomentum                  = "SDL_MAC_SCROLL_MOMENTUM"
	HintMainCallbackRate                   = "SDL_MAIN_CALLBACK_RATE"
	HintMouseAutoCapture                   = "SDL_MOUSE_AUTO_CAPTURE"
	HintMouseDoubleClickRadius             = "SDL_MOUSE_DOUBLE_CLICK_RADIUS"
	HintMouseDoubleClickTime               = "SDL_MOUSE_DOUBLE_CLICK_TIME"
	HintMouseDefaultSystemCursor           = "SDL_MOUSE_DEFAULT_SYSTEM_CURSOR"
	HintMouseEmulateWarpWithRelative       = "SDL_MOUSE_EMULATE_WARP_WITH_RELATIVE"
	HintMouseFocusClickthrough             = "SDL_MOUSE_FOCUS_CLICKTHROUGH"
	HintMouseNormalSpeedScale              = "SDL_MOUSE_NORMAL_SPEED_SCALE"
	HintMouseRelativeModeCenter            = "SDL_MOUSE_RELATIVE_MODE_CENTER"
	HintMouseRelativeSpeedScale            = "SDL_MOUSE_RELATIVE_SPEED_SCALE"
	HintMouseRelativeSystemScale           = "SDL_MOUSE_RELATIVE_SYSTEM_SCALE"
	HintMouseRelativeWarpMotion            = "SDL_MOUSE_RELATIVE_WARP_MOTION"
	HintMouseRelativeCursorVisible         = "SDL_MOUSE_RELATIVE_CURSOR_VISIBLE"
	HintMouseTouchEvents                   = "SDL_MOUSE_TOUCH_EVENTS"
	HintMuteConsoleKeyboard                = "SDL_MUTE_CONSOLE_KEYBOARD"
	HintNoSignalHandlers                   = "SDL_NO_SIGNAL_HANDLERS"
	HintOpenglLibrary                      = "SDL_OPENGL_LIBRARY"
	HintEglLibrary                         = "SDL_EGL_LIBRARY"
	HintOpenglEsDriver                     = "SDL_OPENGL_ES_DRIVER"
	HintOpenvrLibrary                      = "SDL_OPENVR_LIBRARY"
	HintOrientations                       = "SDL_ORIENTATIONS"
	HintPollSentinel                       = "SDL_POLL_SENTINEL"
	HintPreferredLocales                   = "SDL_PREFERRED_LOCALES"
	HintQuitOnLastWindowClose              = "SDL_QUIT_ON_LAST_WINDOW_CLOSE"
	HintRenderDirect3DThreadsafe           = "SDL_RENDER_DIRECT3D_THREADSAFE"
	HintRenderDirect3D11Debug              = "SDL_RENDER_DIRECT3D11_DEBUG"
	HintRenderVulkanDebug                  = "SDL_RENDER_VULKAN_DEBUG"
	HintRenderGPUDebug                     = "SDL_RENDER_GPU_DEBUG"
	HintRenderGPULowPower                  = "SDL_RENDER_GPU_LOW_POWER"
	HintRenderDriver                       = "SDL_RENDER_DRIVER"
	HintRenderLineMethod                   = "SDL_RENDER_LINE_METHOD"
	HintRenderMetalPreferLowPowerDevice    = "SDL_RENDER_METAL_PREFER_LOW_POWER_DEVICE"
	HintRenderVsync                        = "SDL_RENDER_VSYNC"
	HintReturnKeyHidesIme                  = "SDL_RETURN_KEY_HIDES_IME"
	HintRogGamepadMice                     = "SDL_ROG_GAMEPAD_MICE"
	HintRogGamepadMiceExcluded             = "SDL_ROG_GAMEPAD_MICE_EXCLUDED"
	HintRpiVideoLayer                      = "SDL_RPI_VIDEO_LAYER"
	HintScreensaverInhibitActivityName     = "SDL_SCREENSAVER_INHIBIT_ACTIVITY_NAME"
	HintShutdownDbusOnQuit                 = "SDL_SHUTDOWN_DBUS_ON_QUIT"
	HintStorageTitleDriver                 = "SDL_STORAGE_TITLE_DRIVER"
	HintStorageUserDriver                  = "SDL_STORAGE_USER_DRIVER"
	HintThreadForceRealtimeTimeCritical    = "SDL_THREAD_FORCE_REALTIME_TIME_CRITICAL"
	HintThreadPriorityPolicy               = "SDL_THREAD_PRIORITY_POLICY"
	HintTimerResolution                    = "SDL_TIMER_RESOLUTION"
	HintTouchMouseEvents                   = "SDL_TOUCH_MOUSE_EVENTS"
	HintTrackpadIsTouchOnly                = "SDL_TRACKPAD_IS_TOUCH_ONLY"
	HintTvRemoteAsJoystick                 = "SDL_TV_REMOTE_AS_JOYSTICK"
	HintVideoAllowScreensaver              = "SDL_VIDEO_ALLOW_SCREENSAVER"
	HintVideoDisplayPriority               = "SDL_VIDEO_DISPLAY_PRIORITY"
	HintVideoDoubleBuffer                  = "SDL_VIDEO_DOUBLE_BUFFER"
	HintVideoDriver                        = "SDL_VIDEO_DRIVER"
	HintVideoDummySaveFrames               = "SDL_VIDEO_DUMMY_SAVE_FRAMES"
	HintVideoEglAllowGetdisplayFallback    = "SDL_VIDEO_EGL_ALLOW_GETDISPLAY_FALLBACK"
	HintVideoForceEgl                      = "SDL_VIDEO_FORCE_EGL"
	HintVideoMacFullscreenSpaces           = "SDL_VIDEO_MAC_FULLSCREEN_SPACES"
	HintVideoMacFullscreenMenuVisibility   = "SDL_VIDEO_MAC_FULLSCREEN_MENU_VISIBILITY"
	HintVideoMinimizeOnFocusLoss           = "SDL_VIDEO_MINIMIZE_ON_FOCUS_LOSS"
	HintVideoOffscreenSaveFrames           = "SDL_VIDEO_OFFSCREEN_SAVE_FRAMES"
	HintVideoSyncWindowOperations          = "SDL_VIDEO_SYNC_WINDOW_OPERATIONS"
	HintVideoWaylandAllowLibdecor          = "SDL_VIDEO_WAYLAND_ALLOW_LIBDECOR"
	HintVideoWaylandModeEmulation          = "SDL_VIDEO_WAYLAND_MODE_EMULATION"
	HintVideoWaylandModeScaling            = "SDL_VIDEO_WAYLAND_MODE_SCALING"
	HintVideoWaylandPreferLibdecor         = "SDL_VIDEO_WAYLAND_PREFER_LIBDECOR"
	HintVideoWaylandScaleToDisplay         = "SDL_VIDEO_WAYLAND_SCALE_TO_DISPLAY"
	HintVideoWinD3Dcompiler                = "SDL_VIDEO_WIN_D3DCOMPILER"
	HintVideoX11NetWmBypassCompositor      = "SDL_VIDEO_X11_NET_WM_BYPASS_COMPOSITOR"
	HintVideoX11NetWmPing                  = "SDL_VIDEO_X11_NET_WM_PING"
	HintVideoX11Nodirectcolor              = "SDL_VIDEO_X11_NODIRECTCOLOR"
	HintVideoX11ScalingFactor              = "SDL_VIDEO_X11_SCALING_FACTOR"
	HintVideoX11Visualid                   = "SDL_VIDEO_X11_VISUALID"
	HintVideoX11WindowVisualid             = "SDL_VIDEO_X11_WINDOW_VISUALID"
	HintVideoX11Xrandr                     = "SDL_VIDEO_X11_XRANDR"
	HintVitaEnableBackTouch                = "SDL_VITA_ENABLE_BACK_TOUCH"
	HintVitaEnableFrontTouch               = "SDL_VITA_ENABLE_FRONT_TOUCH"
	HintVitaModulePath                     = "SDL_VITA_MODULE_PATH"
	HintVitaPvrInit                        = "SDL_VITA_PVR_INIT"
	HintVitaResolution                     = "SDL_VITA_RESOLUTION"
	HintVitaPvrOpengl                      = "SDL_VITA_PVR_OPENGL"
	HintVitaTouchMouseDevice               = "SDL_VITA_TOUCH_MOUSE_DEVICE"
	HintVulkanDisplay                      = "SDL_VULKAN_DISPLAY"
	HintVulkanLibrary                      = "SDL_VULKAN_LIBRARY"
	HintWaveFactChunk                      = "SDL_WAVE_FACT_CHUNK"
	HintWaveChunkLimit                     = "SDL_WAVE_CHUNK_LIMIT"
	HintWaveRiffChunkSize                  = "SDL_WAVE_RIFF_CHUNK_SIZE"
	HintWaveTruncation                     = "SDL_WAVE_TRUNCATION"
	HintWindowActivateWhenRaised           = "SDL_WINDOW_ACTIVATE_WHEN_RAISED"
	HintWindowActivateWhenShown            = "SDL_WINDOW_ACTIVATE_WHEN_SHOWN"
	HintWindowAllowTopmost                 = "SDL_WINDOW_ALLOW_TOPMOST"
	HintWindowFrameUsableWhileCursorHidden = "SDL_WINDOW_FRAME_USABLE_WHILE_CURSOR_HIDDEN"
	HintWindowsCloseOnAltF4                = "SDL_WINDOWS_CLOSE_ON_ALT_F4"
	HintWindowsEnableMenuMnemonics         = "SDL_WINDOWS_ENABLE_MENU_MNEMONICS"
	HintWindowsEnableMessageloop           = "SDL_WINDOWS_ENABLE_MESSAGELOOP"
	HintWindowsGameinput                   = "SDL_WINDOWS_GAMEINPUT"
	HintWindowsRawKeyboard                 = "SDL_WINDOWS_RAW_KEYBOARD"
	HintWindowsForceSemaphoreKernel        = "SDL_WINDOWS_FORCE_SEMAPHORE_KERNEL"
	HintWindowsIntresourceIcon             = "SDL_WINDOWS_INTRESOURCE_ICON"
	HintWindowsIntresourceIconSmall        = "SDL_WINDOWS_INTRESOURCE_ICON_SMALL"
	HintWindowsUseD3D9Ex                   = "SDL_WINDOWS_USE_D3D9EX"
	HintWindowsEraseBackgroundMode         = "SDL_WINDOWS_ERASE_BACKGROUND_MODE"
	HintX11ForceOverrideRedirect           = "SDL_X11_FORCE_OVERRIDE_REDIRECT"
	HintX11WindowType                      = "SDL_X11_WINDOW_TYPE"
	HintX11XcbLibrary                      = "SDL_X11_XCB_LIBRARY"
	HintXinputEnabled                      = "SDL_XINPUT_ENABLED"
	HintAssert                             = "SDL_ASSERT"
	HintPenMouseEvents                     = "SDL_PEN_MOUSE_EVENTS"
	HintPenTouchEvents                     = "SDL_PEN_TOUCH_EVENTS"
)

type HintCallback uintptr

func NewHintCallback(callback func(userdata unsafe.Pointer, name, oldValue, newValue string)) HintCallback {
	cb := purego.NewCallback(func(userdata unsafe.Pointer, name, oldValue, newValue *byte) uintptr {
		callback(userdata, convert.ToString(name), convert.ToString(oldValue), convert.ToString(newValue))
		return 0
	})

	return HintCallback(cb)
}

// SetHint sets a hint with normal priority.
func SetHint(name, value string) bool {
	return sdlSetHint(name, value)
}

func AddHintCallback(name string, callback HintCallback, userdata unsafe.Pointer) bool {
	return sdlAddHintCallback(name, callback, userdata)
}

func GetHint(name string) string {
	return sdlGetHint(name)
}

func GetHintBoolean(name string, defaultValue bool) bool {
	return sdlGetHintBoolean(name, defaultValue)
}

func RemoveHintCallback(name string, callback HintCallback, userdata unsafe.Pointer) {
	sdlRemoveHintCallback(name, callback, userdata)
}

func ResetHint(name string) bool {
	return sdlResetHint(name)
}

func ResetHints() {
	sdlResetHints()
}

func SetHintWithPriority(name string, value string, priority HintPriority) bool {
	return sdlSetHintWithPriority(name, value, priority)
}
