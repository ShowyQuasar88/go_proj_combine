package service

import (
	"context"
	"math/rand"
	pb "verify_code/api/v1"
)

type VerifyCodeService struct {
	pb.UnimplementedVerifyCodeServer
}

func NewVerifyCodeService() *VerifyCodeService {
	return &VerifyCodeService{}
}

func (s *VerifyCodeService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeRequest) (*pb.GetVerifyCodeReply, error) {
	return &pb.GetVerifyCodeReply{
		Code: GenerateRandCode(req.Length, req.Type),
	}, nil
}

func GenerateRandCode(length int32, codeType pb.TYPE) string {
	switch codeType {
	case pb.TYPE_DEFAULT, pb.TYPE_DIGIT:
		return generateRandCode(length, "0123456789")
	case pb.TYPE_LETTER:
		return generateRandCode(length, "abcdefghijklnmopqrstuvwxyz")
	case pb.TYPE_MIXED:
		return generateRandCode(length, "0123456789abcdefghijklnmopqrstuvwxyz")
	default:
		return generateRandCode(length, "0123456789")
	}
}

func generateRandCode(length int32, codeType string) string {
	strLen := len(codeType)
	result := make([]byte, length)
	for i := 0; i < int(length); i++ {
		randIdx := rand.Intn(strLen)
		result[i] = codeType[randIdx]
	}
	return string(result)
}
