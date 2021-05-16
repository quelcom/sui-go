# Sui-go

Sui-go builds on top of the great [sui](https://github.com/jeroenpardon/sui) startpage and expands its capabilities with a focus on simplicity and performance.

## Enhancements

### Easier configuration

Apps, links and providers information is rendered in server side and the configuration is moved to `config.toml`. Handlebars dependency is removed.

### Modules

Modules allow to show dynamic information periodically which is then rendered in server side when the page is requested. The running interval is configured by the user and uses goroutines and channels in the background. The common scenario is fetching information from other services.

Current modules are Pi-hole stats and weather information. You can use the same module several times (for example weather status for different locations). Check `config.toml` for examples.

The inspiration was [Heimdall's enhanced apps](https://apps.heimdall.site/applications/enhanced) functionality, but the approach here is more minimalistic (although not as nice looking).

### Easier deployment

It runs as a single binary: HTML template and other sui assets are embedded to the binary thanks to the [go-embed](https://golang.org/pkg/embed/) directive. You can copy both binary and config file and you are good to go (no need to serve HTML files through Docker or use the basic HTTP server provided by your language of choice).

### Efficient

Not that Handlebars templates are having performance issues, but the rendering time with sui-go is in a different order of magnitude: from ~12 milliseconds to ~250 microseconds (on a 6 years old laptop).

Moreover, there are no filesystem read or write operations. Config file is read once when the application starts, but that's about it.

## Installation

Clone the repo and run the following command to generate the `sui-go` binary:

```bash
make build
```

You will need to have `go` installed.

## Usage

Copy the binary and edit `config.toml`. Then run:

```bash
./sui-go
```

Please note there is no authentication. You might want to run this behind a web server with reverse proxy capabilities.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

Run `make help` for information on linting, tests, etc.

## Thank you / dependencies

- [SUI](https://github.com/jeroenpardon/sui) For the great startpage template
- [wttr.in](https://github.com/chubin/wttr.in) For the weather info
- [go-toml](github.com/pelletier/go-toml) TOML parsing
- [gjson](github.com/tidwall/gjson) JSON parsing

## License
[MIT](https://choosealicense.com/licenses/mit/)