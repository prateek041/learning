This package is for internal use only, to keep track of how many errors does tetragon go through throughout it's lifecycle. Package itself is `errormetrics` and consists of two metrics
- ErrorTotal
- HandlerErrors

Something to do with package [[observer]].
## Errors
- process cache miss on get
- process cache evicted
- process cache miss on remove
- process pid tid mismatch
- event missing process info
- handler error
- event finalize process info failed
- process metadata username failed
- process metadata username ignored not in host namespace

## Event handler error
- unknown opcode
- event handler failed

> [!todo]
> This file needs to be refined.

