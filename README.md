# sigterm

A subset of the `syscall` package for terminating processes. This package provides a `Signal` type that implements
the `encoding.TextMarshaller` and `encoding.TextUnmarshaler` interfaces for text-based serialization and deserialization.
It can be used for configuration management using environment variables[^env_caarlos0] or flags [^flag].

List of termination signals can be found in the GNU libc manual [^gnu_manual].

## Installation

```bash
go get -u github.com/josestg/sigterm
```

[^flag]: https://golang.org/pkg/flag/#TextVar
[^env_caarlos0]: https://github.com/caarlos0/env
[^gnu_manual]: https://www.gnu.org/software/libc/manual/html_node/Termination-Signals.html
