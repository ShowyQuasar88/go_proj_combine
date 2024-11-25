package pkg

import (
	"backend/internal/pkg/auth"
	"backend/internal/pkg/crypto"
	"github.com/google/wire"
)

// ProviderSet 提供基础设施层的依赖
var ProviderSet = wire.NewSet(
	crypto.NewCrypto,
	auth.NewJWTHelper,
	auth.NewJWTConfig,
)
