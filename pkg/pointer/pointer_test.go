// Copyright © 2019-2020 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pointer_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/ronaldotantra/leaderboard-api/pkg/pointer"
	"github.com/stretchr/testify/require"
)

func TestAny(t *testing.T) {
	var _ *bool = pointer.Any(true).(*bool)
	var _ *byte = pointer.Any(byte(1)).(*byte)
	var _ *complex64 = pointer.Any(complex64(1.1)).(*complex64)
	var _ *complex128 = pointer.Any(complex128(1.1)).(*complex128)
	var _ *float32 = pointer.Any(float32(1.1)).(*float32)
	var _ *float64 = pointer.Any(float64(1.1)).(*float64)
	var _ *int = pointer.Any(int(1)).(*int)
	var _ *int8 = pointer.Any(int8(8)).(*int8)
	var _ *int16 = pointer.Any(int16(16)).(*int16)
	var _ *int32 = pointer.Any(int32(32)).(*int32)
	var _ *int64 = pointer.Any(int64(64)).(*int64)
	var _ *rune = pointer.Any(rune(1)).(*rune)
	var _ *string = pointer.Any("ptr").(*string)
	var _ *uint = pointer.Any(uint(1)).(*uint)
	var _ *uint8 = pointer.Any(uint8(8)).(*uint8)
	var _ *uint16 = pointer.Any(uint16(16)).(*uint16)
	var _ *uint32 = pointer.Any(uint32(32)).(*uint32)
	var _ *uint64 = pointer.Any(uint64(64)).(*uint64)
	var _ *uintptr = pointer.Any(uintptr(64)).(*uintptr)

	require.Equal(t, true, *pointer.Any(true).(*bool))
	require.Equal(t, false, *pointer.Any(false).(*bool))
	require.Equal(t, true, *(*pointer.Any(pointer.Bool(true)).(**bool)))
	require.Equal(t, int64(1), *pointer.Any(int64(1)).(*int64))
	require.Equal(t, uint64(1), *pointer.Any(uint64(1)).(*uint64))
	require.Equal(t, struct{}{}, *pointer.Any(struct{}{}).(*struct{}))
	require.Equal(t, string("pointer"), *pointer.Any(string("pointer")).(*string))
	require.Equal(t, reflect.Method{}, *pointer.Any(reflect.Method{}).(*reflect.Method))
	require.Equal(t, reflect.Method{Name: "name"}, *pointer.Any(reflect.Method{Name: "name"}).(*reflect.Method))
}

func TestBool(t *testing.T) {
	require.Equal(t, true, *pointer.Bool(true))
	require.Equal(t, false, *pointer.Bool(false))
}

func TestByte(t *testing.T) {
	require.Equal(t, byte(0), *pointer.Byte(byte(0)))
	require.Equal(t, byte(1), *pointer.Byte(byte(1)))
}

func TestComplex64(t *testing.T) {
	require.Equal(t, complex64(0), *pointer.Complex64(complex64(0)))
	require.Equal(t, complex64(1), *pointer.Complex64(complex64(1)))
}

func TestComplex128(t *testing.T) {
	require.Equal(t, complex128(0), *pointer.Complex128(complex128(0)))
	require.Equal(t, complex128(1), *pointer.Complex128(complex128(1)))
}

func TestFloat32(t *testing.T) {
	require.Equal(t, float32(-1.1), *pointer.Float32(float32(-1.1)))
	require.Equal(t, float32(-1), *pointer.Float32(float32(-1)))
	require.Equal(t, float32(0.0), *pointer.Float32(float32(0.0)))
	require.Equal(t, float32(1), *pointer.Float32(float32(1)))
	require.Equal(t, float32(1.1), *pointer.Float32(float32(1.1)))
}

func TestFloat64(t *testing.T) {
	require.Equal(t, float64(-1.1), *pointer.Float64(float64(-1.1)))
	require.Equal(t, float64(-1), *pointer.Float64(float64(-1)))
	require.Equal(t, float64(0.0), *pointer.Float64(float64(0.0)))
	require.Equal(t, float64(1), *pointer.Float64(float64(1)))
	require.Equal(t, float64(1.1), *pointer.Float64(float64(1.1)))
}

