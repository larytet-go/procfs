//
// maps.Maps describes data in /proc/<pid>/maps.
//
// Use maps.New() to create a new maps.Maps object
// from data in a path.
//
package maps

//
// Copyright Arkady Maisnikov (jandre@gmail.com)
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//

import (
	"bufio"
	"github.com/jandre/procfs/util"
	"os"
	"strings"
)

//
// Abstraction for /proc/<pid>/maps
//
type Maps struct {
	AddressStart uint64 // This is the starting ...
	AddressEnd   uint64 // and ending address of the region in the process's address space
	Perms        string // Describes how pages in the region can be accessed.
	Offset       uint64 // If the region was mapped from a file (using mmap), this is the offset in the file where the mapping begins. If the memory was not mapped from a file, it's just 0.
	Device       string // If the region was mapped from a file, this is the major and minor device number (in hex) where the file lives.
	Inode        int    // If the region was mapped from a file, this is the file number.
	Pathname     string // If the region was mapped from a file, this is the name of the file.
}

type ProcMap struct {
	AddressRange string // This is the starting and ending address of the region in the process's address space
	Perms        string // Describes how pages in the region can be accessed.
	Offset       uint64 // If the region was mapped from a file (using mmap), this is the offset in the file where the mapping begins. If the memory was not mapped from a file, it's just 0.
	Device       string // If the region was mapped from a file, this is the major and minor device number (in hex) where the file lives.
	Inode        int    // If the region was mapped from a file, this is the file number.
	Pathname     string // If the region was mapped from a file, this is the name of the file.
}

func New(path string) ([]*Maps, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	var maps []*Maps

	for scanner.Scan() {
		procMap := &ProcMap{}
		line := scanner.Text()
		columns := strings.Split(string(line), " ")
		err = util.ParseStringsIntoStruct(procMap, columns)
		if err != nil {
			var newMap *Maps = &Maps{Perms: procMap.Perms, Offset: procMap.Offset, Device: procMap.Device, Inode: procMap.Inode, Pathname: procMap.Inode}
			maps = append(maps, newMap)
		}
	}

	return maps, err
}
