package service

import (
	"bytes"
	"context"
	"github.com/bitstored/watermarking-service/pkg/errors"
	"github.com/bitstored/watermarking-service/pkg/watermarking"
	"image/png"

	"github.com/bitstored/watermarking-service/pkg/steganography"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}
func (s *Service) EncodeMessage(ctx context.Context, src []byte, text []byte, fullMessage bool) ([]byte, *errors.Err) {
	reader := bytes.NewReader(src)
	img, err := png.Decode(reader)
	if err != nil {
		return nil, errors.NewError(errors.KindMalformedStructure, err.Error())
	}
	out, err1 := steganography.EncodeImage(ctx, img, text, fullMessage)
	if err1 != nil {
		return nil, err1
	}
	buf := new(bytes.Buffer)
	err = png.Encode(buf, out)
	bytes := buf.Bytes()
	return bytes, nil
}

func (s *Service) DecodeMessage(ctx context.Context, src []byte) ([]byte, []byte, *errors.Err) {
	reader := bytes.NewReader(src)
	img, err := png.Decode(reader)
	if err != nil {
		return nil, nil, errors.NewError(errors.KindMalformedStructure, err.Error())
	}
	out, text, err1 := steganography.DecodeImage(ctx, img)
	if err1 != nil {
		return nil, nil, err1
	}
	buf := new(bytes.Buffer)
	err = png.Encode(buf, out)
	bytes := buf.Bytes()
	return bytes, text, nil
}

func (s *Service) WatermarkImageWithText(ctx context.Context, src []byte, text string) ([]byte, *errors.Err) {
	reader := bytes.NewReader(src)
	img, err := png.Decode(reader)
	if err != nil {
		return nil, errors.NewError(errors.KindMalformedStructure, err.Error())
	}
	out, err1 := watermarking.WatermarkImageWithText(ctx, img, text)
	if err1 != nil {
		return nil, err1
	}
	buf := new(bytes.Buffer)
	err = png.Encode(buf, out)
	bytes := buf.Bytes()
	return bytes, nil
}

func (s *Service) WatermarkImageWithImage(ctx context.Context, src []byte, watermark []byte) ([]byte, *errors.Err) {
	reader1 := bytes.NewReader(src)
	img1, err := png.Decode(reader1)
	if err != nil {
		return nil, errors.NewError(errors.KindMalformedStructure, err.Error())
	}

	reader2 := bytes.NewReader(watermark)
	img2, err := png.Decode(reader2)
	if err != nil {
		return nil, errors.NewError(errors.KindMalformedStructure, err.Error())
	}

	out, err1 := watermarking.WatermarkImageWithImage(ctx, img1, img2)
	if err1 != nil {
		return nil, err1
	}
	buf := new(bytes.Buffer)
	err = png.Encode(buf, out)
	bytes := buf.Bytes()
	return bytes, nil
}
