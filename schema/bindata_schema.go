// Code generated by go-bindata.
// sources:
// schema.json
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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5b\x5f\xaf\xdb\xb6\x15\x7f\xbf\x9f\x82\x70\x07\x14\x28\x64\xdf\x66\x2b\xf2\xd0\x3e\x05\xe9\x36\x0c\x58\xbb\x01\x6d\x9f\x8a\x3c\x50\x12\x6d\x31\x57\x22\x55\x92\xb2\xe3\x05\xf9\x50\xfb\x22\xfb\x4c\x3b\x87\x14\x69\xd9\x96\x64\x2a\x96\xdd\x9b\x22\x7a\x48\x7c\xa9\xc3\xc3\x1f\x0f\x0f\xcf\x3f\x52\xef\x1f\x08\x3c\x8b\x3f\xe9\xac\x60\x15\x5d\x7c\x4b\x16\x85\x31\xf5\xb7\x8f\x8f\x6f\xb5\x14\x4b\xd7\xba\x92\x6a\xf3\x98\x2b\xba\x36\xcb\xaf\xbf\x79\x74\x6d\x5f\x2c\x12\xd7\x33\x67\x3a\x53\xbc\x36\x5c\x0a\xec\xfd\x5a\x56\x75\xc9\x0c\x23\x8e\x8c\xac\xa5\x22\x35\xaf\xd9\x96\xff\x87\x54\x4c\x6b\xba\x61\x7a\xe5\xfb\x9a\x7d\xcd\xb0\x93\x4c\xdf\xb2\xcc\xf8\xd6\x5a\xc9\x9a\x29\xc3\x99\x86\x77\x0e\x9f\x6d\xcf\x4a\xce\x84\x39\x6a\x73\xd8\x15\x5b\x23\x97\x2f\x1e\x73\xb6\xe6\x82\x23\x14\xfd\xd8\x52\x07\xd2\x0f\xc9\x81\x13\x13\x5b\xae\xa4\xa8\x80\x40\x9f\xf3\xf3\xa8\xa8\x52\x74\xbf\x48\x8e\x5f\x56\x5c\xfc\xc3\xb0\x0a\xbb\xbd\x38\x79\xc5\xdb\xf6\xf7\x43\x90\x3a\xc3\x2e\xc8\x87\x5e\x64\xa5\xdc\xf0\x6c\xa9\x0d\x35\xec\x9e\xc8\x3a\xc3\x0e\x21\xcb\xa9\xa1\x9a\xdd\x55\x5e\xed\x90\x43\x88\x40\x4f\x32\xd0\xa8\xbb\x0a\xaa\x1d\x73\x08\x52\x26\xab\x8a\xdf\x55\x46\x6e\xc4\x71\x3c\xcb\x8a\x19\x7a\x77\x4c\x6e\xd4\x01\x60\xfb\xa6\x5a\xd6\x4f\x9b\xdb\x82\xea\x62\xf2\x03\x76\xf1\x3c\x74\x50\x2d\x68\x9e\x5b\x52\x5a\xfe\xbb\x6b\x7f\xd6\xb4\xd4\xac\x25\x51\xec\xb7\x86\x2b\x96\x43\xf3\xaf\xde\x1a\xbd\x09\x66\x30\x8c\x15\x67\xb5\x0c\x37\xa5\x9d\xe9\x6b\x4b\x40\xb8\x00\x53\x59\x51\x6b\x47\x4f\xe6\x76\x62\x62\x5f\x11\x5d\xd3\x8c\x59\xd3\xca\x68\x56\x78\xbb\xba\x04\xe5\xcc\x9b\x8c\x8b\x0d\x71\x83\x12\x23\xc9\x86\x6f\x19\x01\x95\x24\x82\x56\x2c\x21\x5b\xa6\x34\xb0\x49\x08\x15\x39\x29\x65\x46\x4b\x62\x38\xf4\x37\xb4\xaa\x89\x5c\x93\x5d\xc1\x04\xa1\x9e\x23\xd9\x51\x4d\x32\xc5\xc0\x2c\xe4\xab\x53\x54\xbd\xc6\x3b\xbc\x1d\x30\xe2\xe1\x3d\xc2\xe9\x7d\x73\xc4\x5b\x1b\x05\xd3\x39\xe1\x3d\x24\x97\x9f\x0b\x66\x67\x89\xf3\x30\xf0\xdb\x4f\xe2\x20\x16\x23\x25\xcc\xb7\xa0\xc6\x4f\xaa\x4b\xb7\x22\x3f\xab\x3d\x8a\x2c\x65\x24\x95\xa6\x20\x8d\xe0\xbf\x35\xcc\x4a\x2a\x83\x65\xe5\xda\x80\x4c\x57\x8b\x33\x2c\x1f\xce\xe1\x2d\x5a\x39\xcf\x3c\xc3\x57\x7e\xfd\x88\x68\xaa\x94\xa9\xf1\x99\x46\x42\xad\x78\xa6\x24\x2a\xc1\x65\xb0\x5c\x18\xb6\x61\x2a\x12\xed\x2f\x82\xbf\xb3\xda\x95\x80\x72\x13\x3b\x8c\x66\x20\xc9\x5c\x27\x41\xd5\x10\x7d\x47\xf3\x89\x63\x91\xc2\xca\x70\x71\x34\x35\xd4\x44\x99\x6a\xa6\xb6\x2c\x4f\x60\x89\x32\x5c\x68\xa3\x1a\x60\xce\x4c\xb6\x22\x3f\x48\x6d\x42\x80\x41\xa8\x82\x00\xc4\xee\xaa\xaa\x62\x39\x87\xa5\x2e\xf7\xd0\xab\x81\x86\x35\x11\xd2\x24\xc0\x9b\x6b\x0b\x8e\xe8\x42\x36\x65\x8e\xab\x6e\x95\x02\x51\x38\xd0\x5f\x01\xe1\x57\xc8\x26\xb7\x0d\x7d\xe2\x7c\x18\x11\x6e\x84\x45\x09\xa4\x5d\xcb\xd2\xee\x8d\xe4\xa0\x44\x49\x77\x91\xc8\x9b\x4b\x21\xcd\x88\xad\xf9\x6b\x87\xea\x92\x91\x51\xac\x56\x0c\x85\xe8\x64\x02\x2b\x46\x05\xe9\x0c\x43\x96\xa4\x2e\xf6\x9a\x83\x0d\x01\xbb\xc2\x95\x69\xf0\x07\xd8\x24\x58\x62\x43\xb9\x60\x6a\x66\x93\xd1\xf6\x1e\xd0\x51\x06\x5b\xc2\x89\xcf\xa3\xb2\x22\x74\xb8\xf0\x67\x80\x05\x22\x1c\x54\xe0\x35\x6d\x4a\x14\x60\xe8\x18\xb5\x85\x64\x3f\xe2\x13\x5c\x3b\x2e\x72\xb9\xd3\x88\xa5\xe4\xa2\x79\x87\x3f\x72\xaa\xa0\x19\x7f\xad\x15\x63\xa9\xce\xf1\x27\x18\x9e\x77\x51\x20\x2d\x61\x14\x42\xd0\x45\x85\x71\xcb\x20\xcc\x7e\x4f\xee\xbb\x45\x8d\x21\x78\xf6\x34\xbb\xc1\x43\xa6\xd6\xa8\xf3\x1c\x74\x8e\xaf\xf7\xd6\xba\xe1\xe6\xed\xa8\xe2\x8a\xfc\xd8\x92\x81\xbb\x02\x2d\x85\xbd\xdc\x68\xb0\x21\x60\x33\x60\x5a\x4c\x61\x9c\x6f\xdd\xa5\x9b\x9b\xe5\xc0\x48\xc1\x99\xa2\x2a\x2b\x50\x57\x80\xae\xb4\x7a\xae\x0b\x5e\x43\x7f\xb3\x63\xe8\x08\x8f\x15\x1e\x3d\x01\xba\x52\xd4\x24\x64\xe9\x2c\x0a\x58\x09\xcb\x1b\x8c\x47\x49\x04\xf4\x94\xea\xc9\x0b\x8e\xa6\x25\x3b\x62\xad\x01\x2a\x03\x64\xd8\x29\x93\x0a\x68\x6a\x89\xf6\x45\x02\xf3\xbd\x63\x61\xa3\xf0\xa3\x19\xc9\x25\x7b\x07\xde\x07\x4d\x22\x25\x9b\x52\xa6\x38\x10\xbe\x0a\x61\x00\x2d\xcb\x2e\x52\x4d\x9e\x84\xdc\x09\xcf\xd6\xe7\x5f\x5c\x00\x6f\x91\x81\x69\xd3\x12\x1a\xb3\x27\x14\x43\x45\x76\x5c\x83\x79\x8c\xf4\x16\xb0\x49\xb7\xb0\x14\x2a\x7a\xa1\xa3\xb8\x8e\xa6\x3d\x67\x9c\xfb\x62\xc3\x40\x34\x1c\x23\x06\x92\x8f\x4d\x87\x2e\xc8\x65\x20\x19\xb9\x3f\xfc\x9e\x24\x65\x0c\xfa\x60\x66\x77\x7f\xe4\x3d\x19\x5f\x40\x3e\x97\xaf\x85\x1d\xf1\x2f\x1c\xfa\xd7\xb3\x21\xde\x9f\x39\x62\xb4\x68\xe4\x4d\x9f\xd4\xce\x68\xbd\xa9\x44\xf2\x23\xea\x7e\x7f\x3d\x64\x90\xaf\xf4\x91\x05\x84\x41\x73\x84\xd6\x2e\x22\xb3\xe5\x20\xe0\xb8\xb4\x2c\xa3\x94\x89\xd7\xdb\x6f\xe6\x1b\xdc\x72\x8b\x1d\xf7\xe5\xac\xe3\xbe\x9c\x1c\xee\xc5\xeb\x55\x58\xa5\x48\xdd\xb2\x62\x88\xa7\x7d\xd9\xa3\x84\x13\xf6\xcb\x50\x7c\xb9\x84\xa8\xe5\xdc\xc1\x5f\x50\xd8\x13\x87\xfe\x7d\x1b\xdf\x43\x88\x4e\x90\x1d\x49\x29\xb8\xa1\xb5\x92\x15\x34\x80\x5b\xb7\xa9\x8b\xf5\xce\x5d\xc7\xeb\x52\x36\x17\xbf\x69\x70\xc1\x2b\xf2\xe5\xeb\xf6\xaf\x2f\xad\x57\xd4\x21\x14\x05\x36\x10\xb3\x6a\x66\xdb\xa4\xd6\x1c\x3c\x30\x64\x76\x18\x30\xd0\x1a\x76\x0f\xe6\xcb\x02\x3c\xb0\x05\x00\x4e\xd8\x66\x50\xb2\xd9\x14\xe6\x2c\x5a\xbd\xb4\xd7\x3a\x5b\xf8\x52\xfc\x74\x41\x75\x86\x2c\xc9\xb5\x8b\xd6\xf5\x63\x33\x1b\x9a\xd8\x60\x3c\xe5\x82\xaa\xbd\x8b\xbf\x73\xe6\x62\xdf\x54\x61\x53\xc7\x30\x0e\xc8\xc5\x01\xa1\xa6\x70\x22\x3e\xd9\xc4\xa3\x31\x68\x1c\x75\xb9\x51\xb2\xa9\xe3\xe9\x4f\x92\xac\x91\x3a\xaf\xdb\x2b\xfd\x83\xda\xf9\x0f\x09\xb6\x2b\xdc\x39\xfd\xed\xd9\xf4\xe2\xac\x6a\x7e\x19\x64\xaf\xfe\x04\x2a\x29\xd8\x80\x55\xf4\x4f\x3f\xfb\xc0\xe0\x82\x22\x9e\xd1\xb7\x05\xd1\xc3\xa4\xc9\xe8\xac\xbb\x4f\x8f\x04\x8e\x58\x1f\x6f\x53\x5f\x79\x1d\xc8\xd3\x42\xaf\xa8\x5d\x3b\x01\xca\xcc\xf2\xea\x54\xaa\x62\xb4\x24\x12\xa4\x65\x7d\x2c\x2f\x3f\xd0\xa7\x2e\x30\xcd\xaa\xad\xcb\x7f\x6e\x2b\xaf\x76\x9c\x1b\x89\xab\xf7\x4d\x9c\x45\x7e\x46\xf9\x02\xc4\x04\x62\x89\x78\xa6\x67\x0c\xc7\xc2\xb6\x5e\xe6\x7a\x87\x7b\xc0\x33\xd1\xdd\x0e\x9b\xc9\x01\x11\x47\x6a\x6d\xa8\xb3\x47\x39\xb8\x53\xb0\xef\xbb\x9e\xbc\x2d\xa4\xe6\x03\x81\x68\xe8\xab\x9b\x14\x53\xb8\xc9\x63\x62\x71\x59\xd1\xcc\x8c\x95\xcd\x03\xf1\x01\x96\xda\x61\x50\xa1\x64\x5f\x3c\xe1\x9f\xb1\x61\x71\xc9\x7e\x64\xa3\xfe\xdc\xae\xaa\xe8\x4f\x4b\x47\xd8\x9f\x06\xe6\x9d\xf9\x25\xad\x68\x92\xc3\xf8\x83\xf5\xbe\x49\xfb\xbb\x2f\x43\xf8\xac\x40\x3d\xc4\xf3\x2a\x10\x1e\x75\x5c\xd4\x20\x57\xd6\xbd\x91\x0a\xfd\x32\x5a\x33\xfe\xc3\xe8\xd0\x9a\x97\x43\x59\x73\xe8\xf7\x59\x7f\x0e\xcc\x7f\x17\xfd\x89\x28\x7f\x05\x93\x3a\x47\x39\xa1\x36\x52\xb5\xe7\x53\xb2\x31\x1b\x89\x55\x84\xb6\x10\x8f\x55\x03\xc1\xec\x8c\xbf\xc3\x0a\x03\x55\x1b\x66\x42\x5a\x7e\x2f\x27\x7c\xb1\x2a\x77\x36\xfb\xd1\x42\x55\xa0\x8e\xaa\xd2\xf9\x67\x4c\x07\x6b\xa9\x4e\x36\x8c\x3f\xea\x1d\xef\xa6\xa4\x91\x53\x22\xe1\x38\x35\x0d\xe2\x4a\x5a\x64\x89\x1f\xea\xb9\x1a\xb8\xd1\xaa\x67\xa0\xfa\xd8\xc5\xed\xaf\x82\x8e\xcc\xe6\x80\xff\x59\xad\xab\x9d\xc7\xa7\xb4\xa6\xc3\x15\xe5\x40\x75\xc5\x9a\xf6\x54\x98\x47\x66\x73\xc0\xff\xdc\xd6\xf4\xe5\xcd\xd6\x34\xd6\x91\x34\xad\xfb\xbc\x89\x27\x71\xb7\xa4\x70\x88\x7e\x67\x62\x53\xc7\x89\xd5\xd3\x91\xa2\xe6\x43\xbf\x04\x7a\x2e\x3e\xce\x5d\xd2\xfd\xdd\x0e\x61\x3f\xae\x36\xd9\xe8\xc1\x2a\xcc\xad\x4a\xb7\x13\xcb\xc5\xd9\x2e\x8f\x27\xf6\x47\x9f\xd1\x1d\x6a\x9e\xc7\x9b\x81\x45\x69\xaf\xaf\x3d\x83\x92\x0d\xc6\x5e\xcb\x16\xcd\x95\x45\x9b\x23\x85\x4d\x9c\x44\xae\xaf\xe1\x74\x01\xce\x75\xd8\x85\x07\x5a\x8e\x25\x53\x44\xf3\x9c\xb5\x46\xc6\x06\xe4\xcb\x76\x3f\x77\xcc\xcb\x8a\xbc\x6e\x14\x5e\x5b\x29\xf7\xe4\x6d\xa3\x0d\x41\xf3\xaa\xed\x05\x14\x2d\xb3\x27\x86\x77\x4f\x4a\x09\x2d\xc0\xa6\x94\x5b\x66\x8f\xb4\x7a\xee\x68\xde\x2c\xa2\xed\xcb\xcd\xac\x0b\xb8\x90\x9b\x5d\xeb\xbc\xe2\x9c\xf1\x98\x12\x07\xe2\xcb\xca\x1c\x48\x27\xd9\x2a\xff\x5c\xe3\x55\x7d\x66\xf6\x69\x44\x4a\x7d\xca\xd0\xde\x61\x1b\x5f\xd5\x48\x2f\x38\x82\xdc\xf2\x19\x90\x5d\x5f\x49\x37\xf4\x99\x3d\x10\xe9\xd8\xef\x2b\x4c\xc6\x2b\xd2\xf2\xb1\x67\xd6\xe2\xc9\xdf\xf0\x05\x33\x51\xd9\xeb\x61\x78\x8f\xd5\xfd\x4f\x52\x25\x69\x5e\xee\x57\x64\x4b\x15\xa7\xc2\x90\x17\x04\x4f\xb9\x89\x66\xe5\x7a\xa9\x65\xa3\x32\x96\x7b\x76\xdf\x05\xa2\x3f\x3b\xa2\x75\x23\xac\xa5\xc1\x08\xe7\x8c\xe6\x2f\x8e\xa6\x6d\x3f\x5c\xe8\x7c\xe6\x55\x6b\x77\xa9\x7b\xe9\xef\x32\x77\xc3\xed\x3c\xbc\x18\xe5\xb0\x01\x03\xad\xb9\x3e\x55\xe6\xff\xfd\xf7\xfc\x46\x85\x7f\xe2\x94\xd2\xa7\xb3\x7e\x80\xe7\xba\x95\x9f\x81\xe0\xe3\x4c\xfc\xe8\xa9\xf3\xd4\x69\x07\xfa\x49\x57\x3a\xc6\xc7\xb6\xfc\x86\x8c\xc2\xc5\xe9\xc5\x78\xb0\xd0\x29\xde\x93\x85\x2e\x1f\xe5\xd1\xfc\x13\x33\x73\x2d\x68\x7d\x51\x21\x46\xd9\x4c\x3b\x2d\x0d\x25\xbd\x83\xd0\x93\x2e\x8a\x7b\x1e\xa3\x7e\x36\x09\xf3\xe4\x5b\x07\xa8\x53\xee\x00\x85\x5e\xee\x1c\x60\x82\x4d\xb9\x61\xd8\x38\xf9\xba\xab\x7f\xae\xd6\xb0\x7b\x96\x62\xfc\xcd\x9b\x59\x72\x26\x08\x64\x6c\x0c\x63\xc3\x0f\x25\x4b\xe2\xf8\x13\xc7\x66\x45\xfe\x26\x15\x11\x72\x97\xb8\x04\x69\x83\xb7\x02\x61\x22\x29\x4d\x21\x34\x72\x77\xe2\x77\x74\x8f\x9f\x1a\x09\x68\xf8\xe9\xf5\x0f\xda\x5d\x23\x2c\x28\xa4\x4b\xd4\xc8\x8a\x67\x2d\x47\xb0\xe9\xe4\xfb\x57\x7f\x9f\x7c\xff\x4f\x17\xf4\x45\x7c\xba\xae\x58\x2d\x35\x37\x52\xed\x27\xd5\x04\xe2\xb7\xde\x82\x36\xa6\x90\x13\x4a\x22\xb0\x3f\xec\x82\xc4\x57\x1c\xa8\xea\xfd\x62\xfa\x6c\xa5\x63\x8b\x08\x5f\xcf\x52\x1b\x9a\x50\x31\xb0\x4b\x96\xb4\x82\x4d\x82\xc8\x92\x83\x30\x92\xc3\x34\x93\xa3\x45\x9b\xe3\x72\x48\xc4\xc7\xb8\x13\x37\x09\x32\x43\x2b\x42\x68\x2a\x1b\x03\x7b\xa6\xdd\x24\x4b\xc8\x0e\xa8\xc8\x0a\x20\x81\xad\x60\xe8\x06\xf6\x93\xd6\x32\xb3\x07\xf0\x64\xc7\x4d\x41\x38\x7e\x05\xc7\xb4\xc1\x0c\xc0\xd5\x50\xdc\xa7\x74\xb7\xdd\x06\x38\xe2\x4f\xbd\x17\x52\x03\x49\xa7\xa2\x00\x98\x99\xfd\x24\xaa\x66\x22\xb7\xf5\x75\x08\x25\x28\x2f\xed\x25\x82\xa8\x2a\x21\x4e\xfd\x13\x28\x66\xfa\xd5\x7a\x6e\x50\xa7\xee\xad\xeb\x77\xc9\xc5\x2f\xc3\x23\xf3\xe8\x9a\x66\x4f\x74\xc3\x12\xbc\x08\xde\x88\x9c\x29\x6d\xa4\xcc\x49\xba\xb7\xf9\x34\x8c\xe2\x29\x48\x45\x05\xfc\xa7\xdc\x27\x63\x5c\x10\x55\x57\xcb\x94\xe2\x1f\xff\xc4\xcf\xf4\x48\xce\x51\x40\x69\x63\x7d\xf6\xe4\xed\x31\x29\x7c\xba\x70\x1b\xb3\xaf\xae\x5d\xcb\xac\x88\xaf\xa8\xc1\x9a\x95\x0c\xe6\x36\xc1\xab\xa8\x53\xfe\x03\x47\x15\xa3\xca\xd1\xf3\x49\xab\x47\x92\xf8\x49\x24\xed\x60\x1f\xa9\x43\x0f\xee\xdf\x0f\x0f\xff\x0f\x00\x00\xff\xff\xf7\x92\x52\x88\xd3\x44\x00\x00")

func schemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaJson,
		"schema.json",
	)
}

func schemaJson() (*asset, error) {
	bytes, err := schemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"schema.json": schemaJson,
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
	"schema.json": &bintree{schemaJson, map[string]*bintree{}},
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

