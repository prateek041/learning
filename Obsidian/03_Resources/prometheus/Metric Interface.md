A Metric models a single sample value with it's metadata. It implements the following interface.

```go
type Metric interface{
// Desc returns the descriptor for the metric.
Desc() *Desc
// Write encodes the Metric into a "Metric" Protocol Buffer data transmission object.
Write() error
}
```

