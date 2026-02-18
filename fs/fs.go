package main

import (

	"fmt"
	"os"
	"context"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"

)

type FS struct{}

func (FS) Root() (fs.Node, error) {
	return Dir{}, nil;
}

type Dir struct{}

func (Dir) Attr(ctx context.Context, a *fuse.Attr) error {
	
	a.Inode = 1; //inode 1 cuz root
	a.Mode = os.ModeDir | 0o755 //octal perms for rwx r-x r-x
	return nil

}


func main() {

	if len(os.Args) != 2 {
		fmt.Println("Invalid usage. Use ./radFs <mntpoint>")
		return
	}

	mount_point := os.Args[1]

	//c is a fuse connection to dev/fuse
	c, err := fuse.Mount(mount_point)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close() //delay execution of Close

	fs.Serve(c, FS{}) //starts listening for FS reqs
	

}