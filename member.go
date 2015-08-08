package main

import "strings"

type Member struct {
	Name string
	Addr string
	Role string
	Tags map[string]string
}

func AddMember(members []Member, name string, addr string, role string, tags string) ([]Member, error) {
	new_member := Member{
		Name: name,
		Addr: addr,
		Role: role,
		Tags: make(map[string]string),
	}
	tag_list := strings.Split(tags, ",")
	for i := 0; i < len(tag_list); i++ {
		tag := strings.Split(tag_list[i], "=")
		new_member.Tags[tag[0]] = tag[1]
	}
	return append(members, new_member), nil
}
