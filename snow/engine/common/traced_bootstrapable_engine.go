// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"context"

	"go.opentelemetry.io/otel/attribute"

	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/lasthyphen/dijetsnodego/ids"
	"github.com/lasthyphen/dijetsnodego/trace"
)

var _ BootstrapableEngine = (*tracedBootstrapableEngine)(nil)

type tracedBootstrapableEngine struct {
	Engine
	bootstrapableEngine BootstrapableEngine
	tracer              trace.Tracer
}

func TraceBootstrapableEngine(bootstrapableEngine BootstrapableEngine, tracer trace.Tracer) BootstrapableEngine {
	return &tracedBootstrapableEngine{
		Engine:              TraceEngine(bootstrapableEngine, tracer),
		bootstrapableEngine: bootstrapableEngine,
	}
}

func (e *tracedBootstrapableEngine) ForceAccepted(ctx context.Context, acceptedContainerIDs []ids.ID) error {
	ctx, span := e.tracer.Start(ctx, "tracedBootstrapableEngine.ForceAccepted", oteltrace.WithAttributes(
		attribute.Int("numAcceptedContainerIDs", len(acceptedContainerIDs)),
	))
	defer span.End()

	return e.bootstrapableEngine.ForceAccepted(ctx, acceptedContainerIDs)
}

func (e *tracedBootstrapableEngine) Clear() error {
	return e.bootstrapableEngine.Clear()
}
