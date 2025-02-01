package pointer_test

import (
	"testing"
	"time"

	"github.com/ronaldotantra/leaderboard-api/pkg/pointer"
	"github.com/stretchr/testify/require"
)

func TestSafeOfInt64(t *testing.T) {
	type args struct {
		v *int64
	}
	tests := []struct {
		name       string
		args       args
		wantOutput int64
	}{
		{
			name: "Test Case 1",
			args: args{
				v: pointer.Int64(1),
			},
			wantOutput: 1,
		}, {
			name: "Test Case 2",
			args: args{
				v: nil,
			},
			wantOutput: 0,
		}, {
			name: "Test Case 3",
			args: args{
				v: pointer.Int64(99),
			},
			wantOutput: 99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.wantOutput, pointer.SafeOf[int64](tt.args.v))
		})
	}
}

func TestSafeTimeUnixNilToNow(t *testing.T) {
	timeNow := time.Now().UnixMilli()

	type args struct {
		v *int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Case 1 Args Nil",
			args: args{
				v: nil,
			},
		}, {
			name: "Test Case 2 Args Not Nil",
			args: args{
				v: &timeNow,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pointer.SafeTimeUnixNilToNow(tt.args.v)
			require.NotZero(t, result)
		})
	}
}
