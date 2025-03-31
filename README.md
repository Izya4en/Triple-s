# Triple-S

## Overview
Triple-S is a simplified implementation of Amazon S3 (Simple Storage Service). It provides a REST API for managing storage buckets and objects, including creating, listing, deleting buckets, and handling file operations. All responses follow the XML format specified by Amazon S3.

## Features
- **Bucket Management**
  - Create, list, and delete storage buckets.
  - Enforce bucket naming conventions.
  - Persist bucket metadata in a CSV file.
- **Object Storage**
  - Upload, retrieve, and delete files.
  - Store metadata (size, type, last modified) in a CSV file.
  - Ensure proper error handling for storage operations.
- **REST API**
  - Fully HTTP-based operations.
  - XML-formatted responses.
  - Proper handling of HTTP status codes.

## Installation
### Prerequisites
- Go 1.20+
- Linux/macOS/Windows

### Build and Run
```sh
$ go build -o triple-s .
$ ./triple-s -port 8080 -dir ./data
```

## API Endpoints
### Bucket Management
#### 1. Create a Bucket
- **Method**: `PUT`
- **Endpoint**: `/{BucketName}`
- **Response**:
  - `200 OK` if created successfully.
  - `400 Bad Request` for invalid name.
  - `409 Conflict` if the bucket exists.

#### 2. List All Buckets
- **Method**: `GET`
- **Endpoint**: `/`
- **Response**:
  - `200 OK` with a list of buckets in XML format.

#### 3. Delete a Bucket
- **Method**: `DELETE`
- **Endpoint**: `/{BucketName}`
- **Response**:
  - `204 No Content` if deleted successfully.
  - `404 Not Found` if the bucket does not exist.
  - `409 Conflict` if the bucket is not empty.

### Object Operations
#### 1. Upload a File
- **Method**: `PUT`
- **Endpoint**: `/{BucketName}/{ObjectKey}`
- **Headers**:
  - `Content-Type`
  - `Content-Length`
- **Response**:
  - `200 OK` if uploaded.
  - `400 Bad Request` for invalid key.
  - `404 Not Found` if the bucket does not exist.

#### 2. Retrieve a File
- **Method**: `GET`
- **Endpoint**: `/{BucketName}/{ObjectKey}`
- **Response**:
  - `200 OK` with file content.
  - `404 Not Found` if not found.

#### 3. Delete a File
- **Method**: `DELETE`
- **Endpoint**: `/{BucketName}/{ObjectKey}`
- **Response**:
  - `204 No Content` if deleted.
  - `404 Not Found` if not found.

## Metadata Storage
### Bucket Metadata (`buckets.csv`)
| Name | CreationTime | LastModifiedTime |
|------|-------------|-----------------|

### Object Metadata (`objects.csv` per bucket)
| ObjectKey | Size | ContentType | LastModified |
|-----------|------|-------------|-------------|

## Usage
```sh
$ ./triple-s --help  
Simple Storage Service.

**Usage:**
    triple-s [-port <N>] [-dir <S>]  
    triple-s --help

**Options:**
- --help     Show this screen.
- --port N   Port number
- --dir S    Path to the directory
```
