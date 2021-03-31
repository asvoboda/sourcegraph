// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1000000000_init.down.sql (19B)
// 1000000000_init.up.sql (19B)
// 1000000001_init.down.sql (233B)
// 1000000001_init.up.sql (611B)
// 1000000002_add_indexes.down.sql (399B)
// 1000000002_add_indexes.up.sql (369B)
// 1000000003_pg_stat_statement_ext.down.sql (62B)
// 1000000003_pg_stat_statement_ext.up.sql (68B)
// 1000000004_comments.down.sql (16B)
// 1000000004_comments.up.sql (3.631kB)
// 1000000005_diagnostic_counts.down.sql (158B)
// 1000000005_diagnostic_counts.up.sql (690B)

package migrations

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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

var __1000000000_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\x48\xcd\x2d\x28\xa9\x54\xc8\xcd\x4c\x2f\x4a\x2c\xc9\xcc\xcf\xe3\x02\x04\x00\x00\xff\xff\x32\x4d\x68\xbd\x13\x00\x00\x00")

func _1000000000_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000000_initDownSql,
		"1000000000_init.down.sql",
	)
}

func _1000000000_initDownSql() (*asset, error) {
	bytes, err := _1000000000_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000000_init.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9c, 0x46, 0xd1, 0x31, 0xb9, 0x68, 0x19, 0xcc, 0x70, 0xb6, 0x7, 0x20, 0x2e, 0x6a, 0x4d, 0xf1, 0xce, 0xd0, 0xc8, 0xda, 0x50, 0xce, 0x8c, 0xee, 0x52, 0x36, 0x80, 0xd0, 0x5a, 0xd2, 0x7a, 0x82}}
	return a, nil
}

var __1000000000_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\x48\xcd\x2d\x28\xa9\x54\xc8\xcd\x4c\x2f\x4a\x2c\xc9\xcc\xcf\xe3\x02\x04\x00\x00\xff\xff\x32\x4d\x68\xbd\x13\x00\x00\x00")

func _1000000000_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000000_initUpSql,
		"1000000000_init.up.sql",
	)
}

func _1000000000_initUpSql() (*asset, error) {
	bytes, err := _1000000000_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000000_init.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9c, 0x46, 0xd1, 0x31, 0xb9, 0x68, 0x19, 0xcc, 0x70, 0xb6, 0x7, 0x20, 0x2e, 0x6a, 0x4d, 0xf1, 0xce, 0xd0, 0xc8, 0xda, 0x50, 0xce, 0x8c, 0xee, 0x52, 0x36, 0x80, 0xd0, 0x5a, 0xd2, 0x7a, 0x82}}
	return a, nil
}

var __1000000001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x72\x75\xf7\xf4\xb3\xe6\xe2\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\xc8\x29\xce\x4c\x8b\x4f\x49\x2c\x49\x8c\xcf\x4d\x2d\x49\x04\x31\xac\x09\x29\x4c\xc9\x4f\x2e\xcd\x4d\xcd\x2b\x29\x26\xa8\xb2\x28\xb5\xb8\x34\xa7\x24\x3e\x39\xa3\x34\x2f\x9b\xb0\xea\x94\xd4\xb4\xcc\xbc\xcc\x92\xcc\xfc\x3c\x62\x4c\x4e\x4b\x2d\x4a\xcd\x4b\x4e\x2d\xb6\xe6\xe2\x72\xf6\xf7\xf5\xf5\x0c\xb1\xe6\x02\x04\x00\x00\xff\xff\xdc\x1f\x48\x24\xe9\x00\x00\x00")

func _1000000001_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000001_initDownSql,
		"1000000001_init.down.sql",
	)
}

