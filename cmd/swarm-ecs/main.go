package main

import (
	"io"
	"os"
	"text/template"

	"github.com/ejholmes/ecs"
	"github.com/ejholmes/ecs/Godeps/_workspace/src/github.com/codegangsta/cli"
)

var flagCluster = cli.StringFlag{
	Name:  "cluster",
	Value: "",
	Usage: "ECS Cluster to query",
}

var flagFormat = cli.StringFlag{
	Name: "format, f",
	Value: `{{ range . }}<{{ .PrivateIpAddress }}:2375>
{{ end }}`,
	Usage: "Format to use to print the results",
}

var commands = []cli.Command{
	{
		Name: "list",
		Flags: []cli.Flag{
			flagCluster,
			flagFormat,
		},
		Action: runList,
	},
}

func main() {
	app := cli.NewApp()
	app.Commands = commands
	app.Run(os.Args)
}

func runList(c *cli.Context) {
	printInstances(c, os.Stdout)
}

func printInstances(c *cli.Context, w io.Writer) error {
	instances, err := ecs.Instances(c.String("cluster"))
	if err != nil {
		return err
	}

	tmpl, err := template.New("format").Parse(c.String("format"))
	if err != nil {
		return err
	}

	return tmpl.Execute(w, instances)
}
