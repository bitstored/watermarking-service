package watermarking

import (
	"bytes"
	"context"
	"google.golang.org/grpc"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"

	"github.com/bitstored/watermarking-service/pkg/errors"
	"github.com/bitstored/watermarking-service/pkg/watermarking/client"
)

const gRPCPortImgTransform = "localhost:50051"

func WatermarkImageWithImage(ctx context.Context, src image.Image, watermarkImage image.Image) (image.Image, *errors.Err) {
	if src == nil {
		return nil, errors.NewError(errors.KindEmptyImage, "image can't be empty")
	}
	if watermarkImage == nil {
		return nil, errors.NewError(errors.KindEmptyImage, "watermark image  can't be empty")
	}

	src1 := imageToRGBA(src)
	watermarkImage1 := imageToRGBA(watermarkImage)

	out, err := watermark(ctx, src1, watermarkImage1)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func WatermarkImageWithText(ctx context.Context, src image.Image, watermarkText string) (image.Image, *errors.Err) {
	img, err := createImage(ctx, watermarkText)

	if err != nil {
		return nil, err
	}

	return WatermarkImageWithImage(ctx, src, img)
}

func createImage(ctx context.Context, text string) (*image.RGBA, *errors.Err) {
	conn, err := grpc.Dial(gRPCPortImgTransform, grpc.WithInsecure())
	if err != nil {
		return nil, errors.NewError(errors.KindConnectionFailed, err.Error())
	}
	c := client.NewTransformationClient(conn)
	req := client.ToImageRequest{Text: text}
	rsp, err := c.ToImage(ctx, &req)
	if err != nil {
		return nil, errors.NewError(errors.KindInvalidArgument, err.Error())
	}
	reader := bytes.NewReader(rsp.Image)
	img, err := png.Decode(reader)
	if err != nil {
		return nil, errors.NewError(errors.KindMalformedStructure, err.Error())
	}

	return imageToRGBA(img), nil
}

func watermark(ctx context.Context, img *image.RGBA, watermark *image.RGBA) (*image.RGBA, *errors.Err) {
	ib := img.Bounds()
	wb := watermark.Bounds()

	if ib.Dx() < wb.Dx() {
		watermark = watermark.SubImage(image.Rect((wb.Dx()-ib.Dx())/2, 0, wb.Dx()-(wb.Dx()-ib.Dx())/2, wb.Dy())).(*image.RGBA)
	}
	if ib.Dy() < wb.Dy() {
		watermark = watermark.SubImage(image.Rect(0, (wb.Dy()-ib.Dy())/2, wb.Dx(), wb.Dy()-(wb.Dy()-ib.Dy())/2)).(*image.RGBA)
	}

	wb = watermark.Bounds()

	for x := 0; x < ib.Dx(); x += wb.Dx() {
		for y := 0; y < ib.Dy(); y += wb.Dy() {
			if ib.Dx() >= x+wb.Dx() && ib.Dy() >= y+wb.Dy() {
				for i := 0; i < wb.Dx(); i++ {
					for j := 0; j < wb.Dy(); j++ {
						dr, dg, db, da := img.At(x+i, y+j).RGBA()
						sr, sg, sb, sa := watermark.At(i, j).RGBA()
						if sa != 0 {
							da = math.MaxUint8
							dr = sr
							dg = sg
							db = sb
						}
						img.Set(x+i, y+j, color.RGBA{uint8(dr), uint8(dg), uint8(db), uint8(da)})
					}
				}
			}
		}
	}
	return img, nil
}

func imageToRGBA(src image.Image) *image.RGBA {
	if src == nil {
		return nil
	}
	bounds := src.Bounds()

	var m *image.RGBA
	var width, height int

	width = bounds.Dx()
	height = bounds.Dy()

	m = image.NewRGBA(image.Rect(0, 0, width, height))

	draw.Draw(m, m.Bounds(), src, bounds.Min, draw.Src)
	return m
}