func _1000000001_initDownSql() (*asset, error) {
	bytes, err := _1000000001_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000001_init.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe5, 0xa1, 0x76, 0x54, 0xd0, 0xb0, 0xb9, 0xdc, 0x93, 0x84, 0xd, 0xb5, 0xa3, 0x55, 0x91, 0xef, 0xc0, 0xe4, 0x89, 0x62, 0x20, 0x2, 0x6e, 0x92, 0x58, 0x73, 0xb3, 0x9b, 0xcd, 0xde, 0xa, 0x86}}
	return a, nil
}

var __1000000001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\xcf\xc1\x4a\x03\x31\x10\x06\xe0\x7b\x9e\x62\x8e\x0a\xbe\xc1\x9e\xda\x12\x25\xb0\xdd\x82\x8d\xe0\x2d\xc4\x64\xe2\x0e\x36\xd3\x92\x4c\x60\x7d\x7b\x71\x41\x50\x91\x65\x05\xbd\x0d\xc3\x0f\xdf\xff\x6f\xf5\x9d\x19\x3a\xa5\x76\xf7\x7a\x63\x35\xd8\xcd\xb6\xd7\x60\x6e\x61\x38\x58\xd0\x8f\xe6\x68\x8f\x70\xaa\x94\x5c\xf4\xe2\x5d\x46\xf1\xef\x07\x5c\xc5\x96\x2f\x8e\x22\x10\x0b\x3e\x63\x99\xe3\xc3\x43\xdf\xdf\x00\xb7\xec\x0a\xd6\x76\x12\x17\xc6\xc6\x2f\xf5\x23\x73\xdd\xad\x43\xe2\x39\xb4\x8c\x2c\x75\x49\xb9\x78\x19\x41\x70\x92\x4f\xbf\xb9\xd9\xd3\xab\xa0\x5f\x4b\x7d\xed\xb9\xc0\x51\x9c\x7e\xf8\xfe\x1e\x8c\x98\x88\x49\xe8\xcc\x8b\x5c\x0d\x23\x66\xfc\xbe\x8f\x22\xb2\x50\x22\x2c\x7f\xb1\x3c\x61\x41\x0e\xf8\x7f\x3d\xd4\xee\xb0\xdf\x1b\xdb\xa9\xb7\x00\x00\x00\xff\xff\x8f\xd0\x35\x75\x63\x02\x00\x00")

func _1000000001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000001_initUpSql,
		"1000000001_init.up.sql",
	)
}

func _1000000001_initUpSql() (*asset, error) {
	bytes, err := _1000000001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000001_init.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8b, 0x99, 0xc4, 0xe6, 0xb3, 0x64, 0x2, 0x61, 0x25, 0xa1, 0x5, 0x57, 0x5f, 0xe5, 0xb2, 0xd6, 0xf0, 0xf4, 0x4c, 0xbc, 0x78, 0x2f, 0x63, 0x64, 0x2a, 0x95, 0x77, 0x8d, 0xde, 0x83, 0x6b, 0xde}}
	return a, nil
}

var __1000000002_add_indexesDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xce\x4d\xaa\xc3\x20\x14\xc5\xf1\xb9\xab\x70\x1f\x8e\x4c\x9e\x3c\x84\x44\x8b\x75\x2e\x62\xae\x54\x92\x98\x12\x75\xd0\xdd\x97\x42\x3f\x28\xb4\xb7\xb3\x33\xf8\xf3\xe3\x74\xe2\x5f\x2a\x46\x08\x1f\xac\x30\xd4\xf2\x6e\x10\x74\x29\x29\xba\xc9\x57\xef\x56\xa8\xfe\x36\xe8\x9f\xd1\x07\xda\x6b\x75\xb4\x86\x4b\x65\x3f\x24\xee\x3c\xc3\x85\x7d\x71\xa6\x2d\xb4\x15\x72\x2d\x08\xf4\x6c\x50\x69\x87\xd2\x96\xea\xc2\xa9\xe5\x19\xd3\xde\x3a\xfc\x1b\xc4\x94\x53\x4d\x5b\x46\xdf\xbd\xaa\x1f\xff\x22\xec\x90\x03\xe0\xe7\x1e\xd1\xdd\x22\xbd\x1e\x47\x69\x19\xb9\x06\x00\x00\xff\xff\xc1\x8c\x87\x89\x8f\x01\x00\x00")

