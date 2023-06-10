# Remote Import Paths in Go with Gin

Redirect Go module import and Godoc source links. Useful if you want to retain control of your
Go module import domain, but you don't want to host a public VCS server or want the ability
to easily switch the VCS provider of your Go module.

In this example, I redirect my `targetd-provisioner` repository from
`nosoxon.net/targetd-provisioner` to `github.com/nosoxon/targetd-provisioner`.

### Documentation
* [Remote Import Paths](https://pkg.go.dev/cmd/go#hdr-Remote_import_paths) &ndash; `go` command reference
* [Source Code Links](https://github.com/golang/gddo/wiki/Source-Code-Links) &ndash; godoc wiki

<center><img src=curl.png></center>