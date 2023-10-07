package main
import (
    "net/http"
    "flag"
    "github.com/labstack/echo"
    "commom/util"
)


func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))

func main() {
    e := echo.New()
    conf := flag.String("config", "./conf/config.toml", "rpc config file")
    flag.Parse()

    var c common.Config

    config, err := common.util.NewConfigToml(conf)
    if err != nil {
        e.Logger.Fatal("parse toml config error:", err)
	os.Exit(0)
    }

	/*server, err := as.InitServer(config)
	if err != nil {
		log.Error("init server error,errf:", err)
		os.Exit(1)
	}

	err = common.InitConfig(c)
	if err != nil {
		log.Error("init buss config error,err:", err)
		os.Exit(1)
	}*/

    e.Logger.Info("start server,port:", config.Port())
    // 监听服务
    err = server.Serve(config.Port())
    if err != nil {
        e.Logger.Fatal("stop server :%s", err)
    }
}
