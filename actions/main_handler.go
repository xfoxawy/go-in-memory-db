package actions

import (
	"bytes"
	"fmt"
	"strings"
)

func (a *Actions) setHandler() string {
	k := a.StringArray[1]
	var v string
	if len(a.StringArray) == 2 {
		v = "NIL"
	} else {
		v = strings.Join(a.StringArray[2:], "")
	}

	a.Client.Dbpointer.Set(k, v)
	return "OK"
}

func (a *Actions) getHandler() string {
	k := a.StringArray[1]
	v, err := a.Client.Dbpointer.Get(k)
	if err != nil {
		return "NIL"
	}
	return v
}

func (a *Actions) delHandler() string {
	k := a.StringArray[1]
	a.Client.Dbpointer.Del(k)
	return "OK"
}

func (a *Actions) issetHandler() string {
	k := a.StringArray[1]
	if a.Client.Dbpointer.Isset(k) {
		return "TRUE"
	}
	return "FALSE"
}

func (a *Actions) dumpHandler() string {
	return a.Client.Dbpointer.Dump()
}

func (a *Actions) clearHandler() string {
	a.Client.Dbpointer.Clear()
	return "OK"
}

func (a *Actions) whichHandler() string {
	return a.Client.Dbpointer.Name()
}

func (a *Actions) useHandler() string {
	key := a.StringArray[1]
	a.Client.UseNewDatabase(key)
	return "OK"
}

func (a *Actions) showHandler() string {
	var content bytes.Buffer
	Databases := a.Client.GetAllDatabases()
	for name := range Databases {
		if name == a.Client.Dbpointer.Name() {
			name = fmt.Sprintf("-> %s\n", name)
		} else {
			name = fmt.Sprintf("%s\n", name)
		}
		content.WriteString(name)
	}
	return content.String()
}

func (a *Actions) byeHandler() string {
	a.Client.Conn.Close()
	return ""
}

func (a *Actions) helpHandler() string {
	return "\n In Memory DB \n\n" +
		"Use: \n" +
		"SET key value \n" +
		"GET key \n" +
		"DEL key \n" +
		"ISSET key \n" +
		"\nLINKED LIST COMMANDS : \n\n" +
		"LSET key \n" +
		"LGET key \n" +
		"LDEL key \n" +
		"LPUSH key value1 value2 etc \n" +
		"LPOP key \n" +
		"LSHIFT key value \n" +
		"LUNSHIFT key \n" +
		"LRM key value1 value2 etc \n" +
		"LREMOVE key value1 value2 etc \n" +
		"LUNLINK key index1 index2 etc \n" +
		"LSEEK key index \n" +
		"\nEND OF LINKED LIST COMMANDS : \n\n" +
		"\nQUEUE COMMANDS : \n\n" +
		"QSET key \n" +
		"QGET key \n" +
		"QDEL key \n" +
		"QSIZE key \n" +
		"QFRONT key \n" +
		"QDEQ key \n" +
		"QENQ key value1 value2 etc \n" +
		"\nEND OF QUEUE COMMANDS : \n\n" +
		"\nHASH TABLE COMMANDS : \n\n" +
		"HSET key \n" +
		"HGET key HashKey \n" +
		"HDEL key \n" +
		"HPUSH key HashKey value \n" +
		"HRM key HashKey c \n" +
		"HREMOVE HashKey c \n" +
		"\nEND OF HASH TABLE COMMANDS : \n\n" +
		"USE name\n" +
		"WHICH \n" +
		"SHOW \n" +
		"CLEAR \n" +
		"DUMP \n" +
		"BYE \n" +
		"HELP ??\n"
}
