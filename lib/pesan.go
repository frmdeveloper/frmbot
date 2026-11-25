package lib

type GrupMention struct {
	JID    string
	Subjek string
}

type Pesan struct {
	ID            string
	From          string
	Sender        string
	SenderAlt     string
	Text          string
	Command       string
	Prefix        string
	Q             string
	IsGroup       bool
	FromMe        bool
	Type          string
	Media         string
	Mime          string
	Mentions      []string
	GroupMentions []GrupMention
	Quoted        *Pesan
	Full          any
}

type IsiPesan struct {
	Text          string
	Caption       string
	Mentions      []string
	GroupMentions []GrupMention
	Image         []byte
	Video         []byte
	Audio         []byte
	PTT           bool
	Sticker       []byte
	Document      []byte
	FileName      string
	Mimetype      string
}

type Peserta struct {
	JID          string
	Nomor        string
	LID          string
	IsAdmin      bool
	IsSuperAdmin bool
}

type GrupInfo struct {
	JID        string
	Nama       string
	Pemilik    string
	Announce   bool
	Peserta    []Peserta
}

type Tombol struct {
	ID   string
	Teks string
}

type TombolURL struct {
	Teks string
	URL  string
}

type BarisList struct {
	ID        string
	Judul     string
	Deskripsi string
}

type BagianList struct {
	Judul string
	Baris []BarisList
}

type Conn interface {
	SendMessage(jid string, isi IsiPesan, quoted *Pesan) error
	OK(m *Pesan, teks string) error
	DownloadAny(m *Pesan) ([]byte, string, error)
	GetGroupInfo(jid string) (*GrupInfo, error)
	GetJoinedGroups() ([]*GrupInfo, error)
	UpdateGroupParticipants(jid string, peserta []string, aksi string) error
	GetProfilePicture(jid string) (string, error)
	GetGroupInviteLink(jid string) (string, error)
	GetGroupInviteInfo(link string) (*GrupInfo, error)
	SetGroupName(jid string, nama string) error
	SetGroupAnnounce(jid string, announce bool) error
	IsOnWhatsApp(nomor string) (bool, string, error)
	KirimStikerBersuara(jid string, stiker, audio []byte, detik int, quoted *Pesan, mention string) error
	SendPresence(state string) error
	SendChatPresence(jid string, state string) error
	RevokeMessage(jid string, m *Pesan) error
	RelayMentah(jid string, mentah any) error
	React(m *Pesan, emoji string) error
	SendButtons(jid string, teks string, footer string, buttons []Tombol, quoted *Pesan) error
	SendList(jid string, judul, teks, footer, teksTombol string, bagian []BagianList, quoted *Pesan) error
	BotJID() string
	BotLID() string
	BotNomor() string
	Terhubung() bool
}

type Konteks struct {
	Conn  Conn
	M     *Pesan
	User  *DataUser
	Group *DataGroup
}
