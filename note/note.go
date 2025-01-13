package note

import (
	"errors"
	"main/inputReader"
	"time"
)

// Note представляет заметку с заголовком, содержимым и временем создания.
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

// New создаёт новую заметку.
func New(title, content string) Note {
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
}

func GetNoteData() (string, string, error) {
	title, err := inputReader.GetUserInput("Note title: ")
	if err != nil {
		return "", "", err
	}
	if title == "" {
		return "", "", errors.New("note title cannot be empty")
	}

	content, err := inputReader.GetUserInput("Note content: ")
	if err != nil {
		return "", "", err
	}
	if content == "" {
		return "", "", errors.New("note content cannot be empty")
	}

	return title, content, nil
}

func GetNoteIndex() (string, error) {
	index, err := inputReader.GetUserInput("Note index(only number): ")
	if err != nil {
		return "", err
	}
	if index == "" {
		return "", errors.New("note index cannot be empty")
	}
	return index, nil
}
