<h1 align="center">oszimt reconnector</h1>

### synopsis
school wifi needs to be logged-into after 1GB of data used. that's just annoying so this scripts automates that.

upon first start it asks for your OSZ-IMT portal credentials and saves them to `~/.oszimt-credentials` base64 encoded.

### installation
download the latest builds here:
- [windows x64](https://nightly.link/vaaski/oszimt-reconnector/workflows/build/go/oszimt-reconnector%20windows%20amd64.zip)
- [mac arm](https://nightly.link/vaaski/oszimt-reconnector/workflows/build/go/oszimt-reconnector%20darwin%20arm64.zip)
- [mac intel](https://nightly.link/vaaski/oszimt-reconnector/workflows/build/go/oszimt-reconnector%20windows%20amd64.zip)

### build from source
- install go 1.21 or higher
- [install goreleaser](https://goreleaser.com/install/#go-install)
- clone this repo
- run `go generate`
