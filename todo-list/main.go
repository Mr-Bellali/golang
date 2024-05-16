package main

import (
    "database/sql"
    "fmt"
    "time"
    _ "github.com/go-sql-driver/mysql"
    "bufio"
    "os"
)

type task struct {
    Id_task       int    `json:"id_task"`
    Name_task     string `json:"name_task"`
    Description   string `json:"description"`
    Creation_date string `json:"creation_date"`
    Status        bool   `json:"status"`
    Due_date      string `json:"due_date"`
}


func input () string{
	in := bufio.NewReader(os.Stdin)
    line, err := in.ReadString('\n')

    if err != nil {
        panic(err.Error())
    }
    return line
}

func createTask(t *task) {
    db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/Todo")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close() 
    
    insertQuery := `
            INSERT INTO tasks (name_task, description, creation_date, status, due_date)
            VALUES (?, ?, ?, ?, ?)
    `

    _, err = db.Exec(insertQuery, t.Name_task, t.Description, time.Now().Format("2006-01-02"), false, t.Due_date)
    if err != nil {
        panic(err.Error())
    }

    fmt.Println("task created")
}


func getTasks() []task {
    db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/Todo")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    rows, err := db.Query("SELECT * FROM tasks")
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()

    var tasks []task
    for rows.Next() {
        var t task
        err := rows.Scan(&t.Id_task, &t.Name_task, &t.Description, &t.Creation_date, &t.Status, &t.Due_date)
        if err != nil {
            panic(err.Error())
        }
        tasks = append(tasks, t)
    }
    if err := rows.Err(); err != nil {
        panic(err.Error())
    }

    return tasks
}


func checkComplete(id *int) {
    fmt.Printf("%v", getTasks())

    db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/Todo")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

}

func promptOptions() {
    var t task
    for {
        fmt.Printf("Options:\nc. create new task, \ng. get all tasks, \nd. delete a task, \nda. delete all tasks, \nm. modify a task,\ne. exit\n")
        fmt.Printf("Enter an option: ")

        var opt string
        fmt.Scan(&opt)
        switch opt {
        case "c":
            input()
            fmt.Printf("Name: ")
            t.Name_task = input()

            fmt.Printf("Description: ")
            t.Description = input()

            fmt.Printf("Due date (YYYY-MM-DD): ")
            t.Due_date = input()

            createTask(&t)

        case "g":
            tasks := getTasks()
            for _, task := range tasks {
                fmt.Printf("%+v\n", task)
            }

        case "e":
            return

        case "d":
            fmt.Println("Delete a task feature not yet implemented")

        case "da":
            fmt.Println("Delete all tasks feature not yet implemented")

        case "m":
            fmt.Println("Modify a task feature not yet implemented")

        default:
            fmt.Println("Invalid option, please try again")
        }
    }
}



func main() {
    promptOptions()
}
