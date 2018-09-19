package main

import "errors"

type List interface {
	setList(k string ,v []string) bool
	getList(k string) ([]string,error)
	delList(k string) bool
	lPush(k string ,v string) bool
	lDel(k string ,v string) bool
}

type list struct {
	db  *database
}

func (l *list) setList(k string , v []string) bool {
	l.db.dataList[k] = v
	return true
}

func (l *list) getList(k string) ([]string , error) {
	if v, ok := l.db.dataList[k];ok {
		return v, nil
	}
	var empty []string
	return empty, errors.New("not found")
}

func (l *list) delList(k string) bool {
	delete(l.db.dataList,k)
	return true
}

func (l *list) lPush(k string , v string) bool {
	l.db.dataList[k] = append(l.db.dataList[k] , v)
	return true
}

func (l *list) lDel(k string ,v string) bool {
	i := indexOf(l.db.dataList[k] , v)
	l.db.dataList[k] = append(l.db.dataList[k][:i], l.db.dataList[k][i+1:]...)
	return true
}