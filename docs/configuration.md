# Redis Backup Agent Config File

By default redis-buagent will look to /etc/redis/buagent.cfg, but can be
passed a filepath/name on the command line to point it to a specific config
file. This config file will contain all options the program needs to operate. 

## Format

RDB Uploader is configured using a config file which uses the "Git Config"
format - a stricter version of an INI file format.

## Section: Main

The main section stores non-destination driver and non-Redis settings. 

### Destination Driver

Using the `driver` option you specify which destination driver to use such as
'rackspacecf' or 'amazons3'.

### Maximum Uplaod File Size

The `maxfilesize` variable allows you to specify a maximum file size to try to upload. It is
defined as a number of bytes.

### Destination Object Format

You can specify the name format for the remote object using the
`destinationformat` option. By default it will be `YYYY-MM-DD-HH-MM-dump.rdb`


## Section: Redis
This seciton lists information about your Redis server.
It will need to know what redis IP and port to connect to, and if you
only want to backup from a slave server.

There will also be a boolean field and variables for connecting to a
Sentinel and finding the slave to backup, or following a master to an
appropriate slave. These features are not yet available.

## Section: [driver]

Each driver has it's own variables to define connection and driver-specific
variables such as the container name. This is provided for a future feature
where you will be able to specify multiple drivers and have the file uploaded
to multiple locations.

To preserve code minimalism and a common bas eof configuration kowledge, the
names of these variable where common (such as username, apikey, and
containername)  should remain the same.


