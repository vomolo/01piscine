package main

import "os"

func main() {
	for _, vs := range os.Args {
		if vs == "01" || vs == "galaxy" || vs == "galaxy 01" {
			os.Stdout.Write([]byte("Alert!!!\n"))
			os.Exit(0)
		}
	}
}
