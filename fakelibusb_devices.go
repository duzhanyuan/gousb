// Copyright 2017 the gousb Authors.  All rights reserved.
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

package gousb

// fake devices connected through the fakeLibusb stack.
var fakeDevices = []*DeviceDesc{
	// Bus 001 Device 001: ID 9999:0001
	// One config, one interface, one setup,
	// two endpoints: 0x01 OUT, 0x82 IN.
	&DeviceDesc{
		Bus:      1,
		Address:  1,
		Spec:     Version(2, 0),
		Device:   Version(1, 0),
		Vendor:   ID(0x9999),
		Product:  ID(0x0001),
		Protocol: 255,
		Configs: map[int]ConfigDesc{1: {
			Number:   1,
			MaxPower: Milliamperes(100),
			Interfaces: []InterfaceDesc{{
				Number: 0,
				AltSettings: []InterfaceSetting{{
					Number:    0,
					Alternate: 0,
					Class:     ClassVendorSpec,
					Endpoints: map[int]EndpointDesc{
						1: {
							Number:        1,
							Direction:     EndpointDirectionOut,
							MaxPacketSize: 512,
							TransferType:  TransferTypeBulk,
						},
						2: {
							Number:        2,
							Direction:     EndpointDirectionIn,
							MaxPacketSize: 512,
							TransferType:  TransferTypeBulk,
						},
					},
				}},
			}},
		}},
	},
	// Bus 001 Device 002: ID 8888:0002
	// One config, two interfaces. interface #0 with no endpoints,
	// interface #1 with two alt setups with different packet sizes for
	// endpoints. Two isochronous endpoints, 0x05 OUT and 0x86 OUT.
	&DeviceDesc{
		Bus:      1,
		Address:  2,
		Spec:     Version(2, 0),
		Device:   Version(1, 3),
		Vendor:   ID(0x8888),
		Product:  ID(0x0002),
		Protocol: 255,
		Configs: map[int]ConfigDesc{1: {
			Number:   1,
			MaxPower: Milliamperes(100),
			Interfaces: []InterfaceDesc{{
				Number: 0,
				AltSettings: []InterfaceSetting{{
					Number:    0,
					Alternate: 0,
					Class:     ClassVendorSpec,
				}},
			}, {
				Number: 1,
				AltSettings: []InterfaceSetting{{
					Number:    1,
					Alternate: 0,
					Class:     ClassVendorSpec,
					Endpoints: map[int]EndpointDesc{
						5: {
							Number:        5,
							Direction:     EndpointDirectionOut,
							MaxPacketSize: 3 * 1024,
							TransferType:  TransferTypeIsochronous,
							UsageType:     IsoUsageTypeData,
						},
						6: {
							Number:        6,
							Direction:     EndpointDirectionIn,
							MaxPacketSize: 3 * 1024,
							TransferType:  TransferTypeIsochronous,
							UsageType:     IsoUsageTypeData,
						},
					},
				}, {
					Number:    1,
					Alternate: 1,
					Class:     ClassVendorSpec,
					Endpoints: map[int]EndpointDesc{
						5: {
							Number:        5,
							Direction:     EndpointDirectionOut,
							MaxPacketSize: 2 * 1024,
							TransferType:  TransferTypeIsochronous,
						},
						6: {
							Number:        6,
							Direction:     EndpointDirectionIn,
							MaxPacketSize: 2 * 1024,
							TransferType:  TransferTypeIsochronous,
						},
					},
				}, {
					Number:    1,
					Alternate: 2,
					Class:     ClassVendorSpec,
					Endpoints: map[int]EndpointDesc{
						5: {
							Number:        5,
							Direction:     EndpointDirectionIn,
							MaxPacketSize: 1024,
							TransferType:  TransferTypeIsochronous,
						},
						6: {
							Number:        6,
							Direction:     EndpointDirectionIn,
							MaxPacketSize: 1024,
							TransferType:  TransferTypeIsochronous,
						},
					},
				}},
			}},
		}},
	},
}
