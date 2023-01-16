# cll-fw-demo


### A short explanation about the project structure

```
.
├── README.md
├── contrib   // contrib exemplify what'll be all the repositories with network/specific implementations (chainlinkX repos)
│   └── network   // network will aggregate the integration interfaces and initialize network providers that fulfill this interfaces
│       ├── evm
│       │   └── ethereum.go   // ethereum as a provider from chainlinkX repos with the transformation logic between framework generic data and ethereum implementation data/logic
│       ├── network.go
│       └── nonevm
│           └── solana.go   // solana as a provider from chainlinkX repos with the transformation logic between framework generic data and solana implementation data/logic
├── go.mod
├── go.sum
├── myapp1   // myapp1 representing an app as part of a product that will use multiple network implementations by the framework
│   └── cmd.go
└── pkg
    ├── infra   // infra is just a generic infra package with config, logger, tracing, etc...
    │   ├── config
    │   │   └── config.go
    │   └── logger
    │       └── logger.go
    └── integration   // integration is the high-level and generic interfaces and types used by apps to integration in multiple networks
        ├── address.go
        ├── integration.go
        └── transaction.go
```