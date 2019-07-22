package actions

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (a *Actions) hSetHandler() string {
	k := a.StringArray[1]

	h, err := a.Client.Dbpointer.CreateHashTable(k)

	if err != nil {
		return err.Error()
	}

	st := len(a.StringArray[2:])
	if st%2 == 0 {
		for i := 2; i <= st; i = i + 2 {
			h.Insert(a.StringArray[i], a.StringArray[i+1])
		}

	}

	return "OK"
}

// multi k bug
func (a *Actions) hGetHandler() string {
	k := a.StringArray[1]
	mapKey := a.StringArray[2]
	v, err := a.Client.Dbpointer.GetHashTable(k)

	if err != nil {
		return "Hash table Does not Exist"
	}
	if value, ok := v.Values[mapKey]; ok {
		return value.Value().(string)
	}
	return "This Key Does not Exist"
}

/**
* @to-do
* test this function
 */
func (a *Actions) hGetAllHandler() string {
	k := a.StringArray[1]
	v, err := a.Client.Dbpointer.GetHashTable(k)

	if err != nil {
		return "Hash table Does not Exist"
	}

	outs := make(map[string]interface{}, len(v.Values))

	for k, v := range v.Values {
		outs[k] = v.Value().(string)
	}

	out, _ := json.Marshal(outs)
	return fmt.Sprintln(string(out))
}

func (a *Actions) hDelHandler() string {
	k := a.StringArray[1]
	if _, err := a.Client.Dbpointer.GetHashTable(k); err == nil {
		a.Client.Dbpointer.DelHashTable(k)
		return k + " " + "is deleted"
	}
	return "Hash table Does not Exist"
}

func (a *Actions) hPushHandler() string {
	k := a.StringArray[1]
	h, err := a.Client.Dbpointer.GetHashTable(k)

	if err != nil {
		return "Hash table Does not Exist"
	}

	st := len(a.StringArray[2:])
	if st%2 == 0 {
		for i := 2; i <= st; i = i + 2 {
			h.Insert(a.StringArray[i], a.StringArray[i+1])
		}

	}
	return "OK"
}

func (a *Actions) hUpdateHandler() string {
	k := a.StringArray[1]
	h, err := a.Client.Dbpointer.GetHashTable(k)

	if err != nil {
		return "Hash table Does not Exist"
	}

	st := len(a.StringArray[2:])
	if st%2 == 0 {
		for i := 2; i <= st; i = i + 2 {
			h.Update(a.StringArray[i], a.StringArray[i+1])
		}

	}

	return "OK"
}

func (a *Actions) hRemoveHandler() string {
	k := a.StringArray[1]

	h, err := a.Client.Dbpointer.GetHashTable(k)

	if err != nil {
		return "Hash table Does not Exist"
	}

	cst := make([]string, len(a.StringArray[2:]))
	copy(cst, a.StringArray[2:])

	for i := 0; i < len(cst); i++ {
		h.Remove(cst[i])
	}

	return "OK"
}

func (a *Actions) hSizeHandler() string {
	k := a.StringArray[1]

	if hash, err := a.Client.Dbpointer.GetHashTable(k); err == nil {
		return strconv.Itoa(hash.Size())
	}
	return "Hash table Does not Exist"
}
