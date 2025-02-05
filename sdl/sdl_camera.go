package sdl

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/convert"
	"github.com/jupiterrider/purego-sdl3/internal/mem"
)

type CameraPosition uint32

const (
	CameraPositionUnknown CameraPosition = iota
	CameraPositionFrontFacing
	CameraPositionBackFacing
)

type CameraID uint32

type Camera struct{}

type CameraSpec struct {
	Format               PixelFormat
	Colorspace           Colorspace
	Width                int32
	Height               int32
	FramerateNumerator   int32
	FramerateDenominator int32
}

func AcquireCameraFrame(camera *Camera, timestampNS *uint64) *Surface {
	return sdlAcquireCameraFrame(camera, timestampNS)
}

func CloseCamera(camera *Camera) {
	sdlCloseCamera(camera)
}

// func GetCameraDriver(index int32) string {
//	return sdlGetCameraDriver(index)
// }

// func GetCameraFormat(camera *Camera, spec *CameraSpec) bool {
//	return sdlGetCameraFormat(camera, spec)
// }

// func GetCameraID(camera *Camera) CameraID {
//	return sdlGetCameraID(camera)
// }

func GetCameraName(instanceId CameraID) *string {
	ret := sdlGetCameraName(instanceId)
	if ret == nil {
		return nil
	}
	name := convert.ToString(ret)
	return &name
}

// func GetCameraPermissionState(camera *Camera) int32 {
//	return sdlGetCameraPermissionState(camera)
// }

// func GetCameraPosition(instance_id CameraID) CameraPosition {
//	return sdlGetCameraPosition(instance_id)
// }

// func GetCameraProperties(camera *Camera) PropertiesID {
//	return sdlGetCameraProperties(camera)
// }

func GetCameras() []CameraID {
	var count int32
	cameras := sdlGetCameras(&count)
	if cameras == nil {
		return nil
	}
	defer Free(unsafe.Pointer(cameras))
	return mem.Copy(cameras, count)
}

func GetCameraSupportedFormats(instanceId CameraID) []*CameraSpec {
	var count int32
	formats := sdlGetCameraSupportedFormats(instanceId, &count)
	if formats == nil {
		return nil
	}
	defer Free(unsafe.Pointer(formats))
	return mem.Copy(formats, count)
}

// func GetCurrentCameraDriver() string {
//	return sdlGetCurrentCameraDriver()
// }

// func GetNumCameraDrivers() int32 {
//	return sdlGetNumCameraDrivers()
// }

func OpenCamera(instanceId CameraID, spec *CameraSpec) *Camera {
	return sdlOpenCamera(instanceId, spec)
}

func ReleaseCameraFrame(camera *Camera, frame *Surface) {
	sdlReleaseCameraFrame(camera, frame)
}
