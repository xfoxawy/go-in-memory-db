package actions

import (
	"strconv"
	"strings"
)

func (a *Actions) lSetHandler() string {
	k := a.StringArray[1]
	v := a.StringArray[2:]

	list := a.Client.Dbpointer.CreateList(k)

	for i := range v {
		list.Push(v[i])
	}
	return "OK"
}

// test it
func (a *Actions) lGetHandler() string {
	k := a.StringArray[1]
	v, err := a.Client.Dbpointer.GetList(k)

	if err != nil || v.Start == nil {
		return "empty or not exit"
	}
	write(a.Client.Conn, v.Start.Value)
	current := v.Start
	for current.Next != nil {
		current = current.Next
		write(a.Client.Conn, current.Value)
	}
	return ""
}

func (a *Actions) lDelHandler() string {
	k := a.StringArray[1]
	if _, err := a.Client.Dbpointer.GetList(k); err == nil {
		a.Client.Dbpointer.DelList(k)
		return "OK"
	}
	return "List Does not Exist"
}

func (a *Actions) lPushHandler() string {
	k := a.StringArray[1]
	list, err := a.Client.Dbpointer.GetList(k)
	if err != nil {
		return "List Does not Exist"
	}
	values := a.StringArray[2:]
	for i := range values {
		list.Push(values[i])
	}
	return "OK"
}

func (a *Actions) lPopHandler() string {
	k := a.StringArray[1]
	if list, err := a.Client.Dbpointer.GetList(k); err == nil {
		p, err := list.Pop()
		if err != nil {
			return "list is empty"
		}
		return p.Value
	}
	return "List Does not Exist"
}

//test it
func (a *Actions) lShiftHandler() string {
	k := a.StringArray[1]
	var v string

	if len(a.StringArray) == 2 {
		v = "NIL"
	} else {
		v = strings.Join(a.StringArray[2:], "")
	}

	if list, err := a.Client.Dbpointer.GetList(k); err == nil {
		list.Shift(v)
		return "OK"
	}
	return "List Does not Exist"
}

func (a *Actions) lUnShiftHandler() string {
	k := a.StringArray[1]

	if list, err := a.Client.Dbpointer.GetList(k); err == nil {
		unshifted, err := list.Unshift()
		if err != nil {
			return "list is empty"
		}
		return unshifted.Value
	}
	return "List Does not Exist"
}

func (a *Actions) lRemoveHandler() string {
	k := a.StringArray[1]
	values := a.StringArray[2:]
	if list, err := a.Client.Dbpointer.GetList(k); err == nil {
		for i := range values {
			err := list.Remove(values[i])
			if err != nil {
				return "list is empty"
			}
		}
		return "OK"
	}
	return "List Does not Exist"
}

func (a *Actions) lUnlinkHandler() string {
	k := a.StringArray[1]
	values := a.StringArray[2:]
	if list, err := a.Client.Dbpointer.GetList(k); err == nil {

		for i := range values {
			intVal, _ := strconv.Atoi(values[i])
			err := list.Unlink(intVal)
			if err != nil {
				return "LinkedList is empty OR Step Not Exist"
			}
		}
		return "OK"
	}
	return "List Does not Exist"
}

func (a *Actions) lSeekHandler() string {
	k := a.StringArray[1]
	var v string
	if len(a.StringArray) == 2 {
		v = "NIL"
	} else {
		v = strings.Join(a.StringArray[2:], "")
	}
	if list, err := a.Client.Dbpointer.GetList(k); err == nil {
		intVal, err := strconv.Atoi(v)
		if err != nil {
			return "LinkedList is empty OR Step Not Exist"
		}
		value, err := list.Seek(intVal)
		if err != nil {
			return "LinkedList is empty OR Step Not Exist"
		}
		return value
	}
	return "List Does not Exist"
}
