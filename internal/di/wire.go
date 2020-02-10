// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"rpcdemo/internal/dao"
	"rpcdemo/internal/service"
	"rpcdemo/internal/server/grpc"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, grpc.New, NewApp))
}
