// infra/otel/tracer.go
package tracer

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type OtelTracer struct {
	tracer trace.Tracer
}

func NewTracer(tracer trace.Tracer) ports.Tracer {
	return &OtelTracer{
		tracer: tracer,
	}
}

type OtelSpan struct {
	span trace.Span
}

func (s *OtelSpan) End() {
	s.span.End()
}

func (t *OtelTracer) Start(
	ctx context.Context,
	name string,
	_ ...ports.SpanOption,
) (context.Context, ports.Span) {

	if t == nil || t.tracer == nil {
		return ctx, ports.NoopSpan{}
	}

	ctx, span := t.tracer.Start(ctx, name)

	if !span.IsRecording() {
		return ctx, ports.NoopSpan{}
	}

	return ctx, &OtelSpan{span: span}
}
