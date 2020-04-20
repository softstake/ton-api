package server

import tonlib "github.com/tonradar/tonlib-go/v2"

func isRawFullAccountState(t interface{}) bool {
	switch v := t.(type) {
	case *tonlib.RawFullAccountState:
		if v.Type != "raw.fullAccountState" {
			return false
		}
		return true
	default:
		return false
	}
}
