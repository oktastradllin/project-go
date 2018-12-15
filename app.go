package main

import (
	"belajar/project-tokopedia/config"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/vwa/util/render"
)

type user struct {
	User_id     sql.NullString `json:"id"`
	Full_name   sql.NullString `json:"name"`
	msisdn      sql.NullString
	userEmail   sql.NullString
	birth_date  sql.NullString
	create_time sql.NullString
	update_time sql.NullString
}

type users struct {
	data []user
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := make(map[string]interface{})
	data["title"] = "Home"
	render.HTMLRender(w, r, "template.index", data)
}

func main() {

	cfg := &config.MainConfig{}
	data := config.ReadConfig(cfg, "main")

	fmt.Println(data)

	// mw := middleware.New()
	// router := httprouter.New()

	// router.ServeFiles("/assets/*filepath", http.Dir("assets/"))
	// router.GET("/", mw.LoggingMiddleware(indexHandler))
	// router.GET("/index", mw.LoggingMiddleware(indexHandler))

	// user := user.New()
	// user.SetRouter(router)

	// komentar := komentar.New()
	// komentar.SetRouter(router)

	// profile := profile.New()
	// profile.SetRouter(router)

	// s := http.Server{
	// 	Addr:    ":" + util.Cfg.Webport,
	// 	Handler: router,
	// }

	// fmt.Printf("Server running at port %s\n", s.Addr)
	// fmt.Printf("Open this url %s on your browser to access VWA", util.Fullurl)
	// s.ListenAndServe()

	// rows, err := db.Query("SELECT user_id,full_name,msisdn,user_email,birth_date,create_time,update_time FROM ws_user limit 10")
	// checkErr(err)
	// allUser := users{}
	// for rows.Next() {
	// 	u := user{}
	// 	err = rows.Scan(
	// 		&u.User_id,
	// 		&u.Full_name,
	// 		&u.msisdn,
	// 		&u.userEmail,
	// 		&u.birth_date,
	// 		&u.create_time,
	// 		&u.update_time,
	// 	)

	// 	// if err != nil {
	// 	// 	fmt.Println(err)
	// 	// }
	// 	allUser.data = append(allUser.data, u)
	// 	checkErr(err)
	// }

	// t, err := template.ParseFiles("hello/view.html")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// t.Execute(w, t)

	// fmt.Println(allUser)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
