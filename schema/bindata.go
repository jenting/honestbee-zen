// Code generated by go-bindata.
// sources:
// enum.graphql
// input/request.graphql
// interface/article.graphql
// interface/pageInfo.graphql
// mutation.graphql
// query.graphql
// schema.graphql
// type/article.graphql
// type/category.graphql
// type/customType.graphql
// type/searchBodyArticle.graphql
// type/searchTitleArticle.graphql
// type/section.graphql
// type/status.graphql
// type/ticketField.graphql
// type/ticketForm.graphql
// DO NOT EDIT!

package schema

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

var _enumGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xcf\x4a\xc4\x30\x10\xc6\xef\x79\x8a\x01\xef\xfa\x0c\x35\x29\xa6\xfe\x69\x8b\xc9\xba\xb8\x97\x22\x66\x60\x0f\x6e\x66\x49\x53\x65\x59\xfa\xee\xe2\x4c\x0a\xa5\x9e\x7e\xf3\x65\xc2\xef\x0b\xb9\x01\x7f\x44\xc0\x38\x9d\x40\xd3\x14\x73\xba\x68\x0a\x78\xab\xb6\x27\x70\x55\x00\x00\xee\x81\x61\x9f\x18\x7e\xcf\x78\xec\x25\x59\xc6\xcb\x3b\xa3\x31\x8c\xde\xaa\x59\xa9\x55\xcb\x33\x7d\x7e\x7c\x2d\x05\x12\x8a\xbb\x6e\x87\x9d\xe3\xe9\x60\x87\xa2\x3e\xd8\x41\xb7\x52\x52\xad\x4b\x1a\xb3\xd1\x3a\x4a\xf9\xfe\x02\x23\xa5\x0c\x27\xcc\x47\x0a\xa5\xa2\x2c\xa4\xa2\xef\x5c\xe3\x9b\x4e\x8c\xfa\xb5\xae\x7c\x6d\x86\xca\x73\xdc\xf5\x66\x89\xff\xd5\x5d\x0a\x98\xc4\x4e\x7f\xe3\x4a\x2e\x2b\xf1\x57\x4e\x33\x4d\xed\xf4\xc6\xf2\x46\x19\x21\xe1\x39\xe1\x88\x31\x8f\x30\x9d\xef\x02\xfd\x44\xf8\xa6\xbc\xfc\x06\x5f\xb9\x96\xc7\x88\xa7\xdb\xb7\x6a\x56\xbf\x01\x00\x00\xff\xff\x51\x9e\xb8\x5d\xa5\x01\x00\x00")

func enumGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_enumGraphql,
		"enum.graphql",
	)
}

func enumGraphql() (*asset, error) {
	bytes, err := enumGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "enum.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _inputRequestGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x31\x0e\xc2\x30\x0c\x45\xf7\x9c\xc2\xbd\x46\xd7\xa2\x4a\xac\x65\x44\x0c\x69\x63\x21\x43\x9d\x40\xe2\x20\x21\xc4\xdd\x19\x92\xb6\x09\x30\xda\xff\x7d\xfb\x91\xbd\x45\x81\x01\xef\x11\x83\xec\xb4\x68\x78\x29\x00\x00\x9f\x36\xed\x12\x35\xea\xad\x54\x05\x67\x70\x72\xcc\x68\x37\xb0\x4b\x73\x53\x5e\x41\xbf\xc6\xc3\xb2\x49\x40\x88\xe3\x05\x27\x69\xe1\x20\x9e\xec\x39\x2d\x85\xa6\x2b\x4a\xef\x3c\xef\xcd\x92\xa4\x5f\x31\x88\xe3\x9e\x70\x36\xa1\x85\x63\xb7\x8d\xcd\xe9\xc7\x2f\x8b\x64\xcd\xd1\x99\xe7\xf6\xe5\x9b\x5d\xad\x32\x6d\x35\x63\xed\x84\xac\x69\xfe\x73\xa0\x70\xc8\x5d\x32\x75\xf3\xa1\xe7\x88\x65\xf3\x13\x00\x00\xff\xff\xe7\x51\xf9\x82\x73\x01\x00\x00")

func inputRequestGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_inputRequestGraphql,
		"input/request.graphql",
	)
}

