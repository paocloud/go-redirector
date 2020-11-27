package main

import (
    "log"
    "net/http"
    "strings"
)

func redirect(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Server", "PAOCLOUD GO Redirector/v1")
    if strings.Compare(r.Host, "paocloud.tech")  == 0 {
	    http.Redirect(w, r, "https://www.payungsakpk.xyz", 302)
    } else if strings.Compare(r.Host, "thepao.cloud")  == 0 {
           http.Redirect(w, r, "https://www.paocloud.biz", 302)
    } else {
	    http.Redirect(w, r, "https://www.paocloud.co.th", 302)
    }
}

func main() {
    http.HandleFunc("/", redirect)
    err := http.ListenAndServe(":9999", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
