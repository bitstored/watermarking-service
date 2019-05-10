package steganography

import (
	"context"
	"encoding/binary"
	"github.com/bitstored/watermarking-service/pkg/errors"
	"image"
)

func DecodeImage(ctx context.Context, img image.Image) (image.Image, []byte, *errors.Err) {
	if img == nil {
		return img, nil, errors.NewError(errors.KindEmptyImage, "Emage ca'nt be empty")
	}
	b := img.Bounds()

	out := imageToRGBA(img)
	message, err := decodeBytes(out)
	if err != nil {
		return out, nil, err
	}
	header, length := extractHeaderData(ctx, message)
	if header == nil {
		return out, nil, errors.NewError(errors.KindMalformedStructure, "can't get data from image, image is not water marked")
	}
	if length*8 > bitCount(b.Dx(), b.Dy()) {
		return out, nil, errors.NewError(errors.KindMalformedStructure, "can't get data from image, image is not water marked")
	}
	offset := headerLenBits/8 + messageLenBits/8
	return out, message[offset : length+offset], nil
}

func extractHeaderData(ctx context.Context, message []byte) ([]byte, int) {
	header := message[0 : headerLenBits/8]
	if header[0] != headerSignature[0] {
		return nil, -1
	}
	length := int(binary.LittleEndian.Uint16(header[headerLenBits/8 : headerLenBits/8+messageLenBits/8]))

	return header, length
}

/*
  @startIndex is the index where in the image the next bit should be set
*/
func decodeBytes(img *image.RGBA) ([]byte, *errors.Err) {
	b := img.Bounds()
	dx := b.Dx()
	dy := b.Dy()
	out := make([]byte, (dx*2+dy*2-4)/4)
	for index := uint(0); ; index += 4 {
		x, y := getCoordinatesFromIndex(dx, dy, int(index/4))
		if x == -1 && y == -1 {
			break
		}
		r, g, b, a := img.At(x, y).RGBA()
		var aux uint32
		setBit(&aux, r, index%8)
		setBit(&aux, g, index%8+1)
		setBit(&aux, b, index%8+2)
		setBit(&aux, a, index%8+3)
		out[index/8] |= byte(aux)
	}
	return out, nil
}
