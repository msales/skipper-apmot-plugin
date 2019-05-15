package main

import (
	"testing"

	"github.com/zalando/skipper/tracing"
)

const pluginDir = "."

func TestLoadPlugin(t *testing.T) {
	if _, err := tracing.LoadTracingPlugin([]string{pluginDir}, []string{"apmot", "token=123456"}); err != nil {
		t.Errorf("failed to load plugin `apmot`: %s", err)
	}
}