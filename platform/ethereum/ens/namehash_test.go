// Copyright 2017 Weald Technology Trading
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ens

import (
	"encoding/hex"
	"testing"
)

func TestNameHash(t *testing.T) {
	tests := []struct {
		input  string
		output string
		err    error
	}{
		{"", "0000000000000000000000000000000000000000000000000000000000000000", nil},
		{"eth", "93cdeb708b7545dc668eb9280176169d1c33cfd8ed6f04690a0bcc88a93fc4ae", nil},
		{"Eth", "93cdeb708b7545dc668eb9280176169d1c33cfd8ed6f04690a0bcc88a93fc4ae", nil},
		{".eth", "8cc9f31a5e7af6381efc751d98d289e3f3589f1b6f19b9b989ace1788b939cf7", nil},
		{"resolver.eth", "fdd5d5de6dd63db72bbc2d487944ba13bf775b50a80805fe6fcaba9b0fba88f5", nil},
		{"foo.eth", "de9b09fd7c5f901e23a3f19fecc54828e9c848539801e86591bd9801b019f84f", nil},
		{"Foo.eth", "de9b09fd7c5f901e23a3f19fecc54828e9c848539801e86591bd9801b019f84f", nil},
		{"foo..eth", "4143a5b2f547838d3b49982e3f2ec6a26415274e5b9c3ffeb21971bbfdfaa052", nil},
		{"bar.foo.eth", "275ae88e7263cdce5ab6cf296cdd6253f5e385353fe39cfff2dd4a2b14551cf3", nil},
		{"Bar.foo.eth", "275ae88e7263cdce5ab6cf296cdd6253f5e385353fe39cfff2dd4a2b14551cf3", nil},
		{"addr.reverse", "91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e2", nil},
	}

	for _, tt := range tests {
		result, err := NameHash(tt.input)
		if tt.err == nil {
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}
			if tt.output != hex.EncodeToString(result[:]) {
				t.Errorf("Failure: %v => %v (expected %v)\n", tt.input, hex.EncodeToString(result[:]), tt.output)
			}
		} else {
			if err == nil {
				t.Fatalf("missing expected error")
			}
			if tt.err.Error() != err.Error() {
				t.Errorf("unexpected error value %v", err)
			}
		}
	}
}

func TestLabelHash(t *testing.T) {
	tests := []struct {
		input  string
		output string
		err    error
	}{
		{"", "c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470", nil},
		{"eth", "4f5b812789fc606be1b3b16908db13fc7a9adf7ca72641f84d75b47069d3d7f0", nil},
		{"foo", "41b1a0649752af1b28b3dc29a1556eee781e4a4c3a1f7f53f90fa834de098c4d", nil},
	}

	for _, tt := range tests {
		output, err := LabelHash(tt.input)
		if tt.err == nil {
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}
			if tt.output != hex.EncodeToString(output[:]) {
				t.Errorf("Failure: %v => %v (expected %v)\n", tt.input, hex.EncodeToString(output[:]), tt.output)
			}
		} else {
			if err == nil {
				t.Fatalf("missing expected error")
			}
			if tt.err.Error() != err.Error() {
				t.Errorf("unexpected error value %v", err)
			}
		}
	}
}
