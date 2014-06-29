package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ehazlett/phoenix"
	"github.com/ehazlett/phoenix/plugins"
	"github.com/gorilla/mux"
)

type (
	Server struct {
		PluginManager *plugins.Manager
		ListenAddr    string
	}
)

func NewServer(manager *plugins.Manager, listenAddr string) *Server {
	return &Server{
		PluginManager: manager,
		ListenAddr:    listenAddr,
	}
}

func (server *Server) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", server.indexHandler).Methods("POST")
	r.HandleFunc("/info", server.infoHandler)
	http.Handle("/", r)
	logger.Infof("listening on %s", server.ListenAddr)
	http.ListenAndServe(server.ListenAddr, nil)
}

func (server *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	// create Message from slack outgoing webhook post
	r.ParseForm()
	token := r.FormValue("token")
	teamId := r.FormValue("team_id")
	channelId := r.FormValue("channel_id")
	channelName := r.FormValue("channel_name")
	userId := r.FormValue("user_id")
	username := r.FormValue("user_name")
	text := r.FormValue("text")
	triggerWord := r.FormValue("trigger_word")
	t := r.FormValue("timestamp")
	timestamp, err := strconv.ParseFloat(t, 64)
	if err != nil {
		msg := fmt.Sprintf("unable to parse timestamp: %s", err)
		logger.Errorf(msg)
		w.WriteHeader(500)
		w.Write([]byte(msg))
		return
	}
	message := &phoenix.Message{
		Token:       token,
		TeamId:      teamId,
		ChannelId:   channelId,
		ChannelName: channelName,
		Timestamp:   time.Unix(int64(timestamp), 0),
		UserId:      userId,
		Username:    username,
		Text:        text,
		TriggerWord: triggerWord,
	}
	logger.Info(message)
	respText := server.PluginManager.Handle(message)
	resp := phoenix.Response{
		Text: respText,
	}
	b, err := json.Marshal(resp)
	if err != nil {
		msg := fmt.Sprintf("error marshaling json: %s", err)
		logger.Errorf(msg)
		w.WriteHeader(500)
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(200)
	w.Write(b)
}

func (server *Server) infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "phoenix %s", version)
}