func _1000000002_add_indexesDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000002_add_indexesDownSql,
		"1000000002_add_indexes.down.sql",
	)
}

func _1000000002_add_indexesDownSql() (*asset, error) {
	bytes, err := _1000000002_add_indexesDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000002_add_indexes.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x93, 0xee, 0x39, 0x5a, 0x7b, 0x6, 0x8a, 0xa6, 0x8f, 0x3d, 0x19, 0x4b, 0x83, 0x87, 0xb, 0xd4, 0xf3, 0x81, 0x4b, 0x12, 0x4, 0x67, 0x66, 0x43, 0xe6, 0xf1, 0xe4, 0x1, 0x82, 0x3c, 0x2d, 0x8e}}
	return a, nil
}

var __1000000002_add_indexesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xce\x41\xaa\x83\x30\x10\xc6\xf1\x7d\x4e\x31\xcb\xf7\xa0\x37\x70\x15\x6b\x28\x52\x6d\x8b\xb8\x71\x15\x42\x32\xc1\xa1\x26\x8a\x99\x40\x8f\x5f\x3c\x40\x84\xee\xbe\xcd\xf7\xe3\x5f\xab\x5b\xfb\xa8\x84\x90\xdd\xa8\x06\x18\x65\xdd\x29\x58\x12\x79\xed\x0c\x1b\x1d\x90\xcd\x31\x40\x36\x0d\xbc\x86\xb6\x97\xc3\x04\x77\x35\xc1\x9f\xcb\x61\xd3\xe4\xfe\xab\xc2\xd3\xad\x36\x07\x8c\x9c\x8a\xd7\x0b\x6c\x86\xe7\x22\xb0\x63\xca\x0b\x6b\x3b\xe7\xf8\x3e\x43\xc8\x7d\xca\x11\xe8\x29\x12\xd3\x1a\xcf\x84\x64\x67\x0c\x78\x48\x18\x99\x3c\xe1\x7e\x12\xe5\x71\xc7\x68\xf1\x67\x4f\x5c\x9f\x7d\xdf\x8e\x95\xf8\x06\x00\x00\xff\xff\x1a\xdb\xd6\xb5\x71\x01\x00\x00")

func _1000000002_add_indexesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000002_add_indexesUpSql,
		"1000000002_add_indexes.up.sql",
	)
}

func _1000000002_add_indexesUpSql() (*asset, error) {
	bytes, err := _1000000002_add_indexesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000002_add_indexes.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5a, 0xe0, 0x76, 0x5f, 0x41, 0x20, 0x0, 0x49, 0x2b, 0x9f, 0x69, 0xf1, 0x8c, 0x4b, 0x83, 0xe1, 0xd2, 0x8f, 0x69, 0xbf, 0x54, 0xae, 0x3, 0x14, 0xac, 0x3a, 0x66, 0x45, 0x46, 0xd6, 0x5e, 0xfd}}
	return a, nil
}

var __1000000003_pg_stat_statement_extDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x72\x75\xf7\xf4\xb3\xe6\xe2\x72\x09\xf2\x0f\x50\x70\x8d\x08\x71\xf5\x0b\xf6\xf4\xf7\x53\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x48\x8f\x2f\x2e\x49\x2c\x01\x13\xa9\xb9\xa9\x79\x25\xc5\xd6\x5c\x5c\xce\xfe\xbe\xbe\x9e\x21\xd6\x5c\x80\x00\x00\x00\xff\xff\xe2\x7c\xbf\xd5\x3e\x00\x00\x00")

