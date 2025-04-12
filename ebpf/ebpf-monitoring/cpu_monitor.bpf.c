#include <bpf/bpf_helpers.h>
#include <vmlinux.h>

struct {
  __uint(type, BPF_MAP_TYPE_HASH);
  __uint(max_entries, 10240);
  __type(key, u32);
  __type(value, u64);
} cpu_time SEC(".maps")

    SEC("tracepoint/sched/sched_switch") int trace_sched_switch(
        struct trace_event_raw_sched_switch *ctx) {
}
