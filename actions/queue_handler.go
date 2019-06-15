package actions

import (
	"strconv"
)

func (a *Actions) qSetHanlder() {
	k := a.StringArray[1]
	v := a.StringArray[2:]

	queue, _ := a.Client.Dbpointer.CreateQueue(k)

	for i := range v {
		queue.Enqueue(v[i])
	}

}

func (a *Actions) qGetHandler() string {

	k := a.StringArray[1]
	if q, err := a.Client.Dbpointer.GetQueue(k); err == nil {
		write(a.Client.Conn, q.Queue.Start.Value)
		current := q.Queue.Start
		for current.Next != nil {
			current = current.Next
			write(a.Client.Conn, current.Value)
		}
		return ""
	}
	return "Queue Does not Exist"
}

func (a *Actions) qDelHandler() string {
	k := a.StringArray[1]
	if _, err := a.Client.Dbpointer.GetQueue(k); err == nil {
		a.Client.Dbpointer.DelQueue(k)
		return "OK"
	}
	return "Queue Does not Exist"
}

func (a *Actions) qSizeHandler() string {
	k := a.StringArray[1]
	if queue, err := a.Client.Dbpointer.GetQueue(k); err == nil {
		stringVal := strconv.Itoa(queue.Size())
		return stringVal
	}
	return "Queue Does not Exist"
}

func (a *Actions) qFrontHandler() string {
	k := a.StringArray[1]
	if queue, err := a.Client.Dbpointer.GetQueue(k); err == nil {
		return queue.Front()
	}
	return "Queue Does not Exist"
}

func (a *Actions) qDeqHandler() string {
	k := a.StringArray[1]
	if queue, err := a.Client.Dbpointer.GetQueue(k); err == nil {
		return queue.Dequeue()
	}
	return "Queue Does not Exist"
}

func (a *Actions) qEnqHandler() string {
	k := a.StringArray[1]
	v := a.StringArray[2:]

	if queue, err := a.Client.Dbpointer.GetQueue(k); err == nil {
		for i := range v {
			queue.Enqueue(v[i])
		}
		return "Ok"
	}
	return "Queue Does not Exist"
}
