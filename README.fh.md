## StreamingFast Firehose Fork of `Ethereum` (`geth` client)

This is our Firehose instrumented fork of [ethereum/go-ethereum](https://github.com/ethereum/go-ethereum) repository. In this README, you will find instructions about how to work with this repository.

### Protocols

The Firehose instrumentation have a protocol version for the messages exchanges with Firehose on Ethereum binary (`fireeth`). The
protocols we currently develop are:

- Protocol `fh2.3` using the `firehose-fh2.3` branch and `fh2.3` tag(s) suffix
- Protocol `fh3.0` using the `firehose-fh3.0` branch and `fh3.0` tag(s) suffix

  > [!NOTE]
  > This is an on-going development switching Firehose implementation to use Geth Native Tracer that is going to be merged upstream at some point

Read [Branches & Workflow](#branches-&-workflow) section for more details about how we handle branching model and versions.

### Initialization

The tooling and other instructions expect the following project
structure, it's easier to work with the Firehose fork when you use
the same names and settings.

```
cd ~/work
git clone --branch="firehose-fh2.3" git@github.com:streamingfast/go-ethereum.git
cd go-ethereum

git remote rename origin sf

git checkout firehose-fh2.3

git remote add origin https://github.com/ethereum/go-ethereum.git
git remote add polygon https://github.com/maticnetwork/bor.git
git remote add bsc https://github.com/binance-chain/bsc.git

git fetch origin
git fetch polygon
git fetch bsc

git checkout release/geth-1.x-fh2.3
git checkout release/bsc-1.x-fh2.2
git checkout release/polygon-0.x-fh2.3
```

##### Assumptions

For the best result when working with this repository and the scripts it contains:

- The remote `sf` exists on main module and points to `git@github.com:streamingfast/go-ethereum.git`
- The remote `origin` exists on main module and points to https://github.com/ethereum/go-ethereum.git

### Branches & Workflow

Dealing with a big repository like Ethereum that have multiple versions for which we need
to track multiple forks (`Matic`, `BSC`) pose a branch management challenges.

Even more that we have our own set of patches to enable deep data extraction
for Firehose consumption.

We use merging of the branches into one another to make that work manageable.
The first and foremost important rule is that we always put new development
in the `firehose-fh2.3` branch.

This branch must always be tracking the lowest supported version of all. Indeed,
this is our "work" branch for our patches, **new development must go there**. If you
perform our work with newer code, the problem that will arise is that this new
firehose work will not be mergeable into forks or older release that we still
support!

The lowest supported Geth version today is `1.10.1`.

We then create `release/<identifier>` branch that tracks the version of interest
for us, versions that we will manages and deploy.

Currently supported forks & version and the release branch

- `firehose-fh2.3` - Default branch with all Firehose commits in it, based on Geth `1.10.1`.
- [release/geth-1.x-fh2.3](https://github.com/streamingfast/go-ethereum/tree/release/geth-1.x-fh2.3) - Ethereum Geth, latest update for this branch is `1.10.7` ([https://github.com/ethereum/go-ethereum](https://github.com/ethereum/go-ethereum)).
- [release/polygon-0.x-fh2.3](https://github.com/streamingfast/go-ethereum/tree/release/polygon-0.x-fh2.3) - Polygon fork (a.k.a Matic), based on Geth `1.10.3`, latest update for this branch is `v0.3.2` ([https://github.com/maticnetwork/bor](https://github.com/maticnetwork/bor)).
- [release/bsc-1.x-fh2.2](https://github.com/streamingfast/go-ethereum/tree/release/bsc-1.x-fh2.2) - BSC fork (Binance), based on Geth `1.10.22`, latest update for this branch is `v1.1.18` ([https://github.com/binance-chain/bsc](https://github.com/binance-chain/bsc)).

> **Note** To find on which Geth version a particular fork is, you can do `git merge-base sf/release/geth-1.x-fh2.3 origin/master` where `origin/master` is the `master` branch of the original Geth repository (https://github.com/ethereum/go-ethereum).

#### Making New Firehose Changes

Making new changes should be performed on the `firehose-fh2.3` branch. When happy
with the changes, simply merge the `firehose-fh2.3` branch in all the release branches we track
and support.

    git checkout firehose-fh2.3
    git pull -p

    # Perform necessary changes, tests and commit(s)

    git checkout release/geth-1.x-fh2.3
    git pull -p
    git merge firehose-fh2.3

    git checkout release/polygon-0.x-fh2.3
    git pull -p
    git merge firehose-fh2.3

    git checkout release/bsc-1.x-fh2.3
    git pull -p
    git merge firehose-fh2.3

    git push sf firehose-fh2.3 release/geth-1.x-fh2.3 release/polygon-0.x-fh2.3 release/bsc-1.x-fh2.3

### Update to New Upstream Version

We assume you are in the top directory of the repository when performing the following
operations. Here, we outline the rough idea. Extra details and command lines to use
will be completed later if missing.

We are using `v1.10.18` as the example release tag that we want to update to, assuming
`v1.10.17` was the previous latest merged tag. Change
those with your own values.

First step is to checkout the release branch of the series you are currently
updating to:

    git checkout release/geth-1.x-fh2.3
    git pull -p

You first fetch the origin repository new data from Git:

    git fetch origin -p

Then apply the update

    git merge v1.11.4

Solve conflicts if any. Once all conflicts have been resolved, commit then
create a tag with release

    git tag geth-v1.10.18-fh2.3

Then push all that to the repository:

    git push sf release/geth-1.x-fh2.3 geth-v1.10.18-fh2.3

> [!NOTE]
> If you need to issue a Firehose bug fix for an existing version of upstream, for example a Firehose fix on `v1.10.8`, you append `-N` at the end where `N` is 1 then increments further is newer revisions are needed, so you would got tag `geth-v1.10.18-fh2.3-1`

### Development

All the *new* development should happen in the `firehose-fh2.3` branch, this is our own branch
containing our commits.

##### Build Locally

    go install ./cmd/geth

#### Release

   Github actions are automatically created when creating a tag

### View only our commits

**Important** To correctly work, you need to use the right base branch, otherwise, it will be screwed up. The `firehose-v2`
branch was based on `v1.10.1` at time of writing.

* From `gitk`: `gitk --first-parent v1.10.1..firehose-fh2.3`
* From terminal: `git log --decorate --pretty=oneline --abbrev-commit --first-parent=v1.10.1..firehose-fh2.3`
* From `GitHub`: [https://github.com/streamingfast/go-ethereum/compare/v1.10.1...firehose-fh2.3](https://github.com/streamingfast/go-ethereum/compare/v1.9.23...firehose-fh2.3)
* Modified files in our fork: `git diff --name-status v1.10.1..firehose-fh2.3 | grep -E "^M" | cut -d $'\t' -f 2`
