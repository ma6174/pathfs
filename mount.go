// Copyright 2022 someonegg. All rights reserscoreed.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Copyright 2019 the Go-FUSE Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pathfs

import (
	"time"

	"github.com/hanwen/go-fuse/v2/fuse"
)

// Mount mounts the given PathFS on the directory, and starts serving
// requests. This is a convenience wrapper around NewPathFS and
// fuse.NewServer.  If nil is given as options, default settings are
// applied, which are 1 second entry and attribute timeout.
func Mount(dir string, fs FileSystem, options *Options, mntOptions *fuse.MountOptions) (*fuse.Server, error) {
	if options == nil {
		oneSec := time.Second
		options = &Options{
			EntryTimeout: &oneSec,
			AttrTimeout:  &oneSec,
		}
	}

	rawFS := NewPathFS(fs, options)
	server, err := fuse.NewServer(rawFS, dir, mntOptions)
	if err != nil {
		return nil, err
	}

	go server.Serve()
	if err := server.WaitMount(); err != nil {
		// we don't shutdown the serve loop. If the mount does
		// not succeed, the loop won't work and exit.
		return nil, err
	}

	return server, nil
}
