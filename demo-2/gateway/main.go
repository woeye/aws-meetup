package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/woeye/aws-demo/gateway/calc_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

type inputRequest struct {
	A int32 `json:"a"`
	B int32 `json:"b"`
}

func main() {
	// address := "dns:///calc-service"
	// address := "localhost:8000"
	address, found := os.LookupEnv("CALC_SERVICE_ADDRESS")
	if !found {
		address = "localhost:8000"
	}

	opts := []grpc.DialOption{
		// grpc.WithPerRPCCredentials(&tokenCreds{Token: "secret-token"}),
		grpc.WithBalancerName(roundrobin.Name),
		grpc.WithInsecure(),
	}
	opts = append(opts, grpc.WithTimeout(time.Second*10))

	log.Printf("Connecting to server: %s\n", address)
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("Could not connect to server: %v\n", err)
	}
	defer conn.Close()
	log.Printf("Connected!")

	calcClient := calc_service.NewCalcServiceClient(conn)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/calc", func(w http.ResponseWriter, r *http.Request) {
		var input inputRequest
		json.NewDecoder(r.Body).Decode(&input)

		rpcCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		resp, err := calcClient.Multiply(rpcCtx, &calc_service.MultiplyRequest{A: input.A, B: input.B})
		if err != nil {
			log.Printf("CalcService call failed: %+v\n", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json, _ := json.Marshal(resp)
		w.Write(json)
	})

	err = http.ListenAndServe(":5000", router)
	if err != nil {
		log.Fatalf("Could not start HTTP server: %v\n", err)
	}
}
