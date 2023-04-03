package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicação em linha de comando"
	app.Usage = "Busca ips e nomes de servidores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "busca ips de endereços na internet",
			Flags:  flags,
			Action: buscaIP,
		},
		{
			Name:   "server",
			Usage:  "busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscaIP(c *cli.Context) {

	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)

	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	server, err := net.LookupNS(host) // name server

	if err != nil {
		log.Fatal(err)
	}

	for _, servers := range server {
		fmt.Println(servers)
	}
}
