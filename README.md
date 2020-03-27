# Reservoir

Command line tool for reservoir sampling from standard input. 

Implements [Algorithm L](https://en.wikipedia.org/wiki/Reservoir_sampling), which observes that the number of items discarded before the next item enters the reservoir follows a geometric distribution. This reduces the expected running time by only computing random numbers for each item joining the reservoir and not for discarded items.

## Installation

You can download the latest release for linux/amd64 [here](https://github.com/ahlaw/reservoir/releases).

Alternatively, you can download the package with the go tool.

```
$ go get -u github.com/ahlaw/reservoir
```

## Usage

```
$ reservoir <capacity>
```

Reads line by line from standard input and returns sample to standard output one item per line. The sample size `capacity` should be a non-negative integer.

### Example

```
$ cat LOTS_OF_LINES.txt | reservoir 1000
line_42
line_1032
line_793
...
```

## License

[MIT License](LICENSE.md)
