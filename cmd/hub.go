package main

import (
	"github.com/miekg/dns"
	"github.com/mr-karan/doggo/pkg/resolve"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// Hub represents the structure for all app wide functions and structs.
type Hub struct {
	Logger     *logrus.Logger
	Version    string
	QueryFlags QueryFlags
	Questions  []dns.Question
	Resolver   *resolve.Resolver
}

// QueryFlags is used store the value of CLI flags.
type QueryFlags struct {
	QNames      *cli.StringSlice
	QTypes      *cli.StringSlice
	QClasses    *cli.StringSlice
	Nameservers *cli.StringSlice
}

// NewHub initializes an instance of Hub which holds app wide configuration.
func NewHub(logger *logrus.Logger, buildVersion string) *Hub {
	// Initialise Resolver
	hub := &Hub{
		Logger:  logger,
		Version: buildVersion,
		QueryFlags: QueryFlags{
			QNames:      cli.NewStringSlice(),
			QTypes:      cli.NewStringSlice(),
			QClasses:    cli.NewStringSlice(),
			Nameservers: cli.NewStringSlice(),
		},
	}
	return hub
}
