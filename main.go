package torpedo_registry

type RichMessage struct {
	BarColor  string
	Text      string
	Title     string
	TitleLink string
	ImageURL  string
}

func (rm *RichMessage) IsEmpty() bool {
	return rm.Text == "" || rm.ImageURL == ""
}

func (rm *RichMessage) ToGenericAttachment() (msg, url string) {
	msg = rm.Text
	url = rm.ImageURL
	return
}

type BotAPI struct {
	API interface{}
	Bot struct {
		Build struct {
			Build     string
			BuildDate string
			Version   string
		}
		GetCachedItem      func(string) string
		SetCachedItems     func(string, map[int]string) string
		GetCommandHandlers func() map[string]func(*BotAPI, interface{}, string)
		GetHelp            func() map[string]string
		Stats              struct {
			StartTimestamp         int64
			ProcessedMessages      int64
			ProcessedMessagesTotal int64
			ConnectedAccounts      int64
			TotalAccounts          int64
		}
		PostMessage func(interface{}, string, *BotAPI, ...interface{})
	}
	CommandPrefix string
}

var (
	config		= make(map[string]string)
	handlers    = make(map[string]func(*BotAPI, interface{}, string))
	help        = make(map[string]string)
	preparsers  = make(map[string]func())
	postparsers = make(map[string]func())
)

func RegisterHandler(name string, f func(*BotAPI, interface{}, string)) {
	handlers[name] = f
	return
}

func RegisterHelp(name, help_str string) {
	help[name] = help_str
	return
}

func GetHandlers() map[string]func(*BotAPI, interface{}, string) {
	return handlers
}

func GetHelp() map[string]string {
	return help
}

func RegisterPreParser(name string, f func()) {
	preparsers[name] = f
	return
}

func RegisterPostParser(name string, f func()) {
	postparsers[name] = f
	return
}

func GetPreParsers() map[string]func() {
	return preparsers
}

func GetPostParsers() map[string]func() {
	return postparsers
}

func GetConfig() map[string]string {
	return config
}

func SetConfigOption(option, value string) {
	config[option] = value
	return
}