package conta

import (
	"math/rand"
	"time"
)

func (u *Banco) GerarID() int{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := r.Intn(int(9999) + 1)

	if id < 9000{
		id += 1000
	}

	return id
}