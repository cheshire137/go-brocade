# go-brocade

*Work in progress*

## How to run

Use to generate an SVG image of [brocade](https://en.wikipedia.org/wiki/Brocade) patterns.

```sh
% make
% bin/go-brocade -w [WIDTH] -h [HEIGHT] -out [FILE.svg]
```

Example:

```sh
% make && bin/go-brocade
```

This will display usage instructions:

```
go build -o bin/go-brocade go-brocade.go
Usage: bin/go-brocade [options]

  -h int
    	Height of SVG image to produce; defaults to 11" at 125px per inch (default 1375)
  -out string
    	Name of SVG file to create, e.g., my-image.svg
  -w int
    	Width of SVG image to produce; defaults to 8.5" at 125px per inch (default 1063)
```

For example:

```sh
% make && bin/go-brocade -out test.svg
go build -o bin/go-brocade go-brocade.go
Wrote test.svg
```

Use a tool -- like [Gapplin](http://gapplin.wolfrosch.com) on macOS -- to view
the resulting SVG file.

## Thanks

- [ajstarks/svgo](https://github.com/ajstarks/svgo)
- [Hero Patterns](https://www.heropatterns.com)
