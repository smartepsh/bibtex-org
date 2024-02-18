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
		Key string `help:"entry key you find" required:"" short:"k"`
		Value string `help:"file path for the entry you want" required:"" type:"existingfile" short:"v"`
	} `cmd:"0" help:"Fetch citekey from bibtex file by field key"`
	FetchField struct {
		Bibtex string `help:"bibtex file path" type:"existingfile" short:"b" required:""`
		Citekey string `help:"entry key you find" required:"" short:"c"`
		Key string `help:"entry field key you find" required:"" short:"k"`
	} `cmd:"1" help:"Fetch field from bibtex file by citekey"`
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


	switch ctx.Command() {
	case "fetch-citekey":
		file, ferr := os.Open(cli.FetchCitekey.Bibtex)
		if ferr != nil {
			os.Exit(1)
		}
		parsed, rerr := bibtex.Parse(file)
		if rerr != nil {
			os.Exit(1)
		}
		for _, item := range parsed.Entries {
			citekey := item.CiteName
			value := item.Fields[cli.FetchCitekey.Key]
			if value != nil {
				if value.String() == cli.FetchCitekey.Value {
					fmt.Println(citekey)
					break
				}
			}
		}

	case "fetch-field":
		file, ferr := os.Open(cli.FetchField.Bibtex)
		if ferr != nil {
			os.Exit(1)
		}
		parsed, rerr := bibtex.Parse(file)
		if rerr != nil {
			os.Exit(1)
		}
		for _, item := range parsed.Entries {
			if item.CiteName != cli.FetchField.Citekey {
				continue
			}
			value := item.Fields[cli.FetchField.Key]
			if value != nil {
				fmt.Println(value.String())
				break
			}
		}
		fmt.Println("")
	}
}
