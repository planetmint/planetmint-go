# planetmint-go
**planetmint** is a blockchain built using Cosmos SDK and CometBFT written in Go and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```
`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Install
To install the latest version of this blockchain node's binary, execute the following command on your machine:
```
git clone https://github.com/planetmint/planetmint-go.git
ignite chain build
```

## Structure
```
- .github/  ... github workflows
- app/      ... app wiring and tx ante handlers
- clients/  ... clients for interactions with external services
- cmd/      ... entry point, sdk config and top level commands
- config/   ... custom planetmint config
- docs/     ... openapi docs
- errormsg/ ... custom error messages
- lib/      ... tools for interacting with planetmint and trust wallet
- monitor/  ... MQTT monitor
- proto/    ... message and type definitions
- tests/    ... e2e test suites
- testutil/
- tools/    ... sdk tools
- util/
- x/        ... custom planetmint modules
```

## Testing
Follow this [guide](https://docs.cosmos.network/v0.47/build/building-modules/testing) for general testing guidelines.

The E2E-tests found in the `tests/` folder setup a test network of n-Nodes and running transactions on said network. Tools to mock interactions with external services for these tests can be found in `testutil/network/`.

For Tests that involve multiple keepers mocks can be found in `x/<module>/testutil/expected_keepers_mocks.go`. These are manipulated in `testutil/keepers/`.

## Contributing
For contributions refer to the RDDL enhancement proposals repository [here](https://github.com/rddl-network/REPs)

### Adding Module Capabilities
Use the `ignite scaffold` [command](https://docs.ignite.com/references/cli#ignite-scaffold) to easily add modules and messages to the existing project. A more manual approach is to add to the `proto/` folder to setup messages and the corresponding message servers and running the `ignite generate proto-go` command.

### Migrations
Module migrations must be registered in each respective module in the `AppModule.RegisterServices(cfg module.Configurator)` function. For each module that is to be upgraded in a migration the ConsensusVersion must be updated. In addition an `UpgradeHandler` needs to be added to the `App.setupUpgradeHandlers()`. Upgrade handlers have a name that needs to be added to an upgrade proposal which needs to be voted on by participating validators.

For more info see [here](https://docs.cosmos.network/v0.47/learn/advanced/upgrade).

## Learn more

- [Planetmint docs](https://docs.rddl.io)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/uy4CA2Xw54)
- [RDDL enhancement Proposals](https://github.com/rddl-network/REPs)
