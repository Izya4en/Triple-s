package helpers

import "triple-s/arch"

func NtoB(notes [][]string) arch.Buckets {
	var buckets arch.Buckets
	for _, note := range notes {
		bucket := arch.Bucket{
			Name:             note[0],
			CreationTime:     note[1],
			LastModifiedTime: note[2],
			Status:           note[3],
		}
		buckets.Buckets = append(buckets.Buckets, bucket)
	}
	return buckets
}

func BtoN(buckets arch.Buckets) [][]string {
	var notes [][]string
	for _, bucket := range buckets.Buckets {
		note := []string{
			bucket.Name,
			bucket.CreationTime,
			bucket.LastModifiedTime,
			bucket.Status,
		}
		notes = append(notes, note)
	}
	return notes
}

func NtoO(notes [][]string) arch.Objects {
	var objects arch.Objects
	for _, note := range notes {
		object := arch.Object{
			ObjectKey:    note[0],
			ContentType:  note[1],
			Size:         note[2],
			LastModified: note[3],
		}
		objects.Objects = append(objects.Objects, object)
	}
	return objects
}

func OtoN(objects arch.Objects) [][]string {
	var notes [][]string
	for _, object := range objects.Objects {
		note := []string{
			object.ObjectKey,
			object.ContentType,
			object.Size,
			object.LastModified,
		}
		notes = append(notes, note)
	}
	return notes
}
