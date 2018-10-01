package router

import (
	"SKSISAPI/method"
	"net/http"
)


func Routing()  {
	http.HandleFunc("/sksoil/",method.SksoilHandler)
	http.HandleFunc("/sksoil/_schema/",method.SkSchemaHandler)
	http.HandleFunc("/sksoil/_table/",method.SkTableHandler)
	http.HandleFunc("/sksoil/_proc/",method.SkProcHandler)
	http.HandleFunc("/sksoil/_func/",method.SkFuncHandler)
}
