package consts

type ERR int

const (
	OK                       ERR = 0
	ERR_NOT_OWNER            ERR = -1
	ERR_BUSY                 ERR = -4
	ERR_NOT_FOUND            ERR = -5
	ERR_NOT_ENOUGH_RESOURCES ERR = -6
	ERR_INVALID_TARGET       ERR = -7
	ERR_NOT_IN_RANGE         ERR = -9
	ERR_TIRED                ERR = -11
	ERR_NO_BODYPART          ERR = -12
)
