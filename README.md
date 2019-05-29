# Logmock

A simple application that generate random logs.

------
Usage:
--
Define the log level you want and the delay in microseconds, then run:
```sh
export LOGMOCK_LEVEL=DEBUG
export LOGMOCK_INTERVAL=100
logmock
```

Levels available:

- **TRACE**
- **DEBUG**
- **INFO**
- **WARN**
- **ERROR**


_Defining a log level will enable all levels above it._
_If a invalid or no level are defined, it will take **INFO** by default._
_If a invalid or no interval are defined, it will take **500** by default._