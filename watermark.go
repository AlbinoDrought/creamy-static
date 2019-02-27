package main

import (
	"fmt"
	"net/http"
)

type watermarkedFileServer struct {
	http.Handler
	watermark string
}

func (f watermarkedFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Handler.ServeHTTP(w, r)

	if w.Header().Get("Content-Type") == "text/html; charset=utf-8" {
		w.Write([]byte(fmt.Sprintf(`
		<div class="rainbow" style="position: absolute; bottom: 0px; right: 0px; opacity: 0.7; padding: 10px;">
			%s
		</div>
		`, f.watermark) + css))
	}
}

const css = `<style type="text/css">@-webkit-keyframes developing{0%,100%{-webkit-filter:hue-rotate(0) drop-shadow(0 0 1px pink)}50%{-webkit-filter:hue-rotate(360deg) drop-shadow(0 0 4px yellow)}}.rainbow{-webkit-animation:developing 10s linear infinite;color:red}</style>`
