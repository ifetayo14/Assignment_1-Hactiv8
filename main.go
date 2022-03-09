package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	id          int
	studentData StudentData
}

type StudentData struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {

	inputArgs := os.Args[1]

	findStudents(inputArgs)
}

func findStudents(studId string) {

	var temp = 0
	studentId, _ := strconv.Atoi(studId)

	students := []Student{
		{
			id: 1,
			studentData: StudentData{
				nama:      "Alvin",
				alamat:    "Balige",
				pekerjaan: "Mahasiswa",
				alasan:    "Menambah Ilmu",
			},
		},
		{
			id: 2,
			studentData: StudentData{
				nama:      "Joe",
				alamat:    "Medan",
				pekerjaan: "Mahasiswa",
				alasan:    "Ingin Pro",
			},
		},
		{
			id: 3,
			studentData: StudentData{
				nama:      "Doe",
				alamat:    "Jakarta",
				pekerjaan: "Mahasiswa",
				alasan:    "Gabut",
			},
		},
	}

	for _, v := range students {
		if studentId == v.id {
			temp = v.id
		}
	}

	if temp != 0 {
		fmt.Println("Nama\t\t:", students[studentId-1].studentData.nama)
		fmt.Println("Alamat\t\t:", students[studentId-1].studentData.alamat)
		fmt.Println("Pekerjaan\t:", students[studentId-1].studentData.pekerjaan)
		fmt.Println("Alasan\t\t:", students[studentId-1].studentData.alasan)
	} else {
		fmt.Println("Data Tidak Ditemukan")
	}
}
