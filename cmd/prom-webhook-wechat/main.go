package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"git.ucloudadmin.com/uk8s/prometheus-webhook-wechat-public/pkg/webrouter"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	if err := parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			return
		}
		log.Fatalf("Parse error: %s", err)
	}
	var config Config
    config.readConfig()

	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// When a client closes their connection midway through a request, the
	// http.CloseNotifier will cancel the request context (ctx).
	// r.Use(middleware.CloseNotify)

	WechatResource := &webrouter.WechatResource{
		Profileurl: cfg.WechatAPIUrlProfiles,
		HttpClient: &http.Client{
			Timeout: cfg.requestTimeout,
		},
		Chatids:    cfg.WechatProfiles.chatids,
		Corpid:     cfg.corpid,
		Corpsecret: cfg.corpsecret,
		TemplateID: cfg.templateid,
	}
	r.Mount("/wechat", WechatResource.Routes())
	//r.Mount(, WechatResource.Routes())

	log.Printf("Starting webserver on %s", cfg.listenAddress)
	if err := http.ListenAndServe(cfg.listenAddress, r); err != nil {
		log.Panicf("Failed to serve: %s", err)
	}
}
