package actions

import (
	"strconv"
)

func (a *Actions) sSetHanlder() {
	k := a.StringArray[1]
	v := a.StringArray[2:]

	stack, _ := a.Client.Dbpointer.CreateStack(k)

	for i := range v {
		stack.Push(v[i])
	}

}

func (a *Actions) sGetHandler() string {

	k := a.StringArray[1]
	if s, err := a.Client.Dbpointer.GetStack(k); err == nil {
		a.Client.Conn.WriteString(s.Stack.Start.Value)
		current := s.Stack.Start
		for current.Next != nil {
			current = current.Next
			a.Client.Conn.WriteString(current.Value)
		}
		return ""
	}
	return "Stack Does not Exist"
}

func (a *Actions) sDelHandler() string {
	k := a.StringArray[1]
	if _, err := a.Client.Dbpointer.GetStack(k); err == nil {
		a.Client.Dbpointer.DelStack(k)
		return "OK"
	}
	return "Stack Does not Exist"
}

func (a *Actions) sSizeHandler() string {
	k := a.StringArray[1]
	if stack, err := a.Client.Dbpointer.GetStack(k); err == nil {
		stringVal := strconv.Itoa(stack.Size())
		return stringVal
	}
	return "Stack Does not Exist"
}

func (a *Actions) sPopHandler() string {
	k := a.StringArray[1]
	if stack, err := a.Client.Dbpointer.GetStack(k); err == nil {
		return stack.Pop()
	}
	return "Stack Does not Exist"
}

func (a *Actions) sPushHandler() string {
	k := a.StringArray[1]
	v := a.StringArray[2:]

	if stack, err := a.Client.Dbpointer.GetStack(k); err == nil {
		for i := range v {
			stack.Push(v[i])
		}
		return "Ok"
	}
	return "Stack Does not Exist"
}
