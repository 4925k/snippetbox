package main

import (
	"testing"
	"time"
)

func Test_humanDate(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "correct date",
			args: args{
				time.Date(2020, 12, 17, 10, 0, 0, 0, time.UTC),
			},
			want: "17 Dec 2020 at 10:00",
		},
		{
			name: "empty",
			args: args{
				time.Time{},
			},
			want: "",
		},
		{
			name: "utc time line",
			args: args{
				time.Date(2020, 12, 17, 10, 0, 0, 0, time.FixedZone("CET", 1*60*60)),
			},
			want: "17 Dec 2020 at 09:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := humanDate(tt.args.t); got != tt.want {
				t.Errorf("humanDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
