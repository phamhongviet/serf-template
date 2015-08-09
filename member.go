package main

import "strings"
import "errors"

type Member struct {
	Name   string
	Addr   string
	Role   string
	Status string
	Tags   map[string]string
}

type Members struct {
	Index map[string]int
	Data  []Member
}

// check if member already exists
func (mems *Members) CheckMember(name string) bool {
	index := mems.Index[name]
	if index >= 0 {
		if mems.Data[index].Name == name {
			return true
		}
	}
	return false
}

func (mems *Members) AddMember(name string, addr string, role string, status string, tags string) error {
	// process tags
	tag_list := strings.Split(tags, ",")
	len_tag_list := len(tag_list)
	// check if member already exists
	if mems.CheckMember(name) {
		// update new member info
		index := mems.Index[name]
		mems.Data[index].Name = name
		mems.Data[index].Addr = addr
		mems.Data[index].Status = status
		mems.Data[index].Role = role
		for i := 0; i < len_tag_list; i++ {
			tag := strings.Split(tag_list[i], "=")
			mems.Data[index].Tags[tag[0]] = tag[1]
		}
		return nil
	} else {
		// create new member
		new_member := Member{
			Name:   name,
			Addr:   addr,
			Role:   role,
			Status: status,
			Tags:   make(map[string]string),
		}
		for i := 0; i < len_tag_list; i++ {
			tag := strings.Split(tag_list[i], "=")
			new_member.Tags[tag[0]] = tag[1]
		}
		// register new member
		new_member_index := len(mems.Data) + 1
		mems.Index[new_member.Name] = new_member_index
		mems.Data = append(mems.Data, new_member)
		return nil
	}
}

func (mems *Members) RemoveMember(name string) error {
	index := mems.Index[name]
	if mems.CheckMember(name) {
		mems.Data = append(mems.Data[:index], mems.Data[index+1:]...)
		delete(mems.Index, name)
		return nil
	} else {
		return errors.New("No such member: " + name)
	}
}
