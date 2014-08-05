package main

import (
	"flag"
	"fmt"
	"log"

	"code.google.com/p/gcfg"
	client "github.com/therealbill/libredis/client"
	drivers "github.com/therealbill/redis-buagent/drivers"
)

type Config struct {
	// Define the config file structure
	Main struct {
		Driver            string
		Maxfilesize       int64
		DestinationFormat string
	}

	Redis struct {
		Host      string
		Port      int
		SlaveOnly bool
		AuthToken string
	}

	Rackspacecf struct {
		Username      string
		Apikey        string
		Containername string
	}

	Amazons3 struct {
		Username      string
		Apikey        string
		Containername string
	}
	Localfile struct {
		Directory string
	}
}

func getDriver(config Config) drivers.Driver {
	switch config.Main.Driver {
	case "rackspacecf":
		mydriver := new(drivers.CloudFilesDriver)
		mydriver.Name = config.Main.Driver
		mydriver.Username = config.Rackspacecf.Username
		mydriver.Apikey = config.Rackspacecf.Apikey
		mydriver.Authurl = "https://auth.api.rackspacecloud.com/v1.0"
		mydriver.Layout = config.Main.DestinationFormat
		mydriver.Containername = config.Rackspacecf.Containername
		return mydriver

	case "amazons3":
		mydriver := new(drivers.AmazonS3Driver)
		mydriver.Name = config.Main.Driver
		mydriver.Username = config.Amazons3.Username
		mydriver.Apikey = config.Amazons3.Apikey
		mydriver.Layout = config.Main.DestinationFormat
		mydriver.Containername = config.Amazons3.Containername
		return mydriver

	case "localfile":
		mydriver := new(drivers.LocalFileDriver)
		mydriver.Name = config.Main.Driver
		mydriver.Layout = config.Main.DestinationFormat
		mydriver.Containername = config.Localfile.Directory
		return mydriver
	}

	return new(drivers.MissingDriver)
}

func main() {
	// The main stuff happens here
	var conf Config
	var configfilename string
	flag.StringVar(&configfilename, "conf", "/etc/redis/buagent.cfg", "Config file to use")
	flag.Parse()

	err := gcfg.ReadFileInto(&conf, configfilename)
	if err != nil {
		log.Fatal(err)
	}
	td := getDriver(conf)
	td.Connect()
	td.Authenticate()

	connstring := fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port)
	targetConf := client.DialConfig{Address: connstring, Password: conf.Redis.AuthToken}

	r, err := client.DialWithConfig(&targetConf)
	info, _ := r.Info()
	doBackup := false

	if conf.Redis.SlaveOnly {
		if info.Replication.Role == "slave" {
			doBackup = true
		}
	} else {
		doBackup = true
	}
	if doBackup {
		rdb, err := r.ExecuteCommand("SYNC")
		if err != nil {
			fmt.Println("Error on sync:", err)
		}
		rdb_data, err := rdb.BytesValue()
		if err != nil {
			fmt.Println("Error on sync:", err)
		}
		if int64(len(rdb_data)) >= conf.Main.Maxfilesize {
			log.Fatal("RDB Data is too large, aborting")
		}
		datasize := float64(len(rdb_data)) / 1024.0
		log.Printf("Origin data is %.4f Kb\n", float64(datasize))

		td.Upload(rdb_data)
	} else {
		log.Fatal("No suitable Redis servers found to do backup from")
	}
}
