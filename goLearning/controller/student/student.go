package student

import (
	"encoding/json"
	"net/http"
	"strconv"

	"belajar/database"
	"belajar/model/student"

	"github.com/gorilla/mux"
)

func GetStudent(w http.ResponseWriter, r *http.Request) {
    rows, err := database.DB.Query("SELECT * FROM students")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var students []student.Student
    for rows.Next() {
        var s student.Student
        if err := rows.Scan(&s.StudentId,&s.CourseId,&s.Name,&s.Class,&s.Address); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        students = append(students, s)
    }

    if err := rows.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(students)
}

func PostStudent(w http.ResponseWriter, r *http.Request) {
	var ps student.Student
	if err := json.NewDecoder(r.Body).Decode(&ps); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Query untuk memasukan mahasiswa ke dalam table
	query := `
		INSERT INTO students (course_id, name, class, address) 
		VALUES (?, ?,?,?)`

	// Mengeksekusi query
	res, err := database.DB.Exec(query, ps.CourseId, ps.Name, ps.Class, ps.Address)
	if err != nil {
		http.Error(w, "Failed to insert student: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Ambil id terakhir
	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve last insert ID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the newly created ID in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Student added successfully",
		"id":      id,
	})
}

func PutStudent(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari URL
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Decode JSON body
	var ps student.Student
	if err := json.NewDecoder(r.Body).Decode(&ps); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Query ubah mahasiswa
	query := `
		UPDATE students 
		SET course_id=?, name=?,class=?,address=?
		WHERE student_id=?`

	// Execute the SQL statement
	result, err := database.DB.Exec(query,ps.CourseId,ps.Name,ps.Class, ps.Address,id)
	if err != nil {
		http.Error(w, "Failed to update student: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to retrieve affected rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if any rows were updated
	if rowsAffected == 0 {
		http.Error(w, "No rows were updated", http.StatusNotFound)
		return
	}

	// Return success message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Student updated successfully",
	})
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Prepare the SQL statement for deleting a category admin
	query := `
		DELETE FROM students
		WHERE student_id = ?`

	// Execute the SQL statement
	result, err := database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete student: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to retrieve affected rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "No rows were deleted", http.StatusNotFound)
		return
	}

	// Return the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Student deleted successfully",
	})
}

