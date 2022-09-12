package main

import (
	"log"
	"logicful/features/contentblock"
	"logicful/features/email"
	"logicful/features/folder"
	"logicful/features/form"
	"logicful/features/formsubmission"
	"logicful/features/optionset"
	"logicful/features/s3store"
	"logicful/features/team"
	"logicful/features/user"
	"logicful/features/workflow"
	"logicful/handler"
	"logicful/service/db"
	"logicful/service/queue"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {

	db.Instance()
	queue.Setup()
	addCron()

	router := httprouter.New()

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
			header.Set("Access-Control-Allow-Headers", "*")
		}
	})

	addOptionSetHandlers(router)
	addContentBlockHandlers(router)
	addFormHandlers(router)
	addFolderHandlers(router)
	addSubmissionHandlers(router)
	addSetS3Handlers(router)
	addUserRoutes(router)
	addEmailWebhookRoutes(router)
	addTeamRoutes(router)

	router.GET("/robots.txt", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		http.ServeFile(w, r, "public/robots.txt")
	})

	router.ServeFiles("/css/*filepath", http.Dir("public/css"))
	router.ServeFiles("/js/*filepath", http.Dir("public/js"))
	router.ServeFiles("/img/*filepath", http.Dir("public/img"))
	router.ServeFiles("/_dist_/*filepath", http.Dir("public/_dist_"))
	router.ServeFiles("/__snowpack__/*filepath", http.Dir("public/__snowpack__"))
	router.ServeFiles("/web_modules/*filepath", http.Dir("public/web_modules"))
	router.ServeFiles("/assets/*filepath", http.Dir("public/assets"))
	router.ServeFiles("/fonts/*filepath", http.Dir("public/fonts"))
	router.NotFound = handler.FileServer(http.Dir("/public"))

	addQueueSubscribers()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + port,
		Handler:      router,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func addOptionSetHandlers(router *httprouter.Router) {
	router.GET("/api/option-set", optionset.ListHandler)
	router.POST("/api/option-set", optionset.SetHandler)
	router.PUT("/api/option-set/:id", optionset.SetHandler)
}

func addContentBlockHandlers(router *httprouter.Router) {
	router.GET("/api/content-block", contentblock.ListHandler)
	router.POST("/api/content-block", contentblock.SetHandler)
	router.PUT("/api/content-block/:id", contentblock.SetHandler)
}

func addFormHandlers(router *httprouter.Router) {
	router.GET("/api/form", form.ListHandler)
	router.GET("/api/form/:formId", form.GetHandler)
	router.POST("/api/form", form.SetHandler)
	router.PUT("/api/form/:formId", form.SetHandler)
}

func addFolderHandlers(router *httprouter.Router) {
	router.GET("/api/folder", folder.ListHandler)
	router.POST("/api/folder", folder.SetHandler)
	router.PUT("/api/folder/:folderId", folder.SetHandler)
	router.DELETE("/api/folder/:folderId", folder.DeleteHandler)
}

func addSubmissionHandlers(router *httprouter.Router) {
	router.GET("/api/form/:formId/submission", formsubmission.ListHandler)
	router.POST("/api/form/:formId/submission/file/:id", formsubmission.GenerateFileUrlHandler)
	router.POST("/api/form/:formId/submission", formsubmission.AddHandler)
	router.DELETE("/api/form/:formId/submission", formsubmission.DeleteHandler)
	router.PUT("/api/form/:formId/submission/mark/read", formsubmission.MarkReadHandler)
	router.GET("/api/form/:formId/submission/read", formsubmission.GetReadHandler)
}

func addSetS3Handlers(router *httprouter.Router) {
	router.POST("/api/s3/json", s3store.SetJsonHandler)
	router.PUT("/api/s3/json", s3store.SetJsonHandler)
	router.POST("/api/s3/put", s3store.GenerateUrlHandler)
}

func addUserRoutes(router *httprouter.Router) {
	router.POST("/api/user/register", user.RegisterHandler)
	router.POST("/api/user/login", user.LoginHandler)
	router.POST("/api/user/refresh", user.RefreshHandler)
	router.POST("/api/user/login/google", user.LoginFromGoogleHandler)
}

func addTeamRoutes(router *httprouter.Router) {
	router.GET("/api/team", team.TeamHandler)
	router.GET("/api/team/members", team.MembersHandler)
	router.PUT("/api/team", team.SetHandler)
	router.POST("/api/team", team.SetHandler)
	router.DELETE("/api/team/members/remove", team.RemoveMembersHandler)
	router.POST("/api/team/members/invite", team.InviteMembersHandler)
	router.POST("/api/team/members/invite/accept", team.AcceptInviteHandler)
}

func addEmailWebhookRoutes(router *httprouter.Router) {
	router.POST("/api/email/webhook", email.WebhookHandler)
}

func addCron() {
	formsubmission.StartProcessor()
}

func addQueueSubscribers() {
	workflow.Register()
	team.Register()
}
