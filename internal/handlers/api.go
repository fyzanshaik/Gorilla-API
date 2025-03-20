package handlers

import (
	"github.com/fyzanshaik/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
) 

func Handler(r *chi.Mux){
	//Setting up my middleware

	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func (router chi.Router) {

		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
		} )
}