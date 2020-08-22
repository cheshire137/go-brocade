# go-brocade

*Work in progress*

## How to run

### svg-to-svgo

Use to generate Go code that will generate the same SVG.

```bash
make
bin/svg-to-svgo some-svg-file.svg > tmp.go
```

### go-brocade

Use to generate an SVG image of [brocade](https://en.wikipedia.org/wiki/Brocade) patterns.

```bash
make
bin/go-brocade > some-svg-file.svg
```

Use a tool -- like [Gapplin](http://gapplin.wolfrosch.com) on macOS -- to view
the resulting SVG file.

## Thanks

- [ajstarks/svgo](https://github.com/ajstarks/svgo)
- [rustyoz/svg](https://github.com/rustyoz/svg)
