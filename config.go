package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	yaml "gopkg.in/yaml.v2"
)

type RoleList []string

func (rls RoleList) HasRole(role string) bool {
	for _, r := range rls {
		if r == role {
			return true
		}
	}
	return false
}
func (roleList RoleList) HasAnyRole(givenRoles []string) bool {
	for r1 := range roleList {
		for r2 := range givenRoles {
			if r1 == r2 {
				return true
			}
		}
	}
	return false
}

type Configuration struct {
	// server     ServerConf `yaml:"server"`
	Ver    string `yaml:"ver"`
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yml:"server"`

	Users []struct {
		Name     string   `yaml:"name"`
		Password string   `yaml:"password"`
		Roles    RoleList `yaml:"roles"`
	} `yml:"users"`

	Targets map[string]struct {
		Enabled bool     `yaml:"enabled"`
		Roles   RoleList `yaml:"roles"`
	} `yml:"users"`
}

type UserMap map[string]struct {
	password string
	roles    []string
}

func (c *Configuration) GetServerAddress() string {
	return c.Server.Host + ":" + strconv.Itoa(c.Server.Port)
}

func (c *Configuration) GetUserMap() *UserMap {
	userMap := make(UserMap)
	for _, u := range c.Users {
		userMap[u.Name] = struct {
			password string
			roles    []string
		}{u.Password, u.Roles}
	}
	return &userMap
}

func (c *Configuration) Load(configFile string) *Configuration {
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		panic(err)
	}
	return c
}

func (c *Configuration) PrintConf() {
	y, _ := yaml.Marshal(c)
	fmt.Println(string(y))
}
