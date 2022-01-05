package template_method_demo

import "testing"

func TestExampleHTTPDownloader(t *testing.T) {
	var downloader Downloader = NewHTTPDownloader()
	downloader.Download("http://example.com/abc.zip")

}
func TestExampleFTPDownloader(t *testing.T) {
	var downloader Downloader = NewFTPDownloader()
	downloader.Download("http://example.com/abc.zip")

}