func _1000000003_pg_stat_statement_extDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000003_pg_stat_statement_extDownSql,
		"1000000003_pg_stat_statement_ext.down.sql",
	)
}

func _1000000003_pg_stat_statement_extDownSql() (*asset, error) {
	bytes, err := _1000000003_pg_stat_statement_extDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000003_pg_stat_statement_ext.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x15, 0x9a, 0x91, 0x10, 0xf0, 0x2e, 0x95, 0xa2, 0x94, 0x35, 0x4a, 0xc0, 0x3a, 0xa, 0x76, 0xfd, 0xff, 0xb, 0xd9, 0x78, 0x99, 0x44, 0x57, 0xdb, 0x72, 0xa5, 0x75, 0x25, 0x61, 0x7e, 0xeb, 0x4d}}
	return a, nil
}

var __1000000003_pg_stat_statement_extUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x72\x75\xf7\xf4\xb3\xe6\xe2\x72\x0e\x72\x75\x0c\x71\x55\x70\x8d\x08\x71\xf5\x0b\xf6\xf4\xf7\x53\xf0\x74\x53\xf0\xf3\x0f\x51\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x48\x8f\x2f\x2e\x49\x2c\x01\x13\xa9\xb9\xa9\x79\x25\xc5\x20\x3d\xfe\xbe\xbe\x9e\x21\xd6\x5c\x80\x00\x00\x00\xff\xff\xca\x11\x32\x41\x44\x00\x00\x00")

func _1000000003_pg_stat_statement_extUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000003_pg_stat_statement_extUpSql,
		"1000000003_pg_stat_statement_ext.up.sql",
	)
}

func _1000000003_pg_stat_statement_extUpSql() (*asset, error) {
	bytes, err := _1000000003_pg_stat_statement_extUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000003_pg_stat_statement_ext.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x35, 0x68, 0x79, 0x4, 0x23, 0xd, 0x39, 0x6f, 0xf, 0x19, 0xb5, 0x29, 0x6f, 0x21, 0x43, 0xd, 0x1b, 0xfe, 0xb1, 0x4c, 0x22, 0xcb, 0xd5, 0x78, 0x6b, 0x89, 0xbd, 0x3b, 0x3a, 0x31, 0xb8, 0x4}}
	return a, nil
}

var __1000000004_commentsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x72\x75\xf7\xf4\xb3\xe6\xe2\x72\xf6\xf7\xf5\xf5\x0c\xb1\xe6\x02\x04\x00\x00\xff\xff\x80\xfb\x5a\xa5\x10\x00\x00\x00")

func _1000000004_commentsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000004_commentsDownSql,
		"1000000004_comments.down.sql",
	)
}

func _1000000004_commentsDownSql() (*asset, error) {
	bytes, err := _1000000004_commentsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000004_comments.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xec, 0x6c, 0xe7, 0xaa, 0xc6, 0xdb, 0x9b, 0x52, 0x49, 0xad, 0xde, 0x34, 0x6a, 0xa9, 0x90, 0x21, 0x97, 0xcc, 0xcd, 0x35, 0xee, 0xc6, 0xd1, 0x36, 0xf1, 0xa0, 0x6c, 0xe8, 0x73, 0x42, 0x77, 0x41}}
	return a, nil
}

