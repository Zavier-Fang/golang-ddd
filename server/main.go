package main

import "server/infrastructure/container"

func main() {
	var c container.Container
	c.Register(container.SINGLE).Run()
}
