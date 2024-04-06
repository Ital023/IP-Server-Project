package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)


func Generate() *cli.App{
	//Initialize cli app
	app := cli.NewApp()

	//Setting info app
	app.Name = "Cli application"
	app.Usage = "Find IPs and names of serves on internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}

	//Slice of app cli commands, as we can see only got one method called by ip, that have one Slice of Flags, and contains one method called by host, set default by google.com, if in params not null, the action named by FindIps is executed
	app.Commands = []cli.Command{
		{
			Name:"ip",
			Usage:"Search for IPs on internet",
			Flags: flags,
			Action: findIps,
		},
	}
	return app
}

func findIps (c *cli.Context){
	host := c.String("host")

	// Returns Slice of Public IPs and errors, if it throw
	ips, erro := net.LookupIP(host)

	if erro != nil{
		log.Fatal(erro)
	}

	//Loop returns IPs availables in Ips slice
	for _,ip := range ips {
		fmt.Println(ip)
	}
}
