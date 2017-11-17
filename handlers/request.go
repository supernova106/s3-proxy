package request

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/supernova106/s3-proxy/config"
)

// GetScreenShot fetches s3 images
func GetScreenShot(c *gin.Context) {
	// Get configs from gin context
	cfg := c.MustGet("cfg").(*config.Config)
	filename := c.Param("filename")

	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)

	ctx := context.Background()
	var cancelFn func()
	if cfg.S3Timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, cfg.S3Timeout)
	}
	// Ensure the context is canceled to prevent leaking.
	// See context package for more information, https://golang.org/pkg/context/
	defer cancelFn()

	result, err := svc.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(cfg.S3Bucket),
		Key:    aws.String(cfg.S3Prefix + "/" + filename),
	})

	if err != nil {
		// Cast err to awserr.Error to handle specific error codes.
		aerr, ok := err.(awserr.Error)
		if ok && aerr.Code() == s3.ErrCodeNoSuchKey {
			// Specific error code handling
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found!"})
		}
		return
	}

	// Make sure to close the body when done with it for S3 GetObject APIs or
	// will leak connections.
	if result.Body != nil {
		defer result.Body.Close()
	}

	// Read IO reader type to byte[]
	body, err := ioutil.ReadAll(result.Body)

	if err != nil {
		log.Fatal("ERROR:", err)
	}

	if strings.Contains(filename, ".jpg") || strings.Contains(filename, ".jpeg") || strings.Contains(filename, ".png") {
		// Render byte[] data
		c.Data(http.StatusOK, "image/jpeg", body)
	} else if strings.Contains(filename, ".html") {
		c.Data(http.StatusOK, "text/html; charset=utf-8", body)
	} else {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Unsupported media type!"})
	}
}

// Check tells the app is running
func Check(c *gin.Context) {
	c.String(200, "Hello! It's running!")
}
