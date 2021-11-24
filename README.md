# arctis7-battery

A very simple tool to get the battery status of a Steelseries Arctis 7 headset.

This is basically a clone of [arctis7.py](https://gist.github.com/flozz/df45b59d6d3594c4b843e00c5df16dd0) written in go.

## Usage

Currently the only supported flag is `-d` to define the time between outputs.

```shell
arctis7-battery      # prints the battery state once
arctis7-battery -d 5 # prints the battery state every 5 seconds
```

If the headset is connected and no errors occur, the battery percentage is written to stdout.
In case of an error, it is logged to stderr and a non zero exit code is returned.

## Dependencies

This program uses [go-hid](https://github.com/sstallion/go-hid) which requires `hidapi` and its headers.
For details, look at the [installation instructions](https://github.com/sstallion/go-hid#installation).
