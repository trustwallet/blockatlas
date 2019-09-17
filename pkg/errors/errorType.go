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
	TypePlatformApi
	TypeStorageSave
	TypeStorageGet
	TypeLoadConfig
	TypeLoadCoins
	TypeObserver
	TypeAssets
	TypeUtil
	TypeCmd
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
	case TypePlatformApi:
		return "Platform API Error"
	case TypePlatformNormalize:
		return "Platform Normalize Error"
	case TypePlatformUnknown:
		return "Platform Unknown Error"
	case TypeObserver:
		return "Observer Error"
	case TypeStorageSave:
		return "Storage Save Error"
	case TypeStorageGet:
		return "Storage Get Error"
	case TypeLoadConfig:
		return "Load Config Error"
	case TypeLoadCoins:
		return "Load Coins Error"
	case TypeAssets:
		return "Assets Error"
	case TypeUtil:
		return "Util Error"
	case TypeCmd:
		return "Cmd Error"
	case TypeUnknown:
		return "Unknown Error"
	default:
		return fmt.Sprintf("Error: %d", int(e))
	}
}
