
# go-log

Yet another logging utility package for golang.


## alternatives

There are many alternatives available, but when I last looked they included:

- [log](http://golang.org/pkg/log/)
- [syslog](http://golang.org/pkg/log/syslog)
- [go-logging](https://github.com/op/go-logging)

The first is golang's builtin logging utility, but it only features log, error, and panic methods, which does not follow [standard log levels](http://en.wikipedia.org/wiki/Syslog#Internet_standards).

The second is their syslog implementation, which mostly follows the standard log levels, and runs through syslog.  It is probably the closest to what I would prefer using given its size and simplicity.  However, I am not a fan of its initialization using a `New()` method, which has a code-smell from a non-golang way of doing things.

The final package is an excellent feature-filled library following log standards.  However, like the go package, they chose to exit the program for you on error and above, and it contains a massive amount of code for something as basic as logging.


## sales pitch

My library aims to deliver the simplist usable implementation, supporting both traditional `stderr` output, and `syslog` compatible implementation.

To summarize, here is what you get:

- follows log standards
- provides stderr output
- optional syslog output
- line numbers for tracking
- returned struct for further application use
- optional color printing
- thread safe
- basic golang-friendly instantiation
- under 250 lines of code

Here is what you don't:

- application decisions made on your behalf
- rich features
- wild abstractions
- interfaces
- unit tests

Given the size of the project, it's something the average developer should be able to grasp at a glance, which makes it a breeze to understand and a pleasure to use.


## usage

To import my library:

    import "github.com/cdelorme/go-log"

Creating a new logger is this simple (you can supply a `Severity` at your discretion, the default is `Debug`):

    logger := log.Logger{Severity: log.Info}

Using the standard log levels as method names, you can send output, and it will only print the message if the log level is at or above the loggers setting:

    logger.Info("message %v", object) // is displayed
    logger.Debug("message %v", object) // is not displayed

_It will still create and return a struct for your application to work with._

You can choose at anytime to disable colored logging, or to silence the output:

    logger.NoColor = true
    logger.Silent = true
