# What it is

A utility to connect to a given Redis server and upload an snapshot
(Redis dump file) to a remote (non-Redis) destination. 

# Why It Exists

Sometimes you don't want to run backups on the node itself, sometimes
you don't have shell access to the node running your Redis. For these
cases, this backup agent was created.

# How to Use it

Go makes this part easy. Assuming you have your GOPATH environment variable set
up you simply run:

`
go get github.com/TheRealBill/redis-buagent/
`

Now, if $GOPATH/bin is in your path you can run `redis-buagent` and it will
complain about the config file being absent.

## Configuration

To get the default file in place:

`
mkdir /etc/redis 
cp $GOPATH/src/github.com/TheRealBill/redis-buagent/config/buagent.cfg /etc/redis/
`

Now you'll need to modify it to have your credentials and remote-specific
settings. For details on the configuration of redis-buagent see [the config doc](docs/configuration.md).


