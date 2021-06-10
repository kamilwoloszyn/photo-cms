package routes

import "github.com/kamilwoloszyn/photo-cms/configs"

func ApplyRoutes(s *configs.Server) {
	s.Routes = s.Engine.Group("v1")
	{
		s.Routes.GET("/ping")
		s.Routes.POST("/login")
		s.Routes.POST("/registration")
		s.Routes.GET("/dashboard")
		s.Routes.POST("/dashboard/upload")
	}
}
