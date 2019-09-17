package errors

import (
	"fmt"
	"testing"
)

func TestIsType(t *testing.T) {
	var tests = []struct {
		error     error
		errorType Type
		result    bool
	}{
		{fmt.Errorf("test"), TypePlatformRequest, false},
		{&Error{Type: TypePlatformRequest}, TypePlatformRequest, true},
		{&Error{Type: TypePlatformUnmarshal}, TypePlatformRequest, false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestIsType %d", i), func(t *testing.T) {
			s := Is(tt.error, tt.errorType)
			if s != tt.result {
				t.Errorf("got %t, want %t", s, tt.result)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	var tests = []struct {
		err1   error
		err2   error
		result bool
	}{
		{fmt.Errorf("test"), &Error{Type: TypePlatformRequest}, false},
		{&Error{Type: TypePlatformNormalize}, &Error{Type: TypePlatformRequest}, false},
		{&Error{Type: TypePlatformRequest}, &Error{Type: TypePlatformRequest}, true},
		{fmt.Errorf("err1"), fmt.Errorf("err2"), false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestEqual %d", i), func(t *testing.T) {
			s := Equal(tt.err1, tt.err2)
			if s != tt.result {
				t.Errorf("got %t, want %t", s, tt.result)
			}
		})
	}
}
