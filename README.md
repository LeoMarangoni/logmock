# Logmock

A simple application that generate random logs.

------
Usage:
--
Define the log level you want and the delay in microseconds, then run:
```sh
export LOGMOCK_LOG_LEVEL=DEBUG
export LOGMOCK_INTERVAL=100
logmock
```

Instead of directly defining a level, you can define the environment:
```sh
export LOGMOCK_ENV=prd
```

Obviously it's not an app for production, but it simulate that defining the log level to *DEBUG* in the **dev** env and *INFO* in the **prd** env.

Levels available:

- **DEBUG** *default* for dev
- **INFO**  *default* for prd
- **WARN** 
- **ERROR**

Envs available:
- **dev** *default*
- **prd**

_Defining a log level will enable all levels above it and completly overrides the default level for the environment._
_If an invalid level are defined, it will take **INFO** by default._
_If no level are defined, it will take the Environment default._
_If an invalid or no interval are defined, it will take **500** by default._
