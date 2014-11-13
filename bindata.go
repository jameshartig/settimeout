package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func assets_favicon_ico() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x62, 0x60,
		0x60, 0x04, 0x42, 0x01, 0x01, 0x06, 0x30, 0xc8, 0x60, 0x65, 0x60, 0x10,
		0x03, 0xd2, 0x1a, 0x40, 0x0c, 0x12, 0x52, 0x60, 0x00, 0xc9, 0x73, 0x40,
		0x24, 0x19, 0x19, 0x10, 0x00, 0xca, 0xfe, 0xff, 0xff, 0x3f, 0xc3, 0x28,
		0x18, 0x05, 0xa3, 0x60, 0x14, 0x8c, 0x02, 0xd2, 0x01, 0x20, 0x00, 0x00,
		0xff, 0xff, 0xe1, 0xe9, 0x47, 0x67, 0x7e, 0x05, 0x00, 0x00,
	},
		"assets/favicon.ico",
	)
}

func assets_index_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8c, 0x55,
		0xd1, 0x6e, 0xe3, 0x36, 0x10, 0x7c, 0x2f, 0xd0, 0x7f, 0x58, 0x18, 0x45,
		0x7a, 0x07, 0x04, 0x96, 0x53, 0x23, 0x05, 0x9a, 0xea, 0x7c, 0x4d, 0x8b,
		0x16, 0x45, 0x9f, 0x0a, 0xc4, 0x05, 0xda, 0x47, 0x4a, 0x5a, 0x59, 0x4c,
		0x28, 0x2e, 0xcb, 0x5d, 0x55, 0x67, 0xa0, 0x1f, 0x7f, 0x4b, 0x4a, 0x4e,
		0x6c, 0x5f, 0x12, 0xc4, 0x0f, 0xb6, 0x69, 0x72, 0x67, 0x87, 0x33, 0xb3,
		0x72, 0xd9, 0x49, 0xef, 0x36, 0x5f, 0x7f, 0x55, 0x76, 0x68, 0x1a, 0xfd,
		0x04, 0x7d, 0x95, 0x62, 0xc5, 0xe1, 0x86, 0x51, 0xb6, 0xb6, 0x47, 0x1a,
		0x04, 0x0c, 0x83, 0x81, 0x3b, 0x8c, 0xff, 0xd9, 0x1a, 0xcb, 0x62, 0xda,
		0xd6, 0x9a, 0x62, 0x2e, 0x2a, 0x2b, 0x6a, 0xf6, 0x87, 0xe2, 0xee, 0xbb,
		0x97, 0x2b, 0x75, 0x6f, 0x3e, 0x15, 0xe6, 0x2f, 0xe9, 0x75, 0x17, 0xb0,
		0xb6, 0xed, 0x1e, 0xa4, 0x43, 0x30, 0x3d, 0x0d, 0x5e, 0x80, 0x5a, 0x10,
		0x45, 0x80, 0x77, 0x0d, 0xb6, 0x66, 0x70, 0xc2, 0x20, 0x04, 0xbd, 0x75,
		0xce, 0x32, 0xd6, 0xe4, 0x1b, 0x7e, 0x0f, 0x7b, 0x1a, 0x60, 0x34, 0x7a,
		0x36, 0x95, 0xb1, 0x76, 0xc0, 0x98, 0x0e, 0xb1, 0x43, 0x0c, 0x50, 0x61,
		0x4b, 0x11, 0x21, 0x22, 0x07, 0x3d, 0x6d, 0xfd, 0x6e, 0x59, 0x56, 0x11,
		0x8a, 0xe3, 0xa6, 0x43, 0x08, 0x14, 0x15, 0xf8, 0xf7, 0xed, 0xf6, 0x4f,
		0x3d, 0xf8, 0xef, 0x80, 0xac, 0xab, 0xd1, 0x4a, 0x97, 0x11, 0x73, 0xfb,
		0x0a, 0xb5, 0x34, 0x2f, 0x83, 0xd1, 0xdf, 0x29, 0x02, 0x53, 0xfd, 0x80,
		0x02, 0xca, 0xc1, 0x63, 0x2d, 0x96, 0x3c, 0x03, 0x79, 0x48, 0x48, 0x70,
		0x7d, 0xb5, 0x5a, 0x9f, 0xd5, 0x33, 0x2a, 0xbf, 0x96, 0x9c, 0xa3, 0x11,
		0x1b, 0xa8, 0xf6, 0xaa, 0x85, 0xc7, 0xd1, 0x59, 0x8f, 0x5f, 0xf0, 0xf9,
		0x47, 0xaf, 0x53, 0x1b, 0x0f, 0xc6, 0x31, 0xa5, 0xba, 0x46, 0xeb, 0x62,
		0x6f, 0x04, 0x7a, 0x6a, 0x6c, 0x6b, 0x31, 0x32, 0x38, 0xfb, 0x80, 0xf0,
		0xf1, 0x9e, 0xc9, 0x5f, 0xc2, 0xc7, 0x9a, 0x39, 0xbd, 0x1b, 0xe7, 0x2a,
		0x53, 0x3f, 0x7c, 0x10, 0x65, 0x7f, 0x99, 0x36, 0xf5, 0xcd, 0x7a, 0x59,
		0xc2, 0x6f, 0x88, 0x0e, 0xda, 0x88, 0x98, 0x44, 0x19, 0x18, 0x61, 0xec,
		0x8c, 0x60, 0x12, 0x89, 0x87, 0xaa, 0xa1, 0xde, 0x58, 0x7f, 0x24, 0x21,
		0xd1, 0xe5, 0x39, 0x21, 0xdd, 0xaf, 0x8d, 0xd6, 0xe9, 0xa1, 0x6f, 0x55,
		0xca, 0xce, 0x8a, 0x1c, 0xc4, 0xe8, 0xcd, 0x27, 0x15, 0x24, 0x6a, 0x6f,
		0xed, 0x71, 0x2c, 0x85, 0xd6, 0x18, 0xa8, 0x22, 0x8d, 0xea, 0xc7, 0x72,
		0x76, 0xba, 0x48, 0x56, 0x1f, 0xb2, 0xb1, 0xde, 0xfc, 0xa5, 0x90, 0xbf,
		0x28, 0x2e, 0x6b, 0x1a, 0xd6, 0xcf, 0xa5, 0xe1, 0x20, 0x44, 0xe2, 0x2c,
		0x9d, 0xcd, 0xce, 0xb3, 0x5e, 0xc3, 0xb6, 0x50, 0x63, 0x94, 0xc4, 0x5b,
		0x7b, 0xab, 0x55, 0x9a, 0x11, 0xe5, 0x16, 0xc1, 0x84, 0x00, 0x46, 0x19,
		0x56, 0x4e, 0xbd, 0x79, 0xce, 0xea, 0xdb, 0x26, 0x25, 0x40, 0x99, 0xa9,
		0x66, 0xd9, 0x97, 0x14, 0x4c, 0x45, 0x35, 0xc0, 0xb2, 0x77, 0xda, 0xc5,
		0xec, 0x72, 0x47, 0xee, 0x68, 0xcc, 0x9a, 0x68, 0xab, 0x79, 0x8f, 0x3b,
		0xcc, 0x6e, 0x0f, 0xae, 0x99, 0xf0, 0xa7, 0x96, 0x6c, 0x05, 0x55, 0x5c,
		0xea, 0xc1, 0x91, 0xc9, 0xf1, 0x82, 0x3b, 0xd3, 0x67, 0xbe, 0xda, 0xe8,
		0x31, 0x03, 0x06, 0xee, 0xf9, 0x9c, 0x8c, 0x3c, 0x0d, 0x46, 0x0e, 0x6e,
		0xac, 0x93, 0x6a, 0x5c, 0x47, 0x1b, 0x24, 0x31, 0x59, 0xc2, 0x56, 0x7f,
		0x4e, 0x26, 0x03, 0x85, 0xa4, 0x6a, 0xa6, 0x56, 0x61, 0xd2, 0xa3, 0x99,
		0xa0, 0x6f, 0xff, 0xb8, 0xfd, 0x1b, 0x92, 0xf3, 0x59, 0x9b, 0xe4, 0xfc,
		0xc4, 0x2a, 0x6a, 0x6a, 0xd4, 0xdf, 0x1e, 0xa5, 0xa3, 0xe6, 0x51, 0xb7,
		0x74, 0x29, 0xed, 0xb4, 0xcf, 0x54, 0x9f, 0x33, 0x25, 0x6c, 0x2e, 0x7c,
		0xc5, 0xe1, 0xc7, 0x73, 0xa7, 0xf2, 0x60, 0xfc, 0xfa, 0xc9, 0xf4, 0xc1,
		0x9d, 0xbb, 0xd5, 0x89, 0x84, 0x9b, 0xa2, 0xd0, 0x39, 0x9f, 0xaf, 0xb3,
		0xb4, 0x54, 0x94, 0x2c, 0x91, 0xfc, 0x6e, 0x73, 0xbd, 0x5a, 0xad, 0xca,
		0x62, 0x5e, 0xc0, 0x98, 0xb5, 0x9b, 0xa6, 0x52, 0x13, 0x0d, 0x69, 0xf7,
		0x64, 0x92, 0xa7, 0xb6, 0x6f, 0x01, 0xbe, 0x5a, 0x71, 0x8a, 0xfd, 0xcb,
		0xd8, 0x57, 0x2b, 0x98, 0x41, 0xc1, 0xe8, 0x00, 0x45, 0x94, 0x21, 0x7a,
		0xc8, 0x25, 0x6f, 0xee, 0xd1, 0x5d, 0xbf, 0xd6, 0x40, 0xd5, 0x4c, 0x89,
		0x53, 0xf0, 0x6b, 0x38, 0xe1, 0xff, 0x24, 0xdb, 0xdd, 0xf4, 0x84, 0x98,
		0x85, 0x3b, 0xd5, 0x6d, 0x4b, 0xaf, 0x29, 0x71, 0x73, 0x4c, 0x33, 0xe2,
		0xfc, 0xf5, 0x1b, 0x78, 0x87, 0x75, 0x47, 0xb0, 0x48, 0xe7, 0x17, 0x70,
		0x71, 0x31, 0x43, 0xac, 0x96, 0xeb, 0xf7, 0xf0, 0x3f, 0xf8, 0x1a, 0x4e,
		0x6e, 0x92, 0x1f, 0x43, 0x8f, 0x2e, 0x67, 0x94, 0xd7, 0x7c, 0x0e, 0xca,
		0x77, 0x88, 0x35, 0xde, 0x40, 0x69, 0xa0, 0x8b, 0xd8, 0x7e, 0x58, 0xcc,
		0xfa, 0xec, 0x34, 0x6a, 0x43, 0xb5, 0xac, 0xa9, 0x2f, 0x5a, 0xc3, 0x29,
		0x65, 0x3f, 0x7c, 0xbf, 0x3e, 0x52, 0x6d, 0xb1, 0x79, 0xe3, 0xc1, 0xb2,
		0x30, 0x9b, 0x13, 0x03, 0x7e, 0xde, 0x7f, 0xd1, 0x4c, 0x34, 0xd8, 0xa2,
		0xcf, 0x8c, 0x04, 0x72, 0xaf, 0x93, 0xc4, 0x9d, 0x8e, 0xb8, 0xdd, 0x2d,
		0x36, 0x3f, 0x1d, 0xad, 0x9e, 0x80, 0xca, 0x62, 0xfe, 0xc7, 0x51, 0x79,
		0xa7, 0xff, 0xaf, 0xcf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x87, 0xce, 0x24,
		0xee, 0xc8, 0x06, 0x00, 0x00,
	},
		"assets/index.html",
	)
}

func assets_robots_txt() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x0a, 0x2d,
		0x4e, 0x2d, 0xd2, 0x4d, 0x4c, 0x4f, 0xcd, 0x2b, 0xb1, 0x52, 0xd0, 0xe2,
		0xe5, 0x72, 0xcc, 0xc9, 0xc9, 0x2f, 0xb7, 0x52, 0xd0, 0x57, 0xe1, 0xe5,
		0x72, 0xc9, 0x2c, 0x4e, 0x84, 0xf2, 0x78, 0xb9, 0x00, 0x01, 0x00, 0x00,
		0xff, 0xff, 0x48, 0x2d, 0x5d, 0x2b, 0x27, 0x00, 0x00, 0x00,
	},
		"assets/robots.txt",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"assets/favicon.ico": assets_favicon_ico,
	"assets/index.html": assets_index_html,
	"assets/robots.txt": assets_robots_txt,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"assets/favicon.ico": &_bintree_t{assets_favicon_ico, map[string]*_bintree_t{
	}},
	"assets/index.html": &_bintree_t{assets_index_html, map[string]*_bintree_t{
	}},
	"assets/robots.txt": &_bintree_t{assets_robots_txt, map[string]*_bintree_t{
	}},
}}
