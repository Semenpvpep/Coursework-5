package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	pb "cv_service/api/proto" // Замените на путь к вашему сгенерированному пакету
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

const (
	grpcServerAddress = "localhost:50051" // Адрес вашего gRPC сервиса
)

type Server struct {
	client pb.RecruitmentServiceClient
}

func NewServer(conn *grpc.ClientConn) *Server {
	return &Server{
		client: pb.NewRecruitmentServiceClient(conn),
	}
}

// Создание резюме
func (s *Server) CreateResumeHandler(w http.ResponseWriter, r *http.Request) {
	var resume pb.Resume
	if err := json.NewDecoder(r.Body).Decode(&resume); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := &pb.CreateResumeRequest{Resume: &resume}
	res, err := s.client.CreateResume(context.Background(), req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res.Resume)
}

// Получение резюме по ID
func (s *Server) GetResumeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	req := &pb.GetResumeRequest{Id: int32(id)}
	res, err := s.client.GetResume(context.Background(), req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res.Resume)
}

// Получение всех резюме
func (s *Server) GetAllResumesHandler(w http.ResponseWriter, r *http.Request) {
	req := &pb.GetAllResumesRequest{}
	res, err := s.client.GetAllResumes(context.Background(), req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res.Resumes)
}

// Обновление резюме
func (s *Server) UpdateResumeHandler(w http.ResponseWriter, r *http.Request) {
	var resume pb.Resume
	if err := json.NewDecoder(r.Body).Decode(&resume); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := &pb.UpdateResumeRequest{Resume: &resume}
	res, err := s.client.UpdateResume(context.Background(), req)
	if err != nil {
		log.Println(err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res.Resume)
}

// Удаление резюме
func (s *Server) DeleteResumeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	req := &pb.DeleteResumeRequest{Id: int32(id)}
	res, err := s.client.DeleteResume(context.Background(), req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if res.Success {
		w.WriteHeader(http.StatusNoContent)
	} else {
		log.Println(err)
		http.Error(w, "Failed to delete resume", http.StatusInternalServerError)
	}
}

// Создание вакансии
func (s *Server) CreateVacancyHandler(w http.ResponseWriter, r *http.Request) {
	var vacancy pb.Vacancy
	if err := json.NewDecoder(r.Body).Decode(&vacancy); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := &pb.CreateVacancyRequest{Vacancy: &vacancy}
	res, err := s.client.CreateVacancy(context.Background(), req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res.Vacancy)
}

// Получение вакансии по ID
func (s *Server) GetVacancyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	req := &pb.GetVacancyRequest{Id: int32(id)}
	res, err := s.client.GetVacancy(context.Background(), req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res.Vacancy)
}

// Получение всех вакансий
func (s *Server) GetAllVacanciesHandler(w http.ResponseWriter, r *http.Request) {
	req := &pb.GetAllVacanciesRequest{}
	res, err := s.client.GetAllVacancies(context.Background(), req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res.Vacancies)
}

// Обновление вакансии
func (s *Server) UpdateVacancyHandler(w http.ResponseWriter, r *http.Request) {
	var vacancy pb.Vacancy
	if err := json.NewDecoder(r.Body).Decode(&vacancy); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := &pb.UpdateVacancyRequest{Vacancy: &vacancy}
	res, err := s.client.UpdateVacancy(context.Background(), req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res.Vacancy)
}

// Удаление вакансии
func (s *Server) DeleteVacancyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	req := &pb.DeleteVacancyRequest{Id: int32(id)}
	res, err := s.client.DeleteVacancy(context.Background(), req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if res.Success {
		w.WriteHeader(http.StatusNoContent)
	} else {
		log.Println(err)
		http.Error(w, "Failed to delete vacancy", http.StatusInternalServerError)
	}
}

func main() {
	// Подключение к gRPC серверу
	conn, err := grpc.Dial(grpcServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	server := NewServer(conn)

	// Настройка маршрутизации
	r := mux.NewRouter()
	r.HandleFunc("/resumes", server.CreateResumeHandler).Methods("POST")
	r.HandleFunc("/resumes/{id}", server.GetResumeHandler).Methods("GET")
	r.HandleFunc("/resumes", server.GetAllResumesHandler).Methods("GET")
	r.HandleFunc("/resumes", server.UpdateResumeHandler).Methods("PUT")
	r.HandleFunc("/resumes/{id}", server.DeleteResumeHandler).Methods("DELETE")

	r.HandleFunc("/vacancies", server.CreateVacancyHandler).Methods("POST")
	r.HandleFunc("/vacancies/{id}", server.GetVacancyHandler).Methods("GET")
	r.HandleFunc("/vacancies", server.GetAllVacanciesHandler).Methods("GET")
	r.HandleFunc("/vacancies", server.UpdateVacancyHandler).Methods("PUT")
	r.HandleFunc("/vacancies/{id}", server.DeleteVacancyHandler).Methods("DELETE")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Println("Starting API Gateway on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("could not listen on :8080: %v", err)
	}
}
