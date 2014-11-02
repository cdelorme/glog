
# go-log

Yet another logging utility package for golang.


## alternatives

While other libraries exist, including golangs own [log](http://golang.org/pkg/log/) package, and third party utilities like [go-logging](https://github.com/op/go-logging), they didn't agree with my use cases.  They do supply a [syslog extension](http://golang.org/pkg/log/syslog/), but it requires that you supply it with a writer, adding to the complexity of using it without writing an abstraction to cover it.

First, golang's own package only supports three methods that don't match [log-level standards](http://en.wikipedia.org/wiki/Syslog#Internet_standards); `printf`, `panic`, and `fatal`, two of which directly impact the flow of the application by throwing `panic()` or `os.Exit(1)`.

The go-logging library is actually fantastic if you want a feature-filled logging library complete with abstractions and extensible tested code.  My only complaint is the complexity; for something as simple as logging we shouldn't need all that much extra around it.


## sales pitch

My library differs in that it aims for utmost simplicity while returning control to the developer.

It:

- supports linux-standard log-level message types
- has only three options; `logLevel`, `silent`, and `nocolor`
- lets the developer control async operations via goroutines (ex. `go logger.Error(...)`)
- returns a `LogMessage` object if you wanted to do more with it

It does not:

- run operational code on behalf of the developer (like throwing `panic()` or `os.Exit(1)`)
- have wildly extensible abstraction layers or a myriad of extra features
- include test files which increase the binary size (and golang binaries are far from "tiny")

For size comparison, I compiled three simple main packages in golang 1.3, each using a logging tool with a single line of output.  The respective binary sizes were:

- 2.8M go-logging
- 1.8M go-log
- 1.7M golang log

So the **golang log** package wins in smallest size, but not utility, while my `go-log` package shaved off over a megabyte of executable size by not including multiple abstraction layers and test files.


## usage

Using my library is simple:

    import "github.com/cdelorme/go-log"

The package name is `log`, and you can get a new logger and begin using it with:

    logger := log.Logger{Level: "Warning"}
    logger.Warning("message %v", object)

_You may customize the prefix format by setting the `.Format` property of the logger.  However, it will always pass the same arguments [in order](https://github.com/cdelorme/go-log/blob/master/logger.go#L28-L32)._
