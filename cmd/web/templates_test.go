package main

import (
	"testing"
	"time"

	"github.com/Surak1at/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2021, 12, 31, 12, 30, 0, 0, time.UTC),
			want: "31 Dec 2021 at 12:30",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2021, 12, 31, 12, 30, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "31 Dec 2021 at 11:30",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}
