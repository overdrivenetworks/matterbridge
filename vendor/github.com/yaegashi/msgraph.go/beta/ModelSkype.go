// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SkypeForBusinessActivityCounts undocumented
type SkypeForBusinessActivityCounts struct {
	// Entity is the base model of SkypeForBusinessActivityCounts
	Entity
	// PeerToPeer undocumented
	PeerToPeer *int `json:"peerToPeer,omitempty"`
	// Organized undocumented
	Organized *int `json:"organized,omitempty"`
	// Participated undocumented
	Participated *int `json:"participated,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// SkypeForBusinessActivityUserCounts undocumented
type SkypeForBusinessActivityUserCounts struct {
	// Entity is the base model of SkypeForBusinessActivityUserCounts
	Entity
	// PeerToPeer undocumented
	PeerToPeer *int `json:"peerToPeer,omitempty"`
	// Organized undocumented
	Organized *int `json:"organized,omitempty"`
	// Participated undocumented
	Participated *int `json:"participated,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// SkypeForBusinessActivityUserDetail undocumented
type SkypeForBusinessActivityUserDetail struct {
	// Entity is the base model of SkypeForBusinessActivityUserDetail
	Entity
	// TotalPeerToPeerSessionCount undocumented
	TotalPeerToPeerSessionCount *int `json:"totalPeerToPeerSessionCount,omitempty"`
	// TotalOrganizedConferenceCount undocumented
	TotalOrganizedConferenceCount *int `json:"totalOrganizedConferenceCount,omitempty"`
	// TotalParticipatedConferenceCount undocumented
	TotalParticipatedConferenceCount *int `json:"totalParticipatedConferenceCount,omitempty"`
	// PeerToPeerLastActivityDate undocumented
	PeerToPeerLastActivityDate *Date `json:"peerToPeerLastActivityDate,omitempty"`
	// OrganizedConferenceLastActivityDate undocumented
	OrganizedConferenceLastActivityDate *Date `json:"organizedConferenceLastActivityDate,omitempty"`
	// ParticipatedConferenceLastActivityDate undocumented
	ParticipatedConferenceLastActivityDate *Date `json:"participatedConferenceLastActivityDate,omitempty"`
	// PeerToPeerIMCount undocumented
	PeerToPeerIMCount *int `json:"peerToPeerIMCount,omitempty"`
	// PeerToPeerAudioCount undocumented
	PeerToPeerAudioCount *int `json:"peerToPeerAudioCount,omitempty"`
	// PeerToPeerAudioMinutes undocumented
	PeerToPeerAudioMinutes *int `json:"peerToPeerAudioMinutes,omitempty"`
	// PeerToPeerVideoCount undocumented
	PeerToPeerVideoCount *int `json:"peerToPeerVideoCount,omitempty"`
	// PeerToPeerVideoMinutes undocumented
	PeerToPeerVideoMinutes *int `json:"peerToPeerVideoMinutes,omitempty"`
	// PeerToPeerAppSharingCount undocumented
	PeerToPeerAppSharingCount *int `json:"peerToPeerAppSharingCount,omitempty"`
	// PeerToPeerFileTransferCount undocumented
	PeerToPeerFileTransferCount *int `json:"peerToPeerFileTransferCount,omitempty"`
	// OrganizedConferenceIMCount undocumented
	OrganizedConferenceIMCount *int `json:"organizedConferenceIMCount,omitempty"`
	// OrganizedConferenceAudioVideoCount undocumented
	OrganizedConferenceAudioVideoCount *int `json:"organizedConferenceAudioVideoCount,omitempty"`
	// OrganizedConferenceAudioVideoMinutes undocumented
	OrganizedConferenceAudioVideoMinutes *int `json:"organizedConferenceAudioVideoMinutes,omitempty"`
	// OrganizedConferenceAppSharingCount undocumented
	OrganizedConferenceAppSharingCount *int `json:"organizedConferenceAppSharingCount,omitempty"`
	// OrganizedConferenceWebCount undocumented
	OrganizedConferenceWebCount *int `json:"organizedConferenceWebCount,omitempty"`
	// OrganizedConferenceDialInOut3rdPartyCount undocumented
	OrganizedConferenceDialInOut3rdPartyCount *int `json:"organizedConferenceDialInOut3rdPartyCount,omitempty"`
	// OrganizedConferenceCloudDialInOutMicrosoftCount undocumented
	OrganizedConferenceCloudDialInOutMicrosoftCount *int `json:"organizedConferenceCloudDialInOutMicrosoftCount,omitempty"`
	// OrganizedConferenceCloudDialInMicrosoftMinutes undocumented
	OrganizedConferenceCloudDialInMicrosoftMinutes *int `json:"organizedConferenceCloudDialInMicrosoftMinutes,omitempty"`
	// OrganizedConferenceCloudDialOutMicrosoftMinutes undocumented
	OrganizedConferenceCloudDialOutMicrosoftMinutes *int `json:"organizedConferenceCloudDialOutMicrosoftMinutes,omitempty"`
	// ParticipatedConferenceIMCount undocumented
	ParticipatedConferenceIMCount *int `json:"participatedConferenceIMCount,omitempty"`
	// ParticipatedConferenceAudioVideoCount undocumented
	ParticipatedConferenceAudioVideoCount *int `json:"participatedConferenceAudioVideoCount,omitempty"`
	// ParticipatedConferenceAudioVideoMinutes undocumented
	ParticipatedConferenceAudioVideoMinutes *int `json:"participatedConferenceAudioVideoMinutes,omitempty"`
	// ParticipatedConferenceAppSharingCount undocumented
	ParticipatedConferenceAppSharingCount *int `json:"participatedConferenceAppSharingCount,omitempty"`
	// ParticipatedConferenceWebCount undocumented
	ParticipatedConferenceWebCount *int `json:"participatedConferenceWebCount,omitempty"`
	// ParticipatedConferenceDialInOut3rdPartyCount undocumented
	ParticipatedConferenceDialInOut3rdPartyCount *int `json:"participatedConferenceDialInOut3rdPartyCount,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// IsDeleted undocumented
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// DeletedDate undocumented
	DeletedDate *Date `json:"deletedDate,omitempty"`
	// LastActivityDate undocumented
	LastActivityDate *Date `json:"lastActivityDate,omitempty"`
	// AssignedProducts undocumented
	AssignedProducts []string `json:"assignedProducts,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// SkypeForBusinessDeviceUsageDistributionUserCounts undocumented
type SkypeForBusinessDeviceUsageDistributionUserCounts struct {
	// Entity is the base model of SkypeForBusinessDeviceUsageDistributionUserCounts
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// Windows undocumented
	Windows *int `json:"windows,omitempty"`
	// WindowsPhone undocumented
	WindowsPhone *int `json:"windowsPhone,omitempty"`
	// AndroidPhone undocumented
	AndroidPhone *int `json:"androidPhone,omitempty"`
	// IPhone undocumented
	IPhone *int `json:"iPhone,omitempty"`
	// IPad undocumented
	IPad *int `json:"iPad,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// SkypeForBusinessDeviceUsageUserCounts undocumented
type SkypeForBusinessDeviceUsageUserCounts struct {
	// Entity is the base model of SkypeForBusinessDeviceUsageUserCounts
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// Windows undocumented
	Windows *int `json:"windows,omitempty"`
	// WindowsPhone undocumented
	WindowsPhone *int `json:"windowsPhone,omitempty"`
	// AndroidPhone undocumented
	AndroidPhone *int `json:"androidPhone,omitempty"`
	// IPhone undocumented
	IPhone *int `json:"iPhone,omitempty"`
	// IPad undocumented
	IPad *int `json:"iPad,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// SkypeForBusinessDeviceUsageUserDetail undocumented
type SkypeForBusinessDeviceUsageUserDetail struct {
	// Entity is the base model of SkypeForBusinessDeviceUsageUserDetail
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// LastActivityDate undocumented
	LastActivityDate *Date `json:"lastActivityDate,omitempty"`
	// UsedWindows undocumented
	UsedWindows *bool `json:"usedWindows,omitempty"`
	// UsedWindowsPhone undocumented
	UsedWindowsPhone *bool `json:"usedWindowsPhone,omitempty"`
	// UsedAndroidPhone undocumented
	UsedAndroidPhone *bool `json:"usedAndroidPhone,omitempty"`
	// UsediPhone undocumented
	UsediPhone *bool `json:"usediPhone,omitempty"`
	// UsediPad undocumented
	UsediPad *bool `json:"usediPad,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// SkypeForBusinessOrganizerActivityCounts undocumented
type SkypeForBusinessOrganizerActivityCounts struct {
	// Entity is the base model of SkypeForBusinessOrganizerActivityCounts
	Entity
	// Im undocumented
	Im *int `json:"im,omitempty"`
	// AudioVideo undocumented
	AudioVideo *int `json:"audioVideo,omitempty"`
	// AppSharing undocumented
	AppSharing *int `json:"appSharing,omitempty"`
	// Web undocumented
	Web *int `json:"web,omitempty"`
	// DialInOut3rdParty undocumented
	DialInOut3rdParty *int `json:"dialInOut3rdParty,omitempty"`
	// DialInOutMicrosoft undocumented
	DialInOutMicrosoft *int `json:"dialInOutMicrosoft,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// SkypeForBusinessOrganizerActivityMinuteCounts undocumented
type SkypeForBusinessOrganizerActivityMinuteCounts struct {
	// Entity is the base model of SkypeForBusinessOrganizerActivityMinuteCounts
	Entity
	// AudioVideo undocumented
	AudioVideo *int `json:"audioVideo,omitempty"`
	// DialInMicrosoft undocumented
	DialInMicrosoft *int `json:"dialInMicrosoft,omitempty"`
	// DialOutMicrosoft undocumented
	DialOutMicrosoft *int `json:"dialOutMicrosoft,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// SkypeForBusinessOrganizerActivityUserCounts undocumented
type SkypeForBusinessOrganizerActivityUserCounts struct {
	// Entity is the base model of SkypeForBusinessOrganizerActivityUserCounts
	Entity
	// Im undocumented
	Im *int `json:"im,omitempty"`
	// AudioVideo undocumented
	AudioVideo *int `json:"audioVideo,omitempty"`
	// AppSharing undocumented
	AppSharing *int `json:"appSharing,omitempty"`
	// Web undocumented
	Web *int `json:"web,omitempty"`
	// DialInOut3rdParty undocumented
	DialInOut3rdParty *int `json:"dialInOut3rdParty,omitempty"`
	// DialInOutMicrosoft undocumented
	DialInOutMicrosoft *int `json:"dialInOutMicrosoft,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// SkypeForBusinessParticipantActivityCounts undocumented
type SkypeForBusinessParticipantActivityCounts struct {
	// Entity is the base model of SkypeForBusinessParticipantActivityCounts
	Entity
	// Im undocumented
	Im *int `json:"im,omitempty"`
	// AudioVideo undocumented
	AudioVideo *int `json:"audioVideo,omitempty"`
	// AppSharing undocumented
	AppSharing *int `json:"appSharing,omitempty"`
	// Web undocumented
	Web *int `json:"web,omitempty"`
	// DialInOut3rdParty undocumented
	DialInOut3rdParty *int `json:"dialInOut3rdParty,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// SkypeForBusinessParticipantActivityMinuteCounts undocumented
type SkypeForBusinessParticipantActivityMinuteCounts struct {
	// Entity is the base model of SkypeForBusinessParticipantActivityMinuteCounts
	Entity
	// Audiovideo undocumented
	Audiovideo *int `json:"audiovideo,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// SkypeForBusinessParticipantActivityUserCounts undocumented
type SkypeForBusinessParticipantActivityUserCounts struct {
	// Entity is the base model of SkypeForBusinessParticipantActivityUserCounts
	Entity
	// Im undocumented
	Im *int `json:"im,omitempty"`
	// AudioVideo undocumented
	AudioVideo *int `json:"audioVideo,omitempty"`
	// AppSharing undocumented
	AppSharing *int `json:"appSharing,omitempty"`
	// Web undocumented
	Web *int `json:"web,omitempty"`
	// DialInOut3rdParty undocumented
	DialInOut3rdParty *int `json:"dialInOut3rdParty,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// SkypeForBusinessPeerToPeerActivityCounts undocumented
type SkypeForBusinessPeerToPeerActivityCounts struct {
	// Entity is the base model of SkypeForBusinessPeerToPeerActivityCounts
	Entity
	// Im undocumented
	Im *int `json:"im,omitempty"`
	// Audio undocumented
	Audio *int `json:"audio,omitempty"`
	// Video undocumented
	Video *int `json:"video,omitempty"`
	// AppSharing undocumented
	AppSharing *int `json:"appSharing,omitempty"`
	// FileTransfer undocumented
	FileTransfer *int `json:"fileTransfer,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// SkypeForBusinessPeerToPeerActivityMinuteCounts undocumented
type SkypeForBusinessPeerToPeerActivityMinuteCounts struct {
	// Entity is the base model of SkypeForBusinessPeerToPeerActivityMinuteCounts
	Entity
	// Audio undocumented
	Audio *int `json:"audio,omitempty"`
	// Video undocumented
	Video *int `json:"video,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// SkypeForBusinessPeerToPeerActivityUserCounts undocumented
type SkypeForBusinessPeerToPeerActivityUserCounts struct {
	// Entity is the base model of SkypeForBusinessPeerToPeerActivityUserCounts
	Entity
	// Im undocumented
	Im *int `json:"im,omitempty"`
	// Audio undocumented
	Audio *int `json:"audio,omitempty"`
	// Video undocumented
	Video *int `json:"video,omitempty"`
	// AppSharing undocumented
	AppSharing *int `json:"appSharing,omitempty"`
	// FileTransfer undocumented
	FileTransfer *int `json:"fileTransfer,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}