func inputRequestGraphql() (*asset, error) {
	bytes, err := inputRequestGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "input/request.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _interfaceArticleGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcf\xb1\x4e\x03\x31\x0c\x06\xe0\xfd\x9e\xc2\x88\x9d\x07\xc8\x76\xb4\xcb\x49\x88\xa5\x30\x21\x06\x5f\xe2\x72\x96\x92\xb8\x72\x1c\xa4\x0a\xf1\xee\xa8\x0d\x15\x6d\x8e\x2d\xfa\x7e\x5b\xf1\x7f\x0f\x63\x06\xce\x46\xba\x47\x4f\x60\x0b\x1a\x04\x2a\x5e\x79\xa6\x02\xa3\x1a\xfb\x48\xd3\x25\x7f\x18\xfe\x46\xfb\x0c\xbe\x06\x00\x00\x0e\x0e\xa6\xed\xdd\xf9\x8d\xd5\x16\xd1\x29\x38\xd8\x99\x72\xfe\x68\xea\x25\x25\xca\x56\xb6\x5c\x70\x8e\xe4\xe0\x51\x24\x12\xe6\x96\x06\xc5\xbd\x75\x76\x50\x49\x62\x14\x7a\x96\xc2\xc6\x92\x1d\x4c\xd9\x1a\x7d\x8a\xd1\xae\xa6\x4e\x36\x52\xb3\x5d\x99\x57\x42\xa3\x30\x9a\x83\x17\x4e\xd4\xb0\x1e\xc2\x1a\x8b\x54\xf5\xf4\x24\x1e\x4f\x87\x5e\xb7\x90\x6a\xe7\xf9\xee\xa6\x0b\xb7\x95\xe2\xe0\xed\x77\xe9\xbd\xe5\x14\x78\xf5\x49\xc4\x99\xe2\x33\xa6\x7f\xc6\xfd\xe9\x72\x3d\x6e\x24\x74\xdf\x57\x8d\xb7\xb0\x58\x8a\xaf\x3d\x66\x4c\xdd\x9e\xb1\xf5\x4d\x66\x09\xc7\x5b\x89\x5d\xdf\xef\xe1\x27\x00\x00\xff\xff\x3e\x63\xc9\x9b\x27\x02\x00\x00")

func interfaceArticleGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_interfaceArticleGraphql,
		"interface/article.graphql",
	)
}

func interfaceArticleGraphql() (*asset, error) {
	bytes, err := interfaceArticleGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "interface/article.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _interfacePageinfoGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x70\xcc\x53\xc8\xcc\x2b\x49\x2d\x4a\x4b\x4c\x4e\x55\x28\xc9\x48\x2c\x51\x48\x49\x2d\x4e\x2e\xca\x4c\x4a\x2d\x56\x08\x48\x4c\x4f\xf5\xcc\x4b\xcb\xd7\xe3\x42\x28\x81\x89\x29\x54\x73\x29\x28\x28\x28\x14\x24\xa6\xa7\x5a\x29\x78\xe6\x95\x28\x42\xb8\xa9\x45\x01\x68\x22\x89\xe9\xa9\xce\xf9\xa5\x79\x25\x48\x62\xc9\x48\xfc\x5a\x2e\x40\x00\x00\x00\xff\xff\x28\xe7\x48\x09\x84\x00\x00\x00")

func interfacePageinfoGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_interfacePageinfoGraphql,
		"interface/pageInfo.graphql",
	)
}

func interfacePageinfoGraphql() (*asset, error) {
	bytes, err := interfacePageinfoGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "interface/pageInfo.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mutationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\x41\x4b\xfb\x40\x10\xc5\xef\xf9\x14\x2f\xf4\xd2\x42\xf8\xff\xef\x0b\x3d\x94\x56\x24\xa0\x22\x46\xbd\xca\xb8\x3b\xb5\x0b\x71\x27\xee\x4e\x2c\x41\xfc\xee\x92\x4d\x8c\xde\x3c\x2d\xf3\x7b\x6f\xdf\xbc\x59\xe1\xfe\xc4\xb8\xee\x95\xd4\x4b\x80\x0e\x1d\x23\x72\x17\x39\x71\xd0\x04\x6a\x5b\xc8\x11\x7a\x62\x70\xd0\x38\xa0\x13\x3f\x72\x1f\x54\x32\xdd\xdd\xd6\xff\x8a\xfc\x6b\xc9\xf8\x28\x00\x60\x85\x86\x83\x83\x8d\x4c\x3a\x46\xbe\xf5\x9c\x34\x2b\x13\xba\x9b\xc8\xda\x4a\x3f\x06\xef\xc5\xb1\xc1\xfe\x67\xc0\x16\xcd\x65\x05\x47\x4a\x06\xb3\xf9\x40\x4a\xe5\xc6\xa0\xd1\xe8\xc3\x4b\xb1\xec\x51\x50\x54\x6f\x5b\xc6\xbb\x28\xa3\xef\xfe\x3b\x39\x07\x3c\x0f\xf0\x63\x57\x97\x8d\xa3\xb4\x9b\x6c\xeb\xd9\x5e\x3b\x83\xfa\x50\x56\x59\x33\x78\x14\xe5\xb2\xc2\x1f\x8d\x5a\xb1\xd4\xb2\xc1\x55\x7e\xb1\xc5\xc5\xcd\xd3\x43\xb3\x31\x98\xc3\x7f\xd7\x3a\x4a\xb4\x8c\x34\x04\x9b\x61\x1e\x9b\x21\xd8\x75\x9f\x38\x06\x7a\xe5\xef\x5b\xca\x0a\x1d\xa5\x74\x96\xe8\x16\xb4\xc1\x72\xe9\x67\xf1\x15\x00\x00\xff\xff\x0a\xfe\xd1\x1f\xa8\x01\x00\x00")

func mutationGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_mutationGraphql,
		"mutation.graphql",
	)
}

func mutationGraphql() (*asset, error) {
	bytes, err := mutationGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mutation.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _queryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x94\x41\x8f\xd3\x30\x10\x85\xef\xf9\x15\xaf\xda\x03\x59\x29\x5a\x81\xb8\x45\xca\xa1\x5b\x60\x15\x81\xda\xa2\x94\x13\x42\xc8\x9b\xcc\xb6\xd6\xa6\x99\xe0\x4c\x0f\x16\xe2\xbf\x23\x3b\x4e\x93\x76\x57\x08\x51\x0e\x7b\x72\x3c\x33\x1e\xbf\xf7\x79\xda\x2b\x6c\x76\x84\xcf\x07\x32\x16\x62\x5b\x82\xa1\xd6\x50\x47\x8d\x74\x50\x75\x0d\x7e\x80\xec\x08\xd4\x88\xb1\x68\x59\xbb\xb8\x6e\x84\x7d\x74\xbe\xce\x6f\x22\x7f\xaa\x6f\xf0\x33\x02\x80\x2b\xdc\x91\xf8\xc3\xa5\x12\xda\xb2\xd1\xd4\xdd\xf8\x8c\xaa\xeb\xc5\x31\x14\x97\x7c\x70\x5d\x17\x5c\x51\x8a\xc5\xb8\x41\x86\xe2\x2e\x41\xcd\xa5\xaa\x29\xc5\x27\xbf\x22\xc3\xfb\xe5\xf7\x2f\x45\x82\x96\xcc\x5a\x6d\x29\x45\xde\x08\x32\xbc\x7d\x9d\xa0\x9d\xec\xdf\x24\xe8\xd8\xc8\xad\x4d\x51\xf8\x15\x19\xd6\xab\x22\xdf\xe4\xab\x65\x9f\x5a\x99\x8a\x4c\x9f\xf5\x9f\xc8\x30\x2f\x16\xd7\x29\x46\x6d\xb3\x89\x91\x60\xc2\xe2\xde\x42\x3b\xf7\x15\xd8\xe0\x91\x6c\xa3\xf6\xd4\xfb\xe2\x86\xc2\x59\x1b\x0f\xe5\x79\xb5\x32\x1f\xfb\xa2\x14\xf9\xbb\x59\x82\x7f\xf2\x3b\xca\xb2\xfe\xaa\x33\xc2\x1d\x95\xa2\xb9\xe9\x06\xbc\x45\xd8\xbf\x3c\xb8\x83\xb2\x29\xda\xa0\x7e\x24\x7b\xc4\x19\xaa\xe3\x50\x91\x57\x17\x32\x0c\xfd\xa2\x33\x7c\xca\x88\x2e\xeb\xc9\x78\xce\x43\xe0\xe5\xf1\x1b\x94\x4d\xf9\x09\xb7\xcb\x33\x0f\xc2\xed\xd1\x83\x4b\xfb\x9b\x2f\x00\xf7\x35\x74\x9b\x7d\x9b\xa2\xeb\x63\xcf\xbc\x5b\xa8\x8e\x43\xc5\xc5\xef\x16\xfa\x4d\xdf\x4d\x74\xf9\x48\x82\x07\x36\xfb\xee\x19\x05\x1b\x9f\xfe\xc0\x66\x1f\xbb\x92\xa0\xe0\x3a\xc5\x98\x98\x45\x27\x33\xa8\x4c\xb9\x1b\x2c\xbd\xea\x20\x5a\x6a\xf2\x05\x7d\x6a\xe3\xf6\x47\xa6\x3f\xdc\x1f\x5d\x8a\x42\x8c\x6e\xb6\x97\x70\x2d\x9e\x34\x3f\x41\xfc\x44\xd6\x3d\x57\x76\xa2\xea\x96\x2b\xfb\x7f\x45\xfd\xed\xf0\xfe\xe1\x17\x7e\x2e\xec\x84\xb3\x28\x39\x84\x21\xed\xbf\x9d\x60\xb7\xce\xa2\x5f\xd1\xef\x00\x00\x00\xff\xff\xb7\x2c\x05\x1a\x87\x06\x00\x00")

func queryGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_queryGraphql,
		"query.graphql",
	)
}

