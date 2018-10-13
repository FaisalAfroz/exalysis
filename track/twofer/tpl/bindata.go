// Code generated by go-bindata. DO NOT EDIT.
// sources:
// comment-format.md (1.157kB)
// generalize-name.md (166B)
// minimal-conditional.md (103B)
// missing-comment.md (101B)
// plus-used.md (246B)
// stub.md (374B)
// use-string-placeholder.md (423B)
// wrong-comment-format.md (108B)

package tpl

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _commentFormatMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xc1\xaa\xec\x36\x0c\x5d\xd7\x5f\x21\xe8\xa2\x6f\x20\x89\x6f\x5b\xba\xe8\xdb\x3e\xda\x4b\x17\x85\x47\xb9\xbb\xa1\xdc\x71\x6c\x25\x16\xe3\x58\xc1\x52\x26\x9d\x96\xfe\x7b\x71\x26\x33\xbc\xa1\x50\xe8\x26\x10\x4b\x96\xce\x39\x3a\xb2\x69\xdb\xd6\xbc\x32\x44\x27\x70\x1c\x0b\x3a\x85\x71\xa1\x80\x89\x32\xca\xef\x1f\xa2\xea\x2c\x1f\xad\x1d\x39\xb9\x3c\x76\x5c\x46\x1b\xd8\x5b\x1c\x06\xf4\x4a\x17\x7c\x1f\xb9\x8b\x3a\xa5\x83\x71\x3d\x2f\x0a\x91\x57\x50\x86\xb5\x90\x22\x78\x9e\x26\xcc\x2a\x9d\x31\x6f\x11\x61\x76\xfe\xec\xc6\xc7\x31\x48\xe4\x25\x05\x10\x75\x45\x61\x25\x8d\xa0\x11\x61\xe5\x12\xe0\xfd\xf3\x2d\xf7\x1d\x06\x4e\x89\x57\x0c\xd0\x5f\xb7\xf0\xbd\x48\x76\x13\x76\xc6\x9c\x4e\x27\x63\x2d\x7c\x7e\x94\xe6\x33\xcc\x85\x2f\x14\x50\x20\xba\x1c\xae\xe0\x39\x5f\xb0\x08\x71\x86\x09\x35\x72\x10\x18\xb8\xc0\x92\x49\x05\xf4\x3a\x93\x77\x29\x5d\x61\x11\x0c\x40\x19\x0a\x7a\x9a\x51\x3a\x33\x7f\x51\x73\xeb\x63\x7e\xe6\x02\xf8\xc7\xcc\x45\x31\xc0\xb0\x64\xaf\xc4\x59\x9a\x7b\xd9\xa6\xb6\x12\x75\x59\xa5\x01\x97\xc3\x03\xeb\xc5\x15\x72\x7d\x42\x69\x36\x0a\xff\xcd\xff\x99\xd8\x5b\x2f\x6f\xfc\x6b\xda\x49\x54\xc0\x5b\xa1\x99\x39\x4b\x15\x7a\xa2\x94\x28\x91\x62\x91\xce\x54\x48\xf7\x1b\x1f\xb4\x17\xa0\xac\x87\xfa\x81\xbf\xcc\x57\xd6\x42\xd7\x75\xe6\xef\x1b\x95\x3a\x8f\xe3\x27\x0e\xf8\x1b\x5e\x08\xd7\x4f\xfb\xa4\xbe\x18\x38\x69\x5c\xfa\xce\xf3\xb4\xcf\xde\x8e\x6c\x57\x3a\x93\xfd\xf7\xad\xaf\x77\x46\xad\x60\x56\xcc\x1e\xe5\x00\x35\x15\x86\xa5\x68\xc4\x02\x32\xa3\xa7\x81\x50\x40\xa3\x53\x08\xec\x1f\xde\xb8\xab\xd0\x23\x0c\x4b\x4a\xf0\x28\xd1\x00\xe6\x40\x79\xbc\x29\xe3\x60\xc6\x42\x1c\x3a\x63\x7e\xc9\xe0\x42\xa0\xaa\x7d\x55\x80\x67\xa5\x89\xfe\xac\x99\x75\xac\x05\x5d\x70\x7d\x95\xe4\xda\xec\xe6\xa9\xa1\xaa\x6c\xc1\x5b\xd3\xe0\xb6\xb9\x01\x4d\xd5\x29\xd5\x27\xbc\xc2\x95\x97\x02\x9e\x03\x9a\xc4\x7c\xae\xd2\xc1\xf1\x95\xbf\x91\x0a\x76\xa9\x48\xdd\xde\x90\xd3\xae\xd2\x47\x6b\xd7\x48\xf3\x8c\x45\xd4\xf9\x33\x96\x4d\xac\xef\x5e\xbe\xfd\xc1\xbe\xfc\x68\xbf\x7f\xb1\x23\xb7\x4f\x97\xdb\x91\x03\xfb\xa7\x2f\x97\xb1\x75\x39\xb4\xb7\x54\x7b\x30\xb2\xf8\x08\xdb\x2a\xd6\x78\xdd\xb7\x47\xb3\xc7\xc9\xa1\x81\xd3\xf6\x73\xba\xd9\xec\x34\x72\x45\x79\xaa\x6b\xe6\xce\x08\x0e\x26\xde\x0c\xa6\xbc\xc9\xb1\x91\x17\xdc\xdc\x0a\x9c\x9f\xe5\xa7\xbc\x85\x79\x18\xc8\x93\x4b\x20\x7a\x4d\x78\x7b\x02\x1a\x73\xfc\xe9\xbe\xe7\xf0\xca\xff\xe3\x31\xa8\x9d\x05\x71\x57\xdd\x09\x67\xca\xa3\xe9\x31\x52\xde\xd0\x08\x82\x8f\x4c\xbe\xae\x59\x7d\x80\xfe\x09\x00\x00\xff\xff\x2c\xca\x57\x85\x85\x04\x00\x00")

