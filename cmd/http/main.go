package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	pbRepo "github.com/sapawarga/api-orchestration/repository/phonebook"
	pbUsecase "github.com/sapawarga/api-orchestration/usecase/phonebook"

	transportPb "github.com/sapawarga/api-orchestration/transport/http/phonebook"

	kitlog "github.com/go-kit/kit/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	viper.SetConfigFile(`./config.json`)
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Cannot find config file, %s", err)
	}
}

func main() {
	httpAddress := flag.String("http.addr", viper.GetString(`API_HTTP_ADDRESS`), `HTTP listen address`)

	// setting grpc connection
	phonebookHost := viper.GetString(`PHONEBOOK_HOST`)
	phonebookPort := viper.GetString(`PHONEBOOK_PORT`)
	phonebookConn, err := grpc.Dial(fmt.Sprintf("%s:%s", phonebookHost, phonebookPort), grpc.WithInsecure())
	checkError("connection_to_phonebook", err)

	// setting repository
	phonebookRepo := pbRepo.NewProvider(phonebookConn)

	// setting usecase
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)
	phonebookUC := pbUsecase.NewUsecase(phonebookRepo, logger)

	errChan := make(chan error)

	// initiating transport http
	go func() {
		log.Println("[HTTP][Info] Starting API")
		logger.Log("transport", "http", "address", *httpAddress, "msg", "listening")
		mux := http.NewServeMux()
		ctx := context.Background()
		mux.Handle("/api/phonebook/", transportPb.MakeHandler(ctx, phonebookUC, logger))
		errChan <- http.ListenAndServe(*httpAddress, accessControl(mux))
	}()

	err = <-errChan
	checkError("initiating transport http", err)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Access-Control-Allow-Origin", "*")
		r.Header.Set("Access-Control-Allow-Methods", "GET, PUT, PATCH, POST, OPTIONS")
		r.Header.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin, Access-Control-Allow-Headers, scope, state, hd, code")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, PATCH, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin, Access-Control-Allow-Headers, scope, state, hd, code")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func checkError(section string, err error) {
	if err != nil {
		log.Printf("%s:%v", section, err)
		panic(err)
	}
}
