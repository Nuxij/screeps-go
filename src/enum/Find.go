package enum

type FIND int

const (
	FIND_EXIT_TOP                   FIND = 1
	FIND_EXIT_RIGHT                 FIND = 3
	FIND_EXIT_BOTTOM                FIND = 5
	FIND_EXIT_LEFT                  FIND = 7
	FIND_EXIT                       FIND = 10
	FIND_CREEPS                     FIND = 101
	FIND_MY_CREEPS                  FIND = 102
	FIND_HOSTILE_CREEPS             FIND = 103
	FIND_SOURCES_ACTIVE             FIND = 104
	FIND_SOURCES                    FIND = 105
	FIND_DROPPED_RESOURCES          FIND = 106
	FIND_STRUCTURES                 FIND = 107
	FIND_MY_STRUCTURES              FIND = 108
	FIND_HOSTILE_STRUCTURES         FIND = 109
	FIND_FLAGS                      FIND = 110
	FIND_CONSTRUCTION_SITES         FIND = 111
	FIND_MY_SPAWNS                  FIND = 112
	FIND_HOSTILE_SPAWNS             FIND = 113
	FIND_MY_CONSTRUCTION_SITES      FIND = 114
	FIND_HOSTILE_CONSTRUCTION_SITES FIND = 115
	FIND_MINERALS                   FIND = 116
	FIND_NUKES                      FIND = 117
	FIND_TOMBSTONES                 FIND = 118
	FIND_POWER_CREEPS               FIND = 119
	FIND_MY_POWER_CREEPS            FIND = 120
	FIND_HOSTILE_POWER_CREEPS       FIND = 121
	FIND_DEPOSITS                   FIND = 122
	FIND_RUINS                      FIND = 123
)
