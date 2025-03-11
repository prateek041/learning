# Prometheus docs

## Fundamental metrics

There are 4 fundamental metrics types:

- Gauge
- Counter
- Summary
- Histogram

## Metric vectors

- GaugeVec
- CounterVec
- SummaryVec
- HistogramVec

> [!note]
> The basic difference between fundamental metrics and metric vectors are labels.

There are two interfaces, [[Metric Interface]] and the [[Collector]] interface.

Fundamental metrics implement both, metric Interface as well the Collector interface. but Vector versions do not implement Metric Interface.

> Note: Remember, there is something very important with **Opts** struct, GaugeOpts, CounterOps.

When you create a new metric, you are most probably implementing your own collector interface.

> Note: check goCollector, expvarCollector and processCollector.

A [[Collector]] manages collection of metrics.
The collector and the metric have to describe themselves to a registry, to make sure there is no run-time error.

## Collector

In prometheus, Collector is the interface that is implemented by anything that can be used to collect metrics

To implement the collector interface, any type must have two methods

- Describe: This gives information about the metric itself.
- Collect: Metrics are collected through this.

## Descriptor

A descriptor is the metadata related to a metric, they probably tell what the metrics look like. It is the immutable meta-data of a metric.

When it comes to metrics, there are four things:
- fully qualified domain names (FQDN)
- help
- variableLabels
- constLabels

Descriptors registered with the same registry have to fulfill certain consistency and uniqueness if they share the same full-qualified name. They must have the same help string and label names in each constLabels and variableLabels but they must differ in the value of the const labels.

> [!important]
> Descriptors that share the same fully-qualified names and the same label values of their constLabels are considered equal.




