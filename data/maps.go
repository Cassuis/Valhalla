package data

import (
	"sync"

	"github.com/Hucaru/Valhalla/nx"

	"github.com/Hucaru/Valhalla/interfaces"
)

type mMap map[uint32]*mapleMap

var mapleMaps = make(mMap)
var mapleMapsMutex = &sync.RWMutex{}

// GetMapsPtr -
func GetMapsPtr() mMap {
	return mapleMaps
}

// GenerateMapsObject -
func GenerateMapsObject() {
	for mapID, stage := range nx.Maps {
		m := &mapleMap{}

		m.SetReturnMap(stage.ReturnMap)

		for _, life := range stage.Life {
			if life.IsMob {
				l := &mapleMob{}

				l.SetID(life.ID)
				l.SetX(life.X)
				l.SetY(life.Y)
				l.SetFoothold(life.Fh)
				l.SetFace(life.F)
				l.SetMobTime(life.MobTime)

				mon := nx.Mob[life.ID]

				l.SetBoss(mon.Boss)
				l.SetEXP(mon.Exp)
				l.SetMaxHp(mon.MaxHp)
				l.SetHp(mon.MaxHp)
				l.SetMaxMp(mon.MaxMp)
				l.SetMp(mon.MaxMp)
				l.SetLevel(mon.Level)

				m.AddMob(l)

			} else {
				l := &mapleNpc{}

				l.SetID(life.ID)
				l.SetX(life.X)
				l.SetY(life.Y)
				l.SetRx0(life.Rx0)
				l.SetRx1(life.Rx1)
				l.SetFoothold(life.Fh)
				l.SetFace(life.F)

				m.AddNpc(l)
			}
		}

		for _, portal := range stage.Portals {
			p := &maplePortal{}

			p.SetName(portal.Name)
			p.SetX(portal.X)
			p.SetY(portal.Y)
			p.SetIsSpawn(portal.IsSpawn)
			p.SetToMap(portal.Tm)
			p.SetToPortal(portal.Tn)

			m.AddPortal(p)
		}

		mapleMapsMutex.Lock()
		mapleMaps[mapID] = m
		mapleMapsMutex.Unlock()
	}
}

func (mM mMap) GetMap(mapID uint32) interfaces.Map {
	mapleMapsMutex.RLock()
	result := mapleMaps[mapID]
	mapleMapsMutex.RUnlock()

	return result
}

type mapleMap struct {
	npcs         []interfaces.Npc
	mobs         []interfaces.Mob
	forcedReturn uint32
	returnMap    uint32
	mobRate      float64
	isTown       bool
	portals      []interfaces.Portal
	mutex        sync.RWMutex
	players      []interfaces.ClientConn
}

func (m *mapleMap) GetNpcs() []interfaces.Npc {
	m.mutex.RLock()
	result := m.npcs
	m.mutex.RUnlock()

	return result
}

func (m *mapleMap) AddNpc(npc interfaces.Npc) {
	m.mutex.Lock()
	m.npcs = append(m.npcs, npc)
	m.mutex.Unlock()
}

func (m *mapleMap) GetMobs() []interfaces.Mob {
	m.mutex.RLock()
	result := m.mobs
	m.mutex.RUnlock()

	return result
}

func (m *mapleMap) AddMob(mob interfaces.Mob) {
	m.mutex.Lock()
	m.mobs = append(m.mobs, mob)
	m.mutex.Unlock()
}

func (m *mapleMap) GetReturnMap() uint32 {
	m.mutex.RLock()
	result := m.returnMap
	m.mutex.RUnlock()

	return result
}

func (m *mapleMap) SetReturnMap(mapID uint32) {
	m.mutex.Lock()
	m.returnMap = mapID
	m.mutex.Unlock()
}

func (m *mapleMap) GetPortals() []interfaces.Portal {
	m.mutex.RLock()
	result := m.portals
	m.mutex.RUnlock()

	return result
}

func (m *mapleMap) AddPortal(portal interfaces.Portal) {
	m.mutex.Lock()
	m.portals = append(m.portals, portal)
	m.mutex.Unlock()
}

func (m *mapleMap) GetPlayers() []interfaces.ClientConn {
	m.mutex.RLock()
	result := m.players
	m.mutex.RUnlock()

	return result
}

func (m *mapleMap) AddPlayer(player interfaces.ClientConn) {
	m.mutex.Lock()
	m.players = append(m.players, player)
	m.mutex.Unlock()
}

func (m *mapleMap) RemovePlayer(player interfaces.ClientConn) {
	index := -1

	m.mutex.RLock()
	for i, v := range m.players {
		if v == player {
			index = i
			break
		}
	}
	m.mutex.RUnlock()

	if index < 0 {
		return
	}

	m.mutex.Lock()
	m.players[index] = m.players[len(m.players)-1]
	m.players[len(m.players)-1] = nil
	m.players = m.players[:len(m.players)-1]
	m.mutex.Unlock()
}
