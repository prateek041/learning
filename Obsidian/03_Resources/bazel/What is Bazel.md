It is a build and test tool from Google. It is used to automate build and test processes for complex software projects

## Key Features
- **Incremental Builds**: It only rebuilds the parts of the project that have changed.
- **Reproducibility**: Builds are consistent and reproducible.
- **Multi language support**: Supports multiple languages like Go, C++, Python etc.

## BUILD.bazel
It is also known as the Build file and it defines the build and test rules. It specifies the targets to build, their dependencies and other build-related information.

### Example file
```bazel
load("@io_bazel_rules_go//go:def.bzl", "go_library)

go_library(
name="sanitycheck",
srcs=["sanitycheck.go"],
importpath="github.com/sourcegraph/sourcegraph/internal/sanitycheck",
visibility=["//:__subpackages__"],
)
```

## Components
### Load Statement
The load statement allows you to reuse rules or macros defined in other files, typically provided by external repositories or other parts of your bazel project

```bazel
load("//path/to:file.bzl, "rule_or_macro")
```

- **//path/to:file.bzl**: This is the path to the bazel file that contains the rule or macro you want to load.
> [!note]
> Path starting with "//" refers to the root of the workspace

- **rule_or_macro**: the name of the rule or macro you want to import.
From the above example for Source graph [[Architecture]]
```bazel
load("@io_bazel_rules_go//go:def.bzl", "go_library)
```
**Breaking it down**:

1: **Repository Reference (`@io_bazel_rules_go`)**
	- `@io_bazel_rules_go` is an external repository.
	- External repositories are usually defined in [Workspace](https://github.com/sourcegraph/sourcegraph/blob/main/WORKSPACE) file of your Bazel project. They allow you to use rules and packages maintained outside your project.
> [!note]
> Read this to know more about [[WORKSPACE]]

2: **Path to File (`//go:def.bzl`)**
	- This specifies the file within the `@io_bazel_rules_go` repository, which contains the rule you want to load.
	- Here, `def.bzl` is located in `go` directory, which is true and the example is referring to [this](https://github.com/bazelbuild/rules_go/blob/master/go/def.bzl#L152) rule to be loaded.
3: **Rule or Macro name ([[go_library]])**
	- rule that is being imported from `def.bzl` file.
	- `go_library` is a predefined rule in the `rules_go` repository that provides necessary logic to build Go library targets.

## How load works in Practice
When Bazel processes a `Build.bazel` file, it reads the load statement to understand which rule and from where needs to be loaded. This allows us to use imported rules within our `Build.bazel` file as if they were defined locally.