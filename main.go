package main

import (
	"flag"
	"fmt"
	"github.com/gookit/color"
)

func main() {
	var host, file string

	conf := &Confluence{}

	flag.StringVar(&host, "h", "", "ip")
	flag.StringVar(&file, "f", "", "filepath")
	flag.Parse()

	view := `
  ___  _  _  ____       ___    ___   ___   ___        ___    _    __  ___  ___ 
 / __)( \/ )( ___) ___ (__ \  / _ \ (__ \ (__ \  ___ (__ \  / )  /  )(__ )( _ )
( (__  \  /  )__) (___) / _/ ( (_) ) / _/  / _/ (___) / _/ / _ \  )(  (_ \/ _ \
 \___)  \/  (____)     (____) \___/ (____)(____)     (____)\___/ (__)(___/\___/  by:Z92G`
	color.Cyan.Println(view)
	fmt.Println()

	if file != "" && host == "" {
		conf.batchScan(file)
	} else {
		conf.singleScan(host)
	}

}
