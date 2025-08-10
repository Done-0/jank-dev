package wire

import (
	"github.com/google/wire"

	mapperImpl "github.com/Done-0/jank/pkg/serve/mapper/impl"
	serviceImpl "github.com/Done-0/jank/pkg/serve/service/impl"
)

// MapperProviderSet 数据访问层相关的 Provider 集合
var MapperProviderSet = wire.NewSet(
	mapperImpl.NewRBACMapper,
)

// ServiceProviderSet 服务相关的 Provider 集合
var ServiceProviderSet = wire.NewSet(
	serviceImpl.NewPluginService,
	serviceImpl.NewRBACService,
	serviceImpl.NewThemeService,
)

// AllProviderSet 所有 Provider 的集合
var AllProviderSet = wire.NewSet(
	MapperProviderSet,
	ServiceProviderSet,
)
