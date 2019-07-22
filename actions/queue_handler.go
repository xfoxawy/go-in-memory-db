package actions

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (a *Actions) qSetHanlder() string {
	k := a.StringArray[1]
	v := a.StringArray[2:]

	queue, err := a.Client.Dbpointer.CreateQueue(k)

	if err != nil {
		return err.Error()
	}

	for i := range v {
		queue.Enqueue(v[i])
	}

	return "OK"

}

// output bug, multiline
func (a *Actions) qGetHandler() string {
	k := a.StringArray[1]

	if q, err := a.Client.Dbpointer.GetQueue(k); err == nil {

		if q.Size() == 0 {
			return ""
		}

		output := make([]string, 0)
		current := q.Queue.Start
		for current.Next != nil {
			output = append(output, current.Value)
			current = current.Next
		}
		output = append(output, q.Queue.End.Value)
		s, _ := json.Marshal(output)
		return fmt.Sprintln(string(s))
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
