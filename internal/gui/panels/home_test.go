package panels

import "testing"

func Test_updateDate(t *testing.T) {
	type args struct {
		setDate string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateDate(tt.args.setDate)
		})
	}
}
