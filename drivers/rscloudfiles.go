package drivers

import (
	"log"
	"time"

	"github.com/ncw/swift"
)

type CloudFilesDriver struct {
	Name          string
	Username      string
	Apikey        string
	Authurl       string
	Layout        string
	Origin        string
	Containername string
	Connection    swift.Connection
}

func (d *CloudFilesDriver) Connect() bool {
	log.Println("Connecting to cloudfiles")
	d.Connection = swift.Connection{
		UserName: d.Username,
		ApiKey:   d.Apikey,
		AuthUrl:  d.Authurl}
	return true
}

func (d *CloudFilesDriver) Authenticate() bool {
	if d.Connection.Authenticated() {
		log.Println("Connection is authenticated")
		return true
	}
	log.Print("Authenticating")
	d.Connection.Authenticate()
	if d.Connection.Authenticated() {
		log.Println("Authentication Successful")
		return true
	} else {
		log.Fatal("Authentication failed")
	}
	return false
}

func (d *CloudFilesDriver) Upload(data []byte) bool {
	// create an object in the container
	now := time.Now().Local()
	var remotename string
	remotename = now.Format(d.Layout)

	writer, err := d.Connection.ObjectCreate(d.Containername, remotename, false, "", "RedisDump", nil)
	if err != nil {
		log.Fatal(err)
	}
	writer.Write(data)
	writer.Close()

	return false
}
