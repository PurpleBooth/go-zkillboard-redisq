# zKillboard RedisQ Client

This is a go client for the RedisQ service from zKillboard. This is a service that creates a push interface for 
consumers to be notified of kills.

## Installing

```bash
go get -u -v github.com/purplebooth/go-zkillboard-redisq
```

## Running

To listen once

```go
z := zkillboard_redisq.NewZKillboardRedisQClient()
killPackage, apiErrs := z.ListenOnce()
```

To listen continuously

```go
errs := make(chan error)
defer close(errs)
kills := make(chan *zkillboard_redisq.Package)
defer close(kills)

z := zkillboard_redisq.NewZKillboardRedisQClient()
z.Listen(kills, errs)
```

The continuous listen endpoint is largely there for simple apps, and I expect you will need to create your own variant 
to achieve what you'd like for more expressive apps.

## Demo Client

```bash
$ zkillboard-redisq-cli --help
NAME:
   zkillboard-redisq-cli - Print every time a kill comes in. Kills are usually delayed by about 90min

USAGE:
   zkillboard-redisq-cli [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --developmentLogging, -d       Enable more verbose, and more human readable logging [$VERBOSE]
   --solarSystem value, -s value  Solar Systems to watch for kills in [$SOLAR_SYSTEMS]
   --character value, -c value    Characters to watch for kills in [$CHARACTERS]
   --corporation value, -o value  Corporations to watch for kills in [$CORPORATIONS]
   --help, -h                     show help
   --auto-complete                
   --version, -V                  print only the version

```

## Dependencies

We use the following dependencies:

```bash
go get -u -v github.com/parnurzeal/gorequest
go get -u -v github.com/urfave/cli           # Only for the demo CLI
go get -u -v go.uber.org/zap                 # Only for the demo CLI
```


## License and Contribution

See [LICENSE.md](LICENSE.md)
See [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md)
