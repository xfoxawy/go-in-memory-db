package actions

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
		return value
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
		write(a.Client.Conn, k+" "+v)
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
	hash.Values = hash.Push(mapKey, value)
	return "OK"
}

func (a *Actions) hRemoveHandler() string {
	k := a.StringArray[1]
	mapKey := a.StringArray[2]

	if hash, err := a.Client.Dbpointer.GetHashTable(k); err == nil {

		hash.Values = hash.Remove(mapKey)
		return "OK"
	}
	return "Hash table Does not Exist"
}