func queryGraphql() (*asset, error) {
	bytes, err := queryGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "query.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x50\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x60\x81\xdc\xd2\x92\xc4\x92\xcc\xfc\x3c\x2b\x05\x5f\x28\x8b\xab\x96\x0b\x10\x00\x00\xff\xff\xf7\xd1\xd7\x38\x33\x00\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeArticleGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\xc1\x6a\xc3\x30\x0c\x86\xef\x7e\x0a\x95\xdd\xf7\x00\xb9\x75\xed\x25\x30\x46\xa1\xdb\x69\xf4\xa0\xd8\x6a\x6a\xb0\xad\x60\xcb\x83\x32\xfa\xee\xa3\xb1\x0b\xa9\x5b\xd8\x2d\xfa\x24\xe5\xff\xf5\xfb\x05\xd6\x20\xe7\x89\x40\x4e\x28\x60\x28\xe9\x68\x07\x4a\xb0\x8e\x62\xb5\xa3\xf4\xaa\xe6\xee\xad\x04\xeb\x27\x47\x9e\x82\x24\xd8\xe1\x48\x7d\x38\x32\xfc\x2a\x00\x80\x09\x47\xea\xa0\x0f\xb2\x2a\x25\xc5\x5d\x43\x70\xa4\x0d\xe7\x20\x0b\xa6\x9b\x1a\xab\x4e\x07\xdf\x55\x72\x75\x50\x17\xa5\xfe\xb1\x79\xef\x72\x69\xb2\xa2\x3e\x08\xc5\x23\x6a\xaa\x66\xad\xe9\xa0\xdf\x56\xcd\x2c\x27\x8e\xbd\xe9\x60\x2f\xd1\x86\xf1\xe6\xcc\xcf\x7f\xd8\xda\x84\x83\xa3\x0e\xde\x98\x1d\x61\x28\x5d\x13\xf1\x28\x0d\x9b\x22\x7b\x16\x32\x2d\xe6\x64\xc5\x72\x58\x5c\xf9\xc3\x42\xfb\xec\x1b\xf2\x90\x4d\x24\x14\x32\x6b\xe9\xe0\xd3\x7a\x2a\x30\x4f\xe6\x11\x26\xce\x51\xd3\x3b\x6b\xbc\x1a\x5d\x5e\xc1\x59\xe6\xf9\xc6\xd3\x0d\x97\x95\x6b\xda\x75\xe9\x50\xfa\x64\xec\x83\x88\xc3\x81\xdc\x07\xfa\x27\xe3\xf3\x2b\xc6\xf3\x86\x4d\x23\x9f\xa3\xbb\x07\x27\xf1\xee\xab\x85\x01\x7d\xb3\x27\x56\xda\x4b\x06\x36\xe7\x7b\xe2\x9e\xdc\xab\x51\x68\xe4\xab\x95\x10\x48\x97\xd8\x37\x95\x95\xa8\x0a\x5d\xf6\xf7\xe5\x43\x5d\xd4\x5f\x00\x00\x00\xff\xff\xc0\x88\x6d\xdf\x0d\x03\x00\x00")

func typeArticleGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeArticleGraphql,
		"type/article.graphql",
	)
}

