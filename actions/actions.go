package actions

import (
	"io"
	"strings"

	"github.com/go-in-memory-db/clients"
)

type Actions struct {
	StringArray []string
	Client      *clients.Client
}

func TakeAction(ch chan *Actions) {
	for {
		data := <-ch
		command := data.StringArray
		conn := data.Client.Conn
		if len(command) < 1 {
			write(conn, "please type somthing :D")
			continue
		}
		switch strings.ToLower(command[0]) {
		case "help":
			write(conn, data.helpHandler())
		case "qset":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			data.qSetHanlder()
			write(conn, "OK")

		case "qget":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			qget := data.qGetHandler()
			write(conn, qget)

		case "qdel":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			qdel := data.qDelHandler()
			write(conn, qdel)

		case "qsize":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			qsize := data.qSizeHandler()
			write(conn, qsize)

		case "qfront":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			qfront := data.qFrontHandler()
			write(conn, qfront)

		case "qdeq":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			qdeq := data.qDeqHandler()
			write(conn, qdeq)

		case "qenq":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			qenq := data.qEnqHandler()
			write(conn, qenq)

		case "hset":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			hset := data.hSetHandler()
			write(conn, hset)

		case "hget":
			if len(command) < 3 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			hget := data.hGetHandler()
			write(conn, hget)

		case "hgetall":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			hgetall := data.hGetAllHandler()
			write(conn, hgetall)

		case "hdel":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			hdel := data.hDelHandler()
			write(conn, hdel)

		case "hpush":
			if len(command) < 4 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			hpush := data.hPushHandler()
			write(conn, hpush)

		case "hrm", "hremove":
			if len(command) < 3 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			hrm := data.hRemoveHandler()
			write(conn, hrm)

		case "lset":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			lset := data.lSetHandler()
			write(conn, lset)

		case "lget":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			lget := data.lGetHandler()
			write(conn, lget)

		case "ldel":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			ldel := data.lDelHandler()
			write(conn, ldel)

		case "lpush":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			if len(command) < 3 {
				write(conn, "TYPE VALUES TO PUSH IT")
				continue
			}
			lpush := data.lPushHandler()
			write(conn, lpush)

		case "lpop":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			lpop := data.lPopHandler()
			write(conn, lpop)

		case "lshift":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			lshift := data.lShiftHandler()
			write(conn, lshift)

		case "lunshift":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			lunshift := data.lUnShiftHandler()
			write(conn, lunshift)

		case "lrm", "lremove":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			lremove := data.lRemoveHandler()
			write(conn, lremove)

		case "lunlink":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			lunlink := data.lUnlinkHandler()
			write(conn, lunlink)

		case "lseek":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			lseek := data.lSeekHandler()
			write(conn, lseek)

		case "set":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			set := data.setHandler()
			write(conn, set)

		case "get":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			get := data.getHandler()
			write(conn, get)

		case "del":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			del := data.delHandler()
			write(conn, del)

		case "isset":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
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
				continue
			}
			use := data.useHandler()
			write(conn, use)

		case "show", "ls":
			show := data.showHandler()
			write(conn, show)

		case "publish":
			publish := data.publishHandler()
			write(conn, publish)

		// case "subscribe":
		// 	subscribe := data.subscribeHandler()
		// 	write(conn, subscribe)
		//
		case "bye":
			data.byeHandler()

		default:
			write(conn, "please use help command to know the commands you can use")
		}
	}
}

func write(w io.Writer, s string) {
	io.WriteString(w, s+"\n")
}
