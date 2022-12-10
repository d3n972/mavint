package scheduledTasks

import (
	"archive/zip"
	"encoding/base64"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/d3n972/mavint/domain"
)

func (g GTFSUpdaterTask) basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
func (g GTFSUpdaterTask) unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())
		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}
			err = os.MkdirAll(fdir, f.Mode())
			if err != nil {
				log.Fatal("F: " + fpath + "  " + err.Error())
				return err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (g GTFSUpdaterTask) exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

type GTFSUpdaterTask struct {
	interval time.Duration
}

func (g GTFSUpdaterTask) GetInterval() time.Duration {
	return 6 * time.Hour
}
func (g GTFSUpdaterTask) Handler(ctx domain.AppContext) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.mavcsoport.hu/gtfs/gtfsMavMenetrend.zip", nil)
	req.Header.Add("Authorization", "Basic "+g.basicAuth("mark.adonyi@gmail.com", "j9i1pwCTNcEx9eRd"))
	resp, err := client.Do(req)
	if err != nil {
		panic("E:" + err.Error())
	}
	if zipBytes, e := io.ReadAll(resp.Body); e == nil {
		os.WriteFile("/tmp/mavgtfs.zip", zipBytes, 0664)
		g.unzip("/tmp/mavgtfs.zip", "/var/lib/gtfs")
		if _, err := g.exists("/var/lib/gtfs"); err == fs.ErrNotExist {
			err := os.MkdirAll("/var/lib/gtfs", 0665)
			if err != nil {
				fmt.Printf("[ERR] %s\n", err.Error())
			}
		}
	}
}