func commentFormatMdBytes() ([]byte, error) {
	return bindataRead(
		_commentFormatMd,
		"comment-format.md",
	)
}

func commentFormatMd() (*asset, error) {
	bytes, err := commentFormatMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "comment-format.md", size: 1157, mode: os.FileMode(420), modTime: time.Unix(1538510115, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x62, 0xc9, 0x45, 0xe1, 0xe0, 0x59, 0xe7, 0x99, 0x6e, 0xf1, 0x3c, 0x61, 0xd, 0x9d, 0x49, 0x84, 0x48, 0xb7, 0x68, 0x70, 0x3b, 0x8f, 0xb6, 0x30, 0xab, 0x12, 0x7, 0x3e, 0x89, 0xd1, 0x5c, 0x90}}
	return a, nil
}

var _generalizeNameMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcd\x31\xae\xc2\x30\x10\x45\xd1\xfe\xaf\xe2\x6d\xe0\x67\x0f\xd0\xa5\xa6\xa0\xf5\xc4\x7e\x91\x47\x38\x33\xc8\x9e\x10\x65\xf7\x48\x11\xfd\xb9\xba\xff\x98\x31\x48\x9c\xbe\xa3\x4a\x2f\xd9\x0b\x0b\xa2\x12\x26\x1b\x07\xd2\xad\x69\x66\x82\x58\x41\xba\xfb\x92\x26\xcc\x81\xc3\xf7\x56\xb0\x10\xa6\x99\xd0\xf5\x0a\xd2\xa3\x4a\xe7\x53\xa3\x26\xac\xbb\xe5\x50\xb7\x9f\x3c\xbc\xbf\xf0\xb7\x7a\x07\x3f\xec\x27\xde\x1e\xb4\x50\x69\xd7\x05\xc3\x31\x23\x8b\x41\xda\x70\xa8\x0d\xf6\xc0\x76\xc2\x0f\xbb\xc0\xf4\x0d\x00\x00\xff\xff\x0d\x1c\xed\x31\xa6\x00\x00\x00")

