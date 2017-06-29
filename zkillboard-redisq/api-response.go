package zkillboard_redisq

type SolarSystem struct {
	IDStr string `json:"id_str"`
	Href  string `json:"href"`
	ID    int64  `json:"id"`
	Name  string `json:"name"`
}

type Alliance struct {
	IDStr string `json:"id_str"`
	Href  string `json:"href"`
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Icon  Icon   `json:"icon"`
}

type Icon struct {
	Href string `json:"href"`
}

type ShipType struct {
	IDStr string `json:"id_str"`
	Href  string `json:"href"`
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Icon  Icon   `json:"icon"`
}

type Corporation struct {
	IDStr string `json:"id_str"`
	Href  string `json:"href"`
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Icon  Icon   `json:"icon"`
}

type Character struct {
	IDStr string `json:"id_str"`
	Href  string `json:"href"`
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Icon  Icon   `json:"icon"`
}

type WeaponType struct {
	IDStr string `json:"id_str"`
	Href  string `json:"href"`
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Icon  Icon   `json:"icon"`
}

type Attacker struct {
	Alliance       Alliance    `json:"alliance"`
	ShipType       ShipType    `json:"shipType"`
	Corporation    Corporation `json:"corporation"`
	Character      Character   `json:"character"`
	DamageDoneStr  string      `json:"damageDone_str"`
	WeaponType     WeaponType  `json:"weaponType"`
	FinalBlow      bool        `json:"finalBlow"`
	SecurityStatus float64     `json:"securityStatus"`
	DamageDone     int64       `json:"damageDone"`
}

type Position struct {
	Y float64 `json:"y"`
	X float64 `json:"x"`
	Z float64 `json:"z"`
}

type War struct {
	Href  string `json:"href"`
	ID    int64  `json:"id"`
	IDStr string `json:"id_str"`
}

type ItemType struct {
	IDStr string `json:"id_str"`
	Href  string `json:"href"`
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Icon  Icon   `json:"icon"`
}

type Item struct {
	Singleton            int64    `json:"singleton"`
	ItemType             ItemType `json:"itemType"`
	QuantityDestroyedStr string   `json:"quantityDestroyed_str,omitempty"`
	Flag                 int64    `json:"flag"`
	FlagStr              string   `json:"flag_str"`
	SingletonStr         string   `json:"singleton_str"`
	QuantityDestroyed    int64    `json:"quantityDestroyed,omitempty"`
	QuantityDroppedStr   string   `json:"quantityDropped_str,omitempty"`
	QuantityDropped      int64    `json:"quantityDropped,omitempty"`
}

type Victim struct {
	Alliance       Alliance    `json:"alliance"`
	DamageTaken    int64       `json:"damageTaken"`
	Items          []Item      `json:"items"`
	DamageTakenStr string      `json:"damageTaken_str"`
	Character      Character   `json:"character"`
	ShipType       ShipType    `json:"shipType"`
	Corporation    Corporation `json:"corporation"`
	Position       Position    `json:"position"`
}

type Killmail struct {
	SolarSystem      SolarSystem `json:"solarSystem"`
	KillID           int64       `json:"killID"`
	KillTime         string      `json:"killTime"`
	Attackers        []Attacker  `json:"attackers"`
	AttackerCount    int64       `json:"attackerCount"`
	Victim           Victim      `json:"victim"`
	KillIDStr        string      `json:"killID_str"`
	AttackerCountStr string      `json:"attackerCount_str"`
	War              War         `json:"war"`
}

type Zkb struct {
	LocationID  int64   `json:"locationID"`
	Hash        string  `json:"hash"`
	FittedValue float64 `json:"fittedValue"`
	TotalValue  float64 `json:"totalValue"`
	Points      int64   `json:"points"`
	Npc         bool    `json:"npc"`
	Href        string  `json:"href"`
}

type Package struct {
	KillID   int64    `json:"killID"`
	Killmail Killmail `json:"killmail"`
	Zkb      Zkb      `json:"zkb"`
}

type ApiResponse struct {
	Package *Package `json:"package"`
}
