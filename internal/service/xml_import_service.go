package service

import (
	"encoding/xml"
	"errors"
	"io"
	"strconv"
	"strings"
	"uptc/sisgestion/internal/models"
	"uptc/sisgestion/internal/repository"
)

type XMLStudentData struct {
	XMLName            xml.Name `xml:"G_INSCRITOSCURSO"`
	IDEstudiante       string   `xml:"NRO_IDEN_EST"`
	CodigoEstudiante   string   `xml:"NRO_ESTUDIANTE"`
	ApellidoEstudiante string   `xml:"APELLIDO_ESTUDIANTE"`
	NombreEstudiante   string   `xml:"NOMBRE_ESTUDIANTE"`
}

type XMLCursoData struct {
	XMLName          xml.Name `xml:"G_CURSOS"`
	NroAsignatura    string   `xml:"NRO_ASIGNATURA"`
	NombreAsignatura string   `xml:"NOMBRE_ASIGNATURA"`
	NroCurso         string   `xml:"NRO_CURSO"`
	EstudiantesList  struct {
		Estudiantes []XMLStudentData `xml:"G_INSCRITOSCURSO"`
	} `xml:"LIST_G_INSCRITOSCURSO"`
}

type XMLRPRelistaCedula struct {
	XMLName     xml.Name `xml:"RPRELISTA_CEDULA"`
	ListGCursos struct {
		Cursos []XMLCursoData `xml:"G_CURSOS"`
	} `xml:"LIST_G_CURSOS"`
}

type XMLImportService struct {
	studentRepo *repository.StudentRepository
	subjectRepo *repository.SubjectRepository
}

func NewXMLImportService(
	studentRepo *repository.StudentRepository,
	subjectRepo *repository.SubjectRepository,
) *XMLImportService {
	return &XMLImportService{
		studentRepo: studentRepo,
		subjectRepo: subjectRepo,
	}
}

func (s *XMLImportService) ImportStudentsFromXML(reader io.Reader, subjectID uint) ([]models.Student, error) {
	subject, err := s.subjectRepo.GetByID(subjectID)
	if err != nil {
		return nil, errors.New("subject not found")
	}

	xmlData, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.New("error reading XML content: " + err.Error())
	}

	xmlString := string(xmlData)
	xmlString = strings.ReplaceAll(xmlString, "�", "Ñ") // Reemplazar � con Ñ
	xmlString = strings.ReplaceAll(xmlString, "�", "ñ") // Reemplazar � con ñ

	decoder := xml.NewDecoder(strings.NewReader(xmlString))

	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return input, nil
	}

	var rpRelista XMLRPRelistaCedula
	if err := decoder.Decode(&rpRelista); err != nil {
		return nil, errors.New("error parsing XML: " + err.Error())
	}

	if len(rpRelista.ListGCursos.Cursos) == 0 {
		return nil, errors.New("no courses found in XML")
	}

	curso := rpRelista.ListGCursos.Cursos[0]
	estudiantes := rpRelista.ListGCursos.Cursos[0].EstudiantesList.Estudiantes
	if len(estudiantes) == 0 {
		return nil, errors.New("no students found in XML")
	}

	groupNum, err := strconv.Atoi(curso.NroCurso)
	if err != nil {
		return nil, errors.New("invalid course number: " + err.Error())
	}

	var studentsCreated []models.Student
	for _, xmlStudent := range estudiantes {
		nombre := strings.TrimSpace(xmlStudent.NombreEstudiante)
		apellido := strings.TrimSpace(xmlStudent.ApellidoEstudiante)
		codigo := strings.TrimSpace(xmlStudent.CodigoEstudiante)

		existingStudent, _ := s.studentRepo.GetByCode(codigo)

		var student models.Student
		if existingStudent != nil {
			if existingStudent.Name != nombre+" "+apellido {
				existingStudent.Name = nombre + " " + apellido
				if err := s.studentRepo.Update(existingStudent); err != nil {
					return nil, errors.New("error updating student: " + err.Error())
				}
			}
			student = *existingStudent
		} else {
			newStudent := models.Student{
				Code: codigo,
				Name: nombre + " " + apellido,
			}

			if err := s.studentRepo.Create(&newStudent); err != nil {
				return nil, errors.New("error creating student: " + err.Error())
			}
			student = newStudent
		}

		studentsCreated = append(studentsCreated, student)

		if err := s.studentRepo.AddSubject(student.ID, subject.ID); err != nil {
			return nil, errors.New("error adding subject to student: " + err.Error())
		}

		subjectGroupStudent := models.SubjectGroupStudent{
			SubjectID: subject.ID,
			GroupNum:  groupNum,
			StudentID: student.ID,
		}

		if err := s.subjectRepo.AddSubjectGroupStudent(&subjectGroupStudent); err != nil {
			return nil, errors.New("error adding subject-group-student relation: " + err.Error())
		}
	}

	return studentsCreated, nil
}
