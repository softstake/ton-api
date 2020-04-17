package server

import (
	tonlib "github.com/mercuryoio/tonlib-go/v2"
)

type CustomTvmTuple struct {
	tonlib.TvmTuple
	Elements []tonlib.TvmStackEntryNumber `json:"elements"`
}

type _CustomTvmStackEntryTuple struct {
	tonlib.TvmStackEntryTuple
	Tuple CustomTvmTuple `json:"tuple"`
}

type CustomTvmStackEntryTuple struct {
	tonlib.TvmStackEntryTuple
	Tuple tonlib.TvmTuple `json:"tuple"`
}

type CustomTvmList struct {
	tonlib.TvmList
	Elements []CustomTvmStackEntryTuple `json:"elements"`
}

type CustomTvmStackEntry struct {
	tonlib.TvmStackEntryList
	List CustomTvmList `json:"list"`
}
