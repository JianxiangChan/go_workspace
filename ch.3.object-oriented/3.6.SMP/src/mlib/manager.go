package mlib

import "errors"

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManger struct {
	musics []MusicEntry
}

//实例化对象 返回实例化地址
func NewMusicManger() *MusicManger {
	return &MusicManger{make([]MusicEntry, 0)}
}

//返回对象长度
func (m *MusicManger) len() int {
	return len(m.musics)
}

//获取对象成员
func (m *MusicManger) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range")
	}
	return &m.musics[index], nil
}

//寻找对象里面某个成员
func (m *MusicManger) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

//对象成员添加信息
func (m *MusicManger) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

//去除对象某个成员变量
func (m *MusicManger) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	removedMusic := &m.musics[index]
	//打散以后传递
	m.musics = append(m.musics[:index], m.musics[index+1:]...)

	return removedMusic
}
