package main

import (
	"fmt"
	"sync"
	"time"
)

type palillo struct{ sync.Mutex }

type filosofo struct {
	id                               int
	palilloIzquierdo, palilloDerecho *palillo
}

func (f filosofo) comer() {
	for j := 0; j < 3; j++ {
		f.palilloIzquierdo.Lock()
		f.palilloDerecho.Lock()

		decir("comenzando a comer: ", f.id)
		time.Sleep(time.Second)

		f.palilloDerecho.Unlock()
		f.palilloIzquierdo.Unlock()

		decir("terminando de comer: ", f.id)
		time.Sleep(time.Second)
	}
	comerEnGrupo.Done()
}

func decir(action string, id int) {
	fmt.Printf("filosofo #%d esta %s\n", id+1, action)
}

var comerEnGrupo sync.WaitGroup

func main() {
	cont := 5

	forks := make([]*palillo, cont)
	for i := 0; i < cont; i++ {
		forks[i] = new(palillo)
	}

	filosofos := make([]*filosofo, cont)
	for i := 0; i < cont; i++ {
		filosofos[i] = &filosofo{
			id: i, palilloIzquierdo: forks[i], palilloDerecho: forks[(i+1)%cont]}
		comerEnGrupo.Add(1)
		go filosofos[i].comer()

	}
	comerEnGrupo.Wait()

}
