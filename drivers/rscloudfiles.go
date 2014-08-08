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
	Logger        *log.Logger
}

func (d *CloudFilesDriver) Connect() bool {
	d.Logger.Println("Connecting to cloudfiles")
	d.Connection = swift.Connection{
		UserName: d.Username,
		ApiKey:   d.Apikey,
		AuthUrl:  d.Authurl}
	return true
}

func (d *CloudFilesDriver) Authenticate() bool {
	if d.Connection.Authenticated() {
		d.Logger.Println("Connection is authenticated")
		return true
	}
	d.Logger.Print("Authenticating")
	d.Connection.Authenticate()
	if d.Connection.Authenticated() {
		d.Logger.Println("Authentication Successful")
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
	d.Logger.Println("Ensuring container is present")
	err := d.Connection.ContainerCreate(d.Containername, nil)
	if err != nil {
		d.Logger.Println("Create container error:", err)
	}
	d.Logger.Printf("Saving to '%s:%s'", d.Containername, remotename)
	writer, err := d.Connection.ObjectCreate(d.Containername, remotename, false, "", "RedisDump", nil)
	if err != nil {
		log.Fatal(err)
	}
	writer.Write(data)
	writer.Close()

	return false
}
