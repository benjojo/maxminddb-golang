// +build appengine

package maxminddb

import (
	"io/ioutil"
	"os"
)

// Open takes a string path to a MaxMind DB file and returns a Reader
// structure or an error. The database file is opened using a memory map,
// except on Google App Engine where mmap is not supported; there the database
// is loaded into memory. Use the Close method on the Reader object to return
// the resources to the system.
func Open(file string) (*Reader, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return FromBytes(bytes)
}

// OpenGzip takes a string path to a MaxMind DB file that is gziped and returns
// a Reader structure or an error. The database file is opened using a memory
// map,except on Google App Engine where mmap is not supported; there the
// database is loaded into memory. Use the Close method on the Reader object
// to return the resources to the system.
func OpenGzip(file string) (*Reader, error) {
	filehandle, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer filehandle.Close()

	gziphandle, err := gzip.NewReader(filehandle)
	if err != nil {
		return nil, err
	}
	defer gziphandle.Close()

	bytes, err := ioutil.ReadAll(gziphandle)
	if err != nil {
		return nil, err
	}

	return FromBytes(bytes)
}

// Close unmaps the database file from virtual memory and returns the
// resources to the system. If called on a Reader opened using FromBytes
// or Open on Google App Engine, this method does nothing.
func (r *Reader) Close() {
}
