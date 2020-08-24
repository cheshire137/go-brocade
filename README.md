# go-brocade

This is a command-line tool for use in generating hopefully interesting SVG images based on
[brocade](https://en.wikipedia.org/wiki/Brocade) patterns. It's a work in progress.

## How to run

This was built on macOS using Go 1.13.4.

```sh
% make
% bin/go-brocade
```

### Example: Listing available options

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
    	Comma-separated list of pattern names to include in the generated image.
    	Ignored when -random is >0. (default "flowerAndStem,swirlyStem,fleur")
  -random int
    	Number of patterns to randomly include. Set to >0 to use, 0 to specify patterns
    	yourself.
  -w int
    	Width of SVG image to produce; defaults to 8.5" at 125px per inch (default 1063)
  -xoffsets string
    	Comma-separated list of X-axis offset values, in pixels, for each pattern.
    	If omitted, will default to 0px.
  -yoffsets string
    	Comma-separated list of Y-axis offset values, in pixels, for each pattern.
    	If omitted, will default to 0px.
```

### Example: Listing available patterns

```sh
% make && bin/go-brocade -list
go build -o bin/go-brocade go-brocade.go
Patterns: fleur, flowerAndStem, jigsaw, jupiter, overcast, sCurve, swirlyStem, yyy
```

### Example: SVG generation with default patterns

```sh
% make && bin/go-brocade -out test.svg -colors "#efefef" -xoffsets "0,-5,-3" -yoffsets "0,30,-65"
go build -o bin/go-brocade go-brocade.go
Using options:
Dimensions: 1063x1375
Output: test.svg
Background: #efefef
Patterns:
- FlowerAndStem #7da955 0,0
- SwirlyStem #d26950 -5,30
- Fleur #ab6ac8 0,0
Wrote test.svg
```

### Example: SVG generation with random patterns

```sh
% make && bin/go-brocade -random 5 -out test.svg
go build -o bin/go-brocade go-brocade.go
Using options:
Dimensions: 1063x1375
Output: test.svg
Background: #d16098
Patterns:
- FlowerAndStem #64ac5a 0,0
- YYY #7989d1 0,0
- YYY #b7953f 0,0
- Jupiter #af51d9 0,0
- FlowerAndStem #d55d45 0,0
Wrote test.svg
```

Use a tool -- like [Gapplin](http://gapplin.wolfrosch.com) on macOS -- to view
the resulting SVG file.

## Thanks

- [ajstarks/svgo](https://github.com/ajstarks/svgo)
- [Hero Patterns](https://www.heropatterns.com)
- [lucasb-eyer/go-colorful](https://github.com/lucasb-eyer/go-colorful)
- [Vectorizer](https://www.vectorizer.io)
