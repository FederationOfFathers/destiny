// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/FederationOfFathers/destiny"
	"github.com/spf13/cobra"
)

func downloadManifest(r *http.Response, path string) error {
	defer r.Body.Close()
	if r.StatusCode != 200 {
		return fmt.Errorf("Expected 200 OK, got %s", r.Status)
	}
	zipFP, err := ioutil.TempFile("", ".tmp.destiny-cli.")
	if err != nil {
		return err
	}
	os.Remove(zipFP.Name())
	io.Copy(zipFP, r.Body)
	zipFP.Sync()
	zipFP.Seek(0, 0)
	zr, err := zip.NewReader(zipFP, r.ContentLength)
	if len(zr.File) < 1 {
		return fmt.Errorf("Expected a zip file with a zipped file inside")
	}
	zfp, err := zr.File[0].Open()
	if err != nil {
		return err
	}
	defer zfp.Close()
	fp, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = io.Copy(fp, zfp)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

// manifestDownloadCmd represents the manifestDownload command
var manifestDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download the destiny manifest data",
	Run: func(cmd *cobra.Command, args []string) {
		manifest, err := api.Client.Manifest()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("downloading mobileAssetContentPath.sqlite")
		rsp, err := manifest.Get(manifest.AssetContentPath)
		if err != nil {
			log.Fatalf("Error downloading %s: %s", rsp.Request.URL.String(), err.Error())
		}
		if err := downloadManifest(rsp, "mobileAssetContentPath.sqlite"); err != nil {
			log.Fatalf("Error downloafing %s: %s", rsp.Request.URL.String(), err.Error())
		}

		var wg sync.WaitGroup

		for lang, path := range manifest.WorldContentPaths {
			wg.Add(1)
			go func(lang, path string) {
				defer wg.Done()
				toPath := fmt.Sprintf("WorldContentPaths-%s.sqlite", lang)
				fmt.Println("starting download", toPath)
				rsp, err := manifest.Get(path)
				if err != nil {
					log.Fatalf("Error downloading %s: %s", rsp.Request.URL.String(), err.Error())
				}
				if err := downloadManifest(rsp, toPath); err != nil {
					log.Fatalf("Error downloafing %s: %s", rsp.Request.URL.String(), err.Error())
				}
				fmt.Println("finished", toPath)
			}(lang, path)
		}
		wg.Wait()

		for _, db := range manifest.GearAssetDataBases {
			wg.Add(1)
			go func(db destiny.ManifestAssetDatabase) {
				defer wg.Done()
				toPath := fmt.Sprintf("GearAssetDataBases-%d.sqlite", db.Version)
				fmt.Println("starting download", toPath)
				rsp, err := manifest.Get(db.Path)
				if err != nil {
					log.Fatalf("Error downloading %s: %s", rsp.Request.URL.String(), err.Error())
				}
				if err := downloadManifest(rsp, toPath); err != nil {
					log.Fatalf("Error downloafing %s: %s", rsp.Request.URL.String(), err.Error())
				}
				fmt.Println("finished", toPath)
			}(db)
		}
		wg.Wait()
	},
}

func init() {
	manifestCmd.AddCommand(manifestDownloadCmd)
}
