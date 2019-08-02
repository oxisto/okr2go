# okr2go

[![CircleCI](https://circleci.com/gh/oxisto/okr2go.svg?style=shield)](https://circleci.com/gh/oxisto/okr2go) [![](https://godoc.org/github.com/oxisto/okr2go?status.svg)](https://godoc.org/github.com/oxisto/okr2go)


okr2go (as in *to go*)  is a simple tracker for your Objective and Key Results (OKR) that you can use locally just using a Markdown file and thus is available *to go*.

Ideally, this Markdown file is stored in a git repository, if you want to synchronize it with others. It also includes a simple Angular-based web frontent to view (and in the future, edit) your objectives.

## Build

okr2go requires golang >= 1.11, [packr](https://github.com/gobuffalo/packr), and yarn installed.

### Install golang

On Ubuntu 16.04, 18.04, 18.10 and 19.04, use the `longsleep/golang-backports` ppa to install Go 1.12:

```
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo apt-get install golang-go
```

On macOS, use `brew` to install Go 1.12:

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

### Install packr.

Install `packr`, which is used to include the web frontend in the final binary.

```
go get -u github.com/gobuffalo/packr/packr
```

### Run build script

Run the included build script to install all dependencies (Angular and Go) and build an executable with the packed web frontend. This will also run `go install`, so that your system-specific executable is available globally, if you have the `go/bin` directory in your path.

```
./build.sh
```

## Launch

Just launch the `okr2go` executable from the `bin` folder of your system, i.e. `bin/darwin-amd64`. A browser window with the web ui should open automatically.
