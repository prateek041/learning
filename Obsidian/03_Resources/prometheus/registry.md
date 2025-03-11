registry registers Prometheus collectors, collects their metrics and gather them into [[metricFamilies]] for exposition.

It implements the [[Gatherer]], [[Collector]] and [[registry#Registerer | Registerer]]

## Registerer
```go
type Registerer interface{
	// Registerer registers a new Collector to be included in metrics collection.
	Register(Collector) error
	// Similar to Register but works for any number of Collectors.
	MustRegister(...Collector)
	// Unregister unregistgers the Collector that equals the Collector passed in as an argument.
	Unregister(Collector) bool
}
```