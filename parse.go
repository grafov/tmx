package tmx

import (
	"encoding/xml"
	"io"
	"io/ioutil"
)

// TMXURL is the URL to your TMX file. If it uses external files, the sources
// given are relative to the location of the TMX file. This should be set if
// you use external tilesets.
var TMXURL string

// Parse returns the Map encoded in the reader
func Parse(r io.Reader) (Map, error) {
	var m Map
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return m, err
	}
	err = xml.Unmarshal(d, &m)
	return m, err
}

// Parse returns the Tileset encoded in the reader. Could be used for
// parsing TSX files.
func ParseTileset(r io.Reader) (Tileset, error) {
	var t Tileset
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return t, err
	}
	err = xml.Unmarshal(d, &t)
	return t, err
}
