package service

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"strings"
	"uptc/sisgestion/internal/models"
	"uptc/sisgestion/internal/repository"
)

type CSVImportService struct {
	subjectRepo *repository.SubjectRepository
}

func NewCSVImportService(subjectRepo *repository.SubjectRepository) *CSVImportService {
	return &CSVImportService{
		subjectRepo: subjectRepo,
	}
}

func (s *CSVImportService) ImportSubjectsFromCSV(reader io.Reader, isElective bool) ([]models.Subject, error) {
	csvReader := csv.NewReader(reader)

	csvReader.LazyQuotes = true
	csvReader.TrimLeadingSpace = true

	headers, err := csvReader.Read()
	if err != nil {
		return nil, errors.New("error reading CSV headers: " + err.Error())
	}

	normalizedHeaders := make([]string, len(headers))
	for i, header := range headers {
		normalizedHeaders[i] = strings.ToLower(strings.TrimSpace(header))
	}

	codeIndex := findIndex(normalizedHeaders, "codigo")
	nameIndex := findIndex(normalizedHeaders, "asignatura")
	creditsIndex := findIndex(normalizedHeaders, "creditos")
	semesterIndex := findIndex(normalizedHeaders, "semestre")

	if codeIndex == -1 || nameIndex == -1 || creditsIndex == -1 || semesterIndex == -1 {
		return nil, errors.New("CSV must contain 'codigo', 'asignatura', 'creditos', and 'semestre' columns")
	}

	var subjects []models.Subject
	var line int = 1

	for {
		line++
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.New("error reading CSV line " + strconv.Itoa(line) + ": " + err.Error())
		}

		if len(record) <= max(codeIndex, nameIndex, creditsIndex, semesterIndex) {
			return nil, errors.New("insufficient fields on line " + strconv.Itoa(line))
		}

		code := strings.TrimSpace(record[codeIndex])
		name := strings.TrimSpace(record[nameIndex])

		creditsStr := strings.TrimSpace(record[creditsIndex])
		credits, err := strconv.Atoi(creditsStr)
		if err != nil {
			return nil, errors.New("invalid credits value on line " + strconv.Itoa(line) + ": " + creditsStr)
		}

		semesterStr := strings.TrimSpace(record[semesterIndex])
		semester, err := strconv.Atoi(semesterStr)
		if err != nil {
			return nil, errors.New("invalid semester value on line " + strconv.Itoa(line) + ": " + semesterStr)
		}

		if code == "" || name == "" {
			return nil, errors.New("code and name are required on line " + strconv.Itoa(line))
		}

		subject := models.Subject{
			Code:       code,
			Name:       name,
			Credits:    credits,
			Semester:      semester,
			IsElective: isElective,
		}

		existing, _ := s.subjectRepo.GetByCode(code)
		if existing != nil {
			subject.ID = existing.ID
			if err := s.subjectRepo.Update(&subject); err != nil {
				return nil, errors.New("error updating subject on line " + strconv.Itoa(line) + ": " + err.Error())
			}
		} else {
			if err := s.subjectRepo.Create(&subject); err != nil {
				return nil, errors.New("error creating subject on line " + strconv.Itoa(line) + ": " + err.Error())
			}
		}

		subjects = append(subjects, subject)
	}

	return subjects, nil
}

func findIndex(slice []string, value string) int {
	for i, item := range slice {
		if item == value {
			return i
		}
	}
	return -1
}

func max(values ...int) int {
	maxValue := values[0]
	for _, v := range values[1:] {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}