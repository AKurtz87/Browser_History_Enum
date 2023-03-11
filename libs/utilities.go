package libs

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
)

func FindAndCopyFile(pathArr []string, filenameArr []string) error {
    for _, path := range pathArr {
        for _, filename := range filenameArr {
            err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
                if err != nil {
                    fmt.Println(err)
                }
                if info.Name() == filename {
                    currentDir, err := os.Getwd()
                    if err != nil {
                        fmt.Println(err)
                    }
                    destPath := filepath.Join(currentDir, "/browser_history", filename)
                    if err := CopyFile(path, destPath); err != nil {
                        fmt.Println(err)
                    }
                    fmt.Printf("File %s found and copied to %s\n", filename, destPath)
                    return filepath.SkipDir
                }
                return nil
            })
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func CopyFile(sourceFile string, destFile string) (err error) {
    source, err := os.Open(sourceFile)
    if err != nil {
        return err
    }
    defer source.Close()

    dest, err := os.Create(destFile)
    if err != nil {
        return err
    }
    defer dest.Close()

    _, err = io.Copy(dest, source)
    if err != nil {
        return err
    }

    err = dest.Sync()
    if err != nil {
        return err
    }

    fileInfo, err := os.Stat(sourceFile)
    if err != nil {
        return err
    }

    err = os.Chmod(destFile, fileInfo.Mode())
    if err != nil {
        return err
    }

    return nil
}

func GetUsername() string {
    currentUser, err := user.Current()
    if err != nil {
        panic(err)
    }
    return currentUser.Username
}

func GeneratePaths(username string) []string {
    return []string{
		// MACOS
"/Users/" + username + "/Library/Application Support/Google/Chrome/Default", //Google Chrome file: History
"/Users/" + username + "/Library/Application Support/Firefox/Profiles/iphgfuqm.default-release", //Mozilla Firefox file: places.sqlite
"/Users/" + username + "/Library/Safari", //Safari file: History.db SIP BLOCK REQ
    }
}

// WINDOWS
//C:\Users\{Username}\AppData\Local\Google\Chrome\User Data\Default\History //Google Chrome
//C:\Users\{Username}\AppData\Roaming\Mozilla\Firefox\Profiles\{Profile_Name}\places.sqlite //Mozilla Firefox
//C:\Users\{Username}\AppData\Local\Microsoft\Edge\User Data\Default\History //Microsoft Edge
//C:\Users\{Username}\AppData\Local\Microsoft\Windows\History\History.IE5\index.dat //Internet Explorer

//LINUX
//~/.config/google-chrome/Default/History //Google Chrome
//~/.mozilla/firefox/{Profile_Name}/places.sqlite //Mozilla Firefox
//~/.config/chromium/Default/History //Chromium