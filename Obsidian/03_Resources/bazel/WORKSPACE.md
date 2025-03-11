The `WORKSPACE` file in a bazel project defines the external dependencies your project relies on. It serves as the entry point for Bazel to understand how to fetch and manage these dependencies.

There are two functions of the file
- Declaring external Dependencies
- Loading and using External Dependencies

Let's consider this example
```bazel
http_archive( name = "io_bazel_rules_go", url = "https://github.com/bazelbuild/rules_go/releases/download/v0.23.3/rules_go-v0.23.3.tar.gz", strip_prefix = "rules_go-0.23.3", ) load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains") go_rules_dependencies() go_register_toolchains()
```

### Declaring an external repository
```bazel
http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/v0.23.3/rules_go-v0.23.3.tar.gz",
    strip_prefix = "rules_go-0.23.3",
)

```
- **http_archive**: A bazel function that specifies an external repository hosted as a tarball on the web.
	- **name**: Identifier for the repository in your Bazel project.
	- **url**: Location of URL containing the tarball.
	- **strip_prefix**: Path prefix to remove the files in the tarball.

### Loading Rules
```bazel
load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
```
- the `load` statement imports specific macros from the specified files, in the above example `go_rules_dependencies` and `go_register_toolchain` macros are loaded from `deps.bzl` file, present in the `go` directory of the `io_bazel_rules_go`.

Finally, we use the loaded rules.