func typeArticleGraphql() (*asset, error) {
	bytes, err := typeArticleGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/article.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeCategoryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x91\x41\x6b\xe3\x30\x10\x85\xef\xfa\x15\x2f\xec\x65\x17\xc2\xd2\xd2\x9b\xc0\x07\xc7\xbd\x18\x4a\x12\x70\x7a\x2a\x3d\xa8\xf6\xd4\x11\x95\x25\x23\x8d\x0f\xa2\xe4\xbf\x97\x5a\xa6\x71\x9c\x42\x6f\x3d\x79\xe6\x7b\x63\xde\x1b\xcd\x1f\xe4\xe0\xd8\x13\xf8\xa8\x18\x0d\x85\xda\xeb\x17\x0a\x28\x14\x53\xeb\xbc\xa6\xf0\x5f\x8c\xfa\x19\x40\x77\xbd\xa1\x8e\x2c\x07\xec\x55\x4b\xa5\x7d\x75\x78\x17\x00\xd0\xab\x96\x24\x4a\xcb\xab\xd4\x92\xdf\x2f\x88\x6a\xa9\x70\x83\xe5\x19\xab\x97\xfd\x97\x93\xc4\xd3\x64\x1b\x57\xcf\xe2\x24\xc4\x0f\x69\xe3\x65\xd6\x38\xa5\xd2\x8d\x44\x79\x3f\x05\x70\x41\xb3\x76\x76\xee\xe7\x49\x31\x35\x39\x4b\x1c\x74\x47\x09\x0e\x7d\x73\x0d\x83\x1b\x7c\x4d\x0f\xae\x56\x86\x24\x2a\xf6\xda\xb6\x49\x71\x03\x8f\xf3\x12\x1b\xe7\x0c\x29\x3b\x5b\xcd\xc7\xc2\x35\x8b\xf9\x37\x8a\x5b\xd5\x2d\xe0\xe0\xcd\x25\x38\x72\x67\x1e\x97\xd0\x5e\xfd\x97\xde\xa1\x4f\x6b\xcd\x05\xf3\x4d\xd2\x40\xf5\xe7\x60\x28\x9c\xb5\xa9\xfc\x3b\x3f\x13\x32\xdc\xdd\xac\xcf\x87\x44\x86\xdb\x35\x82\xf3\xbc\x89\x12\xd5\xf8\x45\x86\xfd\xae\x2a\x0f\xe5\x6e\x9b\xa4\x9d\x6f\xc8\x27\x75\x2c\x91\x21\xaf\x8a\x7f\x12\xd5\x64\x36\x3a\x2b\xcf\xba\x36\xf4\x3b\xce\xf9\x64\x26\x4e\xe2\x23\x00\x00\xff\xff\x6c\x9b\xfc\x69\xe3\x02\x00\x00")

func typeCategoryGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeCategoryGraphql,
		"type/category.graphql",
	)
}

func typeCategoryGraphql() (*asset, error) {
	bytes, err := typeCategoryGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/category.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeCustomtypeGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x08\xc9\xcc\x4d\x55\xc8\x2c\x56\x48\x54\x08\x72\x73\x36\x36\x36\xb6\x54\x28\xc9\xcc\x4d\x2d\x2e\x49\xcc\x2d\xd0\xe3\x2a\x4e\x4e\xcc\x49\x2c\x02\xab\xe1\x02\x04\x00\x00\xff\xff\x0d\x9d\xf9\x69\x2b\x00\x00\x00")

func typeCustomtypeGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeCustomtypeGraphql,
		"type/customType.graphql",
	)
}

func typeCustomtypeGraphql() (*asset, error) {
	bytes, err := typeCustomtypeGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/customType.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeSearchbodyarticleGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xcd\x8a\xeb\x30\x0c\x85\xf7\x79\x0a\x95\xbb\xbf\x0f\x90\x5d\x7f\x36\x81\xcb\xa5\x90\x99\xd5\xd0\x85\x62\xab\x89\xc1\xb6\x82\x2d\x0f\x84\xa1\xef\x3e\x34\x4e\x21\x75\x3a\x30\xbb\xf8\x93\x14\x9d\xa3\xf3\x07\xf6\x20\xd3\x48\x20\x03\x0a\x68\x8a\x2a\x98\x8e\x22\xb4\x84\x41\x0d\x07\xd6\xd3\x3e\x88\x51\x96\xe2\xdf\x6a\xee\xdb\x16\xc0\xb8\xd1\x92\x23\x2f\x11\xce\xd8\x53\xe3\xaf\x0c\x5f\x15\x00\xc0\x88\x3d\xd5\xd0\x78\xd9\xe5\x27\x85\x73\x41\xb0\xa7\x23\x27\x2f\x2b\xa6\x8a\x37\x2e\x7b\x6a\xf8\xd8\x2c\xdf\x5d\xaa\x5b\x55\xfd\xda\xc4\x4f\x1e\xd6\x16\x16\xd4\x78\xa1\x70\x45\x45\x8b\x15\xa3\x6b\x68\x4e\x8b\xa2\x24\x03\x87\x46\xd7\xd0\x4a\x30\xbe\x7f\xe8\x76\xf3\x1f\x4e\x26\x62\x67\xa9\x86\x03\xb3\x25\xf4\xb9\xaa\x03\x5e\xa5\x60\x63\x60\xc7\x42\xba\xc4\x1c\x8d\x18\xf6\xab\x1b\x7c\xb2\x50\x9b\x5c\x41\x36\x97\x0b\x84\x42\x7a\x2f\x35\xbc\x19\x47\x19\xa6\x51\x6f\x61\xe4\x14\x14\xfd\x63\x85\x77\xa1\x6b\x17\x9c\x64\xee\x2f\x34\x3d\x70\x1e\x99\xb3\xc8\x43\x97\x5c\x27\x6d\x36\x4b\x2c\x76\x64\xff\xa3\x7b\xd1\x3e\x67\x1c\xa6\x23\xeb\x62\x7d\x0a\xf6\x19\x0c\xe2\xec\x7b\x09\x3d\xba\x62\x4e\x8c\x94\x4e\x3a\xd6\xd3\x33\xb1\x2f\xfc\x46\x6f\xc6\x91\xa4\x88\x12\x85\x7a\xbe\xeb\xf3\x9e\x54\xce\xe2\xb8\xb0\x3c\x95\xe9\xba\xde\xe6\x8f\xea\x56\x7d\x07\x00\x00\xff\xff\x4a\x06\xcb\xf3\x54\x03\x00\x00")

func typeSearchbodyarticleGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeSearchbodyarticleGraphql,
		"type/searchBodyArticle.graphql",
	)
}

