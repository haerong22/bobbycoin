package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/haerong22/bobbycoin/explorer"
	"github.com/haerong22/bobbycoin/rest"
)

func usage() {
	fmt.Printf("Welcome to Bobby Coin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:    Set the PORT of the server\n")
	fmt.Printf("-mode:    Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Sets port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		// start rest api
		rest.Start(*port)
	case "html":
		// start html explorer
		explorer.Start(*port)
	default:
		usage()
	}

	// rest := flag.NewFlagSet("rest", flag.ExitOnError)

	// portFlag := rest.Int("port", 4000, "Sets the port of the server")

	// switch os.Args[1] {
	// case "explorer":
	// 	fmt.Println("Start Explorer")
	// case "rest":
	// 	rest.Parse(os.Args[2:])
	// default:
	// 	usage()
	// }

	// if rest.Parsed() {
	// 	fmt.Println(*portFlag)
	// 	fmt.Println("Start REST API")
	// }
}
