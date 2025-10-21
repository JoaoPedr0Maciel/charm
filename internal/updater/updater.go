package updater

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

const (
	githubAPIURL = "https://api.github.com/repos/JoaoPedr0Maciel/charm/releases/latest"
	repoURL      = "https://github.com/JoaoPedr0Maciel/charm"
)

type Release struct {
	TagName string  `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

type Asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

func Update(currentVersion string) error {
	cyan := color.New(color.FgHiCyan)
	green := color.New(color.FgHiGreen)
	yellow := color.New(color.FgYellow)
	white := color.New(color.FgHiWhite)

	cyan.Println("\n🔄 Checking for updates...")

	// Buscar última versão
	release, err := getLatestRelease()
	if err != nil {
		return fmt.Errorf("failed to check for updates: %w", err)
	}

	latestVersion := strings.TrimPrefix(release.TagName, "v")
	currentVersionClean := strings.TrimPrefix(currentVersion, "v")

	white.Printf("Current version: %s\n", currentVersionClean)
	white.Printf("Latest version:  %s\n", latestVersion)

	if currentVersionClean == latestVersion {
		green.Println("\n✅ You're already running the latest version!")
		return nil
	}

	yellow.Printf("\n📦 New version available: %s -> %s\n", currentVersionClean, latestVersion)
	cyan.Println("⬇️  Downloading update...")

	// Encontrar o asset apropriado
	assetName := getAssetName()
	var downloadURL string
	for _, asset := range release.Assets {
		if asset.Name == assetName {
			downloadURL = asset.BrowserDownloadURL
			break
		}
	}

	if downloadURL == "" {
		return fmt.Errorf("no binary found for %s/%s", runtime.GOOS, runtime.GOARCH)
	}

	// Baixar o arquivo
	tmpFile, err := downloadFile(downloadURL)
	if err != nil {
		return fmt.Errorf("failed to download update: %w", err)
	}
	defer os.Remove(tmpFile)

	cyan.Println("📂 Extracting binary...")

	// Extrair o binário
	binaryPath, err := extractBinary(tmpFile)
	if err != nil {
		return fmt.Errorf("failed to extract binary: %w", err)
	}

	cyan.Println("🔧 Installing update...")

	// Substituir o binário atual
	if err := replaceBinary(binaryPath); err != nil {
		return fmt.Errorf("failed to install update: %w", err)
	}

	green.Printf("\n✨ Successfully updated to version %s!\n", latestVersion)
	white.Println("Run 'charm version' to verify the new version.")

	return nil
}

func getLatestRelease() (*Release, error) {
	resp, err := http.Get(githubAPIURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}

	return &release, nil
}

func getAssetName() string {
	osName := runtime.GOOS
	arch := runtime.GOARCH

	// Mapear arquitetura
	if arch == "amd64" {
		arch = "x86_64"
	} else if arch == "arm64" {
		arch = "aarch64"
	}

	// Formato: charm_Linux_x86_64.tar.gz, charm_Darwin_arm64.tar.gz, charm_Windows_x86_64.zip
	osTitle := strings.Title(osName)
	ext := ".tar.gz"
	if osName == "windows" {
		ext = ".zip"
	}

	return fmt.Sprintf("charm_%s_%s%s", osTitle, arch, ext)
}

func downloadFile(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	tmpFile, err := os.CreateTemp("", "charm-update-*")
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()

	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		os.Remove(tmpFile.Name())
		return "", err
	}

	return tmpFile.Name(), nil
}

func extractBinary(archivePath string) (string, error) {
	if strings.HasSuffix(archivePath, ".tar.gz") {
		return extractTarGz(archivePath)
	}
	// TODO: Adicionar suporte para .zip no Windows
	return "", fmt.Errorf("unsupported archive format")
}

func extractTarGz(archivePath string) (string, error) {
	file, err := os.Open(archivePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	tmpDir, err := os.MkdirTemp("", "charm-extract-*")
	if err != nil {
		return "", err
	}

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		if header.Name == "charm" || header.Name == "charm.exe" {
			target := filepath.Join(tmpDir, header.Name)
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return "", err
			}
			defer f.Close()

			if _, err := io.Copy(f, tr); err != nil {
				return "", err
			}

			return target, nil
		}
	}

	return "", fmt.Errorf("charm binary not found in archive")
}

func replaceBinary(newBinaryPath string) error {
	// Obter o caminho do binário atual
	currentBinary, err := os.Executable()
	if err != nil {
		return err
	}

	// Resolver symlinks
	currentBinary, err = filepath.EvalSymlinks(currentBinary)
	if err != nil {
		return err
	}

	// Criar backup
	backupPath := currentBinary + ".backup"
	if err := os.Rename(currentBinary, backupPath); err != nil {
		return fmt.Errorf("failed to backup current binary: %w", err)
	}

	// Copiar novo binário
	source, err := os.Open(newBinaryPath)
	if err != nil {
		os.Rename(backupPath, currentBinary) // Restaurar backup
		return err
	}
	defer source.Close()

	dest, err := os.OpenFile(currentBinary, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		os.Rename(backupPath, currentBinary) // Restaurar backup
		return err
	}
	defer dest.Close()

	if _, err := io.Copy(dest, source); err != nil {
		os.Rename(backupPath, currentBinary) // Restaurar backup
		return err
	}

	// Remover backup
	os.Remove(backupPath)

	return nil
}
