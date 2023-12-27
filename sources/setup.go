package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func Setup_Start() {
	Setup_CreateFolder()
	Setup_CheckForModPacks()
}

var ModpacksMutex sync.Mutex

func Setup_CreateFolder() {

	_CreateFolder("/storage/")
	_CreateFolder("/storage/public")
	_CreateFolder("/storage/public/modpacks")
	_CreateFolder("/storage/public/resources")

}
func Setup_CheckForModPacks() error {
	fmt.Println("[INFOS] Looking for modpacks..")

	var wg sync.WaitGroup

	err := filepath.Walk(TRAYS_FOLDER+"storage/public/modpacks", func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && filePath != TRAYS_FOLDER+"storage/public/modpacks" {
			if filePath == TRAYS_FOLDER+"storage/public/modpacks" {
				return nil
			}

			wg.Add(1)
			go func(filePath string, info os.FileInfo) {
				defer wg.Done()
				fmt.Print("Calculating checksum of the folder " + info.Name() + "\n")
				hash, err := HASH_CalculateChecksumForDirectory(TRAYS_FOLDER + filePath)
				if err != nil {
					fmt.Println("[ERROR] Failed to calculate hash of", filePath, "{", err, "}")
					return
				}

				files, err := _CheckForFilesInsidePackFolder(filePath)
				if err != nil {
					fmt.Printf("[ERROR] %s\n", err)
				}

				ModpacksMutex.Lock()
				fmt.Print("\nSetting up new modpack data..\n")
				Pack := Pack{
					Name:       info.Name(),
					Path:       strings.ReplaceAll(filePath, "\\", "/"),
					Version:    "testing",
					ApiVersion: "1.16.5-FORGE-UNDEF",
					Files:      files,
					Hash:       hash,
					Icon:       info.Name() + "/icon.png",
					Background: info.Name() + "/background.png",
					Color:      "1E1E1E",
				}
				Modpacks = append(Modpacks, Pack)
				profile, _ := json.MarshalIndent(Pack, "", " ")
				_ = os.WriteFile(TRAYS_FOLDER+filePath+"/trays.json", profile, 0644)
				ModpacksMutex.Unlock()

			}(filePath, info)
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		fmt.Println("[ERROR] Failed to walk through directories:", err)
		return err
	}
	wg.Wait() // Wait for all hash calculations to finish

	if Modpacks != nil {
		fmt.Println("[SUCCESS] Modpacks found:")
		for _, modpack := range Modpacks {
			fmt.Printf("Name: %s\nPath: %s\nVersion: %s\nHash: %s\n\n", modpack.Name, modpack.Path, modpack.Version, modpack.Hash)
		}
	}

	return nil
}

// Shortcut:
func _CreateFolder(path string) {
	if _, err := os.Stat(TRAYS_FOLDER + path); os.IsNotExist(err) {
		fmt.Println("[SETUP] creating ", path)
		_ = os.Mkdir(TRAYS_FOLDER+path, 0700)
	}
}

func _CheckForFilesInsidePackFolder(modpackPath string) (Files, error) {
	var Mods []Mod
	var Config []Object
	var Scripts []Object

	err := filepath.Walk(modpackPath+"/mods", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}
		checksum, err := HASH_CalculateChecksum(path)
		if err != nil {
			return err
		}

		Mods = append(Mods, Mod{File: Object{Name: info.Name(), Path: strings.ReplaceAll(path, "\\", "/"), Hash: checksum}})
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	err = filepath.Walk(modpackPath+"/config", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}
		checksum, err := HASH_CalculateChecksum(path)
		if err != nil {
			return err
		}

		Config = append(Config, Object{
			Name: info.Name(),
			Path: strings.ReplaceAll(path, "\\", "/"),
			Hash: checksum,
		})
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	err = filepath.Walk(modpackPath+"/scripts", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}
		checksum, err := HASH_CalculateChecksum(path)
		if err != nil {
			return err
		}

		Scripts = append(Scripts, Object{
			Name: info.Name(),
			Path: strings.ReplaceAll(path, "\\", "/"),
			Hash: checksum,
		})
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return Files{Mods: Mods, Config: Config, Script: Scripts}, nil
}
