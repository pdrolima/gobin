package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type Weapon struct {
	RequireLevel      int16
	RequireClass      int16
	RequireStrength   int16
	RequireCon        int16
	RequireDex        int16
	RequireInt        int16
	DurabilityMin     int32
	DurabilityMax     int32
	ItemClass         int16
	ItemFlag          byte
	Creator           string
	WeaponType        int32
	WeaponClass       int32
	WeaponLevel       int32
	RequireProjectile int32
	DamageHighMin     int32
	DamageHighMax     int32
	AttackSpeed       int32
	AttackRange       float32
	AttackShortRange  float32
}

func main() {

	o := "64004000320000002c01000084680000487100002c00040c440065006e006e0079007e00010000000d0000000f00000062210000e1040000f508000000000000000000001e0000000000a0410000a04002002d00de070000de07000007000000a42a000016000000a92a0000d9000000af2a000003000000ae2a000005000000f3a1000096000000f3a1000096000000d24600005402000009000000"

	b, _ := hex.DecodeString(o)
	buf := bytes.NewReader(b)

	weapon := Weapon{}
	binary.Read(buf, binary.LittleEndian, &weapon.RequireLevel)
	binary.Read(buf, binary.LittleEndian, &weapon.RequireClass)
	binary.Read(buf, binary.LittleEndian, &weapon.RequireStrength)
	binary.Read(buf, binary.LittleEndian, &weapon.RequireCon)
	binary.Read(buf, binary.LittleEndian, &weapon.RequireDex)
	binary.Read(buf, binary.LittleEndian, &weapon.RequireInt)
	binary.Read(buf, binary.LittleEndian, &weapon.DurabilityMin)
	binary.Read(buf, binary.LittleEndian, &weapon.DurabilityMax)
	binary.Read(buf, binary.LittleEndian, &weapon.ItemClass)
	binary.Read(buf, binary.LittleEndian, &weapon.ItemFlag)
	binary.Read(buf, binary.LittleEndian, &weapon.Creator)
	fmt.Println(weapon)
}
