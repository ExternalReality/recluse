package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

var cmdBuild = &cobra.Command{
	Use:   "build [binary]",
	Short: "Build a Hermitux kernel for a given binary",
	Args:  cobra.MinimumNArgs(1),
	Run:   build,
}

func build(cmd *cobra.Command, args []string) {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	dir := usr.HomeDir + "/.recluse"
	_, err = os.Stat(dir)

	if os.IsNotExist(err) {
		err = os.MkdirAll(dir+"/bin", os.ModePerm)
		err = os.MkdirAll(dir+"/lib", os.ModePerm)
	} else if err != nil {
		panic(err)
	} else {
		return
	}

	//Download cmake archive
	major := "3.7"
	minor := "2"
	platform := "Linux-x86_64"
	archivename := fmt.Sprintf("cmake-%s.%s-%s.tar", major, minor, platform)
	filename := archivename + ".gz"
	filepath := "/tmp/" + filename
	url := fmt.Sprintf("https://cmake.org/files/v%s/%s", major, filename)
	_, err = os.Stat(filepath)
	if os.IsNotExist(err) {
		out, err := os.Create(filepath)
		if err != nil {
			panic(err)
		}
		defer out.Close()
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			panic(fmt.Errorf("bad status: %s", resp.Status))
		}
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}
	err = untar(filepath, dir+"/bin")
	if err != nil {
		panic(err)
	}
}

func untar(source, dst string) error {
	r, err := os.Open(source)
	if err != nil {
		return err
	}
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case header == nil:
			continue
		}

		target := filepath.Join(dst, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}
			f.Close()
		}
	}
}
