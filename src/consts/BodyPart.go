package consts

type BodyPart string

const (
	MOVE          = "move"
	WORK          = "work"
	CARRY         = "carry"
	ATTACK        = "attack"
	RANGED_ATTACK = "ranged_attack"
	TOUGH         = "tough"
	HEAL          = "heal"
	CLAIM         = "claim"
)

// func (bp BodyPart) String() string {
// 	return [...]string{"move", "work", "carry", "attck", "ranged_attack", "tough", "heal", "claim"}[bp]
// }
