
# go-log

Yet another logging utility package for golang.


## alternatives

There are two decent options that I looked at before I decided to write my own.

Golang's [log](http://golang.org/pkg/log/) package is well documented and full featured.  However it does not include [standard log levels](http://en.wikipedia.org/wiki/Syslog#Internet_standards) is non-standard and does not include the normal Severity message types.  It also executes exit code on your behalf when printing error or panic messages.

The third party [go-logging](https://github.com/op/go-logging) library is an excellent feature-filled library with standard log level support.  However, like the go package, they chose to exit the program for you on error and above.  It also contains a rather sizable amount of code for something that should be handling logging.


## sales pitch

My library differs in that it aims for utmost simplicity while returning control to the developer.  It avoids rich features, wild abstractions, interfaces, and unit tests in favor of delivering a small tight package with great flexibility.

To summarize:

- supports standard log levels
- all output goes to Stderr
- has only three properties; `Severity`, `Silent`, and `NoColor`
- returns a struct for further actions or wrappers
- is concurrently safe
- is under 200 lines of code

_To add context regarding the size the [golang log package is over 300 lines of code, 250 if you remove all comments](http://golang.org/src/pkg/log/log.go?m=text), and the `go-logging` package is 1698 lines of code, or if you exclude test files 1148 lines of code._


## usage

Add my package as a dependency:

    import "github.com/cdelorme/go-log"

You can create a new logger struct directly, and optionally set a different log level:

    logger := log.Logger{Level: log.Info}

Using the standard log levels as method names, you can send output, and it will only print the message if the log level is at or above the loggers setting:

    logger.Info("message %v", object) // is displayed
    logger.Debug("message %v", object) // is not displayed

You can optionally turn off the colors anytime (including creation):

    logger.NoColor = true

You can also choose to completely silence the logger:

    logger.Silent = true