var __1000000004_commentsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x96\x6d\x6b\x1c\x37\x10\xc7\xdf\xfb\x53\x0c\xf4\x45\x6c\x70\x6e\x93\x1a\x5a\x88\x29\x34\x71\x4c\x6b\xf0\x03\x34\x2e\x7d\x11\xca\x31\xb7\x9a\x5b\x89\x68\x35\x5b\xcd\xc8\xf6\x7d\xfb\x22\xed\xde\x7a\xaf\xb9\xf6\x7c\x34\xc5\xa1\xef\xf6\x41\x9a\xf9\xcf\x7f\x7e\x7a\x78\x77\xfe\xd3\xc5\xf5\xe9\xc1\xc1\xd9\xcd\xd5\xd5\xf9\xf5\x2d\xdc\x5c\xc3\xed\xdb\x77\x97\xe7\x50\xb3\x21\x17\x94\xfc\x5c\x6a\x4b\x2d\xce\x5b\xd7\x44\x54\xc7\x41\xe0\xe2\x03\xbc\xf8\x99\xbd\x11\x40\x10\x17\x1a\x4f\x50\xb3\x4f\x6d\x00\x51\x8e\x2e\x34\xa0\x96\x40\x14\x35\x09\xf0\xb2\xbc\xb5\x2c\x0a\x91\x6a\x0a\x0a\x63\x28\x40\x55\x6a\x3b\x9d\xbd\x38\x9d\x2a\x38\xbb\xb9\xfc\xf5\xea\xfa\x9f\x24\xcc\xee\x28\x4a\x0e\x90\xa5\xdc\xe6\x64\x65\x04\xac\x3f\xab\x45\x85\x7b\x94\x92\x5a\x31\x36\xa4\xff\x91\x10\xe3\xa2\xae\x8a\x8c\xdf\x2c\xa9\xa5\x08\x1c\x21\xb0\xee\xce\x05\x4b\x74\x9e\x4c\x4e\xf9\xb9\xfd\x5e\xdc\x72\x6e\x50\x71\x6e\xb8\x4e\x2d\x05\xed\x6d\xff\xa0\x1c\x49\x20\xd2\x92\x22\x85\x9a\x8e\xc1\xf2\x1d\x45\x50\x7a\xd0\x63\x68\x39\xb8\x4f\x14\x8f\x01\x83\x01\xe3\xb0\x09\x2c\xea\x6a\xc8\x81\x00\x17\x9c\x14\x10\x3a\x8c\xea\xea\xe4\xb1\x9f\x05\xeb\x04\x70\xef\xd4\x05\x40\x30\xa9\xed\xb6\x1b\xb1\x45\xd5\x2c\x8f\x9e\x3b\x33\x76\xc2\x19\x0a\xea\x96\x2e\x3b\xd1\x3b\x8e\x22\x5c\x3b\x54\x32\x25\x34\xb8\x50\x3e\x97\x60\xa9\xf3\x8c\x46\x40\x71\xe1\x09\x0e\x33\x33\xf4\x43\xcd\x6d\xe7\x49\xc9\x1c\x3d\x5d\x46\x87\x6a\x47\x0d\xe5\x65\xc8\xbe\x59\x63\x24\x8f\xea\xee\x08\x94\xb7\x6a\x8b\xcc\x7f\x03\xc1\xd6\xda\xb3\xaf\x39\xe9\x5b\x68\x78\xf1\x92\x42\x26\xc5\x40\x87\xab\x5c\x15\xd4\x1c\x96\x1c\xdb\xb2\x20\xfa\x74\x1f\xdf\x0f\x73\xdf\xa3\xe2\xef\x87\x56\xb5\x93\x37\x55\x25\x9c\x62\x4d\x4d\xc4\xce\xce\x6a\x6e\xab\xc6\xa9\x4d\x8b\xf2\x38\xf9\x35\x7d\xfe\xf1\x64\xf6\xed\x49\xf5\xb2\x5a\x78\x5e\x54\x14\x94\x62\x17\x9d\x50\x95\x31\x8d\x01\x7d\x35\x32\x5b\x49\x41\xa6\xca\xfa\xcb\x63\xa5\xab\x8e\x64\xd6\xf0\x37\x97\xaf\x4f\xde\x7c\x77\x04\xf9\x7d\x17\x85\x2d\x29\x8e\xd5\x0e\x10\xe6\x82\x42\x6a\x17\x7d\xa7\x23\x49\xf2\x0a\xb5\x4d\xe1\x93\x4c\x7d\xbd\x77\x6a\x9f\x48\xd5\x3a\xcb\xf3\x42\x35\xaa\x08\xa9\x9d\xf7\x65\xcd\x87\xb2\xfa\x5e\x2f\x38\x05\x93\x85\x74\xdc\x25\x5f\x34\xb8\x60\xe8\x81\x64\x43\x46\x09\xb6\x39\xbf\x57\xb4\xe4\xb8\xad\x86\x19\xdc\x5a\x27\x70\x87\x3e\x11\x38\x81\x24\x64\x32\x38\x16\xc5\x4e\x1c\xc8\x49\x06\x9c\xa6\x9e\xf7\x12\xf2\xf8\x7b\xeb\x6a\x9b\xff\xaf\x60\x41\x9e\x43\xb3\xab\xb9\x5b\x6a\x5c\x2b\x93\x75\x0e\x21\xdd\xd0\x50\xba\x9a\x35\x1c\x8e\x6b\x2b\xaf\xba\x63\x88\x18\x9a\x69\xc3\x8e\xa0\x43\x17\xa5\xdf\x92\x73\x13\x58\xa8\x3f\x1f\x68\xe7\x4a\xdb\x10\xf6\xbc\x50\x6c\x4a\x71\xe6\x61\x94\x91\x82\xfb\x23\x6d\xed\x45\xf6\x68\x90\xf0\x59\xab\xcf\xef\x28\xae\xb6\x9b\x0b\x5d\x24\xc9\x86\x8a\xe5\xe4\x4d\xdf\xff\xd2\x71\x27\x43\xe4\xc3\x96\x4d\xf2\xfc\x24\x68\xf7\x2c\x6d\xff\x3d\xed\x97\x32\xff\x2c\x4f\xff\x3a\xb6\xb5\xef\x5f\x3d\x75\x5b\x33\xb4\x74\xc1\x3d\xde\x6a\x26\xdc\x8f\x5c\x0f\x48\xaf\x39\x1e\xc1\x77\x6d\xc7\x51\xd7\x67\xae\xe4\x63\x1d\x6b\xdb\x2f\xd9\xb2\x38\xf3\xac\x9d\x67\xc9\xa3\x80\x67\x3e\x49\x27\x42\xca\x2d\x87\x46\x1d\x43\x85\xfd\x1d\x6b\xaf\x8a\x26\x15\xfc\x35\xd8\xe3\xaf\xbd\x2c\xda\x07\x4e\x0c\x80\x31\xe2\x2a\xbb\xf7\xf1\x92\xeb\x72\xfb\xfa\x3a\x08\x7d\xfd\x6a\x44\x54\x76\xef\xce\xc3\x5d\x6f\x7f\x44\xe9\xe1\x0b\x20\xfa\x98\xff\xb9\x77\xe0\x51\xc7\xbf\x02\x74\x12\xe7\x8b\xf0\x39\xf5\xe7\xff\x8c\xe7\xc5\xed\xe9\xc1\x9f\x01\x00\x00\xff\xff\xec\xa9\xb1\x9b\x2f\x0e\x00\x00")

