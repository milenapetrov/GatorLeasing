package server

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/milenapetrov/GatorLeasing/gator-leasing-server/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/config"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/entity"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/handler"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/server/middleware"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/service"
)

type Server struct {
	config            *config.ServerConfig
	leaseHandler      *handler.LeaseHandler
	tenantUserService service.ITenantUserService
	userContext       *entity.UserContext
}

func NewServer(config *config.ServerConfig, leaseHandler *handler.LeaseHandler, tenantUserService service.ITenantUserService, userContext *entity.UserContext) *Server {
	return &Server{
		config:            config,
		leaseHandler:      leaseHandler,
		tenantUserService: tenantUserService,
		userContext:       userContext,
	}
}

func (s *Server) Run() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully waits for existing connections to finish -e.g. 15s or 1m")

	httpServer := &http.Server{
		Addr:         s.config.Address,
		Handler:      s.handler(),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	httpServer.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}

func (s *Server) handler() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.CorsMiddleware)

	s.handle(r, "/leases", "GET", s.leaseHandler.GetAllLeases, false)
	s.handle(r, "/leases", "POST", s.leaseHandler.PostLease, true)
	s.handle(r, "/leases/{id}", "PUT", s.leaseHandler.PutLease, true)
	s.handle(r, "/leases/{id}", "DELETE", s.leaseHandler.DeleteLease, true)

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(os.Getenv("HOST_URL") + "/swagger/doc.json"), //The url pointing to API definition
	)).Methods(http.MethodGet)

	r.PathPrefix("/").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			return
		})

	return r
}

func (s *Server) handle(router *mux.Router, path string, method string, f func(w http.ResponseWriter, r *http.Request), requiresAuth bool) {
	if requiresAuth {
		router.Handle(path, middleware.EnsureValidToken()(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				err := s.setUserContext(r, s.userContext)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					errorJson, _ := json.Marshal(map[string]string{"error": err.Error()})
					w.Write([]byte(errorJson))
				}
				f(w, r)
			},
		))).Methods(method)
	} else {
		router.Handle(path, http.HandlerFunc(f)).Methods(method)
	}
}

func (s *Server) setUserContext(r *http.Request, userContext *entity.UserContext) error {
	claims := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	userContext.UserID = claims.RegisteredClaims.Subject
	tenantUser, err := s.tenantUserService.GetOrCreateUser()
	if err != nil {
		return err
	}

	userContext.ID = tenantUser.ID
	userContext.InvitedAs = tenantUser.InvitedAs

	return nil
}
