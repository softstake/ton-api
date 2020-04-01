package server

import tonlib "github.com/mercuryoio/tonlib-go/v2"

func isRawFullAccountState(t interface{}) bool {
	switch t.(type) {
	case *tonlib.RawFullAccountState:
		return true
	default:
		return false
	}
}
