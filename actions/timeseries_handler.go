package actions

func (a *Actions) tsSetHandler() string {
	k := a.StringArray[1]
	a.Client.Dbpointer.CreateTimeseries(k)
	return k
}
