package estimate

import "testing"

func TestEstimateValue(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test small",
			args: args{
				value: 9,
			},
			want: "small",
		},
		{
			name: "Test medium",
			args: args{
				value: 99,
			},
			want: "medium",
		},
		{
			name: "Test big",
			args: args{
				value: 100,
			},
			want: "big",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EstimateValue(tt.args.value); got != tt.want {
				t.Errorf("EstimateValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
