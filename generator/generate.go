/*Package generator is for generating a new project */
package generator

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Generate
This function creates a golang project
*/
func Generate(name, temp string) error {
	if err := template(name, temp); err != nil {
		return errors.New("Error: " + fmt.Sprintf("%s", err))
	}
	return nil
}

func template(name, tm string) error {
	switch temp := strings.ToLower(tm); temp {
	case "mvc":
		if err := createMVC(name); err != nil {
			return errors.New("Error: " + fmt.Sprintf("%s", err))
		}
	case "basic":
		if err := createFolder(name); err != nil {
			return errors.New("Error: " + fmt.Sprintf("%s", err))
		}
		if err := createBasic(name); err != nil {
			log.Fatal(err)
		}
	default:
		if err := createFolder(name); err != nil {
			return errors.New("Error: " + fmt.Sprintf("%s", err))
		}
	}
	return nil
}

/* Private function */
func createBasic(name string) error {

	var (
		main   = name + "/main.go"
		readme = name + "/README.md"
		git    = name + "/.gitignore"
	)
	if name != "" {
		f, err := createFile(main)
		if err != nil {
			return err
		}
		writeToFile(f)
		f, err = createFile(git)
		if err != nil {
			return err
		}
		writeToFile(f)
		f, err = createFile(readme)
		if err != nil {
			return err
		}
		writeToFile(f)
		return nil
	}
	return nil
}

/* Private function */
func createFolder(name string) error {
	if name != "" {
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	return errors.New("Error: -new flag not set")
}

func createFile(fileName string) (*os.File, error) {
	f, err := os.Create(fileName)
	if err != nil {
		return nil, nil
	}
	return f, nil
}

/* Private function
write to file is in no way dynamic at this time but provides simple starter files
*/
func writeToFile(file *os.File) {
	name, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	switch name := name.Name(); name {
	case "main.go":
		data := []byte("package main\n\r\n\r\nfunc init() {\n\r\n\r\n}\n\r\n\r\nfunc main() {\n\r\n\r\n}")

		_, err = file.Write(data)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	case "controller.go":
		data := []byte("package controllers\n\r\n\r\nfunc init() {\n\r\n\r\n}\n\r\n\r\nfunc main() {\n\r\n\r\n}")
		_, err = file.Write(data)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	case "model.go":
		data := []byte("package models\n\r\n\r\nfunc init() {\n\r\n\r\n}\n\r\n\r\nfunc main() {\n\r\n\r\n}")
		_, err = file.Write(data)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	case "index.html":
		data := []byte(`<!doctype html>
			
			<html lang="en">
			<head>
				<meta charset="utf-8">
			
				<title>title here</title>
				<meta name="description" content="content discription here">
				<meta name="author" content="SitePoint">
			
				<link rel="stylesheet" href="css/styles.css?v=1.0">
			
				<!--[if lt IE 9]>
					<script src="https://cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.3/html5shiv.js"></script>
				<![endif]-->
			</head>
			
			<body>
				<script src="js/scripts.js"></script>
			</body>
			</html>`)

		_, err = file.Write(data)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	case ".gitignore":
		data := []byte(`# General
			.DS_Store
			.vscode
			.AppleDouble
			.LSOverride
			
			# Icon must end with two \r
			Icon
			
			
			# Thumbnails
			._*
			
			# Files that might appear in the root of a volume
			.DocumentRevisions-V100
			.fseventsd
			.Spotlight-V100
			.TemporaryItems
			.Trashes
			.VolumeIcon.icns
			.com.apple.timemachine.donotpresent
			
			# Directories potentially created on remote AFP share
			.AppleDB
			.AppleDesktop
			Network Trash Folder
			Temporary Items
			.apdisk`)
		_, err = file.Write(data)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	case "README.md":
		data := []byte("#Basic Readme file\n\r\n\r\nfile generated from getGoing!")
		_, err = file.Write(data)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

}

/* Private function */
func createMVC(name string) error {
	var (
		main           = name + "/main.go"
		modelFile      = "model.go"
		controllerFile = "controller.go"
		viewFile       = "index.html"
		readme         = name + "/README.md"
		public         = name + "/public/"
		root           = name + "/app/"
		models         = root + "models/"
		views          = root + "views/"
		controllers    = root + "controllers/"
		git            = name + "/.gitignore"
	)
	if name != "" {
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			return err
		}
		f, err := createFile(main)
		if err != nil {
			return err
		}
		writeToFile(f)
		f, err = createFile(readme)
		if err != nil {
			return err
		}
		writeToFile(f)
		f, err = createFile(git)
		if err != nil {
			return err
		}
		writeToFile(f)
		if err := os.Mkdir(public, os.ModePerm); err != nil {
			return err
		}
		if err := os.Mkdir(root, os.ModePerm); err != nil {
			return err
		}
		if err := os.Mkdir(models, os.ModePerm); err != nil {
			return err
		}
		f, err = createFile(models + modelFile)
		if err != nil {
			return err
		}
		writeToFile(f)
		if err := os.Mkdir(views, os.ModePerm); err != nil {
			return err
		}
		f, err = createFile(views + viewFile)
		if err != nil {
			return err
		}
		writeToFile(f)
		if err := os.Mkdir(controllers, os.ModePerm); err != nil {
			return err
		}
		f, err = createFile(controllers + controllerFile)
		if err != nil {
			return err
		}
		writeToFile(f)
	}
	return nil
}
