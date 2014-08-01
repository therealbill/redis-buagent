package drivers

import (
	"log"
	"time"

	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

type AmazonS3Driver struct {
	Name          string
	Username      string
	Apikey        string
	Layout        string
	Origin        string
	Containername string
	Connection    *s3.S3
}

func (d *AmazonS3Driver) Connect() bool {
	log.Println("AmazonS3 connection call ")
	return true
}

func (d *AmazonS3Driver) Authenticate() bool {
	log.Print("Authenticating")
	return true
}

func (d *AmazonS3Driver) Upload(rdb []byte) bool {
	log.Println("Upload called on:", d.Name)
	// create an object in the container
	now := time.Now().Local()
	var remotename string
	remotename = now.Format(d.Layout)
	log.Println("Saving to", remotename)
	auth := aws.Auth{AccessKey: d.Username, SecretKey: d.Apikey}
	s := s3.New(auth, aws.USEast)
	bucket := s.Bucket(d.Containername)
	log.Print("Putting object")

	err := bucket.Put(remotename, rdb, "text/plain", s3.BucketOwnerFull)
	if err != nil {
		print("Error:", err)
		log.Fatal(err)
	}
	return false
}
