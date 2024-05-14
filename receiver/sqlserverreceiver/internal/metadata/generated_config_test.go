// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestMetricsBuilderConfig(t *testing.T) {
	tests := []struct {
		name string
		want MetricsBuilderConfig
	}{
		{
			name: "default",
			want: DefaultMetricsBuilderConfig(),
		},
		{
			name: "all_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					SqlserverBatchRequestRate:                   MetricConfig{Enabled: true},
					SqlserverBatchSQLCompilationRate:            MetricConfig{Enabled: true},
					SqlserverBatchSQLRecompilationRate:          MetricConfig{Enabled: true},
					SqlserverDatabaseIoReadLatency:              MetricConfig{Enabled: true},
					SqlserverLockWaitRate:                       MetricConfig{Enabled: true},
					SqlserverLockWaitTimeAvg:                    MetricConfig{Enabled: true},
					SqlserverPageBufferCacheHitRatio:            MetricConfig{Enabled: true},
					SqlserverPageCheckpointFlushRate:            MetricConfig{Enabled: true},
					SqlserverPageLazyWriteRate:                  MetricConfig{Enabled: true},
					SqlserverPageLifeExpectancy:                 MetricConfig{Enabled: true},
					SqlserverPageOperationRate:                  MetricConfig{Enabled: true},
					SqlserverPageSplitRate:                      MetricConfig{Enabled: true},
					SqlserverProcessesBlocked:                   MetricConfig{Enabled: true},
					SqlserverPropertiesDbStatus:                 MetricConfig{Enabled: true},
					SqlserverResourcePoolDiskThrottledReadRate:  MetricConfig{Enabled: true},
					SqlserverResourcePoolDiskThrottledWriteRate: MetricConfig{Enabled: true},
					SqlserverTransactionRate:                    MetricConfig{Enabled: true},
					SqlserverTransactionWriteRate:               MetricConfig{Enabled: true},
					SqlserverTransactionLogFlushDataRate:        MetricConfig{Enabled: true},
					SqlserverTransactionLogFlushRate:            MetricConfig{Enabled: true},
					SqlserverTransactionLogFlushWaitRate:        MetricConfig{Enabled: true},
					SqlserverTransactionLogGrowthCount:          MetricConfig{Enabled: true},
					SqlserverTransactionLogShrinkCount:          MetricConfig{Enabled: true},
					SqlserverTransactionLogUsage:                MetricConfig{Enabled: true},
					SqlserverUserConnectionCount:                MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					SqlserverComputerName: ResourceAttributeConfig{Enabled: true},
					SqlserverDatabaseName: ResourceAttributeConfig{Enabled: true},
					SqlserverInstanceName: ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					SqlserverBatchRequestRate:                   MetricConfig{Enabled: false},
					SqlserverBatchSQLCompilationRate:            MetricConfig{Enabled: false},
					SqlserverBatchSQLRecompilationRate:          MetricConfig{Enabled: false},
					SqlserverDatabaseIoReadLatency:              MetricConfig{Enabled: false},
					SqlserverLockWaitRate:                       MetricConfig{Enabled: false},
					SqlserverLockWaitTimeAvg:                    MetricConfig{Enabled: false},
					SqlserverPageBufferCacheHitRatio:            MetricConfig{Enabled: false},
					SqlserverPageCheckpointFlushRate:            MetricConfig{Enabled: false},
					SqlserverPageLazyWriteRate:                  MetricConfig{Enabled: false},
					SqlserverPageLifeExpectancy:                 MetricConfig{Enabled: false},
					SqlserverPageOperationRate:                  MetricConfig{Enabled: false},
					SqlserverPageSplitRate:                      MetricConfig{Enabled: false},
					SqlserverProcessesBlocked:                   MetricConfig{Enabled: false},
					SqlserverPropertiesDbStatus:                 MetricConfig{Enabled: false},
					SqlserverResourcePoolDiskThrottledReadRate:  MetricConfig{Enabled: false},
					SqlserverResourcePoolDiskThrottledWriteRate: MetricConfig{Enabled: false},
					SqlserverTransactionRate:                    MetricConfig{Enabled: false},
					SqlserverTransactionWriteRate:               MetricConfig{Enabled: false},
					SqlserverTransactionLogFlushDataRate:        MetricConfig{Enabled: false},
					SqlserverTransactionLogFlushRate:            MetricConfig{Enabled: false},
					SqlserverTransactionLogFlushWaitRate:        MetricConfig{Enabled: false},
					SqlserverTransactionLogGrowthCount:          MetricConfig{Enabled: false},
					SqlserverTransactionLogShrinkCount:          MetricConfig{Enabled: false},
					SqlserverTransactionLogUsage:                MetricConfig{Enabled: false},
					SqlserverUserConnectionCount:                MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					SqlserverComputerName: ResourceAttributeConfig{Enabled: false},
					SqlserverDatabaseName: ResourceAttributeConfig{Enabled: false},
					SqlserverInstanceName: ResourceAttributeConfig{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricConfig{}, ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}

func TestResourceAttributesConfig(t *testing.T) {
	tests := []struct {
		name string
		want ResourceAttributesConfig
	}{
		{
			name: "default",
			want: DefaultResourceAttributesConfig(),
		},
		{
			name: "all_set",
			want: ResourceAttributesConfig{
				SqlserverComputerName: ResourceAttributeConfig{Enabled: true},
				SqlserverDatabaseName: ResourceAttributeConfig{Enabled: true},
				SqlserverInstanceName: ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				SqlserverComputerName: ResourceAttributeConfig{Enabled: false},
				SqlserverDatabaseName: ResourceAttributeConfig{Enabled: false},
				SqlserverInstanceName: ResourceAttributeConfig{Enabled: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadResourceAttributesConfig(t *testing.T, name string) ResourceAttributesConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	sub, err = sub.Sub("resource_attributes")
	require.NoError(t, err)
	cfg := DefaultResourceAttributesConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
