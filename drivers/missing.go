package drivers

import (
	"github.com/ncw/swift"
	//"io/ioutil"
	"log"
	//"time"
)

type MissingDriver struct {
	Name       string
	Username   string
	Apikey     string
	Authurl    string
	Connection swift.Connection
	Logger     *log.Logger
}

func (d *MissingDriver) Connect() bool {
	log.Println("Connect called on:", d.Name)
	log.Println("Username is:", d.Username)
	return false
}
func (d *MissingDriver) Authenticate() bool {
	log.Println("Authenticate called on:", d.Name)
	return false
}
func (d *MissingDriver) Upload(rdb []byte) bool {
	log.Println("Upload called for", d.Name)
	log.Println("Size to upload:", len(rdb))
	return false
}
