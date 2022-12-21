package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Handler interface {
	Handle(line string, fileSystem *FileSystem)
}

type BashCommandHandler struct {
}

type FileHandler struct {
}

func (h *BashCommandHandler) Handle(line string, fileSystem *FileSystem) error {
	words := strings.Fields(line)
	command := words[1]

	if command == "cd" {
		argument := words[2]
		if argument != ".." {
			fileSystem.ChangeDirectoryIn(argument)
		} else if err := fileSystem.ChangeDirectoryOut(); err != nil {
			return err
		}
	} else if command != "ls" {
		return errors.New(fmt.Sprintf("[BashCommandHandler.Handle] error: unknown bash command %s", command))
	}
	return nil
}

func (h *FileHandler) Handle(line string, fileSystem *FileSystem) error {
	words := strings.Fields(line)
	fileSize, err := strconv.Atoi(words[0])

	if err != nil {
		return err
	}

	if err = fileSystem.AddFileSizeToDirectory(fileSize); err != nil {
		return err
	}

	return nil
}
