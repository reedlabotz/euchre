package main

import (
	"fmt"
)

const LOAD_SCREEN string = `
 _____ _   _ _____  _   _ ______ _____ 
|  ___| | | /  __ \| | | || ___ \  ___|
| |__ | | | | /  \/| |_| || |_/ / |__  
|  __|| | | | |    |  _  ||    /|  __| 
| |___| |_| | \__/\| | | || |\ \| |___ 
\____/ \___/ \____/\_| |_/\_| \_\____/`

func main() {
	fmt.Printf("%s\n\n", LOAD_SCREEN)
	server := NewServer()
	server.Init()
	server.Run()
}
