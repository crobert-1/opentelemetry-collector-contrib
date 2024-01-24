// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package alibabacloudlogserviceexporter

import (
	"context"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"testing"
)

func TestNewMetricsExporter(t *testing.T) {
	set := exportertest.NewNopCreateSettings()
	l := &logServiceMetricsSender{
		logger: set.Logger,
	}

	got, err := newMetricsExporterWithLog(l, set, &Config{
		Endpoint: "us-west-1.log.aliyuncs.com",
		Project:  "demo-project",
		Logstore: "demo-logstore",
	})

	assert.NoError(t, err)
	require.NotNil(t, got)
	defer func() {
		assert.NoError(t, got.Shutdown(context.Background()))
	}()

	// This will put trace data to send buffer and return success.
	err = got.ConsumeMetrics(context.Background(), testdata.GenerateMetricsOneMetric())
	assert.NoError(t, err)

	// Since it's an invalid endpoint, we should block here until the `Fail` callback is complete.
	// pushMetricsData sends the request asynchronously, so the leaking goroutine is just the
	//require.Eventually(t, func() bool {
	//	l.logger
	//}, 5*time.Second, 100*time.Millisecond)
}

func TestNewFailsWithEmptyMetricsExporterName(t *testing.T) {
	got, err := newMetricsExporter(exportertest.NewNopCreateSettings(), &Config{})
	assert.Error(t, err)
	require.Nil(t, got)
}
