//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=../../../api/proto pack_calculator.proto
package pack_calculator

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Yapanyushin/tabeo-challenge/api/proto"
)

const (
	pathIndexHTML = "assets/index.html"
)

type PageData struct {
	OrderQuantity int32
	Packs         []*proto.PacksAmount
	Error         string
}

type HttpServer struct {
	apiHost string
	Server  *http.Server
}

func NewServer(apiHost, port string) *HttpServer {
	mux := http.NewServeMux()
	srv := &HttpServer{
		apiHost: apiHost,
	}

	mux.HandleFunc("/", srv.homeHandler)
	mux.HandleFunc("/calculate", srv.calculateHandler)

	srv.Server = &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	return srv
}

func (s HttpServer) homeHandler(w http.ResponseWriter, _ *http.Request) {
	if err := template.Must(template.ParseFiles(pathIndexHTML)).Execute(w, PageData{}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		if _, err = fmt.Fprintln(w, "Ooops!"); err != nil {
			log.Printf("can't write error : %s", err.Error())
		}
	}
}

func (s HttpServer) calculateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(pathIndexHTML)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := fmt.Fprintln(w, "Ooops!"); err != nil {
			log.Printf("can't write error : %s", err.Error())
		}
	}
	defer func() {
		if err != nil {
			if err := tmpl.Execute(w, PageData{Error: err.Error()}); err != nil {
				if _, err := fmt.Fprintln(w, "Ooops!"); err != nil {
					log.Printf("can't write error : %s", err.Error())
				}
			}
		}
	}()

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	orderQuantityStr := r.FormValue("orderQuantity")
	orderQuantity, err := strconv.Atoi(orderQuantityStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set up a connection to the server
	conn, err := grpc.NewClient(s.apiHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := proto.NewPackCalculatorClient(conn).CalculatePack(ctx, &proto.CalculatePacksAmountRequest{Items: int32(orderQuantity)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := PageData{OrderQuantity: int32(orderQuantity), Packs: resp.Packs}
	if err := tmpl.Execute(w, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
