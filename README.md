Repo for issue https://github.com/moleculerjs/moleculer/issues/688

Error when call Golang moleculer-go from Node.js moleculerjs, run: 

```
make bug
```

Error in console output

```
[2020-03-05T20:19:18.190Z] WARN  imac-denis.local-10872/BROKER: Service 'goMath.add' is not available.
(node:10872) UnhandledPromiseRejectionWarning: ServiceNotAvailableError: Service 'goMath.add' is not available.
    at ServiceBroker.findNextActionEndpoint (/Users/denis/go/src/github.com/nnqq/moleculer-go-and-js/js/node_modules/moleculer/src/service-broker.js:857:13)
    at ServiceBroker.call (/Users/denis/go/src/github.com/nnqq/moleculer-go-and-js/js/node_modules/moleculer/src/service-broker.js:876:25)
    at Timeout._onTimeout (/Users/denis/go/src/github.com/nnqq/moleculer-go-and-js/js/server.js:21:37)
    at listOnTimeout (internal/timers.js:531:17)
    at processTimers (internal/timers.js:475:7)
```
