package main

import backend "backend/internal/1_framework_driver"

func main() {
	a := backend.NewApp()
	a.Start()
}
