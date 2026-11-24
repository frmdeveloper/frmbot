package lib

import (
	"encoding/json"
	"os"
	"sync"
)

type DataUser struct {
	Banned         []string `json:"banned,omitempty"`
	LolosPerangkat []string `json:"lolosperangkat,omitempty"`

	mu     sync.Mutex
	path   string
	mentah map[string]json.RawMessage
}

type DataGroup struct {
	Welcome   map[string]any `json:"welcome"`
	Goodbye   map[string]any `json:"goodbye"`
	BanMember map[string]any `json:"banmember"`
	Absen     map[string]any `json:"absen"`
	Chatai    map[string]any `json:"chatai"`

	mu     sync.Mutex
	path   string
	mentah map[string]json.RawMessage
}

func MuatUser(path string) *DataUser {
	d := &DataUser{path: path}
	if b, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(b, d)
		_ = json.Unmarshal(b, &d.mentah)
	}
	return d
}

func MuatGroup(path string) *DataGroup {
	d := &DataGroup{
		Welcome:   map[string]any{},
		Goodbye:   map[string]any{},
		BanMember: map[string]any{},
		Absen:     map[string]any{},
		Chatai:    map[string]any{},
		path:      path,
	}
	if b, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(b, d)
		_ = json.Unmarshal(b, &d.mentah)
		if d.Welcome == nil {
			d.Welcome = map[string]any{}
		}
		if d.Goodbye == nil {
			d.Goodbye = map[string]any{}
		}
		if d.BanMember == nil {
			d.BanMember = map[string]any{}
		}
		if d.Absen == nil {
			d.Absen = map[string]any{}
		}
		if d.Chatai == nil {
			d.Chatai = map[string]any{}
		}
	}
	return d
}

func simpanGabung(path string, milik []string, nilai any) error {
	gabung := map[string]json.RawMessage{}
	if b, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(b, &gabung)
	}
	for _, k := range milik {
		delete(gabung, k)
	}
	b, err := json.Marshal(nilai)
	if err != nil {
		return err
	}
	milikBaru := map[string]json.RawMessage{}
	if err := json.Unmarshal(b, &milikBaru); err != nil {
		return err
	}
	for k, v := range milikBaru {
		gabung[k] = v
	}
	keluar, err := json.Marshal(gabung)
	if err != nil {
		return err
	}
	return os.WriteFile(path, keluar, 0o644)
}

func (d *DataUser) Simpan() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return simpanGabung(d.path, []string{"banned", "lolosperangkat"}, d)
}

func (d *DataGroup) Simpan() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return simpanGabung(d.path, []string{"welcome", "goodbye", "banmember", "absen", "chatai"}, d)
}
