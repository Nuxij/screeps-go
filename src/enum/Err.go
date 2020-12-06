package enum

var Errors = Enum{[]EnumItem{
	{0, "ERR_OK"},
	{-1, "ERR_NOT_OWNER"},
	{-4, "ERR_BUSY"},
	{-5, "ERR_NOT_FOUND"},
	{-6, "ERR_NOT_ENOUGH_RESOURCES"},
	{-7, "ERR_INVALID_TARGET"},
	{-9, "ERR_NOT_IN_RANGE"},
	{-11, "ERR_TIRED"},
	{-12, "ERR_NO_BODYPART"},
}}
