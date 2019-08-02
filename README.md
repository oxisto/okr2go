# okr2go

okr2go (as in *to go*)  is a simple tracker for your Objective and Key Results (OKR) that you can use locally just using a Markdown file and thus is available *to go*.

Ideally, this Markdown file is stored in a git repository, if you want to synchronize it with others. It also includes a simple Angular-based web frontent to view (and in the future, edit) your objectives.

## Build

Install `packr`, which is used to include the web frontend in the final binary.

```
go get -u github.com/gobuffalo/packr/packr
```

Run the included build script to install all dependencies (Angular and Go) and build an executable with the packed web frontend.

```
./build.sh
```

## Launch

Just launch `okr2go` directly. A browser window with the web ui should open automatically.
