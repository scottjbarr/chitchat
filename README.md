# Chit Chat

A little Redis experiment in Go.

Basically publish text to a Redis topic, and the subscribers speak the phrase
via the `/usr/bin/say` command.

For amusement create mutliple subscribers on multiple machines, with different
voices.

## Build

See `build.sh`

## Usage

### Publisher

The `voice` parameter is ignored for the publisher.

    $ ./bin/chitchatpub --help
    Usage of ./bin/chitchatpub:
    -db=0: Redis database
    -host="127.0.0.1": Redis host
    -key="speech": Key to broadcast to
    -port=6379: Redis port
    -voice="": Voice to use for speech

### Subscriber

The `voice` parameter is optional. If not supplied the `say` command will be
executed without a voice specified so the system default will be used.

If a `voice` is specified it is assumed to exist on your system. Specifiying
a voice that does not exist will result in silence when speaking a message.

    $ ./bin/chitchatsub --help
    Usage of ./bin/chitchatsub:
    -db=0: Redis database
    -host="127.0.0.1": Redis host
    -key="speech": Key to broadcast to
    -port=6379: Redis port
    -voice="": Voice to use for speech

### Examples

Subscribe to messages, using voice `alex`

    $ ./bin/chitchatsub -v alex --host 192.168.x.y

Publish a message

    $ ./bin/chitchatsub -v alex --host 192.168.x.y

Publish a message.

    $ ./bin/chitchatpub --host 192.168.x.y This is the sound of noise

Listen to the voices, but take care when following their advice :)

## License

The MIT License (MIT)

Copyright (c) 2015 Scott Barr

See [LICENSE.md](LICENSE.md)
