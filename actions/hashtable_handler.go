package actions

import (
	"strconv"
)

func (a *Actions) hSetHandler() string {
	k := a.StringArray[1]
	a.Client.Dbpointer.CreateHashTable(k)
	return k
}

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

	for k, v := range v.Values {
		write(a.Client.Conn, k+" "+v.Value().(string))
	}
	return ""
}

func (a *Actions) hDelHandler() string {
	k := a.StringArray[1]
	if _, err := a.Client.Dbpointer.GetHashTable(k); err == nil {
		a.Client.Dbpointer.DelHashTable(k)
		return k + " " + "has deleted"
	}
	return "Hash table Does not Exist"
}

func (a *Actions) hPushHandler() string {
	k := a.StringArray[1]
	mapKey := a.StringArray[2]

	hash, err := a.Client.Dbpointer.GetHashTable(k)
	if err != nil {
		return "Hash table Does not Exist"
	}

	value := a.StringArray[3]
	hash.Values = hash.Insert(mapKey, value).Values
	return "OK"
}

func (a *Actions) hUpdateHandler() string {
	k := a.StringArray[1]
	mapKey := a.StringArray[2]

	hash, err := a.Client.Dbpointer.GetHashTable(k)
	if err != nil {
		return "Hash table Does not Exist"
	}

	value := a.StringArray[3]
	hash.Values = hash.Update(mapKey, value).Values
	return "OK"
}

func (a *Actions) hRemoveHandler() string {
	k := a.StringArray[1]
	mapKey := a.StringArray[2]

	if hash, err := a.Client.Dbpointer.GetHashTable(k); err == nil {

		hash.Values = hash.Remove(mapKey).Values
		return "OK"
	}
	return "Hash table Does not Exist"
}

func (a *Actions) hSeekHandler() string {
	k := a.StringArray[1]
	seekingValue := a.StringArray[2]

	if hash, err := a.Client.Dbpointer.GetHashTable(k); err == nil {

		return hash.Get(seekingValue).Value().(string)
	}
	return "Hash table Does not Exist"
}

func (a *Actions) hSizeHandler() string {
	k := a.StringArray[1]

	if hash, err := a.Client.Dbpointer.GetHashTable(k); err == nil {
		return strconv.Itoa(hash.Size())
	}
	return "Hash table Does not Exist"
}

func (a *Actions) hFindHandler() string {
	k := a.StringArray[1]
	value := a.StringArray[2]

	if hash, err := a.Client.Dbpointer.GetHashTable(k); err == nil {
		return hash.Find(value)
	}
	return "Hash table Does not Exist"
}
