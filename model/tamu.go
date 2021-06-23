package model

import "time"

type Tamu struct {
	ID        int64
	Name      string
	Keperluan string
	Tanggal   *time.Time
}
