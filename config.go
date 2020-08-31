package api_doc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	Title int = iota
	Url
	Header
	Method
	Param
	Hint
	Return
)

var (
	baseKeys = map[string]int{
		"@title":  Title,
		"@url":    Url,
		"@header": Header,
		"@method": Method,
		"@param":  Param,
		"@return": Return,
		"@hint":   Hint,
	}
)

type Config struct {
	Addr        string
	Port        string
	Files       []string
	ApiServ     string
	MatchKeys   []map[string]string
	defaultPath string
	defaultPort string
	defaultLogs string
	defaultAddr string
}

func NewConfig(path string) *Config {
	c := &Config{
		defaultPath: "config.json",
		defaultLogs: "",
		defaultPort: "8888",
		defaultAddr: "0.0.0.0",
	}
	if path != "" {
		c.defaultPath = path
	}
	c.readConfigFile()
	c.collectMatchKey()
	return c
}

func (c *Config) readConfigFile() {
	cfg, err := ioutil.ReadFile(c.defaultPath)
	checkError(err, true)
	if err := json.Unmarshal(cfg, c); err != nil {
		panic(err)
	}
}

func (c *Config) collectMatchKey() {
	for _, key := range c.MatchKeys {
		for k, v := range key {
			switch k {
			case "@title":
				baseKeys[v] = Title
			case "@url":
				baseKeys[v] = Url
			case "@header":
				baseKeys[v] = Header
			case "@method":
				baseKeys[v] = Method
			case "@param":
				baseKeys[v] = Param
			case "@return":
				baseKeys[v] = Return
			case "@hint":
				baseKeys[v] = Hint
			}
		}
	}

	fmt.Println(baseKeys)
}

func (c *Config) files() []string {
	return c.Files
}

func (c *Config) handlerConfigUpdated() map[string]int {
	toMap := make(map[string]int)
	return toMap
}
