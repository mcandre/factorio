# DEMO APPS

# EXAMPLES

```console
$ hello
Hello World!
```

# REQUIREMENTS

* [Go](https://golang.org/) 1.21.5+

## Recommended

* POSIX compatible [tar](https://pubs.opengroup.org/onlinepubs/7908799/xcu/tar.html)

# BUILD & INSTALL

```console
$ go install ./...
```

# UNINSTALL

```console
$ rm "${GOPATH}/bin/hello"
```

# PORT

```console
$ FACTORIO_BANNER=hello-0.0.1 factorio
$ sh -c "cd bin && tar czvf hello-0.0.1.tgz hello-0.0.1"
```

# CLEAN

```console
$ rm -rf bin
```
