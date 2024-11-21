package storage

import "cloud.google.com/go/storage/internal/apiv2/storagepb"

var RegisterStorageServer = storagepb.RegisterStorageServer

type UnimplementedStorageServer = storagepb.UnimplementedStorageServer
