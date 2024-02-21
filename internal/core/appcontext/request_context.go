package appcontext

import (
	coremodels "backend/internal/core/models"
	"context"
)

type requestCtxType string

const requestCtxKey requestCtxType = "requestCtx"

func WithValue(ctx context.Context, info *coremodels.RequestContext) context.Context {
	return context.WithValue(ctx, requestCtxKey, info)
}

func Value(ctx context.Context) *coremodels.RequestContext {
	v, ok := ctx.Value(requestCtxKey).(*coremodels.RequestContext)
	if ok {
		return v
	}
	return nil
}

func GetRequestId(ctx context.Context) string {
	v := Value(ctx)
	if v == nil {
		return ""
	}
	return v.RequestId
}