func TestInt(t *testing.T) {
	require.Equal(t, int(-1), *pointer.Int(int(-1)))
	require.Equal(t, int(0), *pointer.Int(int(0)))
	require.Equal(t, int(1), *pointer.Int(int(1)))
}

func TestInt8(t *testing.T) {
	require.Equal(t, int8(-1), *pointer.Int8(int8(-1)))
	require.Equal(t, int8(0), *pointer.Int8(int8(0)))
	require.Equal(t, int8(1), *pointer.Int8(int8(1)))
}

func TestInt16(t *testing.T) {
	require.Equal(t, int16(-1), *pointer.Int16(int16(-1)))
	require.Equal(t, int16(0), *pointer.Int16(int16(0)))
	require.Equal(t, int16(1), *pointer.Int16(int16(1)))
}

func TestInt32(t *testing.T) {
	require.Equal(t, int32(-1), *pointer.Int32(int32(-1)))
	require.Equal(t, int32(0), *pointer.Int32(int32(0)))
	require.Equal(t, int32(1), *pointer.Int32(int32(1)))
}

func TestInt64(t *testing.T) {
	require.Equal(t, int64(-1), *pointer.Int64(int64(-1)))
	require.Equal(t, int64(0), *pointer.Int64(int64(0)))
	require.Equal(t, int64(1), *pointer.Int64(int64(1)))
}

func TestRune(t *testing.T) {
	require.Equal(t, rune(0), *pointer.Rune(rune(0)))
	require.Equal(t, rune(1), *pointer.Rune(rune(1)))
}

func TestString(t *testing.T) {
	require.Equal(t, string(""), *pointer.String(string("")))
	require.Equal(t, string("pointer"), *pointer.String(string("pointer")))
}

func TestUint(t *testing.T) {
	require.Equal(t, uint(0), *pointer.Uint(uint(0)))
	require.Equal(t, uint(1), *pointer.Uint(uint(1)))
}

func TestUint8(t *testing.T) {
	require.Equal(t, uint8(0), *pointer.Uint8(uint8(0)))
	require.Equal(t, uint8(1), *pointer.Uint8(uint8(1)))
}

func TestUint16(t *testing.T) {
	require.Equal(t, uint16(0), *pointer.Uint16(uint16(0)))
	require.Equal(t, uint16(1), *pointer.Uint16(uint16(1)))
}

func TestUint32(t *testing.T) {
	require.Equal(t, uint32(0), *pointer.Uint32(uint32(0)))
	require.Equal(t, uint32(1), *pointer.Uint32(uint32(1)))
}

func TestUint64(t *testing.T) {
	require.Equal(t, uint64(0), *pointer.Uint64(uint64(0)))
	require.Equal(t, uint64(1), *pointer.Uint64(uint64(1)))
}

func TestUintptr(t *testing.T) {
	require.Equal(t, uintptr(0), *pointer.Uintptr(uintptr(0)))
	require.Equal(t, uintptr(1), *pointer.Uintptr(uintptr(1)))
}

func TestTime(t *testing.T) {
	require.Equal(t, time.Time{}, *pointer.Time(time.Time{}))
	require.Equal(t, time.Time{}.Add(time.Hour), *pointer.Time(time.Time{}.Add(time.Hour)))
}

func TestDuration(t *testing.T) {
	require.Equal(t, -1*time.Hour, *pointer.Duration(-1 * time.Hour))
	require.Equal(t, time.Duration(0), *pointer.Duration(time.Duration(0)))
	require.Equal(t, 127*time.Hour, *pointer.Duration(127 * time.Hour))
}

func TestConvertToNil(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "success",
			args: args{
				v: "test",
			},
			want: pointer.String("test"),
		}, {
			name: "nil if empty",
			args: args{
				v: "",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, pointer.ConvertToNil[string](tt.args.v))
		})
	}
}
