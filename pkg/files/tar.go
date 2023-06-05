package files

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)


func ExtractTarGz(filename, targetDir string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(f)
	return extract(tr, targetDir)
}

func ExtractTar(filename, targetDir string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	tr := tar.NewReader(f)
	return extract(tr, targetDir)
}

func extract(tr *tar.Reader, targetDir string) error {
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // reached end of archive
		} else if err != nil {
			return err
		}

		targetPath := filepath.Join(targetDir, hdr.Name)

		info := hdr.FileInfo()
		if info.IsDir() {
			err = os.MkdirAll(targetPath, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}
		} else {
			// Create the file's parent directories if needed
			err = os.MkdirAll(filepath.Dir(targetPath), 0755)
			if err != nil {
				return err
			}

			outputFile, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}

			_, err = io.Copy(outputFile, tr)
			outputFile.Close()

			if err != nil {
				return err
			}
		}
	}

	return nil
}