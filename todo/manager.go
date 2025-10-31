package todo

import "fmt"

func AddTask(title string) error {
	tasks, err := LoadTask()
	if err != nil {
		return err
	}

	newTask := Task{
		ID:        len(tasks) + 1,
		Title:     title,
		Completed: false,
	}

	tasks = append(tasks, newTask)
	if err := SaveTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("Tugas '%s' telah ditambahkan.\n", title)
	return nil
}

func ListTasks() error {
	tasks, err := LoadTask()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("Belum ada tugas.")
		return nil
	}

	for _, t := range tasks {
		status := "❌"
		if t.Completed {
			status = "✅"
		}
		fmt.Printf("%d. %s [%s]\n", t.ID, t.Title, status)
	}
	return nil
}

func CompleteTask(id int) error {
	tasks, err := LoadTask()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = true
			if err := SaveTasks(tasks); err != nil {
				return err
			}
			fmt.Printf("Tugas '%s' selesai!\n", t.Title)
			return nil
		}
	}

	fmt.Println("Tugas tidak ditemukan.")
	return nil
}