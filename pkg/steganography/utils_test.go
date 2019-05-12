package steganography

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetCoordinatesFromIndex(t *testing.T) {
	ts := []struct {
		Name  string
		X     int
		Y     int
		Index int
		Dx    int
		Dy    int
	}{
		{
			Name:  "OverflowBoth",
			Dx:    10,
			Dy:    7,
			Index: 48,
			X:     -1,
			Y:     -1,
		},
		{
			Name:  "OverflowX",
			Dx:    10,
			Dy:    15,
			Index: 8*4 + 2,
			X:     9,
			Y:     10,
		},
		{
			Name:  "OverflowY",
			Dy:    10,
			Dx:    15,
			Index: 8*4 + 3,
			Y:     0,
			X:     10,
		},
		{
			Name:  "Ok1",
			Dy:    16,
			Dx:    15,
			Index: 4 + 4 + 0,
			Y:     3,
			X:     0,
		},
		{
			Name:  "Ok2",
			Dy:    16,
			Dx:    15,
			Index: 4 + 4 + 1,
			Y:     3,
			X:     14,
		},
		{
			Name:  "Ok3",
			Dy:    16,
			Dx:    15,
			Index: 4 + 4 + 2,
			Y:     0,
			X:     3,
		},
		{
			Name:  "Ok4",
			Dy:    16,
			Dx:    15,
			Index: 4 + 4 + 3,
			Y:     15,
			X:     3,
		},
		{
			Name:  "Ok5",
			Dy:    16,
			Dx:    15,
			Index: 0,
			Y:     1,
			X:     0,
		},
	}
	for _, tc := range ts {
		t.Run(tc.Name, func(t *testing.T) {
			x, y := getCoordinatesFromIndex(tc.Dx, tc.Dy, tc.Index)
			require.Equal(t, tc.X, x)
			require.Equal(t, tc.Y, y)

		})
	}
}
