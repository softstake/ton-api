package server

import "github.com/mercuryoio/tonlib-go"

func isRawFullAccountState(t interface{}) bool {
	switch t.(type) {
	case tonlib.RawFullAccountState:
		return true
	default:
		return false
	}
}
