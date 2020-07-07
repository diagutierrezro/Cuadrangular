package controllers

import (
	"github.com/diagutierrezro/Cuadrangular/models"
	"fmt"

	"github.com/astaxie/beego"
)

// PartidosController operations for Equipo
type PartidosController struct {
	beego.Controller
}

// URLMapping ...
func (c *PartidosController) URLMapping() {
	c.Mapping("GenerarPartidos", c.GenerarPartidos)
}

// GenerarPartidos ...
// @Title Generar Partidos
// @Description generar partidos
// @Param	body		body 	models.RegistrarPartidos	true		"body for RegistrarPartido content"
// @Success 200 {object} models.RegistrarPartidos
// @Failure 403
// @router / [get]
func (c *PartidosController) GenerarPartidos() {
	fmt.Println("Entró a generar partidos")
	models.ObtenerEquipos()
}
