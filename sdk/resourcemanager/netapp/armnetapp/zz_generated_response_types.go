//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

// AccountBackupsClientDeleteResponse contains the response from method AccountBackupsClient.Delete.
type AccountBackupsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountBackupsClientGetResponse contains the response from method AccountBackupsClient.Get.
type AccountBackupsClientGetResponse struct {
	Backup
}

// AccountBackupsClientListResponse contains the response from method AccountBackupsClient.List.
type AccountBackupsClientListResponse struct {
	BackupsList
}

// AccountsClientCreateOrUpdateResponse contains the response from method AccountsClient.CreateOrUpdate.
type AccountsClientCreateOrUpdateResponse struct {
	Account
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.Delete.
type AccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	Account
}

// AccountsClientListBySubscriptionResponse contains the response from method AccountsClient.ListBySubscription.
type AccountsClientListBySubscriptionResponse struct {
	AccountList
}

// AccountsClientListResponse contains the response from method AccountsClient.List.
type AccountsClientListResponse struct {
	AccountList
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.Update.
type AccountsClientUpdateResponse struct {
	Account
}

// BackupPoliciesClientCreateResponse contains the response from method BackupPoliciesClient.Create.
type BackupPoliciesClientCreateResponse struct {
	BackupPolicy
}

// BackupPoliciesClientDeleteResponse contains the response from method BackupPoliciesClient.Delete.
type BackupPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// BackupPoliciesClientGetResponse contains the response from method BackupPoliciesClient.Get.
type BackupPoliciesClientGetResponse struct {
	BackupPolicy
}

// BackupPoliciesClientListResponse contains the response from method BackupPoliciesClient.List.
type BackupPoliciesClientListResponse struct {
	BackupPoliciesList
}

// BackupPoliciesClientUpdateResponse contains the response from method BackupPoliciesClient.Update.
type BackupPoliciesClientUpdateResponse struct {
	BackupPolicy
}

// BackupsClientCreateResponse contains the response from method BackupsClient.Create.
type BackupsClientCreateResponse struct {
	Backup
}

// BackupsClientDeleteResponse contains the response from method BackupsClient.Delete.
type BackupsClientDeleteResponse struct {
	// placeholder for future response values
}

// BackupsClientGetResponse contains the response from method BackupsClient.Get.
type BackupsClientGetResponse struct {
	Backup
}

// BackupsClientGetStatusResponse contains the response from method BackupsClient.GetStatus.
type BackupsClientGetStatusResponse struct {
	BackupStatus
}

// BackupsClientGetVolumeRestoreStatusResponse contains the response from method BackupsClient.GetVolumeRestoreStatus.
type BackupsClientGetVolumeRestoreStatusResponse struct {
	RestoreStatus
}

// BackupsClientListResponse contains the response from method BackupsClient.List.
type BackupsClientListResponse struct {
	BackupsList
}

// BackupsClientUpdateResponse contains the response from method BackupsClient.Update.
type BackupsClientUpdateResponse struct {
	Backup
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PoolsClientCreateOrUpdateResponse contains the response from method PoolsClient.CreateOrUpdate.
type PoolsClientCreateOrUpdateResponse struct {
	CapacityPool
}

// PoolsClientDeleteResponse contains the response from method PoolsClient.Delete.
type PoolsClientDeleteResponse struct {
	// placeholder for future response values
}

// PoolsClientGetResponse contains the response from method PoolsClient.Get.
type PoolsClientGetResponse struct {
	CapacityPool
}

// PoolsClientListResponse contains the response from method PoolsClient.List.
type PoolsClientListResponse struct {
	CapacityPoolList
}

// PoolsClientUpdateResponse contains the response from method PoolsClient.Update.
type PoolsClientUpdateResponse struct {
	CapacityPool
}

// ResourceClientCheckFilePathAvailabilityResponse contains the response from method ResourceClient.CheckFilePathAvailability.
type ResourceClientCheckFilePathAvailabilityResponse struct {
	CheckAvailabilityResponse
}

// ResourceClientCheckNameAvailabilityResponse contains the response from method ResourceClient.CheckNameAvailability.
type ResourceClientCheckNameAvailabilityResponse struct {
	CheckAvailabilityResponse
}

// ResourceClientCheckQuotaAvailabilityResponse contains the response from method ResourceClient.CheckQuotaAvailability.
type ResourceClientCheckQuotaAvailabilityResponse struct {
	CheckAvailabilityResponse
}

// ResourceQuotaLimitsClientGetResponse contains the response from method ResourceQuotaLimitsClient.Get.
type ResourceQuotaLimitsClientGetResponse struct {
	SubscriptionQuotaItem
}

// ResourceQuotaLimitsClientListResponse contains the response from method ResourceQuotaLimitsClient.List.
type ResourceQuotaLimitsClientListResponse struct {
	SubscriptionQuotaItemList
}

// SnapshotPoliciesClientCreateResponse contains the response from method SnapshotPoliciesClient.Create.
type SnapshotPoliciesClientCreateResponse struct {
	SnapshotPolicy
}

// SnapshotPoliciesClientDeleteResponse contains the response from method SnapshotPoliciesClient.Delete.
type SnapshotPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// SnapshotPoliciesClientGetResponse contains the response from method SnapshotPoliciesClient.Get.
type SnapshotPoliciesClientGetResponse struct {
	SnapshotPolicy
}

// SnapshotPoliciesClientListResponse contains the response from method SnapshotPoliciesClient.List.
type SnapshotPoliciesClientListResponse struct {
	SnapshotPoliciesList
}

// SnapshotPoliciesClientListVolumesResponse contains the response from method SnapshotPoliciesClient.ListVolumes.
type SnapshotPoliciesClientListVolumesResponse struct {
	SnapshotPolicyVolumeList
}

// SnapshotPoliciesClientUpdateResponse contains the response from method SnapshotPoliciesClient.Update.
type SnapshotPoliciesClientUpdateResponse struct {
	SnapshotPolicy
}

// SnapshotsClientCreateResponse contains the response from method SnapshotsClient.Create.
type SnapshotsClientCreateResponse struct {
	Snapshot
}

// SnapshotsClientDeleteResponse contains the response from method SnapshotsClient.Delete.
type SnapshotsClientDeleteResponse struct {
	// placeholder for future response values
}

// SnapshotsClientGetResponse contains the response from method SnapshotsClient.Get.
type SnapshotsClientGetResponse struct {
	Snapshot
}

// SnapshotsClientListResponse contains the response from method SnapshotsClient.List.
type SnapshotsClientListResponse struct {
	SnapshotsList
}

// SnapshotsClientRestoreFilesResponse contains the response from method SnapshotsClient.RestoreFiles.
type SnapshotsClientRestoreFilesResponse struct {
	// placeholder for future response values
}

// SnapshotsClientUpdateResponse contains the response from method SnapshotsClient.Update.
type SnapshotsClientUpdateResponse struct {
	Snapshot
}

// SubvolumesClientCreateResponse contains the response from method SubvolumesClient.Create.
type SubvolumesClientCreateResponse struct {
	SubvolumeInfo
}

// SubvolumesClientDeleteResponse contains the response from method SubvolumesClient.Delete.
type SubvolumesClientDeleteResponse struct {
	// placeholder for future response values
}

// SubvolumesClientGetMetadataResponse contains the response from method SubvolumesClient.GetMetadata.
type SubvolumesClientGetMetadataResponse struct {
	SubvolumeModel
}

// SubvolumesClientGetResponse contains the response from method SubvolumesClient.Get.
type SubvolumesClientGetResponse struct {
	SubvolumeInfo
}

// SubvolumesClientListByVolumeResponse contains the response from method SubvolumesClient.ListByVolume.
type SubvolumesClientListByVolumeResponse struct {
	SubvolumesList
}

// SubvolumesClientUpdateResponse contains the response from method SubvolumesClient.Update.
type SubvolumesClientUpdateResponse struct {
	SubvolumeInfo
}

// VaultsClientListResponse contains the response from method VaultsClient.List.
type VaultsClientListResponse struct {
	VaultList
}

// VolumeGroupsClientCreateResponse contains the response from method VolumeGroupsClient.Create.
type VolumeGroupsClientCreateResponse struct {
	VolumeGroupDetails
}

// VolumeGroupsClientDeleteResponse contains the response from method VolumeGroupsClient.Delete.
type VolumeGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// VolumeGroupsClientGetResponse contains the response from method VolumeGroupsClient.Get.
type VolumeGroupsClientGetResponse struct {
	VolumeGroupDetails
}

// VolumeGroupsClientListByNetAppAccountResponse contains the response from method VolumeGroupsClient.ListByNetAppAccount.
type VolumeGroupsClientListByNetAppAccountResponse struct {
	VolumeGroupList
}

// VolumesClientAuthorizeReplicationResponse contains the response from method VolumesClient.AuthorizeReplication.
type VolumesClientAuthorizeReplicationResponse struct {
	// placeholder for future response values
}

// VolumesClientBreakReplicationResponse contains the response from method VolumesClient.BreakReplication.
type VolumesClientBreakReplicationResponse struct {
	// placeholder for future response values
}

// VolumesClientCreateOrUpdateResponse contains the response from method VolumesClient.CreateOrUpdate.
type VolumesClientCreateOrUpdateResponse struct {
	Volume
}

// VolumesClientDeleteReplicationResponse contains the response from method VolumesClient.DeleteReplication.
type VolumesClientDeleteReplicationResponse struct {
	// placeholder for future response values
}

// VolumesClientDeleteResponse contains the response from method VolumesClient.Delete.
type VolumesClientDeleteResponse struct {
	// placeholder for future response values
}

// VolumesClientGetResponse contains the response from method VolumesClient.Get.
type VolumesClientGetResponse struct {
	Volume
}

// VolumesClientListResponse contains the response from method VolumesClient.List.
type VolumesClientListResponse struct {
	VolumeList
}

// VolumesClientPoolChangeResponse contains the response from method VolumesClient.PoolChange.
type VolumesClientPoolChangeResponse struct {
	// placeholder for future response values
}

// VolumesClientReInitializeReplicationResponse contains the response from method VolumesClient.ReInitializeReplication.
type VolumesClientReInitializeReplicationResponse struct {
	// placeholder for future response values
}

// VolumesClientReplicationStatusResponse contains the response from method VolumesClient.ReplicationStatus.
type VolumesClientReplicationStatusResponse struct {
	ReplicationStatus
}

// VolumesClientResyncReplicationResponse contains the response from method VolumesClient.ResyncReplication.
type VolumesClientResyncReplicationResponse struct {
	// placeholder for future response values
}

// VolumesClientRevertResponse contains the response from method VolumesClient.Revert.
type VolumesClientRevertResponse struct {
	// placeholder for future response values
}

// VolumesClientUpdateResponse contains the response from method VolumesClient.Update.
type VolumesClientUpdateResponse struct {
	Volume
}
