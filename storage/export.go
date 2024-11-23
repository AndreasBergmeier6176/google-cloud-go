package storage

import "cloud.google.com/go/storage/internal/apiv2/storagepb"

var RegisterStorageServer = storagepb.RegisterStorageServer

type UnimplementedStorageServer = storagepb.UnimplementedStorageServer
type Storage_BidiWriteObjectServer = storagepb.Storage_BidiWriteObjectServer
type Storage_ReadObjectServer = storagepb.Storage_ReadObjectServer
type StartResumableWriteRequest = storagepb.StartResumableWriteRequest
type Storage_WriteObjectServer = storagepb.Storage_WriteObjectServer

type Bucket = storagepb.Bucket
type CancelResumableWriteRequest = storagepb.CancelResumableWriteRequest
type CancelResumableWriteResponse = storagepb.CancelResumableWriteResponse
type ComposeObjectRequest = storagepb.ComposeObjectRequest
type CreateBucketRequest = storagepb.CreateBucketRequest
type DeleteBucketRequest = storagepb.DeleteBucketRequest
type DeleteObjectRequest = storagepb.DeleteObjectRequest
type GetBucketRequest = storagepb.GetBucketRequest
type GetObjectRequest = storagepb.GetObjectRequest
type ListBucketsRequest = storagepb.ListBucketsRequest
type ListBucketsResponse = storagepb.ListBucketsResponse
type ListObjectsRequest = storagepb.ListObjectsRequest
type ListObjectsResponse = storagepb.ListObjectsResponse
type LockBucketRetentionPolicyRequest = storagepb.LockBucketRetentionPolicyRequest
type Object = storagepb.Object
type QueryWriteStatusRequest = storagepb.QueryWriteStatusRequest
type QueryWriteStatusResponse = storagepb.QueryWriteStatusResponse
type RestoreObjectRequest = storagepb.RestoreObjectRequest
type ReadObjectRequest = storagepb.ReadObjectRequest
type ReadObjectResponse = storagepb.ReadObjectResponse
type RewriteObjectRequest = storagepb.RewriteObjectRequest
type RewriteResponse = storagepb.RewriteResponse
type StartResumableWriteResponse = storagepb.StartResumableWriteResponse
type UpdateBucketRequest = storagepb.UpdateBucketRequest
type UpdateObjectRequest = storagepb.UpdateObjectRequest

type ChecksummedData = storagepb.ChecksummedData
type ObjectChecksums = storagepb.ObjectChecksums
type ContentRange = storagepb.ContentRange
