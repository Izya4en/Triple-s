package helpers

import (
	"log"
	"path/filepath"
	"triple-s/arch"
	"triple-s/base"
)

func ReadBucket() (arch.Buckets, error) {
	bucketPath := filepath.Join(base.Dir, "buckets.csv")
	log.Printf("Reading buckets metadate: %s\n", bucketPath)

	notes, err := ReadCSV(bucketPath)
	if err != nil {
		return arch.Buckets{}, err
	}
	return NtoB(notes), nil
}

func WriteBucket(bucketData arch.Buckets) error {
	bucketPath := filepath.Join(base.Dir, "buckets.csv")
	notes := BtoN(bucketData)

	return WriteCSV(bucketPath, base.BucketsHeader, notes)
}
