package main

func main() {
	initLogger()
	mainLogger.Info("Starting server-ctl!")
	startWeb()
}