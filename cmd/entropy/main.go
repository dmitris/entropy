// Entropy estimates shannon entropy of the standard input byte stream.

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/lazybeaver/entropy"
)

const entropyThreashold = 3.0

func main() {
	var reader io.Reader
	minOpt := flag.Float64("min", entropyThreashold, "threshold enthropy value for outputting filenames")
	allOpt := flag.Bool("all", false, "print entropy values for all given files")
	flag.Parse()

	if len(flag.Args()) == 0 {
		estimator := entropy.NewShannonEstimator()
		if _, err := io.Copy(estimator, reader); err != nil {
			fmt.Printf("IO Error: %s", err)
			return
		}
		fmt.Println(estimator.Value())
		return
	}
	for _, fname := range flag.Args() {
		f, err := os.Open(fname)
		if err != nil {
			log.Printf("unable to open file %s: %v", fname, err)
			continue
		}
		estimator := entropy.NewShannonEstimator()
		if _, err := io.Copy(estimator, f); err != nil {
			fmt.Printf("IO Error: %s", err)
			continue
		}
		ent := estimator.Value()
		if *allOpt || ent >= *minOpt {
			fmt.Println(ent, fname)
		}
	}

}