func typeSearchbodyarticleGraphql() (*asset, error) {
	bytes, err := typeSearchbodyarticleGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/searchBodyArticle.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeSearchtitlearticleGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x70\x54\x28\xa9\x2c\x48\x55\x28\xc9\x48\x2c\x51\x48\x49\x2d\x4e\x2e\xca\x4c\x4a\x2d\x56\x08\x4e\x4d\x2c\x4a\xce\x08\xc9\x2c\xc9\x49\x75\x2c\x2a\xc9\x4c\xce\x49\xd5\xe3\x02\xab\xc3\x94\x50\xa8\xe6\x52\x50\x50\x50\x28\x01\x09\x59\x29\x04\x97\x14\x65\xe6\xa5\x2b\x82\x85\x92\x13\x4b\x52\xd3\xf3\x8b\x2a\x43\x30\xa5\x4a\x8b\x72\x10\x02\xb5\x5c\x80\x00\x00\x00\xff\xff\xc0\x5c\x15\xb6\x87\x00\x00\x00")

func typeSearchtitlearticleGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeSearchtitlearticleGraphql,
		"type/searchTitleArticle.graphql",
	)
}

func typeSearchtitlearticleGraphql() (*asset, error) {
	bytes, err := typeSearchtitlearticleGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/searchTitleArticle.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeSectionGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\xc1\x6a\xf3\x30\x10\x84\xef\x7a\x8a\x09\xff\xe5\x2f\x84\xd2\xd2\x9b\xc0\x07\xc7\xbd\x18\x4a\x13\x70\x7a\x2a\x3d\xa8\xf2\xd6\x11\xc8\x92\x91\xd6\x87\x50\xf2\xee\x25\x96\x4a\x1c\xb7\xd0\x93\xbd\xdf\x8c\x99\x59\xef\x3f\x94\xe0\xe3\x40\xe0\x83\x62\xb4\x14\x75\x30\xef\x14\xd1\x90\x66\xe3\x5d\xbc\x15\x93\xfa\x3d\xc2\xf4\x83\xa5\x9e\x1c\x47\xec\x54\x47\xb5\xfb\xf0\xf8\x14\x00\x30\xa8\x8e\x24\x6a\xc7\xab\x34\x52\xd8\x2d\x88\xea\xa8\xf2\xa3\xe3\x19\xd3\x8b\x39\xe6\x1c\x89\xd7\x1c\xb9\x7a\x13\x27\x21\xfe\xa8\x79\xdd\x32\x17\x32\xad\x44\xfd\x98\xb3\x7d\x34\x67\x69\x1e\x1d\x48\x31\xb5\x25\x4b\xec\x4d\x4f\x09\x8e\x43\xfb\x13\x46\x3f\x06\x4d\x4f\x5e\x2b\x4b\x12\x0d\x07\xe3\xba\xa4\xf8\x91\x27\xbf\xc4\xc6\x7b\x4b\xca\xcd\xb6\x0a\xc7\xca\xb7\x0b\xff\x18\xec\x35\x38\x70\x6f\x5f\x96\xd0\xa9\x7e\xf1\x5d\x5a\x78\x48\x1b\xcc\x05\xfb\x4b\x29\xad\x98\x3a\x7f\x8e\x77\x2e\xfd\x10\x89\x2a\xb3\xc9\xa0\x02\x1b\x6d\x29\x5e\x0c\xff\xe7\xd7\x42\x81\x87\xbb\xf5\xe5\x9e\x28\x70\xbf\x46\xf4\x81\x37\x47\x89\x66\x7a\xa2\xc0\x6e\xdb\xd4\xfb\x7a\xfb\x9c\xa4\x6d\x68\x29\x24\x75\x7a\x45\x81\xb2\xa9\x6e\x24\xca\x1c\x26\x4e\xe2\x2b\x00\x00\xff\xff\xb3\x99\x44\x47\x6c\x02\x00\x00")

func typeSectionGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeSectionGraphql,
		"type/section.graphql",
	)
}

