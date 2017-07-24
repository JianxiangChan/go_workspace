package mlib

import "testing"

func TestOps(t *testing.T) {
	//test NewMusicManager
	mm := NewMusicManger()
	if mm == nil {
		t.Error("NewMusicManager failed")
	}
	if mm.Len() != 0 {
		t.Error("NewMusicManager failed")
	}

	//test MusicManager.Add()
	m0 := &MusicEntry{
		"1", "My Heart will go on", "Celion Dion",
		"http://qbox.me/24501234", "MP3"}
	mm.Add(m0)
	if mm.Len() != 1 {
		t.Error("MusicManger.Add() failed.")
	}

	//test MusicManager.Find
	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManger.Find() failed")
	}
	if m.Id != m0.Id || m.Artist != m0.Artist ||
		m.Name != m0.Name || m.Type != m0.Type ||
		m.Source != m0.Source {
		t.Error("MusicManager.Find() failed")
	}

	//test MusicManager.Get
	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager.Get() failed.", err)
	}

	//test MusicManager.Remove
	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove failed", err)
	}
}