func generalizeNameMdBytes() ([]byte, error) {
	return bindataRead(
		_generalizeNameMd,
		"generalize-name.md",
	)
}

func generalizeNameMd() (*asset, error) {
	bytes, err := generalizeNameMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "generalize-name.md", size: 166, mode: os.FileMode(420), modTime: time.Unix(1539192870, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6, 0x80, 0xb6, 0x8d, 0xbf, 0x31, 0x5b, 0x5b, 0x96, 0x0, 0x2e, 0xcd, 0x97, 0xea, 0x38, 0x7f, 0xb4, 0x72, 0x34, 0x81, 0xb7, 0x87, 0xd0, 0x2e, 0x25, 0x31, 0xb8, 0xc8, 0xd7, 0x4c, 0x47, 0x47}}
	return a, nil
}

var _minimalConditionalMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xcb\x51\x0d\xc3\x30\x0c\x04\x50\x2a\x47\x60\x25\x31\x24\x51\xe3\x76\x96\x12\x5f\xe5\x9e\x3f\xc2\x7e\x52\x00\xbc\x0f\x16\x0b\x27\x6b\x74\x28\x17\x44\x5c\x1e\x1d\xfa\x19\xa6\x87\xcf\x36\x70\x32\xba\xcb\x19\x6d\x1c\xf8\x56\xa6\x85\xc6\xda\xb0\xa5\x21\xed\xb1\x26\x8f\x1b\x93\xaf\xc0\x6b\x63\x96\x9e\x12\x5e\xa5\xc7\x7d\xfc\x03\x00\x00\xff\xff\xb7\xa4\xcc\x80\x67\x00\x00\x00")

func minimalConditionalMdBytes() ([]byte, error) {
	return bindataRead(
		_minimalConditionalMd,
		"minimal-conditional.md",
	)
}

func minimalConditionalMd() (*asset, error) {
	bytes, err := minimalConditionalMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "minimal-conditional.md", size: 103, mode: os.FileMode(420), modTime: time.Unix(1538507702, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x59, 0xb2, 0x85, 0x17, 0xd0, 0x29, 0xbf, 0x5b, 0x1e, 0x5f, 0xc0, 0x80, 0xaf, 0x9a, 0x69, 0x9e, 0xaa, 0xb5, 0xfb, 0xb6, 0x10, 0x15, 0xc4, 0x77, 0x78, 0xcf, 0xa, 0x3f, 0x68, 0x2b, 0xd0, 0xbc}}
	return a, nil
}

var _missingCommentMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcb\xc1\x0d\xc2\x30\x0c\x86\xd1\x55\xbe\x0b\x47\x3a\x06\x7b\x98\xd4\x22\x11\x89\x7f\xa9\x76\x41\x6c\xcf\xa9\xf7\xf7\xee\xdc\x92\x11\xbc\x44\x76\x9d\x73\xa7\xdb\xc7\x31\x9a\xd6\xf2\xa8\x8d\x87\x0e\x96\x0e\x67\xf7\xb2\x31\x93\x9f\x4e\x9a\xc5\x05\xa7\xf4\xc6\x8a\xea\x7e\xa5\x24\xbd\xd5\x50\xf0\xf4\xa9\xef\xf6\x0f\x00\x00\xff\xff\x04\x3d\xae\x58\x65\x00\x00\x00")

func missingCommentMdBytes() ([]byte, error) {
	return bindataRead(
		_missingCommentMd,
		"missing-comment.md",
	)
}

