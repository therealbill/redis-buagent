package drivers

import (
	"fmt"
	"log"
	"os"
	"time"
)

// LocalFileDriver provides a means to store the dump data locally
// Containername is the directory to store the dump in
type LocalFileDriver struct {
	Name          string
	Layout        string
	Containername string
	Logger        *log.Logger
}

func (d *LocalFileDriver) Connect() bool {
	return true
}

func (d *LocalFileDriver) Authenticate() bool {
	dstat, err := os.Stat(d.Containername)

	if err == nil {
		if dstat.IsDir() {
			// check for perms
		} else {
			d.Logger.Printf("Destination '%s' exists but is not a directory. Bailing to avoid overwriting important data.", d.Containername)
			return false
		}
	} else {
		d.Logger.Printf("Destination doesn't exist, creating it.")
		err = os.MkdirAll(d.Containername, 0700)
		if err != nil {
			d.Logger.Printf("Unable to create '%s', error:%s", d.Containername, err)
			return false
		}
		d.Logger.Println("Destination created")
	}

	return true
}

func (d *LocalFileDriver) Upload(data []byte) bool {
	// create an object in the container
	now := time.Now().Local()
	filename := now.Format(d.Layout)
	destination := fmt.Sprintf("%s/%s", d.Containername, filename)
	d.Logger.Printf("Writing to %s", destination)
	writer, err := os.Create(destination)
	if err != nil {
		d.Logger.Fatal(err)
	}
	writer.Write(data)
	writer.Close()

	return false
}
