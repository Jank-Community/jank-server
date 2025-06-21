package service

func InitPlugin() {
	go StartPluginServer()

	go StartPluginHost()
}
