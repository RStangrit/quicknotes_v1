package fileOps

import (
	"encoding/json"
	"fmt"
	"main/errHandler"
	"main/note"
	"os"
	"strconv"
	"strings"
)

const FileName = "storage.json"

func CheckFileExistence() bool {
	_, err := os.Stat(FileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Storage file not found")
			return false
		}
		fmt.Println("Error checking file:", err)
		return false
	}
	return true
}

func ReadNotesFromFile() {
	notes, err := readNotesFromFile()
	if err != nil {
		errHandler.HandleError(err)
		return
	}

	titles := extractTitles(notes)
	fmt.Println("Available notes:", strings.Join(titles, ", "))
}

func ReadNoteFromFile(i string) {
	notes, err := readNotesFromFile()
	if err != nil {
		errHandler.HandleError(err)
		return
	}
	index, _ := strconv.Atoi(i)
	content := extractContent(notes, index)
	fmt.Println("Note content:", content)
}

func WriteNoteToFile(title string, content string) {
	newNote := note.New(title, content)

	if CheckFileExistence() {
		fmt.Println("File already exists. Appending the note...")
		err := appendNoteToFile(newNote)
		if err != nil {
			errHandler.HandleError(err)
		}
	} else {
		fmt.Println("File not found. Creating a new file...")
		err := createNewFileWithNotes([]note.Note{newNote})
		if err != nil {
			errHandler.HandleError(err)
		}
	}
}

func readNotesFromFile() ([]note.Note, error) {
	fileContent, err := os.ReadFile(FileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var notes []note.Note
	err = json.Unmarshal(fileContent, &notes)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling file content: %w", err)
	}

	return notes, nil
}

func appendNoteToFile(note note.Note) error {
	notes, err := readNotesFromFile()
	if err != nil {
		return err
	}

	notes = append(notes, note)
	return writeNotesToFile(notes)
}

func createNewFileWithNotes(notes []note.Note) error {
	return writeNotesToFile(notes)
}

func writeNotesToFile(notes []note.Note) error {
	content, err := json.Marshal(notes)
	if err != nil {
		return fmt.Errorf("error marshalling notes: %w", err)
	}

	err = os.WriteFile(FileName, content, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	fmt.Println("File written successfully.")
	return nil
}

func extractTitles(notes []note.Note) []string {
	var titles []string
	for i, note := range notes {
		titles = append(titles, fmt.Sprintf("Index%d: %s", i, note.Title))
	}
	return titles
}

func extractContent(notes []note.Note, i int) string {
	if i < 0 || i >= len(notes) {
		return "invalid index"
	}
	return notes[i].Content
}

func updateNoteField(notes []note.Note, index int, title string, content string) error {
	notes[index].Title = title
	notes[index].Content = content
	writeNotesToFile(notes)
	return nil
}

func deleteNote(notes []note.Note, index int) []note.Note {
	return append(notes[:index], notes[index+1:]...)

}

func UpdateNoteInFile(index string, title string, content string) error {
	notes, err := readNotesFromFile()
	if err != nil {
		errHandler.HandleError(err)
		// return
	}
	i, _ := strconv.Atoi(index)
	err = updateNoteField(notes, i, title, content)
	return err
}

func DeleteNoteInFile(index string) error {
	notes, err := readNotesFromFile()
	if err != nil {
		errHandler.HandleError(err)
	}
	i, _ := strconv.Atoi(index)
	notes = deleteNote(notes, i)
	writeNotesToFile(notes)
	return err
}
