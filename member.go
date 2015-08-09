package main

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Member struct {
	Name     string
	Addr     string
	Port     int
	Tags     map[string]string
	Status   string
	Protocol map[string]int
}

type SerfOutput struct {
	Members []Member
}

func ParseMembers(serf_output []byte) ([]Member, error) {
	var result SerfOutput
	err := json.Unmarshal(serf_output, &result)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(result.Members); i++ {
		result.Members[i].Addr = strings.Split(result.Members[i].Addr, ":"+strconv.Itoa(result.Members[i].Port))[0]
	}
	return result.Members, nil
}
