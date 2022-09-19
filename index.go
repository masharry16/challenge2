package main

import (
    "html/template"
    "net/http"
	"fmt"	
	"strings"
)

type student struct {
	id        int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
	username  string
	password  string
	errorMsg  string
	state	  int
}


func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        var data = map[string]string{
            "Name":    "john wick",
            "Message": "have a nice day",
        }

        var t, err = template.ParseFiles("login.html")

        if err != nil {
            fmt.Println(err.Error())
            return
        }

        t.Execute(w, data)

    })

	http.HandleFunc("/login", func(rw http.ResponseWriter, req *http.Request) {

		var username 	= strings.ToLower(req.FormValue("username"))
		var password    = req.FormValue("password")

		var uname		= []string{"Hari", "Thomas", "Brian", "Fitri"}
		var pswd   		= []string{"Hari123", "Thomas123","Brian123", "Fitri12345"}
		var nama 		= []string{"Hari", "Thomas", "Brian", "Fitri"}
		var alamat 		= []string{"Jalan Tentara Pelajar No 1", "Jalan Tentara Pelajar No 2", "Jalan Tentara Pelajar No 3","Jalan Tentara Pelajar No 4"}
		var pekerjaan 	= []string{"Programmer", "Programmer", "Programmer", "Programmer"}
		var alasan 		= []string{"Make money", "Better life future", "Better life future", "Better life future"}
	
		var std student
		var data = map[string]string{
			"nama":    		"",
            "alamat": 		"",
			"pekerjaan":    "",
            "alasan": 		"",
			"state" :		"",
			"msg" :			"",
			"uname" :		"",
		}

		var found int = 0
	
		for i := 0; i < len(username); i++ {
	
			if strings.ToLower(uname[i]) == username {

				std.id = i
				std.nama = nama[i]
				std.alamat = alamat[i]
				std.pekerjaan = pekerjaan[i]
				std.alasan = alasan[i]
				std.username = uname[i]
				std.password = pswd[i]

				data["nama"] 		= nama[i]
				data["alamat"] 		= alamat[i]
				data["pekerjaan"] 	= pekerjaan[i]
				data["alasan"] 		= alasan[i]
				data["uname"] 		= uname[i]

				found = 1
			}
	
		}

		if found == 1 {
			if std.password != password {
				data["msg"] 	= "Password Salah"
				data["state"]   = "0"
			}else{
				data["state"]   = "1"
			}
		} else {
			data["msg"] 	= "User Tidak ditemukan"
			data["state"]   = "0"		
		}

		fmt.Println(std.errorMsg)

		if data["state"] == "0" {
			var t, err = template.ParseFiles("error.html")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			t.Execute(rw, data)
		}else{
			var t, err = template.ParseFiles("response.html")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			t.Execute(rw, data)
		}

    })

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}

