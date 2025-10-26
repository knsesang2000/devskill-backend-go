package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes wires MVP routes to the router group
func RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/problems", getProblems)
	api.POST("/submissions", postSubmission)
}

func getProblems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"items": []gin.H{
			{
				"id":          "runtime-server-hello",
				"title":       "Run a server using Node.js / Deno / Bun",
				"description": "Create a minimal HTTP server that responds with 'ok' on GET /health. You may use Node.js, Deno, or Bun.",
				"runtimeOptions": []string{"node", "deno", "bun"},
				"accepts":        []string{"js", "ts"},
				"example": gin.H{
					"node": "const http = require('http'); http.createServer((req,res)=>{ if(req.url==='\\/health'){res.writeHead(200,{'Content-Type':'application/json'}); res.end(JSON.stringify({status:'ok'}));} else {res.writeHead(404); res.end();}}).listen(3000);",
					"deno": "import { serve } from 'https://deno.land/std@0.224.0/http/server.ts'; serve((req)=> req.url.endsWith('/health') ? new Response(JSON.stringify({status:'ok'}), {headers:{'content-type':'application/json'}}): new Response('',{status:404}));",
					"bun":  "import { serve } from 'bun'; serve({ port: 3000, fetch(req){ if(new URL(req.url).pathname === '/health'){ return new Response(JSON.stringify({status:'ok'}), {headers:{'content-type':'application/json'}});} return new Response('',{status:404}); } });",
				},
			},
		},
	})
}

type submissionReq struct {
	ProblemID string `json:"problemId"`
	Language  string `json:"language"`
	Runtime   string `json:"runtime"`
	Code      string `json:"code"`
}

func postSubmission(c *gin.Context) {
	var req submissionReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// MVP simulated evaluation: minimal pattern checks
	passed := false
	if req.Runtime == "node" && (contains(req.Code, "createServer") || contains(req.Code, "http.createServer")) {
		passed = true
	}
	if req.Runtime == "deno" && contains(req.Code, "serve(") {
		passed = true
	}
	if req.Runtime == "bun" && contains(req.Code, "serve(") {
		passed = true
	}

	status := "failed"
	score := 0
	if passed {
		status = "passed"
		score = 100
	}

	c.JSON(http.StatusOK, gin.H{
		"submissionId": time.Now().UnixNano(),
		"problemId":    req.ProblemID,
		"status":       status,
		"score":        score,
		"notes":        "MVP stub evaluation; runtime execution will be added later.",
	})
}
