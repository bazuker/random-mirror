# truerandom-mirror

Simple API server that has only four end-points:
1. GET **/random** returns a unique number from the cache
2. GET **/numbers** returns all buffered numbers.
3. GET **/count** returns a number that represents amount of numbers in the cache**
4. POST **/yourSecretHandler*** updates the cache with new numbers

_yourSecretHandler* is a key provided as a flag with the program launch._

## Installation
Build from the source
```bash
$ go get -u github.com/go-playground/lars
$ go get -u github.com/kisulken/truerandom-mirror
```
Or check out [pre-built binaries](https://github.com/kisulken/truerandom-mirror/releases)

## Example of use
```bash
$ truerandom-mirror --key="mySecrectHandler" -port=7890 -stack=1000
```

[Here is a python client](https://github.com/kisulken/videorand) that generates random numbers from the video input
and submits them to a server over HTTP.
