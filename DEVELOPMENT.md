# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.17+

## Recommended

* [snyk](https://www.npmjs.com/package/snyk) 1.893.0 (`npm install -g snyk@1.893.0`)
* [zip](https://linux.die.net/man/1/zip)

# SECURITY AUDIT

```console
$ snyk test
```

# INSTALL

```console
$ go install ./...
```

# UNINSTALL

```console
$ rm "$GOPATH/src/factorio"
```

# TEST

```console
$ factorio
```

# PORT

```console
$ FACTORIO_BANNER=factorio-0.0.1 factorio

$ cd bin

$ zip -r factorio-0.0.1.zip factorio-0.0.1
```

# CLEAN

```console
$ rm -rf bin
```
