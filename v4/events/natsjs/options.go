package natsjs

import (
	"crypto/tls"

	"go-micro.dev/v4/logger"
)

// Options which are used to configure the nats stream.
type Options struct {
	ClusterID   string
	ClientID    string
	Address     string
	NkeyConfig  string
	TLSConfig   *tls.Config
	Logger      logger.Logger
	SyncPublish bool
	Name        string
}

// Option is a function which configures options.
type Option func(o *Options)

// ClusterID sets the cluster id for the nats connection.
func ClusterID(id string) Option {
	return func(o *Options) {
		o.ClusterID = id
	}
}

// ClientID sets the client id for the nats connection.
func ClientID(id string) Option {
	return func(o *Options) {
		o.ClientID = id
	}
}

// Address of the nats cluster.
func Address(addr string) Option {
	return func(o *Options) {
		o.Address = addr
	}
}

// TLSConfig to use when connecting to the cluster.
func TLSConfig(t *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = t
	}
}

// Nkey string to use when connecting to the cluster.
func NkeyConfig(nkey string) Option {
	return func(o *Options) {
		o.NkeyConfig = nkey
	}
}

// Logger sets the underlyin logger
func Logger(log logger.Logger) Option {
	return func(o *Options) {
		o.Logger = log
	}
}

// SynchronousPublish allows using a synchronous publishing instead of the default asynchronous
func SynchronousPublish(sync bool) Option {
	return func(o *Options) {
		o.SyncPublish = sync
	}
}

// Name allows to add a name to the natsjs connection
func Name(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}
