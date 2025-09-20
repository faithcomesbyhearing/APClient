package main

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"strings"
)

func enqueueYaml(filename string, yamlData []byte, queueBucket string, window fyne.Window) {
	ctx := context.Background()
	awsConfig, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-west-2"))
	if err != nil {
		dialog.ShowError(err, window)
		return
	}
	client := s3.NewFromConfig(awsConfig)
	objectKey := "input/" + filename
	input := &s3.PutObjectInput{
		Bucket: &queueBucket,
		Key:    &objectKey,
		Body:   strings.NewReader(string(yamlData)),
	}
	_, err = client.PutObject(ctx, input)
	if err != nil {
		dialog.ShowError(err, window)
		return
	}
	dialog.ShowInformation("Success", "Request uploaded to queue "+queueBucket, window)
}
