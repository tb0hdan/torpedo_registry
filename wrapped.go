package torpedo_registry

func (self *ConfigStruct) RegisterHelpAndHandler(name, help_str string, handler func(*BotAPI, interface{}, string)) {
	self.RegisterHandler(name, handler)
	self.RegisterHelp(name, help_str)
}

func (self *ConfigStruct) RegisterParser(name string, preparser func(cfg *ConfigStruct), postparser func(cfg *ConfigStruct)) {
	self.RegisterPreParser(name, preparser)
	self.RegisterPostParser(name, postparser)
}