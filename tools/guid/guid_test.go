package guid_test

import (
	"github.com/microsoft/abstrakt/tools/guid"
	"testing"
)

func TestGUID_Equals(t *testing.T) {
	type args struct {
		RHS guid.GUID
	}
	tests := []struct {
		name string
		LHS  guid.GUID
		args args
		want bool
	}{
		{
			name: "GUID.01", // Exact match
			LHS:  guid.GUID("d6e4a5e9-696a-4626-ba7a-534d6ff450a5"),
			args: args{RHS: guid.GUID("d6e4a5e9-696a-4626-ba7a-534d6ff450a5")},
			want: true,
		},
		{
			name: "GUID.02", // Differing case
			LHS:  guid.GUID("d6e4a5e9-696a-4626-ba7a-534d6ff450a5"),
			args: args{RHS: guid.GUID("D6E4a5e9-696a-4626-bA7a-534d6ff450a5")},
			want: true,
		},
		{
			name: "GUID.03", // Not equivalent
			LHS:  guid.GUID("d6e4a5e9-696a-4626-ba7a-534d6ff450a5"),
			args: args{RHS: guid.GUID("aaa4a5e9-696a-4626-ba7a-534d6ff450a5")},
			want: false,
		},
		{
			name: "GUID.04",
			LHS:  guid.Empty,
			args: args{RHS: guid.GUID("d6e4a5e9-696a-4626-ba7a-534d6ff450a5")},
			want: false,
		},
		{
			name: "GUID.05",
			LHS:  guid.GUID("d6e4a5e9-696a-4626-ba7a-534d6ff450a5"),
			args: args{RHS: guid.Empty},
			want: false,
		},
		{
			name: "GUID.06",
			LHS:  guid.Empty,
			args: args{RHS: guid.Empty},
			want: true,
		},
		{
			name: "GUID.07",
			LHS:  guid.Empty,
			args: args{RHS: guid.GUID("some junk")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.LHS.Equals(tt.args.RHS); got != tt.want {
				t.Errorf("GUID.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
