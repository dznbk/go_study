package main

type Gundam int

const (
	RX78 Gundam = iota
	zeta
	zz
)

//go:generate stringer -type=Gundam
