package cmd

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// MoveSelected finds any photos that are in the selected folder, but not
// in the edits/raw folder and moves those from the /raw folder to the
// /edits/raw folder (and creates the target directory if needed)
func MoveSelected(targetPath string) (int, error) {
	selectedDirectoryName := viper.GetString("SelectedDirectory")
	fullSelectedPath := filepath.Join(targetPath, selectedDirectoryName)
	selectedPhotos, err := ioutil.ReadDir(fullSelectedPath)
	if err != nil {
		return 0, err
	}

	rawDirectoryName := viper.GetString("RawDirectory")
	fullRawPath := filepath.Join(targetPath, rawDirectoryName)
	rawPhotos, err := ioutil.ReadDir(fullRawPath)
	if err != nil {
		return 0, err
	}

	rawPhotoMap := make(map[string]fs.FileInfo)
	for _, rawPhoto := range rawPhotos {
		imageID := getImageID(rawPhoto.Name())
		rawPhotoMap[imageID] = rawPhoto
	}

	log.Printf("found %v images to move", len(selectedPhotos))

	editsDirectoryName := viper.GetString("EditsDirectory")
	moved := 0
	for _, selectedFile := range selectedPhotos {
		imageID := getImageID(selectedFile.Name())

		rawPhoto, ok := rawPhotoMap[imageID]
		if ok {
			log.Info().Msgf("moving %v to the edits folder", imageID)
			origin := filepath.Join(fullRawPath, rawPhoto.Name())
			destination := filepath.Join(targetPath, editsDirectoryName, rawDirectoryName, rawPhoto.Name())
			err = os.Rename(origin, destination)
			if err != nil {
				return moved, err
			}
			moved += 1
		}
	}
	log.Info().Msgf("completed selecting %v raw images", moved)
	return moved, nil
}

func getImageID(imageName string) string {
	parts := strings.Split(imageName, ".")
	if len(parts) != 2 {
		return ""
	}
	return parts[0]
}
