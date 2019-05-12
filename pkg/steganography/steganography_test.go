package steganography

import (
	"context"
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPerserved(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	changeByte(img)
	require.EqualValues(t, img.At(0, 0), color.RGBA{8, 8, 8, 8})
}

func changeByte(img *image.RGBA) {
	img.Set(0, 0, color.RGBA{8, 8, 8, 8})
}

func TestEncode(t *testing.T) {
	msg := []byte("ana are mere")
	ctx := context.Background()
	in := image.NewRGBA(image.Rect(0, 0, 100, 100))
	out, err := EncodeImage(ctx, in, msg, false)
	require.NoError(t, err.Error())
	b := out.Bounds()
	count := uint32(0)
	for x := 0; x < b.Dx(); x++ {
		for y := 0; y < b.Dy(); y++ {
			r, g, b, a := out.At(x, y).RGBA()
			count += (1 & r) + (1 & g) + (1 & b) + (1 & a)
		}
	}
	require.NotZero(t, count)
}

func TestDecode(t *testing.T) {
	msg := []byte("ana are mere")
	ctx := context.Background()
	in := image.NewRGBA(image.Rect(0, 0, 100, 100))
	out, err := EncodeImage(ctx, in, msg, false)
	require.NoError(t, err.Error())
	_, message, err := DecodeImage(ctx, out)
	require.NoError(t, err.Error())
	require.Equal(t, msg, message)

}
