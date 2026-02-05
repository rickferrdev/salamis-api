package handlers

import (
	"github.com/rickferrdev/salamis-api/internal/adapters/inbound/http/handlers/auth"
	"github.com/rickferrdev/salamis-api/internal/adapters/inbound/http/handlers/me"
	"github.com/rickferrdev/salamis-api/internal/adapters/inbound/http/handlers/middlewares"
	"go.uber.org/fx"
)

var Module = fx.Module("handlers", fx.Provide(
	auth.NewAuthHandler,
	me.NewMeHandler,
	middlewares.NewGuardMiddleware,
))
