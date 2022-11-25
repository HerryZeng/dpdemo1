package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/HerryZeng/dpdemo1/MyLog"
	"github.com/HerryZeng/dpdemo1/config"
	"github.com/HerryZeng/dpdemo1/handler"
	"github.com/HerryZeng/dpdemo1/middleware"
	"github.com/HerryZeng/dpdemo1/model"
	example "github.com/arl/statsviz/_example"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	cfg                  = flag.String("config", "", "")
	DiscountHandler      handler.DiscountHandler
	RestaurantNavHandler handler.RestaurantNavHandler
	HotelDetailHandler   handler.HotelDetailHandler
	TeamDetailHandler    handler.TeamDetailHandler
	OrderSeatHandler     handler.OrderSeatHandler
	TakeOutHandler       handler.TakeOutHandler
	MeHandler            handler.MeHandler
	SuggestFoodHandler   handler.SuggestFoodHandler
	SuggestHandler       handler.SuggestHandler
	GuessHandler         handler.GuessHandler
	NavHandler           handler.NavHandler
	PostTeamOrderHandler handler.PostTeamOrderHandler
)

func init() {
	initViper()
	initDB()
	initHandler()
}

func main() {
	go example.Work()
	fmt.Println("Point your browser to http://localhost:9090/debug/statsviz/")
	flag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	model.DB.Init()
	defer model.DB.Close()
	r := gin.New()
	Load(
		r,
		middleware.ProcessTraceID(),
		middleware.Logging(),
	)
	port := viper.GetString("addr")
	MyLog.Log.Info("开始监听http地址", port)
	// log.Printf(http.ListenAndServe(port, r).Error())
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf(err.Error())
	}
}
