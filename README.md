# arctis7-battery

A very simple tool to get the battery status of a Steelseries Arctis 7 headset.

## Usage

Currently there is no support for any command line arguments.
If the headset is connected and no errors occur, the battery percentage is written to stdout.
In case of an error, it is logged to stderr and an exit code is returned.

## Dependencies

This program uses [go-hid](https://github.com/sstallion/go-hid) which requires `hidapi` and its headers.
For details, look at the [installation instructions](https://github.com/sstallion/go-hid#installation).