func typeSectionGraphql() (*asset, error) {
	bytes, err := typeSectionGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/section.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeStatusGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x70\x54\x28\xa9\x2c\x48\x55\x28\xc9\x48\x2c\x51\x48\x49\x2d\x4e\x2e\xca\x4c\x4a\x2d\x56\x08\x2e\x49\x2c\x29\x2d\xd6\xe3\x02\xcb\x41\x38\x0a\xd5\x5c\x0a\x0a\x0a\x0a\xe9\xf9\x61\xa9\x45\xc5\x99\xf9\x79\x56\x0a\xc1\x25\x45\x99\x79\xe9\x8a\x60\xe1\xc4\x82\x02\xac\xe2\xc5\xa9\x45\x65\xa9\x45\x21\x99\xb9\xa9\x56\x0a\x20\x52\x91\xab\x96\x0b\x10\x00\x00\xff\xff\x88\x6b\xed\x9c\x75\x00\x00\x00")

func typeStatusGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeStatusGraphql,
		"type/status.graphql",
	)
}

func typeStatusGraphql() (*asset, error) {
	bytes, err := typeStatusGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/status.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeTicketfieldGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xc1\x4e\xfb\x30\x0c\xc6\xef\x7d\x0a\x4f\xff\xfb\xff\x01\x7a\x1b\x4c\x93\x76\x01\xa4\x4d\x5c\xd0\x0e\x5e\x63\x15\x8b\x34\x09\x89\x5b\x98\x10\xef\x8e\x96\x0e\xa9\x8b\x57\xb4\x5b\xf5\xfb\xec\xfa\x73\xfc\xfd\x83\x25\xc8\x31\x10\xc8\x2b\x0a\x18\x4a\x4d\xe4\x03\x25\xd8\x71\xf3\x46\xb2\x66\xb2\xe6\x7f\x95\x0b\x26\x04\xbe\x2a\x00\x00\x36\x35\x6c\x56\x8b\xfc\xdd\x47\x5b\xc3\x56\x22\xbb\x76\x04\xa7\x9e\x82\xb0\xd8\x02\x45\xfc\xd8\x69\x3a\xba\x08\xc2\xde\xa9\xf2\xd5\x9c\x16\x7c\xe2\x91\x6e\x9c\x8c\x08\x1b\xe1\x81\x6a\xb8\xf3\xde\x12\xba\xf3\x3f\xe8\xbd\xe7\x48\xa6\xc0\x8d\xb7\x16\x43\x22\xb3\xf6\x71\xd9\x92\x93\xa4\xfa\x5a\xfa\x0c\x6b\x1f\x9f\xd1\xb2\x41\x6d\x20\xaf\xb7\x71\x4f\x3e\x0a\xda\xeb\x6b\x5e\x57\x07\x4e\x7c\x98\x8a\x17\x73\xc9\xb0\xe0\xbc\xfc\xbb\xce\x8c\x2c\xd8\x5e\x0e\x6b\x22\xa1\x90\x59\x4a\x0d\x3b\xee\xe8\x7c\xbc\x60\x34\x8c\xd4\xf9\xe1\x34\xb9\x7c\xa9\x3e\x89\xef\x72\x0e\x1e\xf3\x1d\x52\x0d\x2f\x93\x70\xdc\x97\xfa\x62\x9f\xfb\xd2\x31\x09\xfd\xd1\xb7\x2d\xf5\xc5\xbe\xfa\xae\xaa\x1b\x02\xaa\x26\xea\xc4\xaa\x92\x2b\x11\x76\xd8\xe9\x78\x3e\x28\x38\xa0\xed\x27\xe8\x36\x8b\x6a\x39\x6d\x51\x95\x9c\x2d\x6a\x5b\xca\xc1\x4f\x00\x00\x00\xff\xff\x0b\x4c\xfc\x70\xc5\x03\x00\x00")

func typeTicketfieldGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeTicketfieldGraphql,
		"type/ticketField.graphql",
	)
}

