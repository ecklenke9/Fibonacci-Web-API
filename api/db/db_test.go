package db

import "testing"

func TestConnectDataBase(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Successful Connect Database",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
