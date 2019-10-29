// routes.go

package routes

import (
	"github.com/gin-gonic/gin"
	"octopus.com/gin-test/mhandlers"
	"octopus.com/gin-test/middleware"
)

func Register(router *gin.Engine) {

	authMiddleware := middleware.MiddlewareAuth{}

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(authMiddleware.SetUserStatus())

	// Handle the index route
	router.GET("/", mhandlers.ShowIndexPage)

	// Group user related routes together
	userRoutes := router.Group("/u")
	{
		// Handle the GET requests at /u/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/login", authMiddleware.EnsureNotLoggedIn(), mhandlers.ShowLoginPage)

		// Handle POST requests at /u/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", authMiddleware.EnsureNotLoggedIn(), mhandlers.PerformLogin)

		// Handle GET requests at /u/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", authMiddleware.EnsureLoggedIn(), mhandlers.Logout)

		// Handle the GET requests at /u/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", authMiddleware.EnsureNotLoggedIn(), mhandlers.ShowRegistrationPage)

		// Handle POST requests at /u/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", authMiddleware.EnsureNotLoggedIn(), mhandlers.Register)
	}

	// Group article related routes together
	articleRoutes := router.Group("/article")
	{
		// Handle GET requests at /article/view/some_article_id
		articleRoutes.GET("/view/:article_id", mhandlers.GetArticle)

		// Handle the GET requests at /article/create
		// Show the article creation page
		// Ensure that the user is logged in by using the middleware
		articleRoutes.GET("/create", authMiddleware.EnsureLoggedIn(), mhandlers.ShowArticleCreationPage)

		// Handle POST requests at /article/create
		// Ensure that the user is logged in by using the middleware
		articleRoutes.POST("/create", authMiddleware.EnsureLoggedIn(), mhandlers.CreateArticle)
	}
}
