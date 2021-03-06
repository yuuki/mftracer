// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../data/schema/flows.sql
package data

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _DataSchemaFlowsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x52\xc1\x8e\x9b\x30\x10\xbd\xf3\x15\x73\x0b\x48\x1c\xf6\xd0\x5b\xd4\x95\x5c\x98\xb4\x68\xc1\x9b\x82\x91\x92\x93\x45\xc0\x8d\x2c\x25\x80\xb0\x93\xf6\xf3\x2b\x1b\x08\x21\x4d\xd4\x54\x5d\x2e\x58\x9a\x99\xf7\xde\xbc\x79\x41\x8a\x84\x21\xb0\xed\x1a\xe1\xc7\xa1\xf9\xc9\x2b\xd9\x89\x52\xcb\xa6\x06\x92\x01\xd2\x3c\x01\x77\x51\x94\x5a\x9e\xc5\xc2\x87\x45\x5b\x28\x65\x9e\xde\xd2\x71\xc6\x51\xf2\x25\x46\x88\x56\x40\xdf\x19\xe0\x26\xca\x58\x06\x75\x53\x09\x05\xae\x03\x00\xf6\xcd\x65\x05\x3b\xb9\x57\xa2\x93\xc5\xc1\x36\xd2\x3c\x8e\x61\x9d\x46\x09\x49\xb7\xf0\x86\x5b\xdf\xf6\xca\xf6\xfc\xc9\xfe\x6b\xa1\x2f\x6d\x7d\xa9\x6d\x3a\xdd\x97\xb4\xd8\x8b\x6e\x02\x09\xbe\x61\xf0\x06\xae\xad\xbf\x7e\x86\x17\xcf\xf1\x96\xa3\xb6\x9c\x46\xdf\x73\x84\x88\x86\xb8\xb9\x27\x91\x1b\x42\x6e\x47\xdf\xe9\xa0\x3a\xcf\x22\xfa\x15\x76\xba\x13\x02\x5c\x53\xf7\x2d\xf7\x5f\x36\x36\xde\x8d\x1b\x5b\x1f\x65\x05\xb7\xdf\x33\x0e\x4c\xf6\xcf\xbf\x9b\xdb\xcc\xbd\x51\xcd\xa9\x2b\x05\x1f\x9d\xbe\x26\x94\xf5\x64\x24\xa4\xb8\xc2\x14\x69\x80\x97\x0b\x0d\x23\x9e\x59\x3f\xc4\x18\x19\x42\x40\xb2\x80\x84\x38\xc8\x11\x4a\xcb\xba\x30\x9c\x33\xf8\xff\x47\x2e\x9b\xba\xee\x97\x51\xb3\x45\x1f\xdd\xf7\xba\xff\x15\x5e\xbc\x01\xa5\x13\x85\x16\x7f\x5a\xad\xe5\x51\x28\x5d\x1c\xdb\x09\x27\xc4\x15\xc9\x63\x06\x41\x9e\xa6\x48\x19\x67\x51\x82\x19\x23\xc9\xba\x47\x3a\xb5\xd5\x07\x20\x59\xa8\x21\x75\xee\xfc\x2c\xfe\x3d\x33\xfd\xe9\xe0\x4f\xe6\xd6\x06\x8d\x0f\xd0\x06\x71\x4a\x05\x97\xd5\x2f\x63\x77\x9f\xc5\x59\x8e\xff\x51\xca\x45\xc8\x63\x05\x96\x7a\x80\x7d\xcc\x7b\x97\x67\x2e\xc6\x5b\x3a\xbf\x03\x00\x00\xff\xff\xde\xf4\x30\xec\x85\x04\x00\x00")

func DataSchemaFlowsSqlBytes() ([]byte, error) {
	return bindataRead(
		_DataSchemaFlowsSql,
		"../data/schema/flows.sql",
	)
}

func DataSchemaFlowsSql() (*asset, error) {
	bytes, err := DataSchemaFlowsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../data/schema/flows.sql", size: 1157, mode: os.FileMode(420), modTime: time.Unix(1521764498, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"../data/schema/flows.sql": DataSchemaFlowsSql,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"..": &bintree{nil, map[string]*bintree{
		"data": &bintree{nil, map[string]*bintree{
			"schema": &bintree{nil, map[string]*bintree{
				"flows.sql": &bintree{DataSchemaFlowsSql, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

