package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

const PartOneSizeLimit = 100000
const FileSystemSizeLimit = 70000000
const SpaceRequiredForUpdate = 30000000

func getFileSystem(reader io.Reader) (*FileSystem, error) {
	fileSystem := NewFileSystem()

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	bashCommandRegex, _ := regexp.Compile("^\\$.*")
	directoryRegex, _ := regexp.Compile("^(dir ).*")
	fileRegex, _ := regexp.Compile("^[0-9]+ .*")

	for scanner.Scan() {
		line := scanner.Text()

		if bashCommandRegex.MatchString(line) {
			bashCommandHandler := &BashCommandHandler{}
			if err := bashCommandHandler.Handle(line, fileSystem); err != nil {
				return nil, err
			}
		} else if directoryRegex.MatchString(line) {
			continue
		} else if fileRegex.MatchString(line) {
			fileHandler := &FileHandler{}
			if err := fileHandler.Handle(line, fileSystem); err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New(fmt.Sprintf("[getFileSystem] error no handler for line input :%s", line))
		}
	}

	// have to pop out to update the total size of the filesystem
	fileSystem.ChangeDirectoryOut()

	return fileSystem, nil
}

func calculateSumDirectorySizeForPartOne(currentDirectory *Directory) int {
	totalSize := 0

	if currentDirectory.Size < PartOneSizeLimit {
		totalSize += currentDirectory.Size
	}

	for _, subDirectory := range currentDirectory.SubDirectoryList {
		totalSize += calculateSumDirectorySizeForPartOne(subDirectory)
	}

	return totalSize
}

func findMinDirectorySizeToRemoveForUpdate(currentDirectory *Directory, fileSystemSize int) int {
	minDirectorySize := currentDirectory.Size

	for _, subDirectory := range currentDirectory.SubDirectoryList {
		subDirectoryMinSize := findMinDirectorySizeToRemoveForUpdate(subDirectory, fileSystemSize)
		unusedSpace := FileSystemSizeLimit - (fileSystemSize - subDirectoryMinSize)
		if unusedSpace > SpaceRequiredForUpdate && subDirectoryMinSize < minDirectorySize {
			minDirectorySize = subDirectoryMinSize
		}
	}

	return minDirectorySize
}

func main() {
	file, err := os.Open("/home/ec2-user/go/src/github.com/iamwillzhu/adventofcode2022day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileSystem, err := getFileSystem(file)

	if err != nil {
		log.Fatal(err)
	}

	sumOfSizeDirectoryLessThanPartOneThreshold := calculateSumDirectorySizeForPartOne(fileSystem.RootDirectory)

	minDirectorySizeToRemoveForUpdate := findMinDirectorySizeToRemoveForUpdate(fileSystem.RootDirectory, fileSystem.RootDirectory.Size)

	fmt.Printf("The sum of the total sizes of directories of at most %d is %d\n", PartOneSizeLimit, sumOfSizeDirectoryLessThanPartOneThreshold)
	fmt.Printf("The total size of the directory to be deleted to free up enough space for the update is %d\n", minDirectorySizeToRemoveForUpdate)

}