func missingCommentMd() (*asset, error) {
	bytes, err := missingCommentMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "missing-comment.md", size: 101, mode: os.FileMode(420), modTime: time.Unix(1538428246, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x18, 0x1e, 0x1f, 0x6f, 0xb5, 0xfc, 0x5c, 0x55, 0x7c, 0x5f, 0x99, 0x68, 0x72, 0xf8, 0x35, 0x9b, 0x61, 0x26, 0x20, 0x27, 0xdf, 0x91, 0x9, 0x7e, 0x9b, 0x17, 0x87, 0x7, 0x5c, 0x97, 0x8, 0x83}}
	return a, nil
}

var _plusUsedMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcf\x4d\x6e\xc4\x30\x08\xc5\xf1\xfd\x9c\xe2\x75\x5d\x35\xf7\xc8\xbe\x07\x08\x65\x48\x8c\xc6\x86\xc8\x10\x55\xbe\x7d\xe5\x7e\xa8\xb3\xfb\x6f\xde\x0f\xf1\x86\x15\x21\x82\xe1\x17\xa8\x0b\xae\x50\x3b\xb0\xbd\x6e\x48\x07\xbb\x31\xa5\x18\xa5\x20\x8b\x20\xb2\xab\x1d\x0b\xde\x0b\x25\x34\x70\x4a\xdf\x85\xb3\x0e\x54\x39\x34\x5f\xb0\x1a\x98\xe2\x5f\x33\x4f\xd0\xe7\x2c\xdf\x6f\x53\xd8\xf6\x96\x1b\x4e\xe2\x07\x1d\x02\x2e\xc2\x0f\x68\xc2\xaf\x5c\xb0\x7e\x9b\xcd\xbb\x80\xbd\x35\xb7\x3a\x70\x85\xdc\xb1\x7b\xff\x3d\x3d\xb3\x51\xe6\x4c\xb2\x3b\x0a\xfd\x0d\xe8\xa4\x0f\xad\x9a\x2a\x71\xcb\x42\x86\xd0\x76\xd6\xf1\xf4\xc2\xdc\xfc\x28\xb1\x7c\x05\x00\x00\xff\xff\x80\x22\x7a\x88\xf6\x00\x00\x00")

func plusUsedMdBytes() ([]byte, error) {
	return bindataRead(
		_plusUsedMd,
		"plus-used.md",
	)
}

func plusUsedMd() (*asset, error) {
	bytes, err := plusUsedMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plus-used.md", size: 246, mode: os.FileMode(420), modTime: time.Unix(1538310351, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf2, 0xb5, 0xf6, 0xb4, 0x29, 0x7f, 0x12, 0x9, 0x5d, 0x70, 0xef, 0x87, 0x17, 0x61, 0xbe, 0x89, 0x47, 0xcd, 0x83, 0x7c, 0x35, 0x4b, 0x79, 0xad, 0xe, 0x83, 0x6e, 0xaa, 0xcd, 0x17, 0x81, 0x15}}
	return a, nil
}

var _stubMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x90\x31\x6b\xe4\x40\x0c\x85\x7b\xff\x8a\x07\xd7\xdc\x16\x67\xf7\xd7\x1e\x47\xfa\xb4\x21\x04\xed\xcc\xb3\x67\xd8\xf1\xc8\x68\x64\x2f\xfe\xf7\xc1\x8e\x03\xe9\xc4\x03\xbd\xef\x93\xfe\xe0\x95\x33\xe7\x3b\x0d\xae\x88\x2c\x74\x42\x6a\x84\x71\x29\x12\x08\x29\x05\x9e\x88\xe6\xeb\x1d\x41\xe7\x99\xd5\x5b\xdf\xfd\xbb\x26\x78\x12\x87\x18\x91\x6b\x50\x33\x06\x87\x1a\xb2\x19\x0b\x37\xa9\x8e\x59\x1e\x44\x76\x24\xb1\xf8\x45\x31\x4a\x3c\x3b\x83\x46\xf6\x5d\x07\xbc\x1c\xe8\x8d\x45\x17\x5a\x83\x1f\x1b\x51\xc3\x7a\x10\xc4\xb3\x56\xbc\x6d\xb4\x1d\x8d\x96\x75\x6d\x65\x7f\xff\x9d\xdc\x97\xf6\x77\x18\x26\x2d\x52\xa7\x5e\x6d\x1a\xa2\x86\x81\xe3\xc8\xe0\x79\xe3\xc7\xa4\x7d\xf2\xb9\xfc\xba\x94\xc5\xf6\x5b\xdf\xfd\x3f\x6b\xae\x08\xb9\x9e\x16\x63\x2e\x44\x4b\xba\x96\x88\x3b\xd1\x74\xa6\xa7\x5c\x27\xec\xba\xe2\x79\xc6\xcf\xe3\x90\x6f\xf3\x43\x67\xd2\xa8\xe1\xa0\xfe\x34\xb9\xa2\x1b\xf2\x08\x4f\xb9\xe1\x49\x3b\x9e\x09\x1d\xc7\x1c\xb2\x14\x2c\x12\x1e\x32\xb1\xff\x0c\x00\x00\xff\xff\x3d\x89\x29\x03\x76\x01\x00\x00")

