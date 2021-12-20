package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ConsultarDisponibilidadesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ConsultarDisponibilidadesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ConsultarDisponibilidadesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ConsultarDisponibilidadesController"],
        beego.ControllerComments{
            Method: "ListarDisponibilidades",
            Router: "/ListaDisponibilidades/:vigencia",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ConsultarDisponibilidadesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ConsultarDisponibilidadesController"],
        beego.ControllerComments{
            Method: "TotalDisponibilidades",
            Router: "/TotalDisponibilidades/:vigencia",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ExpedirResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ExpedirResolucionController"],
        beego.ControllerComments{
            Method: "Cancelar",
            Router: "/cancelar",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ExpedirResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ExpedirResolucionController"],
        beego.ControllerComments{
            Method: "Expedir",
            Router: "/expedir",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ExpedirResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ExpedirResolucionController"],
        beego.ControllerComments{
            Method: "ExpedirModificacion",
            Router: "/expedirModificacion",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ExpedirResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:ExpedirResolucionController"],
        beego.ControllerComments{
            Method: "ValidarDatosExpedicion",
            Router: "/validar_datos_expedicion",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDesvinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDesvinculacionesController"],
        beego.ControllerComments{
            Method: "ActualizarVinculacionesCancelacion",
            Router: "/actualizar_vinculaciones_cancelacion",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDesvinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDesvinculacionesController"],
        beego.ControllerComments{
            Method: "AdicionarHoras",
            Router: "/adicionar_horas",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDesvinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDesvinculacionesController"],
        beego.ControllerComments{
            Method: "AnularAdicionDocente",
            Router: "/anular_adicion",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDesvinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDesvinculacionesController"],
        beego.ControllerComments{
            Method: "AnularModificaciones",
            Router: "/anular_modificaciones",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDesvinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDesvinculacionesController"],
        beego.ControllerComments{
            Method: "ConsultarCategoria",
            Router: "/consultar_categoria",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDesvinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDesvinculacionesController"],
        beego.ControllerComments{
            Method: "ListarDocentesCancelados",
            Router: "/docentes_cancelados",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDesvinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDesvinculacionesController"],
        beego.ControllerComments{
            Method: "ListarDocentesDesvinculados",
            Router: "/docentes_desvinculados",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDesvinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDesvinculacionesController"],
        beego.ControllerComments{
            Method: "ValidarSaldoCDP",
            Router: "/validar_saldo_cdp",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDocumentoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionDocumentoResolucionController"],
        beego.ControllerComments{
            Method: "GetContenidoResolucion",
            Router: "/get_contenido_resolucion",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionPrevinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionPrevinculacionesController"],
        beego.ControllerComments{
            Method: "CalcularTotalSalarios",
            Router: "/Precontratacion/calcular_valor_contratos",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionPrevinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionPrevinculacionesController"],
        beego.ControllerComments{
            Method: "Calcular_total_de_salarios_seleccionados",
            Router: "/Precontratacion/calcular_valor_contratos_seleccionados",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionPrevinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionPrevinculacionesController"],
        beego.ControllerComments{
            Method: "ListarDocentesCargaHoraria",
            Router: "/Precontratacion/docentes_x_carga_horaria",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionPrevinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionPrevinculacionesController"],
        beego.ControllerComments{
            Method: "InsertarPrevinculaciones",
            Router: "/Precontratacion/insertar_previnculaciones",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionPrevinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionPrevinculacionesController"],
        beego.ControllerComments{
            Method: "ListarDocentesPrevinculados",
            Router: "/docentes_previnculados",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionPrevinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionPrevinculacionesController"],
        beego.ControllerComments{
            Method: "ListarDocentesPrevinculadosAll",
            Router: "/docentes_previnculados_all",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionPrevinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionPrevinculacionesController"],
        beego.ControllerComments{
            Method: "GetCdpRpDocente",
            Router: "/rp_docente/:num_vinculacion/:vigencia/:identificacion",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionResolucionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionResolucionesController"],
        beego.ControllerComments{
            Method: "GetResolucionesPorDocente",
            Router: "/consulta_docente/:persona_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionResolucionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionResolucionesController"],
        beego.ControllerComments{
            Method: "GetResolucionesAprobadas",
            Router: "/get_resoluciones_aprobadas",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionResolucionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionResolucionesController"],
        beego.ControllerComments{
            Method: "GetResolucionesInscritas",
            Router: "/get_resoluciones_inscritas",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionResolucionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_docentes_mid/controllers:GestionResolucionesController"],
        beego.ControllerComments{
            Method: "InsertarResolucionCompleta",
            Router: "/insertar_resolucion_completa",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
