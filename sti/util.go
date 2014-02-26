package sti

import (
	"archive/tar"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func writeTar(tw *tar.Writer, path string, fi os.FileInfo) error {
	fr, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fr.Close()

	h := new(tar.Header)
	h.Name = path
	h.Size = fi.Size()
	h.Mode = int64(fi.Mode())
	h.ModTime = fi.ModTime()

	err = tw.WriteHeader(h)
	if err != nil {
		return err
	}

	_, err = io.Copy(tw, fr)
	return err
}

func tarDirectory(dir string) (*os.File, error) {
	fw, err := ioutil.TempFile("", "sti-tar")
	if err != nil {
		return nil, err
	}
	defer fw.Close()

	tw := tar.NewWriter(fw)
	defer tw.Close()

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			err = writeTar(tw, path, info)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return fw, nil
}

func copy(sourcePath string, targetPath string) error {
	cmd := exec.Command("cp", "-ad", sourcePath, targetPath)
	return cmd.Run()
}

func gitCheckout(source string, targetPath string) error {
	return nil
}