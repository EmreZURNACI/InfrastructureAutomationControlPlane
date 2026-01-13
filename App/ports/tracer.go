package ports

import (
	"context"
)

type Tracer interface {
	Start(
		ctx context.Context,
		spanName string,
		opts ...SpanOption,
	) (context.Context, Span)
}
type Span interface {
	End()
}

type SpanOption func(*SpanConfig)

type SpanConfig struct {
	Attributes map[string]any
}
type NoopSpan struct{}

func (NoopSpan) End() {}
