package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Partidos struct {
	Partidos []Partido
}

type Equipos struct {
	Equipos []Equipo
}

func ObtenerEquipos() (err error) {
	var Equipos_Obtenidos []Equipo
	var Partidos_Obtenidos []Partido

	o := orm.NewOrm()
	qe, err := o.QueryTable(new(Equipo)).RelatedSel().All(&Equipos_Obtenidos)
	fmt.Println("PRIMERA PRUEBA")
	fmt.Println(qe)
	fmt.Println(Equipos_Obtenidos)

	fmt.Println("EQUIPOS!!!")
	for _, s := range Equipos_Obtenidos {
		fmt.Println(s.Id)
		fmt.Println(s.Nombre)
	}
	fmt.Println("Inicio")
	qp, err := o.QueryTable(new(Partido)).RelatedSel().All(&Partidos_Obtenidos)
	fmt.Println("Final")
	if qp == 0 {
		fmt.Println("MANDo a imprimir")
		GenerarPartidos(Equipos_Obtenidos)
	}

	return
}

func GenerarPartidos(Equipos_Obtenidos []Equipo) {
	fmt.Println("LLegó a imprimir")
	var Partido Partido
	for _, i := range Equipos_Obtenidos {
		for _, j := range Equipos_Obtenidos {
			if i.Id != j.Id {
				Partido.Id = 0
				Partido.IdEquipoLocal = i.Id
				Partido.IdEquipoVisitante = j.Id
				Partido.GolesLocal = -1
				Partido.GolesVisitante = -1
				fmt.Println("Creo partido")
				AddPartido(&Partido)
			}
		}
	}
	return
}
