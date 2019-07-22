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
		"help":     {0: data.helpHandler},
		"qset":     {2: data.qSetHanlder},
		"qget":     {2: data.qGetHandler},
		"qdel":     {2: data.qDelHandler},
		"qsize":    {2: data.qSizeHandler},
		"qfront":   {2: data.qFrontHandler},
		"qdeq":     {2: data.qDeqHandler},
		"qenq":     {3: data.qEnqHandler},
		"hset":     {2: data.hSetHandler},
		"hget":     {3: data.hGetHandler},
		"hgetall":  {2: data.hGetAllHandler},
		"hdel":     {2: data.hDelHandler},
		"hpush":    {4: data.hPushHandler},
		"hrm":      {3: data.hRemoveHandler},
		"hremove":  {3: data.hRemoveHandler},
		"lset":     {2: data.lSetHandler},
		"lget":     {2: data.lGetHandler},
		"ldel":     {2: data.lDelHandler},
		"lpush":    {3: data.lPushHandler},
		"lpop":     {2: data.lPopHandler},
		"lshift":   {2: data.lShiftHandler},
		"lunshift": {2: data.lUnShiftHandler},
		"lrm":      {2: data.lRemoveHandler},
		"lremove":  {2: data.lRemoveHandler},
		"lunlink":  {2: data.lUnlinkHandler},
		"lseek":    {2: data.lSeekHandler},
		"set":      {2: data.setHandler},
		"get":      {2: data.getHandler},
		"del":      {2: data.delHandler},
		"isset":    {2: data.issetHandler},
		"dump":     {0: data.dumpHandler},
		"clear":    {0: data.clearHandler},
		"which":    {0: data.whichHandler},
		"use":      {2: data.useHandler},
		"show":     {0: data.showHandler},
		"ls":       {0: data.showHandler},
		"bye":      {0: data.byeHandler},
	}
	return decisions
}
