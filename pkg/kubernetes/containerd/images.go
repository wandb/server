package containerd

import (
	"archive/tar"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/containerd/containerd/content"
)

func LoadImage(file string) error {
	return nil
}


func DownloadImage(name string, tag string, filename string) error {
	client := Client()
	defer client.Close()
	ctx := context.Background()

	imageURL := fmt.Sprintf("%s:%s", name, tag)
	image, err := client.Pull(
		ctx,
		imageURL,
	)
	if err != nil {
		return fmt.Errorf("failed to pull Containerd image %s: %v", image, err)
	}

	desc := image.Target()

	tarFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create tar file: %v", err)
	}
	defer tarFile.Close()

    tarWriter := tar.NewWriter(tarFile)
    defer tarWriter.Close()

	reader, err := image.ContentStore().ReaderAt(ctx, desc)
	if err != nil {
		return err
	}
	defer reader.Close()

	hdr := &tar.Header{
		Name:    desc.Digest.Hex(),
		Mode:    0600,
		Size:    reader.Size(),
		ModTime: time.Now(),
	}
	if err := tarWriter.WriteHeader(hdr); err != nil {
		return err
	}

	r := content.NewReader(reader)
	_, err = io.Copy(tarWriter, r)
	return err
}