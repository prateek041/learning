Collector interface is implemented by anything that prometheus can use to collect metrics and it must be registered.

The [[docs#Fundamental metrics|default metrics]] provided by prometheus are also Collectors where they only collect one metric i.e. themselves. But a collector can collect multiple metrics.

To implement the collector interface, any type must have two methods

```go
type Collector interface {
// sends all the possible descriptors of metrics collected by this collector.
Describe(chan <- *Desc)
// Collect is called by the Prometheus Registry when collecting metrics. The implementation sends each collected metric to the channel.
Collect(chan <- Metric)
}
```

> [!note]
> There is something related to `unchecked` Collector. read that at some time.

- Describe: This gives information about the metric itself.
- Collect: Metrics are collected through this.

