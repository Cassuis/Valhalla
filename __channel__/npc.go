package channel

import (
	"github.com/Hucaru/Valhalla/mnet"
	"github.com/Hucaru/Valhalla/packets"
)

type mapleNpc struct {
	id, spawnID              int32
	x, y, rx0, rx1, foothold int16
	face, state              byte
}

func (n *mapleNpc) SetID(id int32) {
	n.id = id
}

func (n *mapleNpc) GetID() int32 {
	return n.id
}

func (n *mapleNpc) SetSpawnID(spawnID int32) {
	n.spawnID = spawnID
}

func (n *mapleNpc) GetSpawnID() int32 {
	return n.spawnID
}

func (n *mapleNpc) SetX(x int16) {
	n.x = x
}

func (n *mapleNpc) GetX() int16 {
	return n.x
}

func (n *mapleNpc) SetY(y int16) {
	n.y = y
}

func (n *mapleNpc) GetY() int16 {
	return n.y
}

func (n *mapleNpc) SetRx0(rx0 int16) {
	n.rx0 = rx0
}

func (n *mapleNpc) GetRx0() int16 {
	return n.rx0
}

func (n *mapleNpc) SetRx1(rx1 int16) {
	n.rx1 = rx1
}

func (n *mapleNpc) GetRx1() int16 {
	return n.rx1
}

func (n *mapleNpc) SetFoothold(y int16) {
	n.foothold = y
}

func (n *mapleNpc) GetFoothold() int16 {
	return n.foothold
}

func (n *mapleNpc) SetFace(face byte) {
	n.face = face
}

func (n *mapleNpc) GetFace() byte {
	return n.face
}

func (n *mapleNpc) SetState(state byte) {
	n.state = state
}

func (n *mapleNpc) GetState() byte {
	return n.state
}

func (n *mapleNpc) Show(conn mnet.MConnChannel) {
	conn.Send(packets.NpcShow(n))
	conn.Send(packets.NPCSetController(n.GetSpawnID(), true))
}

func (n *mapleNpc) Hide(conn mnet.MConnChannel) {
	conn.Send(packets.NPCSetController(n.GetSpawnID(), false))
	conn.Send(packets.NPCRemove(n.GetSpawnID()))
}
