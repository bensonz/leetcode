# !/bin/bash
#

Name=$1
mkdir $Name
mkdir $Name/go-solution
touch $Name/go-solution/main.go
touch $Name/README.md
mkdir $Name/js-solution
touch $Name/js-solution/solve.js
cat > $Name/go-solution/main.go <<EOL
package main

import (
	"fmt"
)

func () {

}

func main() {
	fmt.Println()
}

EOL
