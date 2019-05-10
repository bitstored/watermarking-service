package watermarking

import (
	"context"
	"image"

	"github.com/bitstored/watermarking-service/pkg/errors"
)

func WatermarkImageWithImage(ctx context.Context, src image.Image, watermark image.Image) (image.Image, *errors.Err) {
	return nil, nil
}
