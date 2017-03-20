
# [glog](https://github.com/cdelorme/glog)

An idiomatic and simple go logger.


## alternatives

Some alternative packages:

- [log](http://golang.org/pkg/log/)
- [syslog](http://golang.org/pkg/log/syslog)
- [github.com/op/go-logging](https://github.com/op/go-logging)

Neither of the builtin options support [standard log levels](http://en.wikipedia.org/wiki/Syslog#Internet_standards).

**All three of these solutions override application behavior, forcing `os.Exit()` or `Panic()`.**

The complexity surrounding embedding syslog into the logger, which creates a network dependency and forces `New()` syntax, are a pungent code smell that does not easily cross platforms.

_There are over 1000 lines of code in [`go-logging`](https://github.com/op/go-logging), not counting the tests._


## sales pitch

**If you can read my source at-a-glance then my work is done.**

Easy to read code is also easy to use and implement, and boosts confidence in stability.

**This library:**

- provides standard log levels
- uses the ever-flexible stderr
- prints messages in standard syslog format
- utilizes `sync.Mutex` for thread safe execution
- includes file name and line numbers for debugging
- supplies idiomatic go initialization (eg. `l := log.Logger{}`)
- has configurable severity controlled via `LOG_LEVEL` environment variable
- omits transient dependencies (only a minimal set of standard library packages)
- totals 345 lines of code (_including comments, unit tests, **and benchmarks**_)
- detects compatible terminals and automatically applies color to message prefixes

_While this use-case seems like a perfect fit for channels, numerous benchmarks have indicated that a `sync.Mutex` provides superior performance._

_Eliminating severity configuration from the application developer domain is surprisingly wonderful for everybody (eg. no longer worry about `-v`, `-vv`, `-vvv`, `-d`, `-s`, `-q`, `--verbose`, `--debug`, `--silent`, or `--quiet`)._


## usage

Install it:

	go get github.com/cdelorme/glog

Import it:

    import "github.com/cdelorme/glog"

Create a new logger instance:

    logger := glog.Logger{}

_The severity of messages defaults to `Error`, and can be set using `LOG_LEVEL` as an environment variable string matching the supported `severities` (including `silent`)._

Format your own messages with any data types synchronously or with goroutines:

    logger.Info("message %v", object)
    go logger.Debug("message %v", object)

_Keep in mind goroutines are cheap **but not free**; if you have bursty traffic such as a web application this may work, but real-time applications with a tight loop may simply accrue memory and cpu debt._


### examples

Let's say you have some structure like this:

	type Configuration struct {
		Port     int
		Address  string
		Username string
		Password string
	}

You may already be familiar with the `MarshalJSON()` override (as well as `json:"-"` comment formatting):

	func (self *Configuration) MarshalJSON() ([]byte, error) {
		return []byte(`{"port": ` + strconv.Itoa(self.Port) + `, "address": "` + self.Address + `", "username": "` + self.Username + `"}`), nil
	}

However, did you know the same works with `%s` formatting:

	func (self Configuration) String() string {
		s, _ := self.MarshalJSON()
		return string(s)
	}

Even better, the "golden goose" does exist for `%v` too:

	func (self Configuration) GoString() string {
		return self.String()
	}

_Note the lack of pointer-receivers on the `String()` and `GoString()` methods._

**_It seems go exposes complexity to help enforce good application design, so if you find these solutions to be painful it may indicate some flaws to investigate._**


## testing

You can run tests via:

	go test -v -race -cover

_All logger output is redirected during tests (and benchmarks) to `/dev/null` using `io.Discard`._


## benchmarks

Benchmarks can be run on the project via:

	go test -run=X -bench=.

While biased by system, here are the results of the last run:

	BenchmarkLogger-8   	 1000000	      2292 ns/op
	ok  	github.com/cdelorme/glog	2.319s


# references

- [syslog spec rfc 5424](https://tools.ietf.org/html/rfc5424)
- [godoc tricks](https://godoc.org/github.com/fluhus/godoc-tricks)
