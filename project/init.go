package hello

import (
	"context"
	"encoding/json"
	"expvar"
	"html/template"
	"log"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"gopkg.in/tokopedia/logging.v1"
)

type ServerConfig struct {
	Name string
}

type Response struct {
	Id          int    `json:"id"`
	Name        string `json:"fullname"`
	MSISDN      string `json:"-"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string
}

type Config struct {
	Server ServerConfig
}

type HelloWorldModule struct {
	cfg       *Config
	something string
	stats     *expvar.Int
}

func NewHelloWorldModule() *HelloWorldModule {

	var cfg Config

	ok := logging.ReadModuleConfig(&cfg, "config", "hello") || logging.ReadModuleConfig(&cfg, "files/etc/gosample", "hello")
	if !ok {
		// when the app is run with -e switch, this message will automatically be redirected to the log file specified
		log.Fatalln("failed to read config")
	}

	// this message only shows up if app is run with -debug option, so its great for debugging
	logging.Debug.Println("hello init called", cfg.Server.Name)

	return &HelloWorldModule{
		cfg:       &cfg,
		something: "John Doe",
		stats:     expvar.NewInt("rpsStats"),
	}

}

func (hlm *HelloWorldModule) SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), r.URL.Path)
	defer span.Finish()

	hlm.stats.Add(1)
	hlm.someSlowFuncWeWantToTrace(ctx, w)
}

func (hlm *HelloWorldModule) SayHellotemplate(w http.ResponseWriter, r *http.Request) {
	// span, ctx := opentracing.StartSpanFromContext(r.Context(), r.URL.Path)
	// defer span.Finish()

	hlm.stats.Add(1)
	t, err := template.ParseFiles("hello/view.html")
	if err != nil {
		log.Println(err)
		return
	}

	resp := Response{
		Id:     1,
		Name:   "James",
		MSISDN: "621412312312",
		Email:  "okta.stradllin@gmail.com",
	}

	t.Execute(w, resp)
	// hlm.someSlowFuncWeWantToTrace(ctx, w)
}

func (hlm *HelloWorldModule) someSlowFuncWeWantToTrace(ctx context.Context, w http.ResponseWriter) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "someSlowFuncWeWantToTrace")
	defer span.Finish()

	resp := Response{
		Id:     1,
		Name:   "James",
		MSISDN: "621412312312",
		Email:  "okta.stradllin@gmail.com",
	}

	jsonData, err := json.Marshal(resp)

	if err != nil {
		return
	}

	w.Write([]byte(jsonData))
	// fmt.Println(jsonData)
}
