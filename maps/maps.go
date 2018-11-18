//
// statm.Statm describes data in /proc/<pid>/statm.
//
// Use statm.New() to create a new stat.Statm object
// from data in a path.
//
package statm

//
// Copyright Jen Andre (jandre@gmail.com)
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
	"github.com/jandre/procfs/util"
	"io/ioutil"
	"strings"
)

//
// Abstraction for /proc/<pid>/maps
//
type Maps struct {
	AddressStart uint32 // This is the starting ...
	AddressEnd   uint32 // and ending address of the region in the process's address space
	Perms        string // Describes how pages in the region can be accessed.
	Offset       uint32 // If the region was mapped from a file (using mmap), this is the offset in the file where the mapping begins. If the memory was not mapped from a file, it's just 0.
	Device       string // If the region was mapped from a file, this is the major and minor device number (in hex) where the file lives.
	Inode        int    // If the region was mapped from a file, this is the file number.
	Pathname     string // If the region was mapped from a file, this is the name of the file.
}

type ProcMaps struct {
	AddressRange uint32 // This is the starting and ending address of the region in the process's address space
	Perms        string // Describes how pages in the region can be accessed.
	Offset       uint32 // If the region was mapped from a file (using mmap), this is the offset in the file where the mapping begins. If the memory was not mapped from a file, it's just 0.
	Device       string // If the region was mapped from a file, this is the major and minor device number (in hex) where the file lives.
	Inode        int    // If the region was mapped from a file, this is the file number.
	Pathname     string // If the region was mapped from a file, this is the name of the file.
}

func New(path string) (*ProcMaps, *Maps, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(buf), " ")
	procMaps := &ProcMaps{}
	maps := &Maps{}

	err = util.ParseStringsIntoStruct(procMaps, lines)
	if err != nil {
		// set maps
	}
	return maps, procMaps, err
}
