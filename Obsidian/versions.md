Versions is the package that handles Build related information, it collects information including `go_version`, `commit`, `time` and `modified`. Then exports it using `buildInfoMetrics`.

Below mentioned struct is used to collect Build information.
```go
type BuildInfo struct {
GoVersion string
Commit string
Time string
Modified string
}
```
[[runtime |runtime/debug]] package is used to get the build information for the Binary.

## Metrics
```
full_qualified_name = "tetragon_build_info"
```
The metrics are collected through a Custom Collector named `buildInfoCollector` which collects itself.

> [!important]
> To fully understand how self collecting metrics work, read the source code [here](https://github.com/cilium/tetragon/blob/21ae8bc468a8e97eed5400f1207e4f1a3155f54b/pkg/version/metrics.go#L13)

Since build information are one time and do not change, `NewConstMetric` function is used to create a metric with fixed value.

> [!note]
> Build information is stored in form of labels.

