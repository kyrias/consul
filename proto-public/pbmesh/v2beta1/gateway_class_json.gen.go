// Code generated by protoc-json-shim. DO NOT EDIT.
package meshv2beta1

import (
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON is a custom marshaler for GatewayClass
func (this *GatewayClass) MarshalJSON() ([]byte, error) {
	str, err := GatewayClassMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GatewayClass
func (this *GatewayClass) UnmarshalJSON(b []byte) error {
	return GatewayClassUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for ParametersReference
func (this *ParametersReference) MarshalJSON() ([]byte, error) {
	str, err := GatewayClassMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ParametersReference
func (this *ParametersReference) UnmarshalJSON(b []byte) error {
	return GatewayClassUnmarshaler.Unmarshal(b, this)
}

var (
	GatewayClassMarshaler   = &protojson.MarshalOptions{}
	GatewayClassUnmarshaler = &protojson.UnmarshalOptions{DiscardUnknown: false}
)
