package torpedo_registry

import (
	"fmt"
	"sync"
)

var Config *ConfigStruct
var once sync.Once

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

type UserProfile struct {
	Nick     string
	RealName string
	Timezone string
	Phone    string
	Email    string
	IsBot    bool
	ID       string
	// Required for IRC (?)
	Server string
}

type MessageHistoryItem struct {
	Timestamp int64
	Channel   string
	Sender    string
	Nick      string
	Message   string
}

type BotAPI struct {
	API interface{}
	Bot struct {
		Build struct {
			Build      string
			BuildDate  string
			Version    string
			ProjectURL string
		}
		GetCachedItem      func(string) string
		SetCachedItems     func(string, map[int]string) string
		GetCommandHandlers func() map[string]func(*BotAPI, interface{}, string)
		GetHelp            func() map[string]string
		Stats              struct {
			StartTimestamp         int64
			ProcessedMessages      int64
			ProcessedMessagesTotal int64
			ConnectedAccounts      int32
			TotalAccounts          int32
		}
		PostMessage func(interface{}, string, *BotAPI, ...interface{})
	}
	CommandPrefix string
	UserProfile   *UserProfile
}

type ConfigStruct struct {
	coroutines   map[string]func(cfg *ConfigStruct)
	data         map[string]string
	handlers     map[string]func(*BotAPI, interface{}, string)
	textHandlers map[string]func(*BotAPI, interface{}, string)
	help         map[string]string
	preparsers   map[string]func(cfg *ConfigStruct)
	postparsers  map[string]func(cfg *ConfigStruct)
}

func (self *ConfigStruct) RegisterHandler(name string, f func(*BotAPI, interface{}, string)) {
	self.handlers[name] = f
	return
}

func (self *ConfigStruct) RegisterTextMessageHandler(name string, f func(*BotAPI, interface{}, string)) {
	self.textHandlers[name] = f
	return
}

func (self *ConfigStruct) RegisterHelp(name, help_str string) {
	self.help[name] = help_str
	return
}

func (self *ConfigStruct) GetHandlers() map[string]func(*BotAPI, interface{}, string) {
	return self.handlers
}

func (self *ConfigStruct) GetTextMessageHandlers() map[string]func(*BotAPI, interface{}, string) {
	return self.textHandlers
}

func (self *ConfigStruct) GetHelp() map[string]string {
	return self.help
}

func (self *ConfigStruct) SetConfig(key, val string) {
	self.data[key] = val
	return
}

func (self *ConfigStruct) GetConfig() map[string]string {
	return self.data
}

func (self *ConfigStruct) GetPreParsers() map[string]func(cfg *ConfigStruct) {
	return self.preparsers
}

func (self *ConfigStruct) GetPostParsers() map[string]func(cfg *ConfigStruct) {
	return self.postparsers
}

func (self *ConfigStruct) RegisterPreParser(name string, f func(cfg *ConfigStruct)) {
	self.preparsers[name] = f
	return
}

func (self *ConfigStruct) RegisterPostParser(name string, f func(cfg *ConfigStruct)) {
	self.postparsers[name] = f
	return
}

func (self *ConfigStruct) RegisterCoroutine(name string, f func(cfg *ConfigStruct)) {
	self.coroutines[name] = f
	return
}

func (self *ConfigStruct) GetCoroutines() map[string]func(cfg *ConfigStruct) {
	return self.coroutines
}

func init() {
	fmt.Println("Registry init called...")
	once.Do(func() {
		fmt.Println("Registry once.Do called...")
		Config = &ConfigStruct{}
		Config.coroutines = make(map[string]func(cfg *ConfigStruct))
		Config.data = make(map[string]string)
		Config.handlers = make(map[string]func(*BotAPI, interface{}, string))
		Config.textHandlers = make(map[string]func(*BotAPI, interface{}, string))
		Config.help = make(map[string]string)
		Config.preparsers = make(map[string]func(cfg *ConfigStruct))
		Config.postparsers = make(map[string]func(cfg *ConfigStruct))
	})
}
