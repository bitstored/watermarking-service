package steganography

import (
	"context"
	"encoding/binary"
	"github.com/bitstored/watermarking-service/pkg/errors"
	"image"
	"image/color"
)

/*
	STEGANOGRAPHY RULES
	First $headerLen LSBs are encoded with $header
	Next $messageLenBits are encoded with actual hidden message length
	Message is crypted with a key and after that the mesage is encoded in LSBs of the pixels
	The data is encoded in the upper, lower, left, right corners in succesive mode, the corner pixels are ignored
*/
var (
	headerLenBits   = 8
	messageLenBits  = 32
	headerSignature = []byte("S")
)

/*
	@ctx is the context of grpc request
	@img is the image to be filled with secret text
	@text is the message to  be encoded in the image
	@fullMessage is the option if the user wants the entire message, if the full message can't be encoded an error is returned,
	if the @fullMessage is false the message will be truncated
*/
// EncodeImage is the method for data encoding
func EncodeImage(ctx context.Context, img image.Image, text []byte, fullMessage bool) (image.Image, *errors.Err) {

	if img == nil {
		return nil, errors.NewError(errors.KindEmptyImage, "Emage ca'nt be empty")
	}

	if text == nil || len(text) == 0 {
		return nil, errors.NewError(errors.KindInvalidArgument, "Text for watermarking can't be empty")
	}

	b := img.Bounds()
	if fullMessage && len(text)*8+headerLenBits+messageLenBits > bitCount(b.Dx(), b.Dy()) {
		return nil, errors.NewError(errors.KindInvalidArgument, "Message is too long to be encoded in this image")
	}

	out := imageToRGBA(img)

	data := fillHeader(ctx, len(text))
	data = append(data, []byte(text)...)

	out, err := encodeBytes(out, data)

	if err != nil && err.Kind != errors.KindOverflow {
		return nil, err
	}

	return out, nil
}

func fillHeader(ctx context.Context, messageLen int) []byte {
	header := make([]byte, messageLenBits/8+headerLenBits/8)
	header[0] = headerSignature[0]
	binary.LittleEndian.PutUint32(header[headerLenBits/8:(headerLenBits+messageLenBits)/8], uint32(messageLen))
	return header
}

/*
  @startIndex is the index where in the image the next bit should be set
*/
func encodeBytes(img *image.RGBA, data []byte) (*image.RGBA, *errors.Err) {
	count := uint(len(data) * 8)
	b := img.Bounds()
	dx := b.Dx()
	dy := b.Dy()
	for index := uint(0); index < count; index += 4 {
		x, y := getCoordinatesFromIndex(dx, dy, int(index/4))
		r, g, b, a := img.At(x, y).RGBA()
		d := data[index/8]
		setBit(&r, uint32(d), index%8)
		setBit(&g, uint32(d), index%8+1)
		setBit(&b, uint32(d), index%8+2)
		setBit(&a, uint32(d), index%8+3)
		col := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
		img.SetRGBA(x, y, col)
	}
	return img, nil
}
