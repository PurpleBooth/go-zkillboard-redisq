package main

import (
	"fmt"
	"github.com/purplebooth/go-zkillboard-redisq/zkillboard-redisq"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "zkillboard-redisq-cli"
	app.Usage = "Print every time a kill comes in. Kills are usually delayed by about 90min"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "developmentLogging, d",
			Usage:  "Enable more verbose, and more human readable logging",
			EnvVar: "VERBOSE",
		},
		cli.StringSliceFlag{
			Name:   "solarSystem, s",
			Usage:  "Solar Systems to watch for kills in",
			EnvVar: "SOLAR_SYSTEMS",
		},
		cli.StringSliceFlag{
			Name:   "character, c",
			Usage:  "Characters to watch for kills in",
			EnvVar: "CHARACTERS",
		},
		cli.StringSliceFlag{
			Name:   "corporation, o",
			Usage:  "Corporations to watch for kills in",
			EnvVar: "CORPORATIONS",
		},
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, V",
		Usage: "print only the version",
	}
	cli.BashCompletionFlag = cli.BoolFlag{
		Name:   "auto-complete",
		Hidden: false,
	}

	app.EnableBashCompletion = true
	app.Action = func(c *cli.Context) error {

		var zapLogger *zap.Logger

		if c.Bool("developmentLogging") == true {
			zapLogger, _ = zap.NewDevelopment()
		} else {
			zapLogger, _ = zap.NewProduction()
		}

		corporations := c.StringSlice("corporation")
		characters := c.StringSlice("character")
		solarSystems := c.StringSlice("solarSystem")

		defer zapLogger.Sync() // flushes buffer, if any
		logger := zapLogger.Sugar()

		exitCmd := make(chan *cli.ExitError)
		defer close(exitCmd)
		errs := make(chan error)
		defer close(errs)
		kills := make(chan *zkillboard_redisq.Package)
		defer close(kills)

		go zkillboard_redisq.NewZKillboardRedisQClient().Listen(kills, errs)

		go func(logger *zap.SugaredLogger) {
			for killmail := range kills {

				if len(corporations) > 0 && !stringInSlice(killmail.Killmail.Victim.Corporation.Name, corporations) {
					continue
				}

				if len(characters) > 0 && !stringInSlice(killmail.Killmail.Victim.Character.Name, characters) {
					continue
				}

				if len(solarSystems) > 0 && !stringInSlice(killmail.Killmail.SolarSystem.Name, solarSystems) {
					continue
				}

				kTime, _ := time.Parse("2006.01.02 15:04:05", killmail.Killmail.KillTime)
				logger.Infow("RECEIVED KILL",
					"killTime", kTime,
					"killPackage", killmail,
				)
			}

			return
		}(logger)

		go func() {
			err := <-errs
			logger.Fatalw(
				"Caught error, exiting",
				"error", err,
			)
			cli.HandleExitCoder(err)

			return
		}()

		interrupts := make(chan os.Signal, 1)
		defer close(interrupts)
		signal.Notify(interrupts, os.Interrupt)

		go func() {
			for interrupt := range interrupts {
				exitCmd <- cli.NewExitError(fmt.Sprintf("Caught %s, Exiting", interrupt.String()), 0)
			}
		}()

		return <-exitCmd
	}

	app.Run(os.Args)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
