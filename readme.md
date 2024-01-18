<div align="center">
<h1>oszimt-reconnector</h1>

<br />

<img alt="oszimt-reconnector" width="300" src="https://raw.githubusercontent.com/vaaski/oszimt-reconnector/main/.github/oszimt-reconnector.svg" />

</div>

<br />

> The osz-imt school wifi needs to be logged-into after 1GB of data used. That's just
> annoying so this scripts automates logging in.

### USAGE
Upon first start it asks for your OSZ-IMT portal credentials and saves them to
`~/.oszimt-credentials` base64 encoded.
Every subsequent start will use these credentials to log you in.

After the initial setup, it'll check your online status every 3 seconds. This frequency
seems to be ideal for downloads to not be interrupted, and for the script to not be too
annoying on the network.

### INSTALLATION
Download the latest builds here:
- [windows x64](https://nightly.link/vaaski/oszimt-reconnector/workflows/build/go/oszimt-reconnector%20windows%20amd64.zip)
- [mac arm](https://nightly.link/vaaski/oszimt-reconnector/workflows/build/go/oszimt-reconnector%20darwin%20arm64.zip)
- [mac intel](https://nightly.link/vaaski/oszimt-reconnector/workflows/build/go/oszimt-reconnector%20windows%20amd64.zip)

### BUILD FROM SOURCE
- Install go 1.21 or higher
- [Install goreleaser](https://goreleaser.com/install/#go-install)
- [Install go-winres](https://github.com/tc-hib/go-winres#installation)
- Clone this repo
- Run `go generate`