func typeTicketfieldGraphql() (*asset, error) {
	bytes, err := typeTicketfieldGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/ticketField.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeTicketformGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xc1\x4a\xc3\x40\x10\x86\xef\x79\x8a\x29\x5e\xf4\xe2\x03\x2c\x78\x48\xad\x42\x40\x7a\x69\xea\xa5\x14\x99\xee\x0e\x3a\xb8\x99\x0d\x33\xa3\xa5\x88\xef\x2e\x49\x84\x62\x83\xa7\x5d\xbe\xef\xdf\xe5\xe7\xbf\x82\x1a\xfc\xd4\x13\xf8\x1b\x3a\x24\xb2\xa8\x7c\x20\x83\x96\xe3\x3b\xf9\x63\xd1\xee\xb6\x1a\xfd\x19\xc0\x57\x05\x00\xc0\x29\x40\xb3\x5a\x8c\xf7\x0f\xcd\x01\x36\xae\x2c\xaf\x13\x10\xec\xe8\x2f\x51\x3c\xae\x67\x30\xb1\xf5\x19\x4f\x73\xa1\x78\x5c\xfd\xe7\x48\xd2\xd6\x48\x9f\xd9\xf8\x90\x29\xc0\xb2\x94\x4c\x28\x93\xec\x8b\xb1\x73\x91\x00\x8d\xf8\x84\x30\x3a\x7f\x5e\xe6\x58\xea\x9c\x97\x8a\x92\xec\xc2\x28\x99\x2b\x47\xa7\x34\xea\x66\x08\xec\x86\xcf\xf6\x93\x8f\x4a\xe8\x94\x6a\x0f\xd0\x72\x47\xbf\x03\xf4\x69\x0e\x7d\x9a\x8c\x29\x27\xbb\x2f\x22\x14\x87\x66\xd7\xb9\x44\x1c\x7a\x3f\x8d\x27\xdc\xc1\xc3\xfa\x65\xbb\xb9\x09\xb0\x6b\xcf\x0f\x16\xfb\xea\xbb\xfa\x09\x00\x00\xff\xff\x6d\xa6\x8a\x81\x9d\x01\x00\x00")

func typeTicketformGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeTicketformGraphql,
		"type/ticketForm.graphql",
	)
}

func typeTicketformGraphql() (*asset, error) {
	bytes, err := typeTicketformGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/ticketForm.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"enum.graphql": enumGraphql,
	"input/request.graphql": inputRequestGraphql,
	"interface/article.graphql": interfaceArticleGraphql,
	"interface/pageInfo.graphql": interfacePageinfoGraphql,
	"mutation.graphql": mutationGraphql,
	"query.graphql": queryGraphql,
	"schema.graphql": schemaGraphql,
	"type/article.graphql": typeArticleGraphql,
	"type/category.graphql": typeCategoryGraphql,
	"type/customType.graphql": typeCustomtypeGraphql,
	"type/searchBodyArticle.graphql": typeSearchbodyarticleGraphql,
	"type/searchTitleArticle.graphql": typeSearchtitlearticleGraphql,
	"type/section.graphql": typeSectionGraphql,
	"type/status.graphql": typeStatusGraphql,
	"type/ticketField.graphql": typeTicketfieldGraphql,
	"type/ticketForm.graphql": typeTicketformGraphql,
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
	"enum.graphql": &bintree{enumGraphql, map[string]*bintree{}},
	"input": &bintree{nil, map[string]*bintree{
		"request.graphql": &bintree{inputRequestGraphql, map[string]*bintree{}},
	}},
	"interface": &bintree{nil, map[string]*bintree{
		"article.graphql": &bintree{interfaceArticleGraphql, map[string]*bintree{}},
		"pageInfo.graphql": &bintree{interfacePageinfoGraphql, map[string]*bintree{}},
	}},
	"mutation.graphql": &bintree{mutationGraphql, map[string]*bintree{}},
	"query.graphql": &bintree{queryGraphql, map[string]*bintree{}},
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
	"type": &bintree{nil, map[string]*bintree{
		"article.graphql": &bintree{typeArticleGraphql, map[string]*bintree{}},
		"category.graphql": &bintree{typeCategoryGraphql, map[string]*bintree{}},
		"customType.graphql": &bintree{typeCustomtypeGraphql, map[string]*bintree{}},
		"searchBodyArticle.graphql": &bintree{typeSearchbodyarticleGraphql, map[string]*bintree{}},
		"searchTitleArticle.graphql": &bintree{typeSearchtitlearticleGraphql, map[string]*bintree{}},
		"section.graphql": &bintree{typeSectionGraphql, map[string]*bintree{}},
		"status.graphql": &bintree{typeStatusGraphql, map[string]*bintree{}},
		"ticketField.graphql": &bintree{typeTicketfieldGraphql, map[string]*bintree{}},
		"ticketForm.graphql": &bintree{typeTicketformGraphql, map[string]*bintree{}},
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

