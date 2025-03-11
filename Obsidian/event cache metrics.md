The package `eventcache` manages the entries in event cache.

These are some event cache related metrics, namely
- processInfoErrors (event_cache_process_info_errors_total)
- podInfoErrors (event_cache_pod_info_errors_total)
- EventCacheCount (event_cache_accesses_total)
- eventCacheErrorsTotal (event_cache_error_total)
- eventCacheRetriesTotal (event_cache_retries_total)
- parentInfoErrors (event_cache_parent_info_errors_total)

> [!note]
> There are following event types in Tetragon
> - UNDEF
> - PROCESS_EXEC
> - PROCESS_EXIT
> - PROCESS_KPROBE
> - PROCESS_TRACEPOINT
> - PROCESS_LOADER
> - PROCESS_UPROBE
> - PROCESS_THROTTLE
> - PROCESS_TEST
> - PROCESS_RATE_LIMIT_INFO

## InitMetrics
This method initializes the registry with initial (zero values) for all the above mentioned errors.