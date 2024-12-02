package server

import (
	"context"
	pb "cv_service/api/proto"
	"cv_service/internal/models"
	"cv_service/internal/repository"
)

type Server struct {
	pb.UnimplementedRecruitmentServiceServer
	storage *repository.Storage
}

func NewServer(storage *repository.Storage) *Server {
	return &Server{storage: storage}
}

func (s *Server) GetAllResumes(ctx context.Context, r *pb.GetAllResumesRequest) (*pb.GetAllResumesResponse, error) {
	res, err := s.storage.GetAllResumes()
	if err != nil {

		return nil, err
	}
	resp := &pb.GetAllResumesResponse{
		Resumes: make([]*pb.Resume, 0, len(res)),
	}

	for _, resume := range res {
		resp.Resumes = append(resp.Resumes, &pb.Resume{
			Id:            int32(resume.Id),
			CandidateName: resume.CandidateName,
			Email:         resume.Email,
			Phone:         resume.Phone,
			Experience:    resume.Experience,
			Education:     resume.Education,
		})
	}
	return resp, nil
}

func (s *Server) GetAllVacancies(ctx context.Context, r *pb.GetAllVacanciesRequest) (*pb.GetAllVacanciesResponse, error) {
	res, err := s.storage.GetAllVacancies()
	if err != nil {

		return nil, err
	}
	resp := &pb.GetAllVacanciesResponse{
		Vacancies: make([]*pb.Vacancy, 0, len(res)),
	}

	for _, vacancy := range res {
		resp.Vacancies = append(resp.Vacancies, &pb.Vacancy{
			Id:          int32(vacancy.Id),
			Title:       vacancy.Title,
			Company:     vacancy.Company,
			Location:    vacancy.Location,
			Description: vacancy.Description,
		})
	}
	return resp, nil
}

func (s *Server) CreateResume(ctx context.Context, req *pb.CreateResumeRequest) (*pb.CreateResumeResponse, error) {
	resume := &models.Resume{
		CandidateName: req.Resume.CandidateName,
		Email:         req.Resume.Email,
		Phone:         req.Resume.Phone,
		Experience:    req.Resume.Experience,
		Education:     req.Resume.Education,
	}

	id, err := s.storage.CreateResume(resume)
	if err != nil {

		return nil, err
	}

	resume.Id = id // Устанавливаем ID созданного резюме
	return &pb.CreateResumeResponse{Resume: &pb.Resume{
		Id:            int32(resume.Id),
		CandidateName: resume.CandidateName,
		Email:         resume.Email,
		Phone:         resume.Phone,
		Experience:    resume.Experience,
		Education:     resume.Education,
	}}, nil
}

func (s *Server) GetResume(ctx context.Context, req *pb.GetResumeRequest) (*pb.GetResumeResponse, error) {
	resume, err := s.storage.GetResume(int(req.Id))
	if err != nil {

		return nil, err
	}

	resumePb := &pb.Resume{
		Id:            int32(resume.Id),
		CandidateName: resume.CandidateName,
		Email:         resume.Email,
		Phone:         resume.Phone,
		Experience:    resume.Experience,
		Education:     resume.Education,
	}

	return &pb.GetResumeResponse{
		Resume: resumePb,
	}, nil
}

func (s *Server) UpdateResume(ctx context.Context, req *pb.UpdateResumeRequest) (*pb.UpdateResumeResponse, error) {
	resume := &models.Resume{
		Id:            int(req.Resume.Id),
		CandidateName: req.Resume.CandidateName,
		Email:         req.Resume.Email,
		Phone:         req.Resume.Phone,
		Experience:    req.Resume.Experience,
		Education:     req.Resume.Education,
	}

	err := s.storage.UpdateResume(resume)
	if err != nil {

		return nil, err
	}

	return &pb.UpdateResumeResponse{Resume: req.Resume}, nil
}

func (s *Server) DeleteResume(ctx context.Context, req *pb.DeleteResumeRequest) (*pb.DeleteResumeResponse, error) {
	err := s.storage.DeleteResume(int(req.Id))
	if err != nil {

		return nil, err
	}

	return &pb.DeleteResumeResponse{Success: true}, nil
}

func (s *Server) CreateVacancy(ctx context.Context, req *pb.CreateVacancyRequest) (*pb.CreateVacancyResponse, error) {
	vacancy := &models.Vacancy{
		Title:       req.Vacancy.Title,
		Company:     req.Vacancy.Company,
		Location:    req.Vacancy.Location,
		Description: req.Vacancy.Description,
	}

	id, err := s.storage.CreateVacancy(vacancy)
	if err != nil {

		return nil, err
	}

	vacancy.Id = id // Устанавливаем ID созданной вакансии
	return &pb.CreateVacancyResponse{Vacancy: &pb.Vacancy{
		Id:          int32(vacancy.Id),
		Title:       vacancy.Title,
		Company:     vacancy.Company,
		Location:    vacancy.Location,
		Description: vacancy.Description,
	}}, nil
}

func (s *Server) GetVacancy(ctx context.Context, req *pb.GetVacancyRequest) (*pb.GetVacancyResponse, error) {
	vacancy, err := s.storage.GetVacancy(int(req.Id))
	if err != nil {
		return nil, err
	}

	vacancyPb := &pb.Vacancy{
		Id:          int32(vacancy.Id),
		Title:       vacancy.Title,
		Company:     vacancy.Company,
		Location:    vacancy.Location,
		Description: vacancy.Description,
	}

	return &pb.GetVacancyResponse{
		Vacancy: vacancyPb,
	}, nil
}

func (s *Server) UpdateVacancy(ctx context.Context, req *pb.UpdateVacancyRequest) (*pb.UpdateVacancyResponse, error) {
	vacancy := &models.Vacancy{
		Id:          int(req.Vacancy.Id),
		Title:       req.Vacancy.Title,
		Company:     req.Vacancy.Company,
		Location:    req.Vacancy.Location,
		Description: req.Vacancy.Description,
	}

	err := s.storage.UpdateVacancy(vacancy)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateVacancyResponse{Vacancy: req.Vacancy}, nil
}

func (s *Server) DeleteVacancy(ctx context.Context, req *pb.DeleteVacancyRequest) (*pb.DeleteVacancyResponse, error) {
	err := s.storage.DeleteVacancy(int(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.DeleteVacancyResponse{Success: true}, nil
}
