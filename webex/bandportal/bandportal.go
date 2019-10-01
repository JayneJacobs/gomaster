package bandportal

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"fmt"

	"gotrain/goMaster/blogger"
	"gotrain/goMaster/dbex/dblayer"
	"gotrain/goMaster/webex/configurate"
	"gotrain/goMaster/webex/dbex/dblayer/vault"
	"gotrain/goMaster/webex/restapi"
	"html/template"
	"log"
	"net"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

var bandWebTemplate *template.Template

var historylog = struct {
	logs []string
	*sync.RWMutex
}{RWMutex: new(sync.RWMutex)}

// Run is the entrypoint for the web ui
func Run() error {
	var err error
	// bandWebTemplate, err = template.ParseFiles("webex/bandportal/cover/band/band.html_go", "webex/bandportal/cover/about/about.html_go", "webex/bandportal/cover/chat/chat.html_go", "webex/bandportal/cover/index.html_go")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	conf := struct {
		Filepath  string   `json:"filespath"`
		Templates []string `json:"templates"`
	}{}
	// files, err := ioutil.ReadDir("webex/..")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }
	err = configurate.GetConfiguration(configurate.JSON, &conf, "portalconfi.json")
	if err != nil {
		fmt.Println(err)
	}

	bandWebTemplate, err = template.ParseFiles(conf.Templates...) //ParseFiles is a variatic funcon
	if err != nil {
		fmt.Println(err)
		return err
	}
	restapi.InitializeAPIHandlers()
	fmt.Println(conf.Filepath)
	fs := http.FileServer(http.Dir(conf.Filepath)) //typecast a filepath string to DIR, converts to Filetyp for http.FileSErver
	http.Handle("/", fs)
	http.HandleFunc("/band/", bandhandler)
	http.HandleFunc("/about/", abouthandler)
	http.HandleFunc("/chat/", chathandler)
	http.Handle("/ChatRoom/", websocket.Handler(chatWS))
	go func() {
		err = http.ListenAndServeTLS(":8062", "cert.pem", "key.pem", nil)
		log.Println(err)
	}()
	return http.ListenAndServe(":8061", nil)
}

func bandhandler(w http.ResponseWriter, r *http.Request) {
	dblayer, err := dblayer.ConnectDatabase("mysql", "Hydra:hydraisme@/Hydra")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("In bandhandler")
	all, err := dblayer.AllMembers()
	if err != nil {
		fmt.Println(err)
	}
	// ExecuteTemplate takes iowriter path to file and  a data item to write to the file
	err = bandWebTemplate.ExecuteTemplate(w, "band.html_go", all) //the filename was already cached in ParseFiles above
	if err != nil {
		log.Panicln(err)
	}
}

func abouthandler(w http.ResponseWriter, r *http.Request) {
	about := struct {
		Msg string `json:"message"`
	}{}
	err := configurate.GetConfiguration(configurate.JSON, &about, "bandportal/apiconfig.json")
	if err != nil {
		fmt.Println(err)
	}
	err = bandWebTemplate.ExecuteTemplate(w, "about.html_go", about)
	if err != nil {
		fmt.Println(err)
	}
}

func chathandler(w http.ResponseWriter, r *http.Request) {
	nameStruct := struct{ Name string }{}
	r.ParseForm()
	if len(r.Form) == 0 {
		if cookie, err := r.Cookie("username"); err != nil {
			bandWebTemplate.ExecuteTemplate(w, "login.html_go", nil)
			return
		} else {
			nameStruct.Name = cookie.Value
			bandWebTemplate.ExecuteTemplate(w, "chat.html_go", nameStruct)
			fmt.Printf("Username Verified chatHandler1, \n Thi s is Cookie contents %v \n", cookie.Value)
		}
	}
	logger := blogger.GetInstance("cookies")
	logger.Println("Just before username verify in cookie builder")

	if r.Method == "POST" {
		// user := r.Form["username"][0]
		// pass := r.Form["password"][0]
		fmt.Println("band Portal POST in chathandler")
		var user, pass string
		if v, ok := r.Form["username"]; ok && len(v) > 0 {
			user = v[0]
			fmt.Println("verified user: ", user)
		}

		if v, ok := r.Form["password"]; ok && len(v) > 0 {
			pass = v[0]
			fmt.Println("verifying password !!!!!!: ", pass)
		}

		if !verifyPassword(user, pass) {
			bandWebTemplate.ExecuteTemplate(w, "login.html_go", nil)
			fmt.Println("Still not verified")
			return
		}
		nameStruct.Name = user
		if _, ok := r.Form["rememberme"]; ok {
			cookie := http.Cookie{Name: "username", Value: user}
			http.SetCookie(w, &cookie)
			fmt.Println("Setting cookie")
			fmt.Printf("Username is Verified chatHandler1, \n Thi s is Cookie contents %v \n", cookie.Value)
		}
	}
	bandWebTemplate.ExecuteTemplate(w, "chat.html_go", nameStruct)

}

func verifyPassword(username, pass string) bool {
	db, err := vault.ConnectPasswordVault()
	if err != nil {
		fmt.Println("error in verifyPassword")
		return false
	}

	password, err1 := vault.GetPassword(db, username)
	if pass != password {
		fmt.Printf("The password %v is not equal from GetPassword to %v \n", pass, password)
	}
	defer db.Close()
	fmt.Printf("This is the password from db %v\n", password)
	if err1 != nil {
		fmt.Println("err1 from GetPassword", err1)
		return false
	}

	data, err := vault.GetPasswordBytes(db, username)
	if err != nil {
		fmt.Println("error in Verify Get PasswordBytes", err)
		return false
	}
	defer db.Close()
	hashedPass := md5.Sum([]byte(pass)) //The result of md5
	fmt.Printf("This is the hashed pass %v, as a %T: ", hashedPass, hashedPass)
	result := password == pass || bytes.Equal(hashedPass[:], data)
	fmt.Printf("This is the resutlt of get bytes %v \n", result)

	return result

}
func chatWS(ws *websocket.Conn) {
	conn, err := net.Dial("tcp", "localhost:2100")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	historylog.RLock()
	for _, log := range historylog.logs {
		err = websocket.Message.Send(ws, log)
		if err != nil {
			historylog.Unlock()
			return
		}
	}
	historylog.RUnlock()

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			message := scanner.Text()
			err := websocket.Message.Send(ws, message)
			if err != nil {
				return
			}
		}
	}()

	for {
		var message string
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			return
		}
		_, err = conn.Write([]byte(message)) //typecast to slice of bytes
		if err == nil {
			historylog.Lock()
			if len(historylog.logs) > 20 {

				historylog.logs = historylog.logs[1:]

			}
			historylog.logs = append(historylog.logs, message)
			historylog.Unlock()
		}
	}
}
