package main

import (
	"github.com/opentracing/opentracing-go"
	"go.elastic.co/apm/module/apmot"
)

// InitTracer initialises the opentracing tracer.
func InitTracer(opts []string) (opentracing.Tracer, error) {
	return apmot.New(), nil
}