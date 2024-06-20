// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	conventions "go.opentelemetry.io/collector/semconv/v1.9.0"
)

// AttributeStatus specifies the a value status attribute.
type AttributeStatus int

const (
	_ AttributeStatus = iota
	AttributeStatusBlocked
	AttributeStatusDaemon
	AttributeStatusDetached
	AttributeStatusIdle
	AttributeStatusLocked
	AttributeStatusOrphan
	AttributeStatusPaging
	AttributeStatusRunning
	AttributeStatusSleeping
	AttributeStatusStopped
	AttributeStatusSystem
	AttributeStatusUnknown
	AttributeStatusZombies
)

// String returns the string representation of the AttributeStatus.
func (av AttributeStatus) String() string {
	switch av {
	case AttributeStatusBlocked:
		return "blocked"
	case AttributeStatusDaemon:
		return "daemon"
	case AttributeStatusDetached:
		return "detached"
	case AttributeStatusIdle:
		return "idle"
	case AttributeStatusLocked:
		return "locked"
	case AttributeStatusOrphan:
		return "orphan"
	case AttributeStatusPaging:
		return "paging"
	case AttributeStatusRunning:
		return "running"
	case AttributeStatusSleeping:
		return "sleeping"
	case AttributeStatusStopped:
		return "stopped"
	case AttributeStatusSystem:
		return "system"
	case AttributeStatusUnknown:
		return "unknown"
	case AttributeStatusZombies:
		return "zombies"
	}
	return ""
}

// MapAttributeStatus is a helper map of string to AttributeStatus attribute value.
var MapAttributeStatus = map[string]AttributeStatus{
	"blocked":  AttributeStatusBlocked,
	"daemon":   AttributeStatusDaemon,
	"detached": AttributeStatusDetached,
	"idle":     AttributeStatusIdle,
	"locked":   AttributeStatusLocked,
	"orphan":   AttributeStatusOrphan,
	"paging":   AttributeStatusPaging,
	"running":  AttributeStatusRunning,
	"sleeping": AttributeStatusSleeping,
	"stopped":  AttributeStatusStopped,
	"system":   AttributeStatusSystem,
	"unknown":  AttributeStatusUnknown,
	"zombies":  AttributeStatusZombies,
}

type metricSystemProcessesCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.processes.count metric with initial data.
func (m *metricSystemProcessesCount) init() {
	m.data.SetName("system.processes.count")
	m.data.SetDescription("Total number of processes in each state.")
	m.data.SetUnit("{processes}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemProcessesCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, statusAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("status", statusAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemProcessesCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemProcessesCount) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemProcessesCount(cfg MetricConfig) metricSystemProcessesCount {
	m := metricSystemProcessesCount{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemProcessesCreated struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.processes.created metric with initial data.
func (m *metricSystemProcessesCreated) init() {
	m.data.SetName("system.processes.created")
	m.data.SetDescription("Total number of created processes.")
	m.data.SetUnit("{processes}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricSystemProcessesCreated) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemProcessesCreated) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemProcessesCreated) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemProcessesCreated(cfg MetricConfig) metricSystemProcessesCreated {
	m := metricSystemProcessesCreated{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                       MetricsBuilderConfig // config of the metrics builder.
	startTime                    pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity              int                  // maximum observed number of metrics per resource.
	metricsBuffer                pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                    component.BuildInfo  // contains version information.
	metricSystemProcessesCount   metricSystemProcessesCount
	metricSystemProcessesCreated metricSystemProcessesCreated
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.Settings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:                       mbc,
		startTime:                    pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                pmetric.NewMetrics(),
		buildInfo:                    settings.BuildInfo,
		metricSystemProcessesCount:   newMetricSystemProcessesCount(mbc.Metrics.SystemProcessesCount),
		metricSystemProcessesCreated: newMetricSystemProcessesCreated(mbc.Metrics.SystemProcessesCreated),
	}

	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.SetSchemaUrl(conventions.SchemaURL)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/hostmetricsreceiver/processes")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricSystemProcessesCount.emit(ils.Metrics())
	mb.metricSystemProcessesCreated.emit(ils.Metrics())

	for _, op := range rmo {
		op(rm)
	}

	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordSystemProcessesCountDataPoint adds a data point to system.processes.count metric.
func (mb *MetricsBuilder) RecordSystemProcessesCountDataPoint(ts pcommon.Timestamp, val int64, statusAttributeValue AttributeStatus) {
	mb.metricSystemProcessesCount.recordDataPoint(mb.startTime, ts, val, statusAttributeValue.String())
}

// RecordSystemProcessesCreatedDataPoint adds a data point to system.processes.created metric.
func (mb *MetricsBuilder) RecordSystemProcessesCreatedDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricSystemProcessesCreated.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
