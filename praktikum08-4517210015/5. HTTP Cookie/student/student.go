package main

// Definisikan struct Student
type Student struct {
	Id    string
	Name  string
	Grade int32
}

// Variabel global untuk menyimpan data siswa
var students = []*Student{}

// Fungsi untuk mendapatkan semua data siswa
func GetStudents() []*Student {
	return students
}

// Fungsi untuk memilih data siswa berdasarkan ID
func SelectStudent(id string) *Student {
	// Iterasi melalui slice students
	for _, each := range students {
		// Jika ID sesuai, kembalikan siswa
		if each.Id == id {
			return each
		}
	}

	// Jika tidak ada siswa dengan ID yang cocok, kembalikan nil
	return nil
}

// Fungsi init() untuk menambahkan data siswa awal ke variabel students
func init() {
	students = append(students, &Student{Id: "s001", Name: "bourne", Grade: 2})
	students = append(students, &Student{Id: "s002", Name: "ethan", Grade: 2})
	students = append(students, &Student{Id: "s003", Name: "wick", Grade: 3})
}