func _1000000004_commentsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000004_commentsUpSql,
		"1000000004_comments.up.sql",
	)
}

func _1000000004_commentsUpSql() (*asset, error) {
	bytes, err := _1000000004_commentsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000004_comments.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb8, 0x3c, 0xb7, 0xeb, 0x6c, 0x23, 0xa2, 0xc3, 0x3, 0xc3, 0xf, 0x85, 0x5b, 0x7d, 0xa4, 0xf, 0x87, 0xb4, 0xae, 0xd4, 0x1, 0xe7, 0x2, 0x5b, 0xb0, 0x76, 0xe, 0xfe, 0x33, 0x65, 0xa1, 0x4a}}
	return a, nil
}

var __1000000005_diagnostic_countsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xcc\x51\x0a\x02\x21\x10\x00\xd0\xff\x39\xc5\xdc\xc3\xaf\xdd\xcd\x42\xd0\x35\x76\x0d\xfa\x13\x51\xab\x81\x1c\xa1\x71\x3b\x7f\x77\xe8\x02\x6f\xd6\x17\xb3\x2a\x80\xc9\x06\xbd\x61\x98\x66\xab\xf1\x2d\xf4\x88\x25\x8d\x14\x4b\xcf\x47\xab\x3c\x04\x4f\x9b\xbf\xe2\xe2\xed\xcd\xad\x68\xce\xa8\xef\x66\x0f\x3b\x4a\x7e\xd5\x96\xe2\xb7\x7e\x84\x3a\xab\x3f\x15\x3e\x5a\x2c\x94\x9e\xdc\x65\x50\x16\x05\xb0\x78\xe7\x4c\x50\xf0\x0b\x00\x00\xff\xff\xb4\xf5\x7c\x5f\x9e\x00\x00\x00")

