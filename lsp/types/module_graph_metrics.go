/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package types

import (
	"sync"
	"sync/atomic"
	"time"
)

// ModuleGraphMetrics tracks performance and usage metrics for the module graph
type ModuleGraphMetrics struct {
	// Counter metrics (atomic for thread safety)
	FilesProcessed         int64 // Total files processed
	ExportsFound           int64 // Total exports found
	DependenciesTracked    int64 // Total dependencies tracked
	TransitiveComputations int64 // Total transitive closure computations
	CacheHits              int64 // Cache hit count
	CacheMisses            int64 // Cache miss count
	ParseErrors            int64 // Parse error count

	// Performance timing metrics
	TotalBuildTime    time.Duration // Total time for BuildFromWorkspace
	AvgFileParseTime  time.Duration // Average time per file parse
	AvgTransitiveTime time.Duration // Average time for transitive computation

	// Memory usage metrics
	MaxMemoryUsage int64 // Peak memory usage (if measurable)

	// State metrics
	ActiveModules  int64 // Number of modules currently tracked
	ActiveElements int64 // Number of elements currently tracked
}

// DefaultMetricsCollector implements MetricsCollector with in-memory storage
type DefaultMetricsCollector struct {
	counters   sync.Map // map[string]*int64
	gauges     sync.Map // map[string]*int64
	durations  sync.Map // map[string]*durationList
	histograms sync.Map // map[string]*histogramList
}

// durationList provides thread-safe duration slice operations
type durationList struct {
	mu        sync.Mutex
	durations []time.Duration
}

// add appends a duration to the list in a thread-safe manner
func (dl *durationList) add(duration time.Duration) {
	dl.mu.Lock()
	defer dl.mu.Unlock()
	dl.durations = append(dl.durations, duration)
}

// histogramList provides thread-safe histogram slice operations
type histogramList struct {
	mu     sync.Mutex
	values []float64
}

// add appends a value to the histogram in a thread-safe manner
func (hl *histogramList) add(value float64) {
	hl.mu.Lock()
	defer hl.mu.Unlock()
	hl.values = append(hl.values, value)
}

// NewDefaultMetricsCollector creates a new in-memory metrics collector
func NewDefaultMetricsCollector() *DefaultMetricsCollector {
	return &DefaultMetricsCollector{}
}

// IncrementCounter implements MetricsCollector
func (m *DefaultMetricsCollector) IncrementCounter(name string) {
	counterInterface, _ := m.counters.LoadOrStore(name, new(int64))
	counter := counterInterface.(*int64)
	atomic.AddInt64(counter, 1)
}

// RecordDuration implements MetricsCollector
func (m *DefaultMetricsCollector) RecordDuration(name string, duration time.Duration) {
	durationsInterface, _ := m.durations.LoadOrStore(name, &durationList{durations: make([]time.Duration, 0)})
	durationsList := durationsInterface.(*durationList)
	durationsList.add(duration)
}

// SetGauge implements MetricsCollector
func (m *DefaultMetricsCollector) SetGauge(name string, value int64) {
	gaugeInterface, _ := m.gauges.LoadOrStore(name, new(int64))
	gauge := gaugeInterface.(*int64)
	atomic.StoreInt64(gauge, value)
}

// AddHistogramValue implements MetricsCollector
func (m *DefaultMetricsCollector) AddHistogramValue(name string, value float64) {
	histogramInterface, _ := m.histograms.LoadOrStore(name, &histogramList{values: make([]float64, 0)})
	histogramList := histogramInterface.(*histogramList)
	histogramList.add(value)
}

// GetCounterValue returns the current value of a counter
func (m *DefaultMetricsCollector) GetCounterValue(name string) int64 {
	if counterInterface, exists := m.counters.Load(name); exists {
		counter := counterInterface.(*int64)
		return atomic.LoadInt64(counter)
	}
	return 0
}

// GetGaugeValue returns the current value of a gauge
func (m *DefaultMetricsCollector) GetGaugeValue(name string) int64 {
	if gaugeInterface, exists := m.gauges.Load(name); exists {
		gauge := gaugeInterface.(*int64)
		return atomic.LoadInt64(gauge)
	}
	return 0
}

// NoOpMetricsCollector implements MetricsCollector but discards all metrics
type NoOpMetricsCollector struct{}

// IncrementCounter implements MetricsCollector (no-op)
func (m *NoOpMetricsCollector) IncrementCounter(name string) {}

// RecordDuration implements MetricsCollector (no-op)
func (m *NoOpMetricsCollector) RecordDuration(name string, duration time.Duration) {}

// SetGauge implements MetricsCollector (no-op)
func (m *NoOpMetricsCollector) SetGauge(name string, value int64) {}

// AddHistogramValue implements MetricsCollector (no-op)
func (m *NoOpMetricsCollector) AddHistogramValue(name string, value float64) {}
