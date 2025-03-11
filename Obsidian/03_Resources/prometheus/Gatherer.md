This is an interface that [[registry]] implements. Gatherer is responsible for gathering the collected metrics into a number of metricFamilies.

```go
type Gatherer interface {
	Gather()([]*dto.MetricFamily, error)
}
```

Gather calls the Collect method of the registered Collectors and then gathers the collected metrics into a lexicographically sorted slice of Uniquely named MetricFamily protobufs.
