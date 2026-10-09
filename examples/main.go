package main

import nethersignal "nn"

func main() {
	ns := nethersignal.New()

	ns.Server.Registration("127.0.0.1:19133")

	ns.Listen(":19132")
}
