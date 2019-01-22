package actions

import (
	"strings"

	"github.com/go-in-memory-db/clients"
	"github.com/go-in-memory-db/logging"
	"github.com/tidwall/redcon"
)

type Actions struct {
	StringArray []string
	Client      *clients.Client
}

func TakeAction(data *Actions) {
	command := data.StringArray
	conn := data.Client.Conn

	if len(command) < 1 {
		write(conn, "please type somthing :D", true)
		return
	}

	switch strings.ToLower(command[0]) {

	case "help":
		write(conn, data.helpHandler(), false)
	case "qset":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		data.qSetHanlder()
		write(conn, "OK", false)

	case "qget":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		qget := data.qGetHandler()
		write(conn, qget, false)

	case "qdel":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		qdel := data.qDelHandler()
		write(conn, qdel, false)

	case "qsize":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		qsize := data.qSizeHandler()
		write(conn, qsize, false)

	case "qfront":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		qfront := data.qFrontHandler()
		write(conn, qfront, false)

	case "qdeq":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		qdeq := data.qDeqHandler()
		write(conn, qdeq, false)

	case "qenq":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		qenq := data.qEnqHandler()
		write(conn, qenq, false)

	case "hset":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		hset := data.hSetHandler()
		write(conn, hset, false)

	case "hget":
		if len(command) < 3 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		hget := data.hGetHandler()
		write(conn, hget, false)

	case "hgetall":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		hgetall := data.hGetAllHandler()
		write(conn, hgetall, false)

	case "hdel":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		hdel := data.hDelHandler()
		write(conn, hdel, false)

	case "hpush":
		if len(command) < 4 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		hpush := data.hPushHandler()
		write(conn, hpush, false)

	case "hrm", "hremove":
		if len(command) < 3 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		hrm := data.hRemoveHandler()
		write(conn, hrm, false)

	case "lset":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		lset := data.lSetHandler()
		write(conn, lset, false)

	case "lget":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		lget := data.lGetHandler()
		write(conn, lget, false)

	case "ldel":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		ldel := data.lDelHandler()
		write(conn, ldel, false)

	case "lpush":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		if len(command) < 3 {
			write(conn, "TYPE VALUES TO PUSH IT", true)
			return
		}
		lpush := data.lPushHandler()
		write(conn, lpush, false)

	case "lpop":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		lpop := data.lPopHandler()
		write(conn, lpop, false)

	case "lshift":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		lshift := data.lShiftHandler()
		write(conn, lshift, false)

	case "lunshift":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		lunshift := data.lUnShiftHandler()
		write(conn, lunshift, false)

	case "lrm", "lremove":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		lremove := data.lRemoveHandler()
		write(conn, lremove, false)

	case "lunlink":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		lunlink := data.lUnlinkHandler()
		write(conn, lunlink, false)

	case "lseek":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		lseek := data.lSeekHandler()
		write(conn, lseek, false)

	case "set":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		set := data.setHandler()
		write(conn, set, false)

	case "get":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		get := data.getHandler()
		write(conn, get, false)

	case "del":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		del := data.delHandler()
		write(conn, del, false)

	case "isset":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		isset := data.issetHandler()
		write(conn, isset, false)

	case "dump":
		content := data.dumpHandler()
		write(conn, content, false)

	case "clear":
		clear := data.clearHandler()
		write(conn, clear, false)

	case "which":
		witch := data.witchHandler()
		write(conn, witch, false)

	case "use":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY", true)
			return
		}
		use := data.useHandler()
		write(conn, use, false)

	case "show", "ls":
		show := data.showHandler()
		write(conn, show, false)

	case "bye":
		data.byeHandler()

	default:
		write(conn, "please use help command to know the commands you can use", true)
	}
}

func write(c redcon.Conn, s string, withError bool) {
	if s != "" {
		c.WriteString(s)
	}
	if withError {
		logging.LoggingLog("user", "file", s)
	}
}
