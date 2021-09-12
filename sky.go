package main

import (
	"autosky/pkg"
)

func main() {
	pkg.Readsky()
	pkg.CreateXLSX()

	/*pkg.LoadOrder()
	pkg.LoadNakl()
	pkg.Parse()
	pkg.CreateXLSX()*/

}
