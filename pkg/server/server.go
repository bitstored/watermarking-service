package server

import (
	"context"

	"github.com/bitstored/watermarking-service/pb"
	"github.com/bitstored/watermarking-service/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	Service *service.Service
}

func NewServer(s *service.Service) *Server {
	return &Server{s}
}

func (s *Server) EncodeMessage(ctx context.Context, in *pb.EncodeMessageRequest) (*pb.EncodeMessageResponse, error) {
	img := in.GetImage()
	if img == nil || len(img) == 0 {
		return nil, status.Error(codes.InvalidArgument, "image is empty")
	}
	text := in.GetText()
	if text == nil || len(text) == 0 {
		return nil, status.Error(codes.InvalidArgument, "text is empty")
	}
	full := in.GetFullMessage()
	out, err := s.Service.EncodeMessage(ctx, img, text, full)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Message())
	}
	rsp := pb.EncodeMessageResponse{Image: out}
	return &rsp, nil
}

func (s *Server) DecodeMessage(ctx context.Context, in *pb.DecodeMessageRequest) (*pb.DecodeMessageResponse, error) {
	img := in.GetImage()
	if img == nil || len(img) == 0 {
		return nil, status.Error(codes.InvalidArgument, "image is empty")
	}
	out, text, err := s.Service.DecodeMessage(ctx, img)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Message())
	}
	rsp := pb.DecodeMessageResponse{Image: out, Text: text}
	return &rsp, nil
}

func (s *Server) WatermarkImageWithText(ctx context.Context, in *pb.WatermarkImageWithTextRequest) (*pb.WatermarkImageWithTextResponse, error) {
	img := in.GetImage()
	if img == nil || len(img) == 0 {
		return nil, status.Error(codes.InvalidArgument, "image is empty")
	}
	text := in.GetWatermark()
	if len([]byte(text)) == 0 {
		return nil, status.Error(codes.InvalidArgument, " watermark text is empty")
	}
	out, err := s.Service.WatermarkImageWithText(ctx, img, text)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Message())
	}
	rsp := pb.WatermarkImageWithTextResponse{Image: out}
	return &rsp, nil
}

func (s *Server) WatermarkImageWithImage(ctx context.Context, in *pb.WatermarkImageWithImageRequest) (*pb.WatermarkImageWithImageResponse, error) {
	img := in.GetImage()
	if img == nil || len(img) == 0 {
		return nil, status.Error(codes.InvalidArgument, "image is empty")
	}
	img2 := in.GetWatermark()
	if img2 == nil || len(img2) == 0 {
		return nil, status.Error(codes.InvalidArgument, "watermark image is empty")
	}
	out, err := s.Service.WatermarkImageWithImage(ctx, img, img2)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Message())
	}
	rsp := pb.WatermarkImageWithImageResponse{Image: out}
	return &rsp, nil
}
