It all starts with tetragon main.go file, it consists the part for metrics, the section looks like this

```go
if option.Config.MetricsServer != "" {
go metricsconfig.EnableMetrics(option.Config.MetricsServer)
metricsconfig.InitAllMetrics(metricsconfig.GetRegistry())
go metrics.StartPodDeleteHandler()
// Handler must be registered before the watcher is started
metrics.RegisterPodDeleteHandler()
}
```

Let's try to understand them one by one.

### Enable metrics

- This creates a [[registry]] and uses [[Once]] to make sure that Registry is created only Once per instance of tetragon.
- Then it uses [[promhttp]] for starting a server to expose metrics on `/metrics`, route.
### InitAllMetrics
```go
func InitAllMetrics(registry *prometheus.Registry) {
	initAllHealthMetrics(registry)
	initAllResourcesMetrics(registry)
	initAllEventsMetrics(registry)
}
```

- This in turn executes three more functions that register three types of metrics with the registry we created above.
- Namely:
	- [[HealthMetrics]]
	- [[ResourceMetrics]]
	- [[EventMetrics]]

> [!note]
> Generated metrics link [tetragon Metrics](https://tetragon.io/docs/reference/metrics/)