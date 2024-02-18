package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	// "github.com/jschaf/bibtex"
	"github.com/nickng/bibtex"
	"os"
)

var cli struct {
	FetchCitekey struct {
		Bibtex string `help:"bibtex file path" type:"existingfile" short:"b" required:""`
		Key string `help:"entry key you find" optional:"" short:"k"`
		Value string `help:"file path for the entry you want" optional:"" type:"existingfile" short:"v"`
	} `cmd:"0" help:"Fetch citekey from bibtex file"`
}

func main() {
	ctx := kong.Parse(&cli,
		kong.Name("bibtex_finder"),
		kong.Description("Operations on bibtex files"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))

	file, ferr := os.Open(cli.FetchCitekey.Bibtex)
	if ferr != nil {
		os.Exit(1)
	}
	parsed, rerr := bibtex.Parse(file)
	if rerr != nil {
		os.Exit(1)
	}

	switch ctx.Command() {
	case "fetch-citekey":
		for _, item := range parsed.Entries {
			citekey := item.CiteName
			value := item.Fields[cli.FetchCitekey.Key]
			if value != nil {
				if value.String() == cli.FetchCitekey.Value {
					fmt.Println(citekey)
					return
				}
			}
		}
		return
	}
}
