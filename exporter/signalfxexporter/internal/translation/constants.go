// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package translation // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/translation"

const (
	// DefaultTranslationRulesYaml defines default translation rules that will be applied to metrics if
	// config.TranslationRules not specified explicitly.
	// Keep it in YAML format to be able to easily copy and paste it in config if modifications needed.
	DefaultTranslationRulesYaml = `
translation_rules:
# drops opencensus.resourcetype dimension from metrics generated by receivers written
# using OC data structures. This rule can be removed once the k8s_cluster and kubeletstats
# receivers have been refactored to use pmetric.Metrics. These dimensions are added as a
# result of the conversion here https://github.com/open-telemetry/opentelemetry-collector/blob/v0.22.0/translator/internaldata/oc_to_resource.go#L128.
# Dropping these dimensions will ensure MTSes aren't broken when the receivers are
# refactored and this resource type dimension will cease to exist. 
- action: drop_dimensions
  metric_name: /^(k8s\.|container\.).*/
  dimension_pairs:
    opencensus.resourcetype:
      k8s: true
      container: true

- action: rename_metrics
  mapping:
    # kubeletstats container cpu needed for calculation below
    container.cpu.time: sf_temp.container_cpu_utilization

# compute cpu utilization metrics: cpu.utilization_per_core (excluded by default) and cpu.utilization
- action: delta_metric
  mapping:
    system.cpu.time: sf_temp.system.cpu.delta
- action: copy_metrics
  mapping:
    sf_temp.system.cpu.delta: sf_temp.system.cpu.usage
  dimension_key: state
  dimension_values:
    interrupt: true
    nice: true
    softirq: true
    steal: true
    system: true
    user: true
    wait: true
- action: aggregate_metric
  metric_name: sf_temp.system.cpu.usage
  aggregation_method: sum
  without_dimensions:
    - state
- action: copy_metrics
  mapping:
    sf_temp.system.cpu.delta: sf_temp.system.cpu.total
- action: aggregate_metric
  metric_name: sf_temp.system.cpu.total
  aggregation_method: sum
  without_dimensions:
    - state
- action: calculate_new_metric
  metric_name: cpu.utilization_per_core
  operand1_metric: sf_temp.system.cpu.usage
  operand2_metric: sf_temp.system.cpu.total
  operator: /
- action: copy_metrics
  mapping:
    cpu.utilization_per_core: sf_temp.cpu.utilization
- action: aggregate_metric
  metric_name: sf_temp.cpu.utilization
  aggregation_method: avg
  without_dimensions:
    - cpu
- action: multiply_float
  scale_factors_float:
    sf_temp.cpu.utilization: 100

# convert cpu metrics
- action: split_metric
  metric_name: system.cpu.time
  dimension_key: state
  mapping:
    idle: sf_temp.cpu.idle
    interrupt: sf_temp.cpu.interrupt
    system: sf_temp.cpu.system
    user: sf_temp.cpu.user
    steal: sf_temp.cpu.steal
    wait: sf_temp.cpu.wait
    softirq: sf_temp.cpu.softirq
    nice: sf_temp.cpu.nice
- action: multiply_float
  scale_factors_float:
    sf_temp.container_cpu_utilization: 100
    sf_temp.cpu.idle: 100
    sf_temp.cpu.interrupt: 100
    sf_temp.cpu.system: 100
    sf_temp.cpu.user: 100
    sf_temp.cpu.steal: 100
    sf_temp.cpu.wait: 100
    sf_temp.cpu.softirq: 100
    sf_temp.cpu.nice: 100
- action: convert_values
  types_mapping:
    sf_temp.container_cpu_utilization: int
    sf_temp.cpu.idle: int
    sf_temp.cpu.interrupt: int
    sf_temp.cpu.system: int
    sf_temp.cpu.user: int
    sf_temp.cpu.steal: int
    sf_temp.cpu.wait: int
    sf_temp.cpu.softirq: int
    sf_temp.cpu.nice: int

# compute cpu.num_processors
- action: copy_metrics
  mapping:
    sf_temp.cpu.idle: sf_temp.cpu.num_processors
- action: aggregate_metric
  metric_name: sf_temp.cpu.num_processors
  aggregation_method: count
  without_dimensions:
    - cpu

- action: copy_metrics
  mapping:
    sf_temp.cpu.idle: sf_temp.cpu.idle_per_core
    sf_temp.cpu.interrupt: sf_temp.cpu.interrupt_per_core
    sf_temp.cpu.system: sf_temp.cpu.system_per_core
    sf_temp.cpu.user: sf_temp.cpu.user_per_core
    sf_temp.cpu.wait: sf_temp.cpu.wait_per_core
    sf_temp.cpu.steal: sf_temp.cpu.steal_per_core
    sf_temp.cpu.softirq: sf_temp.cpu.softirq_per_core
    sf_temp.cpu.nice: sf_temp.cpu.nice_per_core

- action: aggregate_metric
  metric_name: sf_temp.cpu.idle
  aggregation_method: sum
  without_dimensions:
    - cpu
- action: aggregate_metric
  metric_name: sf_temp.cpu.interrupt
  aggregation_method: sum
  without_dimensions:
    - cpu
- action: aggregate_metric
  metric_name: sf_temp.cpu.system
  aggregation_method: sum
  without_dimensions:
    - cpu
- action: aggregate_metric
  metric_name: sf_temp.cpu.user
  aggregation_method: sum
  without_dimensions:
    - cpu
- action: aggregate_metric
  metric_name: sf_temp.cpu.steal
  aggregation_method: sum
  without_dimensions:
    - cpu
- action: aggregate_metric
  metric_name: sf_temp.cpu.wait
  aggregation_method: sum
  without_dimensions:
    - cpu
- action: aggregate_metric
  metric_name: sf_temp.cpu.softirq
  aggregation_method: sum
  without_dimensions:
    - cpu
- action: aggregate_metric
  metric_name: sf_temp.cpu.nice
  aggregation_method: sum
  without_dimensions:
    - cpu

# compute memory.total
- action: copy_metrics
  mapping:
    system.memory.usage: sf_temp.memory.total
  dimension_key: state
  dimension_values:
    buffered: true
    cached: true
    free: true
    used: true
- action: aggregate_metric
  metric_name: sf_temp.memory.total
  aggregation_method: sum
  without_dimensions:
    - state

# convert memory metrics
- action: copy_metrics
  mapping:
    system.memory.usage: sf_temp.system.memory.usage

# sf_temp.memory.used needed to calculate memory.utilization
- action: split_metric
  metric_name: sf_temp.system.memory.usage
  dimension_key: state
  mapping:
    used: sf_temp.memory.used

# Translations to derive filesystem metrics
## sf_temp.disk.total, required to compute disk.utilization
- action: copy_metrics
  mapping:
    system.filesystem.usage: sf_temp.disk.total
- action: aggregate_metric
  metric_name: sf_temp.disk.total
  aggregation_method: sum
  without_dimensions:
    - state

## sf_temp.disk.summary_total, required to compute disk.summary_utilization
- action: copy_metrics
  mapping:
    system.filesystem.usage: sf_temp.disk.summary_total
- action: aggregate_metric
  metric_name: sf_temp.disk.summary_total
  aggregation_method: avg
  without_dimensions:
    - mode
    - mountpoint
- action: aggregate_metric
  metric_name: sf_temp.disk.summary_total
  aggregation_method: sum
  without_dimensions:
    - state
    - device
    - type

## sf_temp.df_complex.used needed to calculate disk.utilization
- action: copy_metrics
  mapping:
    system.filesystem.usage: sf_temp.system.filesystem.usage

- action: split_metric
  metric_name: sf_temp.system.filesystem.usage
  dimension_key: state
  mapping:
    used: sf_temp.df_complex.used

## disk.utilization
- action: calculate_new_metric
  metric_name: sf_temp.disk.utilization
  operand1_metric: sf_temp.df_complex.used
  operand2_metric: sf_temp.disk.total
  operator: /
- action: multiply_float
  scale_factors_float:
    sf_temp.disk.utilization: 100

## disk.summary_utilization
- action: copy_metrics
  mapping:
    sf_temp.df_complex.used: sf_temp.df_complex.used_total

- action: aggregate_metric
  metric_name: sf_temp.df_complex.used_total
  aggregation_method: avg
  without_dimensions:
    - mode
    - mountpoint

- action: aggregate_metric
  metric_name: sf_temp.df_complex.used_total
  aggregation_method: sum
  without_dimensions:
    - device
    - type

- action: calculate_new_metric
  metric_name: sf_temp.disk.summary_utilization
  operand1_metric: sf_temp.df_complex.used_total
  operand2_metric: sf_temp.disk.summary_total
  operator: /
- action: multiply_float
  scale_factors_float:
    sf_temp.disk.summary_utilization: 100


# Translations to derive disk I/O metrics.
- action: copy_metrics
  mapping:
    system.disk.io.read: sf_temp.system.disk.io.read
- action: copy_metrics
  mapping:
    system.disk.io.write: sf_temp.system.disk.io.write
- action: calculate_new_metric
  metric_name: sf_temp.system.disk.io.total
  operand1_metric: sf_temp.system.disk.io.read
  operand2_metric: sf_temp.system.disk.io.write
  operator: +

## Calculate extra system.disk.operations.total and system.disk.io.total metrics summing up read/write ops/IO across all devices.
- action: copy_metrics
  mapping:
    system.disk.operations.read: sf_temp.system.disk.operations.read
    system.disk.operations.write: sf_temp.system.disk.operations.write
- action: calculate_new_metric
  metric_name: sf_temp.system.disk.operations.total
  operand1_metric: sf_temp.system.disk.operations.read
  operand2_metric: sf_temp.system.disk.operations.write
  operator: +
- action: copy_metrics
  mapping:
    sf_temp.system.disk.operations.total: sf_temp.disk.ops
- action: aggregate_metric
  metric_name: sf_temp.system.disk.operations.total
  aggregation_method: sum
  without_dimensions:
    - device
- action: aggregate_metric
  metric_name: sf_temp.system.disk.io.total
  aggregation_method: sum
  without_dimensions:
    - device

## Calculate an extra disk_ops.total metric as number all all read and write operations happened since the last report.
- action: copy_metrics
  mapping:
    system.disk.operations: sf_temp.disk.ops
- action: aggregate_metric
  metric_name: sf_temp.disk.ops
  aggregation_method: sum
  without_dimensions:
    - direction
    - device
- action: delta_metric
  mapping:
    sf_temp.disk.ops: disk_ops.total

- action: delta_metric
  mapping:
    system.disk.pending_operations: disk_ops.pending

# Translations to derive Network I/O metrics.

## Calculate extra network I/O metrics system.network.packets.total and system.network.io.total.
- action: copy_metrics
  mapping:
    system.network.packets.receive: sf_temp.system.network.packets.receive
    system.network.packets.transmit: sf_temp.system.network.packets.transmit
- action: calculate_new_metric
  metric_name: sf_temp.system.network.packets.total
  operand1_metric: sf_temp.system.network.packets.receive
  operand2_metric: sf_temp.system.network.packets.transmit
  operator: +
- action: copy_metrics
  mapping:
    system.network.io.receive: sf_temp.system.network.io.receive
    system.network.io.transmit: sf_temp.system.network.io.transmit
- action: calculate_new_metric
  metric_name: sf_temp.system.network.io
  operand1_metric: sf_temp.system.network.io.receive
  operand2_metric: sf_temp.system.network.io.transmit
  operator: +

- action: copy_metrics
  mapping:
    sf_temp.system.network.io: sf_temp.system.network.io.total
- action: aggregate_metric
  metric_name: sf_temp.system.network.packets.total
  aggregation_method: sum
  without_dimensions:
    - device
- action: aggregate_metric
  metric_name: sf_temp.system.network.io.total
  aggregation_method: sum
  without_dimensions:
    - device

## Calculate extra network.total metric.
- action: copy_metrics
  mapping:
    sf_temp.system.network.io: sf_temp.network.total
  dimension_key: direction
  dimension_values:
    receive: true
    transmit: true
- action: aggregate_metric
  metric_name: sf_temp.network.total
  aggregation_method: sum
  without_dimensions:
    - direction
    - device

# memory utilization
- action: calculate_new_metric
  metric_name: sf_temp.memory.utilization
  operand1_metric: sf_temp.memory.used
  operand2_metric: sf_temp.memory.total
  operator: /

- action: multiply_float
  scale_factors_float:
    sf_temp.memory.utilization: 100

# Virtual memory metrics
- action: copy_metrics
  mapping:
    system.paging.operations.in: sf_temp.system.paging.operations.in
    system.paging.operations.out: sf_temp.system.paging.operations.out

- action: split_metric
  metric_name: sf_temp.system.paging.operations.in
  dimension_key: type
  mapping:
    major: vmpage_io.swap.in
    minor: vmpage_io.memory.in

- action: split_metric
  metric_name: sf_temp.system.paging.operations.out
  dimension_key: type
  mapping:
    major: vmpage_io.swap.out
    minor: vmpage_io.memory.out

# process metric
- action: copy_metrics
  mapping:
    process.cpu.time: sf_temp.process.cpu.time
  dimension_key: state
  dimension_values:
    user: true
    system: true

- action: aggregate_metric
  metric_name: sf_temp.process.cpu.time
  aggregation_method: sum
  without_dimensions:
    - state

- action: rename_metrics
  mapping:
    sf_temp.container_cpu_utilization: container_cpu_utilization
    sf_temp.cpu.idle: cpu.idle
    sf_temp.cpu.idle_per_core: cpu.idle
    sf_temp.cpu.interrupt: cpu.interrupt
    sf_temp.cpu.interrupt_per_core: cpu.interrupt
    sf_temp.cpu.nice: cpu.nice
    sf_temp.cpu.nice_per_core: cpu.nice
    sf_temp.cpu.num_processors: cpu.num_processors
    sf_temp.cpu.softirq: cpu.softirq
    sf_temp.cpu.softirq_per_core: cpu.softirq
    sf_temp.cpu.steal: cpu.steal
    sf_temp.cpu.steal_per_core: cpu.steal
    sf_temp.cpu.system: cpu.system
    sf_temp.cpu.system_per_core: cpu.system
    sf_temp.cpu.user: cpu.user
    sf_temp.cpu.user_per_core: cpu.user
    sf_temp.cpu.utilization: cpu.utilization
    sf_temp.cpu.wait: cpu.wait
    sf_temp.cpu.wait_per_core: cpu.wait
    sf_temp.disk.summary_utilization: disk.summary_utilization
    sf_temp.disk.utilization: disk.utilization
    sf_temp.memory.total: memory.total
    sf_temp.memory.utilization: memory.utilization
    sf_temp.network.total: network.total
    sf_temp.system.disk.io.total: system.disk.io.total
    sf_temp.system.disk.operations.total: system.disk.operations.total
    sf_temp.system.network.io.total: system.network.io.total
    sf_temp.system.network.packets.total: system.network.packets.total
    sf_temp.process.cpu.time: process.cpu_time_seconds

# remove redundant metrics
- action: drop_metrics
  metric_names:
    sf_temp.df_complex.used: true
    sf_temp.df_complex.used_total: true
    sf_temp.disk.ops: true
    sf_temp.disk.summary_total: true
    sf_temp.disk.total: true
    sf_temp.memory.used: true
    sf_temp.system.cpu.delta: true
    sf_temp.system.cpu.total: true
    sf_temp.system.cpu.usage: true
    sf_temp.system.disk.io.read: true
    sf_temp.system.disk.io.write: true
    sf_temp.system.disk.operations.read: true
    sf_temp.system.disk.operations.write: true
    sf_temp.system.filesystem.usage: true
    sf_temp.system.memory.usage: true
    sf_temp.system.network.io: true
    sf_temp.system.network.io.receive : true
    sf_temp.system.network.io.transmit: true
    sf_temp.system.network.packets.receive: true
    sf_temp.system.network.packets.transmit: true
    sf_temp.system.paging.operations.in: true
    sf_temp.system.paging.operations.out: true
`
)
