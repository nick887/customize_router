package customize_router

import (
	"github.com/Chronostasys/centralog/centralog"
	"github.com/nick887/customize_router/custom"
	"github.com/nick887/customize_router/log"
	"go.uber.org/zap"
)

// @author Ritchie
// @version 1.0, 2021/10/5
// @since 1.0.0


func InitCustomRouter(logOpt *centralog.LogOptions,cr *custom.CustmoRouter) {
	err := centralog.InitLoggerWithOpt(zap.NewProductionConfig(), logOpt)
	if err != nil {
		panic(err)
	}
	cr.Use(log.Log())
}