func stubMdBytes() ([]byte, error) {
	return bindataRead(
		_stubMd,
		"stub.md",
	)
}

func stubMd() (*asset, error) {
	bytes, err := stubMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "stub.md", size: 374, mode: os.FileMode(420), modTime: time.Unix(1538510253, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xba, 0x55, 0xc2, 0x6b, 0x10, 0xe7, 0x88, 0xb5, 0xa, 0x5d, 0xa3, 0x88, 0x5f, 0xbe, 0x65, 0x54, 0x72, 0x50, 0xbb, 0xb5, 0x54, 0x62, 0x3b, 0xf7, 0xd5, 0x24, 0xc1, 0x2, 0x59, 0x47, 0x20, 0x30}}
	return a, nil
}

var _useStringPlaceholderMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\xb1\x52\xf3\x40\x0c\x84\xfb\x3c\xc5\x36\x99\xbf\x8a\x8b\xbf\xa4\xa3\x81\x92\x86\x86\xce\xe2\xa2\xf3\x69\xe6\x7c\xf2\x9c\x64\x07\xbf\x3d\xe3\x73\x08\xa6\xd6\xee\x7e\xdf\xe8\x02\x5e\xb8\xae\x58\x28\xcf\x0c\x5f\x27\x46\xa0\x82\x4f\x46\xd4\x3a\x92\x3b\x5f\x71\x13\x4f\xf0\xc4\x98\x32\x05\xbe\x24\xcd\x57\xae\xe8\xcf\xe7\xa5\xef\xf0\x56\x18\x1a\xf1\xaa\xff\x0c\xe6\x95\xcb\xe0\xc9\x20\xd6\x0a\x6d\xcf\x56\x73\x1e\xbb\xd3\x87\xce\x48\xb4\x70\xbb\xd0\x75\xa1\xe2\x34\xb0\x3d\xda\x87\x30\xa4\xfc\xf0\xa5\x0c\x20\xc3\x8d\x73\x86\x44\xac\x3a\x63\xb6\x7d\x23\x68\xad\x1c\x7c\xd7\xba\x5b\x45\xad\x38\x31\x85\xd4\xe6\x3a\x3c\x6f\x52\x52\x06\xfb\x93\x12\xdb\xf4\xad\xef\xf0\xfe\xb0\x9c\x38\x48\x94\x70\xcc\x19\x28\x9b\x42\x63\xe4\x8a\x51\x2b\x43\x27\x17\x2d\xb6\x63\x0e\x86\x9e\xc8\xef\xc4\x17\xad\xe0\x2f\x1a\xa7\xcc\xbf\x4f\x3b\xd8\xc5\xac\xe4\x8d\x1f\xfb\xf6\xeb\xc6\xd8\xf1\xeb\xde\xa8\x1c\xc4\x44\xcb\xd3\x96\xea\xfe\xc7\xbe\xfb\x0e\x00\x00\xff\xff\x52\xe5\x7c\x2b\xa7\x01\x00\x00")

