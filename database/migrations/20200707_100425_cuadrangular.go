package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Cuadrangular_20200707_100425 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Cuadrangular_20200707_100425{}
	m.Created = "20200707_100425"

	migration.Register("Cuadrangular_20200707_100425", m)
}

// Run the migrations
func (m *Cuadrangular_20200707_100425) Up() {
	m.SQL("CREATE SEQUENCE id_cuadrangular_equipo_seq START WITH 1 INCREMENT BY 1 NO MINVALUE MAXVALUE 123456789 CACHE 1;")
	m.SQL("CREATE SEQUENCE id_equipo_seq START WITH 1 INCREMENT BY 1 NO MINVALUE MAXVALUE 123456789 CACHE 1;")
	m.SQL("CREATE SEQUENCE id_partido_seq START WITH 1 INCREMENT BY 1 NO MINVALUE MAXVALUE 123456789 CACHE 1;")
	m.SQL("CREATE TABLE equipo ( id_equipo integer DEFAULT nextval('public.id_equipo_seq'::regclass) PRIMARY KEY, nombre VARCHAR NOT NULL);")
	m.SQL("CREATE TABLE partido ( id_partido integer DEFAULT nextval('public.id_partido_seq'::regclass) PRIMARY KEY, id_equipo_local integer NOT NULL, id_equipo_visitante integer NOT NULL, goles_local integer DEFAULT -1, goles_visitante integer DEFAULT -1);")
	m.SQL("CREATE TABLE cuadrangular_equipo ( id_cuadrangular_equipo integer DEFAULT nextval('public.id_cuadrangular_equipo_seq'::regclass) PRIMARY KEY, id_equipo integer NOT NULL, goles_favor integer DEFAULT 0, goles_contra integer DEFAULT 0, diferencia_gol integer DEFAULT 0, puntos integer DEFAULT 0 NOT NULL);")
	m.SQL("ALTER TABLE cuadrangular_equipo ADD CONSTRAINT id_equipo_fk FOREIGN KEY (id_equipo) REFERENCES equipo (id_equipo);")
}

// Reverse the migrations
func (m *Cuadrangular_20200707_100425) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
