package watermarking

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"testing"

	"github.com/bitstored/watermarking-service/pkg/watermarking/client"
)

func TestCreateImageFromText(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fail()
	}
	c := client.NewTransformationClient(conn)
	req := client.ToImageRequest{Text: "Hello world!"}
	rsp, err := c.ToImage(context.Background(), &req)
	require.NoError(t, err)
	reader := bytes.NewReader(rsp.Image)
	img1, err := png.Decode(reader)
	require.NoError(t, err)

	imgfile, err := os.Open("test_data/hello_world.png")
	require.NoError(t, err)

	defer imgfile.Close()
	img2, _, err := image.Decode(imgfile)
	require.NoError(t, err)

	for x := 0; x < img1.Bounds().Dx(); x++ {
		for y := 0; y < img1.Bounds().Dy(); y++ {
			require.Equal(t, img2.At(x, y), img1.At(x, y))
		}
	}
	f, err := os.Create("test_data/hello_world_test.png")
	require.NoError(t, err)

	defer f.Close()
	png.Encode(f, img1)
}

func TestWatermarkWithImage(t *testing.T) {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	imgfile, err := os.Open("test_data/test_no.jpg")
	require.NoError(t, err)
	defer imgfile.Close()

	watermarkfile, err := os.Open("test_data/hello_world.png")
	require.NoError(t, err)
	defer watermarkfile.Close()

	src, _, err := image.Decode(imgfile)
	watermark, _, err := image.Decode(watermarkfile)
	require.NoError(t, err)

	out, err1 := WatermarkImageWithImage(context.Background(), imageToRGBA(src), imageToRGBA(watermark))

	require.Nil(t, err1)

	f, err := os.Create("test_data/test_yes_img.png")
	require.NoError(t, err)

	defer f.Close()
	png.Encode(f, out)
}

func TestWatermarkWithText(t *testing.T) {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	imgfile, err := os.Open("test_data/test_no.jpg")
	require.NoError(t, err)
	defer imgfile.Close()

	src, _, err := image.Decode(imgfile)
	require.NoError(t, err)

	out, err1 := WatermarkImageWithText(context.Background(), imageToRGBA(src), "bejan.diana.andrei@yahoo.com")

	require.Nil(t, err1.Error())

	f, err := os.Create("test_data/test_yes_text.png")
	require.NoError(t, err)

	defer f.Close()
	png.Encode(f, out)
}