func _1000000005_diagnostic_countsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000005_diagnostic_countsDownSql,
		"1000000005_diagnostic_counts.down.sql",
	)
}

func _1000000005_diagnostic_countsDownSql() (*asset, error) {
	bytes, err := _1000000005_diagnostic_countsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000005_diagnostic_counts.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x24, 0xdf, 0x64, 0x43, 0x3, 0x86, 0x84, 0x95, 0xcd, 0x30, 0x3e, 0x35, 0x9c, 0x94, 0x48, 0x25, 0x22, 0x84, 0xd3, 0x6e, 0x3b, 0x68, 0x81, 0x4c, 0xe2, 0xf0, 0xea, 0x44, 0x23, 0xa, 0xe9, 0x4b}}
	return a, nil
}

var __1000000005_diagnostic_countsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\xcd\x6e\x82\x40\x14\x46\xf7\x3c\xc5\xb7\x73\xa5\x69\xd7\xac\x50\x68\x63\x82\xd0\x58\x5c\x93\x91\xb9\xc8\x24\xc3\x1d\x32\x3f\xb6\x7d\xfb\x06\x52\x1a\x35\x26\xb5\x5b\x6e\x38\xdf\x39\xb0\xce\x5e\xb7\x45\x1c\x45\xcb\x25\x5a\xe1\x3c\x59\x78\x03\x17\x86\x41\x7f\x41\x52\x2b\x82\xf6\xf0\x9d\x60\xf4\x82\x83\xd0\x08\x83\x14\x9e\xa2\x24\xaf\xb2\x3d\xaa\x64\x9d\x67\xd0\x4e\xb5\xb5\x14\x5e\xd4\xd2\x34\xa1\x27\xf6\x0e\x49\x9a\x62\x53\xe6\x87\x5d\x01\xd7\x74\xd4\x8b\xfa\x4c\xd6\x29\xc3\x50\xec\x91\x66\x2f\xc9\x21\xaf\xf0\x8c\xa2\xac\x50\x1c\xf2\x3c\xfe\x0f\x91\x43\x5f\x4b\x25\x4e\x6c\x9c\x57\x8d\xbb\x42\x3e\x5d\x20\xc7\x28\x69\xcd\xf0\x1b\x22\xda\x31\x50\x68\x0d\xfa\x54\xce\x2b\x3e\xa1\x31\x3a\xf4\xec\xd0\x89\x33\xe1\x48\xc4\x70\xe4\xff\x96\x99\xee\xf7\x03\xd3\x7d\xf9\x36\xeb\x3c\x90\x75\x49\xba\x0d\xbb\x46\x45\x9b\x72\xb7\xcb\x8a\x0a\x65\x31\xbf\x70\x87\xb8\xba\xd1\xd9\xbe\x63\x51\x75\xf4\x63\x89\xf9\xb1\x69\xe1\x3b\xe5\x60\xcd\x07\x96\x08\x8e\xe4\xf8\xe3\x25\x79\xb2\xbd\x62\xc2\x60\xc9\x11\x37\x04\xc1\x12\xc4\x8d\x91\xe3\xd7\x32\x2d\xc6\xb5\xd5\x22\x7e\x50\xe6\xb6\x68\xb6\xe1\xd0\x1f\xc9\x4e\xbc\x8b\xab\xf3\xc6\x92\x84\x62\xf8\x8e\xa6\x25\xb4\x8a\xb4\x1c\xf7\xa6\xc1\x6d\x15\x47\xdf\x01\x00\x00\xff\xff\x3a\x3d\x23\xfe\xb2\x02\x00\x00")

