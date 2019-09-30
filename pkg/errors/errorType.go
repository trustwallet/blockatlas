package errors

import (
	"fmt"
)

type Type uint16

const (
	TypeNone Type = iota
	TypePlatformUnmarshal
	TypePlatformNormalize
	TypePlatformUnknown
	TypePlatformRequest
	TypePlatformClient
	TypePlatformError
	TypePlatformApi
	TypeUnknown
)

func (e Type) String() string {
	switch e {
	case TypeNone:
		return ""
	case TypePlatformRequest:
		return "Platform Request Error"
	case TypePlatformUnmarshal:
		return "Platform Unmarshal Error"
	case TypePlatformClient:
		return "Platform Client Generic Error"
	case TypePlatformApi:
		return "Platform API Error"
	case TypePlatformNormalize:
		return "Platform Normalize Error"
	case TypePlatformUnknown:
		return "Platform Unknown Error"
	case TypePlatformError:
		return "Custom Platform Error"
	case TypeUnknown:
		return "Unknown Error"
	default:
		return fmt.Sprintf("Error: %d", int(e))
	}
}
