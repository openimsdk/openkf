// Copyright © 2023 OpenIM open source community. All rights reserved.
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

package client

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/openimsdk/openkf/server/internal/config"
	"github.com/openimsdk/openkf/server/pkg/log"
)

var (
	_minioClient *minio.Client
	_bucket      string
)

// InitMinio init minio client.
func InitMinio() {
	endpoint := fmt.Sprintf("%s:%d", config.Config.Minio.Ip, config.Config.Minio.Port)
	accessKeyID := config.Config.Minio.AccessKeyId
	secretAccessKey := config.Config.Minio.SecretAccessKey
	location := config.Config.Minio.Location
	_bucket = config.Config.Minio.Bucket

	// Initialize _minioClient
	minioClient, err := minio.New(
		endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: false,
		},
	)
	if err != nil {
		log.Panicf("Minio", err.Error(), " Open failed ", endpoint)
	}

	// create bucket if not exists
	exists, err := minioClient.BucketExists(context.Background(), _bucket)
	if err != nil {
		log.Panicf("Minio", err.Error(), " Open Bucket failed ", endpoint)
	}
	if !exists {
		if err = minioClient.MakeBucket(
			context.Background(),
			_bucket,
			minio.MakeBucketOptions{Region: location},
		); err != nil {
			log.Panicf("Minio", err.Error(), " Open failed ", endpoint)
		}
	}

	_minioClient = minioClient

	log.Info("MinIO", "connect ok", endpoint)
}

// PutObject put object to minio.
func PutObject(objectName string, r io.Reader, objectSize int64) error {
	_, err := _minioClient.PutObject(
		context.Background(),
		_bucket,
		objectName,
		r,
		objectSize,
		minio.PutObjectOptions{ContentType: "application/octet-stream"},
	)

	return err
}

// GetObject get object from minio.
func GetObject(objectName string) (io.Reader, error) {
	object, err := _minioClient.GetObject(context.Background(), _bucket, objectName, minio.GetObjectOptions{})

	return object, err
}

// GetObjectUrlForADay get object url from minio. Time is one day.
func GetObjectUrlForADay(objectName string) (string, error) {
	// set 0 to get object
	presignedURL, err := _minioClient.PresignedGetObject(context.Background(), _bucket, objectName, time.Hour*24, nil)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}

// GetObejctURL get object url with objectName from minio.
func GetObejctURL(objectName string) string {
	return fmt.Sprintf("%s/%s/%s", _minioClient.EndpointURL(), _bucket, objectName)
}
