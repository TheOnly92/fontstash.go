# Font Stash: Dynamic font glyph cache for OpenGL

This library was ported from [C](https://github.com/akrinke/Font-Stash) to Go.

Font Stash enables easy string rendering in OpenGL applications. It supports truetype fonts and UTF-8 encoded localized strings. All glyphs are cached in OpenGL texture atlases. Font rasterization is done using Sean Barrett's [stb_truetype.h](http://nothings.org/).

Font Stash was originally created and [published](http://digestingduck.blogspot.com/2009/08/font-stash.html) by [Mikko Mononen](http://digestingduck.blogspot.com).

## Installation

Install using the "go get" command:

    go get github.com/TheOnly92/fontstash.go/fontstash

## Documentation

- [API Reference](http://godoc.org/github.com/TheOnly92/fontstash.go/fontstash)
