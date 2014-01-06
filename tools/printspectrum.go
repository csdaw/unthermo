package main

import (
	"bitbucket.org/proteinspector/unthermo"
	"os"
	"strconv"
	"fmt"
)

func main() {
	//Parse arguments
	scannumber,_ := strconv.Atoi(os.Args[1])
	filename := os.Args[2]
	//Execute Scan with argument prettyprint
	unthermo.Scan(filename, uint64(scannumber), prettyprint)
}

//Print time, mz and Intensity of every peak in spectrum
var prettyprint = func(spectrum unthermo.Spectrum) {
	for _, peak := range spectrum {
		fmt.Println(peak.Time, peak.Mz, peak.I)
	}
}
