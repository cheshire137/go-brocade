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
    	Comma-separated list of colors to apply to the pattern, e.g., #ff00ff,#999999.
    	Defaults to randomly chosen colors. The first color will be used for the
    	background color.
  -h int
    	Height of SVG image to produce; defaults to 11" at 125px per inch (default 1375)
  -list
    	Pass this to list available patterns.
  -out string
    	Name of SVG file to create, e.g., my-image.svg
  -patterns string
    	Comma-separated list of pattern names to include in the generated image. (default "flowerAndStem,swirlyStem,fleur")
  -w int
    	Width of SVG image to produce; defaults to 8.5" at 125px per inch (default 1063)
  -xoffsets string
    	Comma-separated list of X-axis offset values, in pixels, for each pattern.
    	If omitted, will default to 0px.
  -yoffsets string
    	Comma-separated list of Y-axis offset values, in pixels, for each pattern.
    	If omitted, will default to 0px.
```

### Listing available patterns example

```sh
% make && bin/go-brocade -list
go build -o bin/go-brocade go-brocade.go
Patterns: fleur, flowerAndStem, jigsaw, jupiter, overcast, swirlyStem
```

### SVG generation example

```sh
% make && bin/go-brocade -out test.svg -colors "#efefef" -xoffsets "0,-5,-3" -yoffsets "0,30,-65" && open test.svg
go build -o bin/go-brocade go-brocade.go
Using colors: #efefef, #7da852, #8a86ca, #c051c4, #d26a4f
Using offsets: 0,0  -5,30  -3,-65  0,0  0,0
Wrote test.svg
```

Use a tool -- like [Gapplin](http://gapplin.wolfrosch.com) on macOS -- to view
the resulting SVG file.

## Thanks

- [ajstarks/svgo](https://github.com/ajstarks/svgo)
- [Hero Patterns](https://www.heropatterns.com)
- [lucasb-eyer/go-colorful](https://github.com/lucasb-eyer/go-colorful)
