package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:CuadrangularEquipoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:CuadrangularEquipoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:CuadrangularEquipoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:CuadrangularEquipoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:CuadrangularEquipoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:CuadrangularEquipoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:CuadrangularEquipoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:CuadrangularEquipoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:CuadrangularEquipoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:CuadrangularEquipoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:EquipoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:EquipoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:EquipoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:EquipoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:EquipoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:EquipoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:EquipoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:EquipoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:EquipoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:EquipoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:PartidoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:PartidoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:PartidoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:PartidoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:PartidoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:PartidoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:PartidoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:PartidoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:PartidoController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:PartidoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:PartidosController"] = append(beego.GlobalControllerRouter["github.com/diagutierrezro/Cuadrangular/controllers:PartidosController"],
		beego.ControllerComments{
			Method:           "GenerarPartidos",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
