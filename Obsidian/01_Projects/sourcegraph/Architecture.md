[Original Documents](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/doc/dev/background-information/architecture/index.md)

Components
- [[Git Server]]: Stores repositories and makes them accessible to other Source-graph services.
- [repo Updater](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/cmd/gitserver/README.md): Makes sure that repositories in Gitserver are up to date.
  
> [!important]
> To learn more about how repos are synchronised, read this: [Life of a Repository](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/doc/dev/background-information/architecture/life-of-a-repository.md)

## Permission Syncing
Code hosts permissions are mirrored into source graph, so the users get a consistent content as on code hosts. This is present in [[Repo Updater]]. [Read more](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/doc/admin/permissions/syncing.md)

## Search
Devs can search across all the code that is connected to Sourcegraph instance.

It is fast because of [zoekt](https://github.com/sourcegraph/zoekt) a trigram based code search, that creates index of the default branch of each connected repository.

There are two major components:
- [zoekt-indexserver](https://sourcegraph.com/github.com/sourcegraph/zoekt/-/tree/cmd/zoekt-sourcegraph-indexserver)
- [zoekt-webserver](https://sourcegraph.com/github.com/sourcegraph/zoekt/-/tree/cmd/zoekt-webserver)
Additionally, there are two types of code, one which isn't indexed yet and another that will never be indexed (non default branch code). Sourcegraph has a fast search path for non-yet indexed code but nothing for code that will never be indexed.
- [Searcher](https://github.com/sourcegraph/sourcegraph/blob/main/cmd/searcher/README.md) is for non-indexed code.

[Syntect server](https://github.com/sourcegraph/sourcegraph/tree/main/docker-images/syntax-highlighter) does syntax highlightingA

>[important]
>To read more about search, read below:
> - [Code search production documentation](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/doc/code_search/index.md).
> - [List of search query](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/doc/dev/background-information/architecture/life-of-a-search-query.md)
> - [Indexed ranking](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/doc/dev/background-information/architecture/indexed-ranking.md)

## Code navigation
By default, Sourcegraph provides imprecise [search-based code navigation](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/doc/code_navigation/explanations/search_based_code_navigation.md), this is less precise because it is fast since it reuses all the architecture that makes search fast, but can result in
- False positives: (finding two definitions or references that aren't references).
- False negatives: (not able to find all definitions or references).
>[!important]
>This is because search is completely text based, but since it works without any configuration (simple regexes), it works for many cases.

Another option is [precise code navigation](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/doc/code_navigation/explanations/precise_code_navigation.md) to implement this, repositories add a step in their build pipeline that computes the indexed for that revision of code and uploads it to sourcegraph. For this, source graph writes language specific indexers, which is a tough thing to do.

> [!Note]
> To know more about code navigation, read following.
> - [code navigation product documentation](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/doc/code_navigation/index.md)DNE
> - [code navigation developer documentation](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/doc/dev/background-information/codeintel/index.md)
> - [available indexers](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/doc/code_navigation/references/indexers.md) DNE

> [!summary]
> Search (Symbol search for basic code navigation)
> Sourcegraph extension API (Hover and definition provider)
> Native integrations (UI of hover tooltips on code hosts)

## Batch changes
Creates and manages large scale code changes across projects, repositories and code hosts.

To create a Batch change, user creates a [[Batch spec]] which specifies the changes to be performed and the repos it should be performed on. Then this spec is executed by [src-cli](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/doc/dev/background-information/architecture/index.md#src-cli) on the client side, which results in [[changeset spec]] which is sent to Sourcegraph and applied to create changesets. A changeset is a PR/MR.

> [!note]
> To learn more about batch changes:
> - [Batch changes product documentation]()
> - [Batch changes design documentation]()
> - [Batch changes developer documentations]()
> - [How src executes a batch spec]()

## Code Insights
They are stored in separate database called `codeinsights-db`. The interaction happens with that database through a GraphQL API. [[Code Insights]]

## Browser Extensions
This brings the features of Sourcegraph directly into UI of code hosts like Github, Gitlab etc.

> [!note]
> To learn more, go through these.
> - [Developing Code host intelligence](https://sourcegraph.com/github.com/sourcegraph/sourcegraph/-/blob/doc/dev/background-information/web/code_host_integrations.md)
> - [Sourcegraph browser extension documentation]()



