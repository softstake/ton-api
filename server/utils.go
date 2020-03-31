package server

import tonlib "github.com/tonradar/tonlib-go/v2"

func isRawFullAccountState(t interface{}) bool {
	switch t.(type) {
	case tonlib.RawFullAccountState:
		return true
	default:
		return false
	}
}
