package actions

import "strconv"

/*
type commandsMapper
Schema {
	"command name" : {
		validation length : execution function
	}
}
*/
type commandsMapper map[string]map[int]func() string

/*
DecisionManager struct
commandsMapper contain all commands with it's own emplementation function
commandName will be setted after check if exist
*/
type DecisionManager struct {
	commandsMapper
	commandName string
}

// NewDecisionManager function
func NewDecisionManager(data *Actions) *DecisionManager {
	return &DecisionManager{
		generateAllDecisions(data),
		"",
	}
}

/*
CheckCommandAvailablity function
if command exist it will override on commandName in struct
*/
func (ad *DecisionManager) CheckCommandAvailablity(commandName string) bool {
	_, ok := ad.commandsMapper[commandName]
	if ok {
		ad.commandName = commandName
	}
	return ok
}

// RunCommand function
func (ad *DecisionManager) RunCommand(commandArray []string) string {
	expectedLength := len(commandArray)
	var message string
	for k, v := range ad.commandsMapper[ad.commandName] {
		if expectedLength < k {
			message += "This command takes " + strconv.FormatInt(int64(k-1), 10) + " params Got " + strconv.FormatInt(int64(expectedLength-1), 10)
		} else {
			message += v()
		}
		break
	}
	return message
}

/*
this function will contain all of our commands
with validation number and exection function
*/
func generateAllDecisions(data *Actions) commandsMapper {
	decisions := map[string]map[int]func() string{
		"help": {0: func() string {
			return data.helpHandler()
		}},
		"qset": {2: func() string {
			data.qSetHanlder()
			return "OK"
		}},
		"qget": {2: func() string {
			return data.qGetHandler()
		}},
		"qdel": {2: func() string {
			return data.qDelHandler()
		}},
		"qsize": {2: func() string {
			return data.qSizeHandler()
		}},
		"qfront": {2: func() string {
			return data.qFrontHandler()
		}},
		"qdeq": {2: func() string {
			return data.qDeqHandler()
		}},
		"qenq": {3: func() string {
			return data.qEnqHandler()
		}},
		"hset": {2: func() string {
			return data.hSetHandler()
		}},
		"hget": {3: func() string {
			return data.hGetHandler()
		}},
		"hgetall": {2: func() string {
			return data.hGetAllHandler()
		}},
		"hdel": {2: func() string {
			return data.hDelHandler()
		}},
		"hpush": {4: func() string {
			return data.hPushHandler()
		}},
		"hrm": {3: func() string {
			return data.hRemoveHandler()
		}},
		"hremove": {3: func() string {
			return data.hRemoveHandler()
		}},
		"lset": {2: func() string {
			return data.lSetHandler()
		}},
		"lget": {2: func() string {
			return data.lGetHandler()
		}},
		"ldel": {2: func() string {
			return data.lDelHandler()
		}},
		"lpush": {3: func() string {
			return data.lPushHandler()
		}},
		"lpop": {2: func() string {
			return data.lPopHandler()
		}},
		"lshift": {2: func() string {
			return data.lShiftHandler()
		}},
		"lunshift": {2: func() string {
			return data.lUnShiftHandler()
		}},
		"lrm": {2: func() string {
			return data.lRemoveHandler()
		}},
		"lremove": {2: func() string {
			return data.lRemoveHandler()
		}},
		"lunlink": {2: func() string {
			return data.lUnlinkHandler()
		}},
		"lseek": {2: func() string {
			return data.lSeekHandler()
		}},
		"set": {2: func() string {
			return data.setHandler()
		}},
		"get": {2: func() string {
			return data.getHandler()
		}},
		"del": {2: func() string {
			return data.delHandler()
		}},
		"isset": {2: func() string {
			return data.issetHandler()
		}},
		"dump": {0: func() string {
			return data.dumpHandler()
		}},
		"clear": {0: func() string {
			return data.clearHandler()
		}},
		"which": {0: func() string {
			return data.witchHandler()
		}},
		"use": {2: func() string {
			return data.useHandler()
		}},
		"show": {0: func() string {
			return data.showHandler()
		}},
		"ls": {0: func() string {
			return data.showHandler()
		}},
		"bye": {0: func() string {
			data.byeHandler()
			return ""
		}},
	}
	return decisions
}
