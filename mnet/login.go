package mnet

import (
	"net"

	"github.com/Hucaru/Valhalla/constant"
	"github.com/Hucaru/Valhalla/database"
	"github.com/Hucaru/Valhalla/mnet/crypt"

	"github.com/Hucaru/Valhalla/mpacket"
)

type MConnLogin interface {
	MConn

	GetLogedIn() bool
	SetLogedIn(bool)
	GetAccountID() int32
	SetAccountID(int32)
	GetGender() byte
	SetGender(byte)
	GetWorldID() byte
	SetWorldID(byte)
	GetChannelID() byte
	SetChannelID(byte)
	GetAdminLevel() int
	SetAdminLevel(int)
}

type login struct {
	baseConn

	logedIn    bool
	accountID  int32
	gender     byte
	worldID    byte
	channelID  byte
	adminLevel int
}

func NewLogin(conn net.Conn, eRecv chan *Event, queueSize int, keySend, keyRecv [4]byte) *login {
	l := &login{}
	l.Conn = conn

	l.eSend = make(chan mpacket.Packet, queueSize)
	l.eRecv = eRecv
	l.endSend = make(chan bool, 1)

	l.cryptSend = crypt.New(keySend, constant.MapleVersion)
	l.cryptRecv = crypt.New(keyRecv, constant.MapleVersion)

	l.reader = func() {
		clientReader(l, l.eRecv, constant.MapleVersion, constant.ClientHeaderSize, l.cryptRecv)
	}

	return l
}

func (l *login) Cleanup() {
	l.baseConn.Cleanup()

	if l.logedIn {
		records, err := database.Handle.Query("UPDATE accounts SET isLogedIn=? WHERE accountID=?", 0, l.accountID)

		defer records.Close()

		if err != nil {
			panic(err)
		}
	}
}

func (l *login) GetLogedIn() bool {
	return l.logedIn
}

func (l *login) SetLogedIn(logedIn bool) {
	l.logedIn = logedIn
}

func (l *login) GetAccountID() int32 {
	return l.accountID
}

func (l *login) SetAccountID(accountID int32) {
	l.accountID = accountID
}

func (l *login) GetGender() byte {
	return l.gender
}

func (l *login) SetGender(gender byte) {
	l.gender = gender
}

func (l *login) GetWorldID() byte {
	return l.worldID
}

func (l *login) SetWorldID(id byte) {
	l.worldID = id
}

func (l *login) GetChannelID() byte {
	return l.channelID
}

func (l *login) SetChannelID(id byte) {
	l.channelID = id
}

func (l *login) GetAdminLevel() int {
	return l.adminLevel
}

func (l *login) SetAdminLevel(level int) {
	l.adminLevel = level
}
