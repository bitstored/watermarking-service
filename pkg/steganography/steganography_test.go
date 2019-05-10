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

func TestEncodeDecode(t *testing.T) {
	msg := []byte("ana are mere")
	ctx := context.Background()
	in := image.NewRGBA(image.Rect(0, 0, 100, 100))
	out, err := EncodeImage(ctx, in, msg, false)
	require.NoError(t, err.Error())
	_, msg1, err := DecodeImage(ctx, out)
	require.NoError(t, err.Error())
	require.EqualValues(t, msg, msg1)

}
