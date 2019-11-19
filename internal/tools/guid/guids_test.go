package guid

import "testing"

func TestGUID_Equals(t *testing.T) {
	type args struct {
		RHS GUID
	}
	tests := []struct {
		name string
		LHS  GUID
		args args
		want bool
	}{
		{
			name: "GUID.01", // Exact match
			LHS:  GUID("d6e4a5e9-696a-4626-ba7a-534d6ff450a5"),
			args: args{RHS: GUID("d6e4a5e9-696a-4626-ba7a-534d6ff450a5")},
			want: true,
		},
		{
			name: "GUID.02", // Differing case
			LHS:  GUID("d6e4a5e9-696a-4626-ba7a-534d6ff450a5"),
			args: args{RHS: GUID("D6E4a5e9-696a-4626-bA7a-534d6ff450a5")},
			want: true,
		},
		{
			name: "GUID.03", // Not equivalent
			LHS:  GUID("d6e4a5e9-696a-4626-ba7a-534d6ff450a5"),
			args: args{RHS: GUID("aaa4a5e9-696a-4626-ba7a-534d6ff450a5")},
			want: false,
		},
		{
			name: "GUID.04",
			LHS:  EmptyGUID,
			args: args{RHS: GUID("d6e4a5e9-696a-4626-ba7a-534d6ff450a5")},
			want: false,
		},
		{
			name: "GUID.05",
			LHS:  GUID("d6e4a5e9-696a-4626-ba7a-534d6ff450a5"),
			args: args{RHS: EmptyGUID},
			want: false,
		},
		{
			name: "GUID.06",
			LHS:  EmptyGUID,
			args: args{RHS: EmptyGUID},
			want: true,
		},
		{
			name: "GUID.07",
			LHS:  EmptyGUID,
			args: args{RHS: GUID("some junk")},
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
