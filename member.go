package main

import (
	"encoding/json"
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
	return result.Members, nil
}
