package server

import (
	tonlib "github.com/tonradar/tonlib-go/v2"
)

type CustomTvmStackEntryNumber struct {
	tonlib.TvmStackEntryNumber
	Number tonlib.TvmNumberDecimal `json:"number"`
}

type CustomTvmTuple struct {
	tonlib.TvmTuple
	Elements []CustomTvmStackEntryNumber `json:"elements"`
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
