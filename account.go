package torpedo_registry

import (
	"sync"
	"fmt"
)

type Account struct {
	API interface{}
	APIKey string
	CommandPrefix string
	Connection struct {
		Connected bool
		Ping float64
		StatusMessage string
		ReconnectCount int64
		DynamicallyDisabled bool
	}
	Description string
	Name string
}

type AccountsStruct struct {
	accounts []*Account
	m     sync.RWMutex
}

func (self *AccountsStruct) AppendAccounts(account *Account) {
	self.m.Lock()
	fmt.Printf("Registering account: %s\n", account)
	self.accounts = append(self.accounts, account)
	self.m.Unlock()
	return
}


func (self *AccountsStruct) GetAccounts() (accounts []*Account) {
	return self.accounts
}


func (self *AccountsStruct) GetAccountByAPIKey(apiKey string) (account *Account) {
	self.m.RLock()
	for _, tmp := range self.accounts {
		if tmp.APIKey == apiKey {
			account = tmp
			break
		}
	}
	self.m.RUnlock()
	return
}