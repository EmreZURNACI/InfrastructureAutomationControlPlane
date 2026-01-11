package main

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/adaptor/client"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/infra/postgres"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/log"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/server"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

func init() {
	log.Logger.Info("Application is starting")
}

func main() {

	db, err := postgres.Connection()
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	if err := db.PrepareDB(); err != nil {
		log.Logger.Error(err.Error())
		return
	}

	ec2Client, err := client.ConnectEC2()
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	tp, err := initTracer()
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	server.Start(server.Dependencies{
		EC2Client: ec2Client,
		DB:        db,
		Tp:        tp,
	})

}
func initTracer() (trace.Tracer, error) {
	headers := map[string]string{
		"content-type": "application/json",
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint("tracing:4318"),
			otlptracehttp.WithURLPath("/v1/traces"),
			otlptracehttp.WithInsecure(),
			otlptracehttp.WithHeaders(headers),
		),
	)
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("InfrastructureAutomationControlPlane"),
			)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp.Tracer("InfrastructureAutomationControlPlane"), nil
}
