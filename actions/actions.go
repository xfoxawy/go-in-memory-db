package actions

import (
	"strings"

	"github.com/go-in-memory-db/clients"
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
		write(conn, "please type somthing :D")
		return
	}

	switch strings.ToLower(command[0]) {

	case "help":
		write(conn, data.helpHandler())
	case "qset":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		data.qSetHanlder()
		write(conn, "OK")

	case "qget":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		qget := data.qGetHandler()
		write(conn, qget)

	case "qdel":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		qdel := data.qDelHandler()
		write(conn, qdel)

	case "qsize":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		qsize := data.qSizeHandler()
		write(conn, qsize)

	case "qfront":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		qfront := data.qFrontHandler()
		write(conn, qfront)

	case "qdeq":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		qdeq := data.qDeqHandler()
		write(conn, qdeq)

	case "qenq":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		qenq := data.qEnqHandler()
		write(conn, qenq)

	case "hset":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		hset := data.hSetHandler()
		write(conn, hset)

	case "hget":
		if len(command) < 3 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		hget := data.hGetHandler()
		write(conn, hget)

	case "hgetall":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		hgetall := data.hGetAllHandler()
		write(conn, hgetall)

	case "hdel":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		hdel := data.hDelHandler()
		write(conn, hdel)

	case "hpush":
		if len(command) < 4 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		hpush := data.hPushHandler()
		write(conn, hpush)

	case "hupdate":
		if len(command) < 4 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		hupdate := data.hUpdateHandler()
		write(conn, hupdate)

	case "hrm", "hremove":
		if len(command) < 3 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		hrm := data.hRemoveHandler()
		write(conn, hrm)

	case "hseek":
		if len(command) < 3 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		hseek := data.hSeekHandler()
		write(conn, hseek)

	case "hfind":
		if len(command) < 3 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		hfind := data.hFindHandler()
		write(conn, hfind)

	case "hsize":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		hsize := data.hSizeHandler()
		write(conn, hsize)

	case "lset":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		lset := data.lSetHandler()
		write(conn, lset)

	case "lget":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		lget := data.lGetHandler()
		write(conn, lget)

	case "ldel":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		ldel := data.lDelHandler()
		write(conn, ldel)

	case "lpush":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		if len(command) < 3 {
			write(conn, "TYPE VALUES TO PUSH IT")
			return
		}
		lpush := data.lPushHandler()
		write(conn, lpush)

	case "lpop":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		lpop := data.lPopHandler()
		write(conn, lpop)

	case "lshift":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		lshift := data.lShiftHandler()
		write(conn, lshift)

	case "lunshift":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		lunshift := data.lUnShiftHandler()
		write(conn, lunshift)

	case "lrm", "lremove":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		lremove := data.lRemoveHandler()
		write(conn, lremove)

	case "lunlink":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		lunlink := data.lUnlinkHandler()
		write(conn, lunlink)

	case "lseek":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		lseek := data.lSeekHandler()
		write(conn, lseek)

	case "set":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		set := data.setHandler()
		write(conn, set)

	case "get":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		get := data.getHandler()
		write(conn, get)

	case "del":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		del := data.delHandler()
		write(conn, del)

	case "isset":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		isset := data.issetHandler()
		write(conn, isset)

	case "dump":
		content := data.dumpHandler()
		write(conn, content)

	case "clear":
		clear := data.clearHandler()
		write(conn, clear)

	case "which":
		witch := data.witchHandler()
		write(conn, witch)

	case "use":
		if len(command) < 2 {
			write(conn, "UNEXPECTED KEY")
			return
		}
		use := data.useHandler()
		write(conn, use)

	case "show", "ls":
		show := data.showHandler()
		write(conn, show)

	case "bye":
		data.byeHandler()

	default:
		write(conn, "please use help command to know the commands you can use")
	}
}

func write(c redcon.Conn, s string) {
	if s != "" {
		c.WriteString(s)
	}
}
