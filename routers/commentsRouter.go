package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["major_league_api_example/controllers:UserController"] = append(beego.GlobalControllerRouter["major_league_api_example/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["major_league_api_example/controllers:UserController"] = append(beego.GlobalControllerRouter["major_league_api_example/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["major_league_api_example/controllers:UserController"] = append(beego.GlobalControllerRouter["major_league_api_example/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["major_league_api_example/controllers:UserController"] = append(beego.GlobalControllerRouter["major_league_api_example/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["major_league_api_example/controllers:UserController"] = append(beego.GlobalControllerRouter["major_league_api_example/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["major_league_api_example/controllers:UserController"] = append(beego.GlobalControllerRouter["major_league_api_example/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["major_league_api_example/controllers:UserController"] = append(beego.GlobalControllerRouter["major_league_api_example/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["major_league_api_example/controllers:WebSocketController"] = append(beego.GlobalControllerRouter["major_league_api_example/controllers:WebSocketController"],
        beego.ControllerComments{
            Method: "ViewChat",
            Router: `/chat`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["major_league_api_example/controllers:WebSocketController"] = append(beego.GlobalControllerRouter["major_league_api_example/controllers:WebSocketController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/ws`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
