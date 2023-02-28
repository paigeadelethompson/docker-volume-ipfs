package fuse

import (
	"context"
	"log"
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	_ "bazil.org/fuse/fs/fstestutil"
)

type FS struct{}

func (FS) Root() (fs.Node, error) {
	return Dir {

	}, nil
}

func (Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	panic("Dir.ReadDirAll")
}

type File struct{}

func (File) Attr(ctx context.Context, a *fuse.Attr) error {
	panic("File.Attr")
}

func (File) ReadAll(ctx context.Context) ([]byte, error) {
	panic("File.ReadAll")
}

type Dir struct{}

func (Dir) Attr(ctx context.Context, a *fuse.Attr) error {
	panic("Dir.Attr")
}

func (Dir) Lookup(ctx context.Context, name string) (fs.Node, error) {
	panic("Dir.Lookup")
}

type FUSEController struct { 
	mountPoint string
	conn *fuse.Conn
}

func New(mountPoint string) FUSEController {
	f := FUSEController {
		mountPoint: mountPoint,
	}

	conn, err := fuse.Mount(
		f.mountPoint,
		fuse.FSName("dockerfs"),
		fuse.Subtype("ipfs"),
	)
	f.conn = conn

	if err != nil {
		log.Fatal(err)
	}

	defer f.conn.Close()
	
	fs.Serve(f.conn, FS{})

	return f
}