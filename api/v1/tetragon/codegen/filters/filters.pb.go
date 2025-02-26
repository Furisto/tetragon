// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

// Code generated by protoc-gen-go-tetragon. DO NOT EDIT

package filters

import (
	fmt "fmt"
	tetragon "github.com/cilium/tetragon/api/v1/tetragon"
	reflect "reflect"
)

func OpCodeForEventType(eventType tetragon.EventType) (reflect.Type, error) {
	var opCode reflect.Type
	switch eventType {
	case tetragon.EventType_PROCESS_EXEC:
		opCode = reflect.TypeOf(&tetragon.GetEventsResponse_ProcessExec{})
	case tetragon.EventType_PROCESS_EXIT:
		opCode = reflect.TypeOf(&tetragon.GetEventsResponse_ProcessExit{})
	case tetragon.EventType_PROCESS_KPROBE:
		opCode = reflect.TypeOf(&tetragon.GetEventsResponse_ProcessKprobe{})
	case tetragon.EventType_PROCESS_TRACEPOINT:
		opCode = reflect.TypeOf(&tetragon.GetEventsResponse_ProcessTracepoint{})
	case tetragon.EventType_PROCESS_DNS:
		opCode = reflect.TypeOf(&tetragon.GetEventsResponse_ProcessDns{})
	case tetragon.EventType_TEST:
		opCode = reflect.TypeOf(&tetragon.GetEventsResponse_Test{})
	default:
		return nil, fmt.Errorf("Unknown EventType %s", eventType)
	}
	return opCode, nil
}
