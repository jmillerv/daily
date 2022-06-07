package panels

import (
	"github.com/matryer/is"
	"testing"
	"time"
)

func Test_updateDate(t *testing.T) {
	is := is.New(t)
	type args struct {
		setDate      string
		expectedDate string
	}
	todayDate := time.Now().Format(defaultDateFormat)
	yesterday := time.Now().AddDate(0, 0, -1).Format(defaultDateFormat)
	tests := []struct {
		name string
		args *args
	}{
		{
			name: "Date Is Today",
			args: &args{setDate: todayDate, expectedDate: todayDate},
		},
		{
			name: "Date Is Not Today",
			args: &args{setDate: yesterday, expectedDate: todayDate},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateDate(tt.args.setDate)
			is.Equal(todayDate, tt.args.expectedDate)
		})
	}
}
