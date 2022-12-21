package main

import "errors"

type FileSystem struct {
	FileSystemCrawler DirectoryStack
	RootDirectory     *Directory
}

type Directory struct {
	Name             string
	SubDirectoryList []*Directory
	Size             int
}

func NewFileSystem() *FileSystem {
	directoryStack := make([]*Directory, 0)
	return &FileSystem{
		FileSystemCrawler: directoryStack,
		RootDirectory:     nil,
	}
}

func (f *FileSystem) ChangeDirectoryIn(newDirectoryName string) {
	newDirectory := NewDirectory(newDirectoryName)

	if f.RootDirectory == nil {
		f.RootDirectory = newDirectory
	}

	if !f.FileSystemCrawler.IsEmpty() {
		currentDirectory, _ := f.FileSystemCrawler.Top()
		currentDirectory.AddSubDirectory(newDirectory)
	}

	f.FileSystemCrawler.Push(newDirectory)
}

func (f *FileSystem) AddFileSizeToDirectory(fileSize int) error {
	currentDirectory, ok := f.FileSystemCrawler.Top()

	if !ok {
		return errors.New("[FileSystem.AddFileSizeToDirectory] error: no directory to add file size to.")
	}
	currentDirectory.Size += fileSize
	return nil
}

func (f *FileSystem) ChangeDirectoryOut() error {
	prevDirectory, ok := f.FileSystemCrawler.Pop()

	if !ok {
		return errors.New("[FileSystem.ChangeDirectoryUp] error: no outer directory to change directory out of.")
	}

	currentDirectory, ok := f.FileSystemCrawler.Top()

	if ok {
		currentDirectory.Size += prevDirectory.Size
	}

	return nil
}

func NewDirectory(newDirectoryName string) *Directory {
	newSubDirectoryList := make([]*Directory, 0)
	return &Directory{
		Name:             newDirectoryName,
		SubDirectoryList: newSubDirectoryList,
		Size:             0,
	}
}

func (d *Directory) AddSubDirectory(subDirectory *Directory) {
	d.SubDirectoryList = append(d.SubDirectoryList, subDirectory)
}
