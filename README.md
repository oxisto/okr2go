# okr2go

[![Actions Status](https://github.com/oxisto/okr2go/workflows/build/badge.svg)](https://github.com/oxisto/okr2go/actions)

okr2go (as in *to go*)  is a simple tracker for your Objective and Key Results (OKR) that you can use locally just using a Markdown file and thus is available *to go*.

Ideally, this Markdown file is stored in a git repository, if you want to synchronize it with others. It also includes a simple React-based web frontend to view (and in the future, edit) your objectives.

## Use a Release

Fetch a binary from the latest GitHub release corresponding to your system architecture and start it using `./okr2go`. A browser window with the web ui should open automatically.

## Build from Source

okr2go requires Go >= 1.16 and yarn installed.

### Install Go

On Linux, follow the instructions at https://golang.org/doc/install or an instruction related to your distribution.

On macOS, use `brew` to install Go:

```
brew install go
```

### Install yarn

On Debian/Ubuntu, the latest yarn version can be installed from the following ppa:

```
curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list
sudo apt-get update && sudo apt-get install yarn
```

On macOS, use `brew` to install yarn:

```
brew install yarn
```

### Run build script

Run the included build script to install all dependencies (React and Go) and build an executable with the packed web frontend. This will also run `go install`, so that your system-specific executable is available globally, if you have the `go/bin` directory in your path.

```
./build.sh
```

### Launch

Just launch the `okr2go` executable from the `bin` folder representing your system, i.e. `bin/darwin-arm64/okr2go`. A browser window with the web ui should open automatically.
