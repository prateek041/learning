```go
func main(){
	sanitycheck.pass()
	svcmain.SingleServiceMain(shared.Service)
}
```

## Sanity Check
It checks if "SANITY_CHECK" environment variable is set to true. If so, it enables testing that the current program is in a runnable state agains the platform it's being executed on.