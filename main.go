package main

import (
	"fmt"
	"github.com/billykore/learn-github-actions/stringutil"
)

func main() {
	i := stringutil.ParseInt("100")
	fmt.Println(i)

	f := stringutil.ParseFloat("5.555")
	fmt.Println(f)
}