func useStringPlaceholderMdBytes() ([]byte, error) {
	return bindataRead(
		_useStringPlaceholderMd,
		"use-string-placeholder.md",
	)
}

func useStringPlaceholderMd() (*asset, error) {
	bytes, err := useStringPlaceholderMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "use-string-placeholder.md", size: 423, mode: os.FileMode(420), modTime: time.Unix(1539198863, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x18, 0xb8, 0x45, 0xe7, 0x4e, 0x3c, 0xf9, 0xa9, 0x3f, 0x85, 0x9a, 0xde, 0x7e, 0x79, 0xdc, 0x78, 0xd2, 0xee, 0x1b, 0xdc, 0x27, 0xd2, 0x47, 0xb4, 0x60, 0x8, 0x31, 0x67, 0xec, 0x68, 0xe7, 0x4a}}
	return a, nil
}

var _wrongCommentFormatMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xca\xc1\x0d\xc2\x30\x0c\x05\xd0\x55\xfe\x85\x23\x1d\x83\x09\x58\x20\x75\x7e\x71\x44\x62\x4b\xb6\x0b\x62\x7b\xa4\x5e\x9f\xde\x1d\x4f\x25\xc4\xd7\xa2\x15\x3a\x53\x62\xec\xc3\x5e\x28\x25\x6e\x89\x54\x3f\x67\x87\xb6\x0f\x2f\x12\x8f\xa0\x14\x0e\x8f\xd5\x6a\xc3\xc3\x03\x9d\x47\x1b\x33\xf1\xf3\x13\xd2\x0c\xa2\x94\xf7\xb5\x93\x52\xc3\x0d\x3b\xa7\x7f\xb7\x7f\x00\x00\x00\xff\xff\x82\x89\x6a\xf2\x6c\x00\x00\x00")

func wrongCommentFormatMdBytes() ([]byte, error) {
	return bindataRead(
		_wrongCommentFormatMd,
		"wrong-comment-format.md",
	)
}

func wrongCommentFormatMd() (*asset, error) {
	bytes, err := wrongCommentFormatMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wrong-comment-format.md", size: 108, mode: os.FileMode(420), modTime: time.Unix(1538428444, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe7, 0xfd, 0x73, 0x49, 0x20, 0x9e, 0x44, 0x25, 0x47, 0x46, 0xe4, 0xe4, 0x61, 0x38, 0x71, 0x5a, 0x9c, 0x19, 0x27, 0x24, 0x3d, 0x4, 0xdd, 0x45, 0xf7, 0x2e, 0x24, 0x9, 0xd9, 0xb8, 0xe8, 0x3c}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"comment-format.md": commentFormatMd,

	"generalize-name.md": generalizeNameMd,

	"minimal-conditional.md": minimalConditionalMd,

	"missing-comment.md": missingCommentMd,

	"plus-used.md": plusUsedMd,

	"stub.md": stubMd,

	"use-string-placeholder.md": useStringPlaceholderMd,

	"wrong-comment-format.md": wrongCommentFormatMd,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"comment-format.md":         &bintree{commentFormatMd, map[string]*bintree{}},
	"generalize-name.md":        &bintree{generalizeNameMd, map[string]*bintree{}},
	"minimal-conditional.md":    &bintree{minimalConditionalMd, map[string]*bintree{}},
	"missing-comment.md":        &bintree{missingCommentMd, map[string]*bintree{}},
	"plus-used.md":              &bintree{plusUsedMd, map[string]*bintree{}},
	"stub.md":                   &bintree{stubMd, map[string]*bintree{}},
	"use-string-placeholder.md": &bintree{useStringPlaceholderMd, map[string]*bintree{}},
	"wrong-comment-format.md":   &bintree{wrongCommentFormatMd, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
