[gitserver github](https://github.com/sourcegraph/sourcegraph/tree/main/cmd/gitserver)

Mirros repositories from their code hosts (github, bitbucket etc.). When other Sourcegraph services need data from git, they talk to gitserver.

Fetch requests however go through [[Repo Updater]].

# Understanding Codebase
[[gitserver main.go]]