// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/xlzd/gotp"
	log "unknwon.dev/clog/v2"

	"github.com/flamego/flamego"
	"github.com/flamego/template"

	"github.com/wuhan005/siesta/assets"
	"github.com/wuhan005/siesta/internal/process"
)

func main() {
	defer log.Stop()
	err := log.NewConsole()
	if err != nil {
		panic(err)
	}

	otpSecret := flag.String("otp", "", "")
	restrictedKeyword := flag.String("process", "minecraft", "")
	flag.Parse()

	// endTime is the time which allows to run the restricted processes.
	endTime := time.Now()
	totp := gotp.NewDefaultTOTP(*otpSecret)

	go func() {
		for {
			if time.Now().Before(endTime) {
				continue
			}
			ctx := context.Background()
			for _, kw := range strings.Split(*restrictedKeyword, ",") {
				if err := process.KillByKeyword(ctx, kw); err != nil {
					log.Error("Failed to kill by keyword: %v", err)
				}
			}
			time.Sleep(5 * time.Second)
		}
	}()

	// Web panel
	f := flamego.Classic()

	embedFS, err := template.EmbedFS(assets.FS, ".", []string{".tmpl"})
	if err != nil {
		log.Fatal("Failed to embed file system: %v", err)
	}
	f.Use(template.Templater(template.Options{FileSystem: embedFS}))

	f.Get("/", func(t template.Template, data template.Data) {
		data["EndTime"] = endTime.Format("2006-01-02 15:04:05")
		t.HTML(http.StatusOK, "index")
	})

	f.Post("/updateEndTime", func(ctx flamego.Context) {
		if err := ctx.Request().ParseForm(); err != nil {
			ctx.ResponseWriter().WriteHeader(http.StatusBadRequest)
			return
		}

		code := ctx.Request().Form.Get("code")
		if totp.Now() != code {
			ctx.ResponseWriter().WriteHeader(http.StatusForbidden)
			return
		}
		minute, err := strconv.Atoi(ctx.Request().Form.Get("minute"))
		if err != nil {
			ctx.ResponseWriter().WriteHeader(http.StatusBadRequest)
			return
		}
		endTime = endTime.Add(time.Duration(minute) * time.Minute)
		ctx.Redirect("/")
	})
	f.Run()
}
