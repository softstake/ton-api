package server

import (
	tonlib "github.com/mercuryoio/tonlib-go/v2"
)

type CustomTvmTuple struct {
	tonlib.TvmTuple
	Elements []tonlib.TvmStackEntryNumber
}

type CustomTvmStackEntryTuple struct {
	tonlib.TvmStackEntryTuple
	Tuple CustomTvmTuple
}

type ActiveBet struct {
	ID         tonlib.TvmStackEntryNumber
	Parameters CustomTvmStackEntryTuple
}

type CustomTvmList struct {
	tonlib.TvmList
	Elements []ActiveBet
}

type CustomTvmStackEntry struct {
	tonlib.TvmStackEntryList
	List CustomTvmList
}
