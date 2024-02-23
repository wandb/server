package utils

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func TarDir(sourceDir, tarballPath string) error {
    tarball, err := os.Create(tarballPath)
    if err != nil {
        return err
    }
    defer tarball.Close()

    gzipWriter := gzip.NewWriter(tarball)
    defer gzipWriter.Close()

    tarWriter := tar.NewWriter(gzipWriter)
    defer tarWriter.Close()

    return filepath.Walk(sourceDir, func(file string, fi os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if sourceDir == file { // Skip the root directory
            return nil
        }
        header, err := tar.FileInfoHeader(fi, "")
        if err != nil {
            return err
        }
        header.Name = filepath.ToSlash(file[len(sourceDir)+1:]) // Ensure correct header name

        if err := tarWriter.WriteHeader(header); err != nil {
            return err
        }
        if !fi.IsDir() {
            data, err := os.Open(file)
            if err != nil {
                return err
            }
            defer data.Close()
            if _, err := io.Copy(tarWriter, data); err != nil {
                return err
            }
        }
        return nil
    })
}