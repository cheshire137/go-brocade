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
go build -o bin/go-brocade go-brocade.go
Usage: bin/go-brocade [options]

  -colors string
    	Comma-separated string of colors to apply to the pattern, e.g., #ff00ff,#999999.
    	Defaults to randomly chosen colors. The first color will be used for the
    	background color.
  -h int
    	Height of SVG image to produce; defaults to 11" at 125px per inch (default 1375)
  -out string
    	Name of SVG file to create, e.g., my-image.svg
  -w int
    	Width of SVG image to produce; defaults to 8.5" at 125px per inch (default 1063)
```

For example:

```sh
% make && bin/go-brocade -out test.svg -colors "#efefef"
go build -o bin/go-brocade go-brocade.go
Using colors: #efefef, #7da852, #8a86ca, #c051c4, #d26a4f
Wrote test.svg
```

Use a tool -- like [Gapplin](http://gapplin.wolfrosch.com) on macOS -- to view
the resulting SVG file.

## Thanks

- [ajstarks/svgo](https://github.com/ajstarks/svgo)
- [Hero Patterns](https://www.heropatterns.com)
- [lucasb-eyer/go-colorful](https://github.com/lucasb-eyer/go-colorful)