func _1000000005_diagnostic_countsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000005_diagnostic_countsUpSql,
		"1000000005_diagnostic_counts.up.sql",
	)
}

func _1000000005_diagnostic_countsUpSql() (*asset, error) {
	bytes, err := _1000000005_diagnostic_countsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000005_diagnostic_counts.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc0, 0x46, 0x73, 0x88, 0x3b, 0xb3, 0x47, 0x8, 0x81, 0x3f, 0x16, 0xcf, 0xf0, 0x61, 0xc9, 0x62, 0xca, 0x49, 0x26, 0xc4, 0x21, 0xc4, 0x4a, 0x99, 0x21, 0xdf, 0xe0, 0xe0, 0xfc, 0xa3, 0x8d, 0x2f}}
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
	"1000000000_init.down.sql":                  _1000000000_initDownSql,
	"1000000000_init.up.sql":                    _1000000000_initUpSql,
	"1000000001_init.down.sql":                  _1000000001_initDownSql,
	"1000000001_init.up.sql":                    _1000000001_initUpSql,
	"1000000002_add_indexes.down.sql":           _1000000002_add_indexesDownSql,
	"1000000002_add_indexes.up.sql":             _1000000002_add_indexesUpSql,
	"1000000003_pg_stat_statement_ext.down.sql": _1000000003_pg_stat_statement_extDownSql,
	"1000000003_pg_stat_statement_ext.up.sql":   _1000000003_pg_stat_statement_extUpSql,
	"1000000004_comments.down.sql":              _1000000004_commentsDownSql,
	"1000000004_comments.up.sql":                _1000000004_commentsUpSql,
	"1000000005_diagnostic_counts.down.sql":     _1000000005_diagnostic_countsDownSql,
	"1000000005_diagnostic_counts.up.sql":       _1000000005_diagnostic_countsUpSql,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
	"1000000000_init.down.sql":                  {_1000000000_initDownSql, map[string]*bintree{}},
	"1000000000_init.up.sql":                    {_1000000000_initUpSql, map[string]*bintree{}},
	"1000000001_init.down.sql":                  {_1000000001_initDownSql, map[string]*bintree{}},
	"1000000001_init.up.sql":                    {_1000000001_initUpSql, map[string]*bintree{}},
	"1000000002_add_indexes.down.sql":           {_1000000002_add_indexesDownSql, map[string]*bintree{}},
	"1000000002_add_indexes.up.sql":             {_1000000002_add_indexesUpSql, map[string]*bintree{}},
	"1000000003_pg_stat_statement_ext.down.sql": {_1000000003_pg_stat_statement_extDownSql, map[string]*bintree{}},
	"1000000003_pg_stat_statement_ext.up.sql":   {_1000000003_pg_stat_statement_extUpSql, map[string]*bintree{}},
	"1000000004_comments.down.sql":              {_1000000004_commentsDownSql, map[string]*bintree{}},
	"1000000004_comments.up.sql":                {_1000000004_commentsUpSql, map[string]*bintree{}},
	"1000000005_diagnostic_counts.down.sql":     {_1000000005_diagnostic_countsDownSql, map[string]*bintree{}},
	"1000000005_diagnostic_counts.up.sql":       {_1000000005_diagnostic_countsUpSql, map[string]*bintree{}},
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