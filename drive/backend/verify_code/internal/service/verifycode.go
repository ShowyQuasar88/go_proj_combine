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
		return generateRandCode(length, 4, "0123456789")
	case pb.TYPE_LETTER:
		return generateRandCode(length, 5, "abcdefghijklnmopqrstuvwxyz")
	case pb.TYPE_MIXED:
		return generateRandCode(length, 6, "0123456789abcdefghijklnmopqrstuvwxyz")
	default:
		return generateRandCode(length, 4, "0123456789")
	}
}

// generateRandCode 优化实现方式
func generateRandCode(length, bits int32, s string) string {
	// 基于 s 的长度计算所需要的比特位【推荐写死】
	// 生成一个掩码
	bitsMask := 1<<bits - 1
	// 63 位可以用多少次
	idxMax := 63 / bits

	result := make([]byte, length)

	// cache 随机位缓存
	// remain 还可以用几次
	for i, cache, remain := int32(0), rand.Int63(), idxMax; i < length; {
		if remain == 0 {
			cache, remain = rand.Int63(), idxMax
		}
		// 利用掩码获取有效部位的随机下标
		if randIndex := int(cache & bitsMask); randIndex < len(s) {
			result[i] = s[randIndex]
			i++
		}
		remain--
		cache >>= bits
	}

	return string(result)
}
