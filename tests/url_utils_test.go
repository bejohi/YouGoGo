package tests

import (
	"github.com/bejohi/YouGoGo/html"
	"net/url"
	"testing"
)

func TestUrlSuffix_UrlIsNil_EmptyStringReturned(t *testing.T){
	// Act
	suffix := html.UriSuffix(nil)

	// Assert
	if suffix != "" {
		t.Error(suffix + " was not empty.")
	}
}

func TestUrlSuffix_UrlIsTrivial_SuffixReturned(t *testing.T){
	// Arrange
	uri,_ := url.Parse("https://github.com/a/suffix")

	// Act
	suffix := html.UriSuffix(uri)

	// Assert
	if suffix != "suffix" {
		t.Error(suffix + " was not suffix.")
	}
}

func TestUrlSuffix_UrlHasParams_SuffixReturned(t *testing.T){
	// Arrange
	uri,_ := url.Parse("https://github.com/a/suffix?b=c")

	// Act
	suffix := html.UriSuffix(uri)

	// Assert
	if suffix != "suffix" {
		t.Error(suffix + " was not suffix.")
	}
}

func TestUrlSuffix_UrlHasHash_SuffixReturned(t *testing.T){
	// Arrange
	uri,_ := url.Parse("https://github.com/a/suffix#b")

	// Act
	suffix := html.UriSuffix(uri)

	// Assert
	if suffix != "suffix" {
		t.Error(suffix + " was not suffix.")
	}
}

func TestUrlSuffix_UrlHasNoSuffix_EmptyStringReturned(t *testing.T){
	// Arrange
	uri,_ := url.Parse("https://github.com")

	// Act
	suffix := html.UriSuffix(uri)

	// Assert
	if suffix != "" {
		t.Error(suffix + " was not empty.")
	}
}

