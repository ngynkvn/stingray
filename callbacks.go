package stingray

import (
	"github.com/golang/protobuf/proto"
	"github.com/ngynkvn/stingray/deadlock"
)

// Callbacks decodes and routes replay events to callback functions
type Callbacks struct {
	onCDemoStop                               []func(*deadlock.CDemoStop) error
	onCDemoFileHeader                         []func(*deadlock.CDemoFileHeader) error
	onCDemoFileInfo                           []func(*deadlock.CDemoFileInfo) error
	onCDemoSyncTick                           []func(*deadlock.CDemoSyncTick) error
	onCDemoSendTables                         []func(*deadlock.CDemoSendTables) error
	onCDemoClassInfo                          []func(*deadlock.CDemoClassInfo) error
	onCDemoStringTables                       []func(*deadlock.CDemoStringTables) error
	onCDemoPacket                             []func(*deadlock.CDemoPacket) error
	onCDemoSignonPacket                       []func(*deadlock.CDemoPacket) error
	onCDemoConsoleCmd                         []func(*deadlock.CDemoConsoleCmd) error
	onCDemoCustomData                         []func(*deadlock.CDemoCustomData) error
	onCDemoCustomDataCallbacks                []func(*deadlock.CDemoCustomDataCallbacks) error
	onCDemoUserCmd                            []func(*deadlock.CDemoUserCmd) error
	onCDemoFullPacket                         []func(*deadlock.CDemoFullPacket) error
	onCDemoSaveGame                           []func(*deadlock.CDemoSaveGame) error
	onCDemoSpawnGroups                        []func(*deadlock.CDemoSpawnGroups) error
	onCDemoAnimationData                      []func(*deadlock.CDemoAnimationData) error
	onCDemoAnimationHeader                    []func(*deadlock.CDemoAnimationHeader) error
	onCNETMsg_NOP                             []func(*deadlock.CNETMsg_NOP) error
	onCNETMsg_SplitScreenUser                 []func(*deadlock.CNETMsg_SplitScreenUser) error
	onCNETMsg_Tick                            []func(*deadlock.CNETMsg_Tick) error
	onCNETMsg_StringCmd                       []func(*deadlock.CNETMsg_StringCmd) error
	onCNETMsg_SetConVar                       []func(*deadlock.CNETMsg_SetConVar) error
	onCNETMsg_SignonState                     []func(*deadlock.CNETMsg_SignonState) error
	onCNETMsg_SpawnGroup_Load                 []func(*deadlock.CNETMsg_SpawnGroup_Load) error
	onCNETMsg_SpawnGroup_ManifestUpdate       []func(*deadlock.CNETMsg_SpawnGroup_ManifestUpdate) error
	onCNETMsg_SpawnGroup_SetCreationTick      []func(*deadlock.CNETMsg_SpawnGroup_SetCreationTick) error
	onCNETMsg_SpawnGroup_Unload               []func(*deadlock.CNETMsg_SpawnGroup_Unload) error
	onCNETMsg_SpawnGroup_LoadCompleted        []func(*deadlock.CNETMsg_SpawnGroup_LoadCompleted) error
	onCNETMsg_DebugOverlay                    []func(*deadlock.CNETMsg_DebugOverlay) error
	onCSVCMsg_ServerInfo                      []func(*deadlock.CSVCMsg_ServerInfo) error
	onCSVCMsg_FlattenedSerializer             []func(*deadlock.CSVCMsg_FlattenedSerializer) error
	onCSVCMsg_ClassInfo                       []func(*deadlock.CSVCMsg_ClassInfo) error
	onCSVCMsg_SetPause                        []func(*deadlock.CSVCMsg_SetPause) error
	onCSVCMsg_CreateStringTable               []func(*deadlock.CSVCMsg_CreateStringTable) error
	onCSVCMsg_UpdateStringTable               []func(*deadlock.CSVCMsg_UpdateStringTable) error
	onCSVCMsg_VoiceInit                       []func(*deadlock.CSVCMsg_VoiceInit) error
	onCSVCMsg_VoiceData                       []func(*deadlock.CSVCMsg_VoiceData) error
	onCSVCMsg_Print                           []func(*deadlock.CSVCMsg_Print) error
	onCSVCMsg_Sounds                          []func(*deadlock.CSVCMsg_Sounds) error
	onCSVCMsg_SetView                         []func(*deadlock.CSVCMsg_SetView) error
	onCSVCMsg_ClearAllStringTables            []func(*deadlock.CSVCMsg_ClearAllStringTables) error
	onCSVCMsg_CmdKeyValues                    []func(*deadlock.CSVCMsg_CmdKeyValues) error
	onCSVCMsg_BSPDecal                        []func(*deadlock.CSVCMsg_BSPDecal) error
	onCSVCMsg_SplitScreen                     []func(*deadlock.CSVCMsg_SplitScreen) error
	onCSVCMsg_PacketEntities                  []func(*deadlock.CSVCMsg_PacketEntities) error
	onCSVCMsg_Prefetch                        []func(*deadlock.CSVCMsg_Prefetch) error
	onCSVCMsg_Menu                            []func(*deadlock.CSVCMsg_Menu) error
	onCSVCMsg_GetCvarValue                    []func(*deadlock.CSVCMsg_GetCvarValue) error
	onCSVCMsg_StopSound                       []func(*deadlock.CSVCMsg_StopSound) error
	onCSVCMsg_PeerList                        []func(*deadlock.CSVCMsg_PeerList) error
	onCSVCMsg_PacketReliable                  []func(*deadlock.CSVCMsg_PacketReliable) error
	onCSVCMsg_HLTVStatus                      []func(*deadlock.CSVCMsg_HLTVStatus) error
	onCSVCMsg_ServerSteamID                   []func(*deadlock.CSVCMsg_ServerSteamID) error
	onCSVCMsg_FullFrameSplit                  []func(*deadlock.CSVCMsg_FullFrameSplit) error
	onCSVCMsg_RconServerDetails               []func(*deadlock.CSVCMsg_RconServerDetails) error
	onCSVCMsg_UserMessage                     []func(*deadlock.CSVCMsg_UserMessage) error
	onCSVCMsg_Broadcast_Command               []func(*deadlock.CSVCMsg_Broadcast_Command) error
	onCSVCMsg_HltvFixupOperatorStatus         []func(*deadlock.CSVCMsg_HltvFixupOperatorStatus) error
	onCUserMessageAchievementEvent            []func(*deadlock.CUserMessageAchievementEvent) error
	onCUserMessageCloseCaption                []func(*deadlock.CUserMessageCloseCaption) error
	onCUserMessageCloseCaptionDirect          []func(*deadlock.CUserMessageCloseCaptionDirect) error
	onCUserMessageCurrentTimescale            []func(*deadlock.CUserMessageCurrentTimescale) error
	onCUserMessageDesiredTimescale            []func(*deadlock.CUserMessageDesiredTimescale) error
	onCUserMessageFade                        []func(*deadlock.CUserMessageFade) error
	onCUserMessageGameTitle                   []func(*deadlock.CUserMessageGameTitle) error
	onCUserMessageHudMsg                      []func(*deadlock.CUserMessageHudMsg) error
	onCUserMessageHudText                     []func(*deadlock.CUserMessageHudText) error
	onCUserMessageColoredText                 []func(*deadlock.CUserMessageColoredText) error
	onCUserMessageRequestState                []func(*deadlock.CUserMessageRequestState) error
	onCUserMessageResetHUD                    []func(*deadlock.CUserMessageResetHUD) error
	onCUserMessageRumble                      []func(*deadlock.CUserMessageRumble) error
	onCUserMessageSayText                     []func(*deadlock.CUserMessageSayText) error
	onCUserMessageSayText2                    []func(*deadlock.CUserMessageSayText2) error
	onCUserMessageSayTextChannel              []func(*deadlock.CUserMessageSayTextChannel) error
	onCUserMessageShake                       []func(*deadlock.CUserMessageShake) error
	onCUserMessageShakeDir                    []func(*deadlock.CUserMessageShakeDir) error
	onCUserMessageWaterShake                  []func(*deadlock.CUserMessageWaterShake) error
	onCUserMessageTextMsg                     []func(*deadlock.CUserMessageTextMsg) error
	onCUserMessageScreenTilt                  []func(*deadlock.CUserMessageScreenTilt) error
	onCUserMessageVoiceMask                   []func(*deadlock.CUserMessageVoiceMask) error
	onCUserMessageSendAudio                   []func(*deadlock.CUserMessageSendAudio) error
	onCUserMessageItemPickup                  []func(*deadlock.CUserMessageItemPickup) error
	onCUserMessageAmmoDenied                  []func(*deadlock.CUserMessageAmmoDenied) error
	onCUserMessageShowMenu                    []func(*deadlock.CUserMessageShowMenu) error
	onCUserMessageCreditsMsg                  []func(*deadlock.CUserMessageCreditsMsg) error
	onCEntityMessagePlayJingle                []func(*deadlock.CEntityMessagePlayJingle) error
	onCEntityMessageScreenOverlay             []func(*deadlock.CEntityMessageScreenOverlay) error
	onCEntityMessageRemoveAllDecals           []func(*deadlock.CEntityMessageRemoveAllDecals) error
	onCEntityMessagePropagateForce            []func(*deadlock.CEntityMessagePropagateForce) error
	onCEntityMessageDoSpark                   []func(*deadlock.CEntityMessageDoSpark) error
	onCEntityMessageFixAngle                  []func(*deadlock.CEntityMessageFixAngle) error
	onCUserMessageCloseCaptionPlaceholder     []func(*deadlock.CUserMessageCloseCaptionPlaceholder) error
	onCUserMessageCameraTransition            []func(*deadlock.CUserMessageCameraTransition) error
	onCUserMessageAudioParameter              []func(*deadlock.CUserMessageAudioParameter) error
	onCUserMessageHapticsManagerPulse         []func(*deadlock.CUserMessageHapticsManagerPulse) error
	onCUserMessageHapticsManagerEffect        []func(*deadlock.CUserMessageHapticsManagerEffect) error
	onCUserMessageUpdateCssClasses            []func(*deadlock.CUserMessageUpdateCssClasses) error
	onCUserMessageServerFrameTime             []func(*deadlock.CUserMessageServerFrameTime) error
	onCUserMessageLagCompensationError        []func(*deadlock.CUserMessageLagCompensationError) error
	onCUserMessageRequestDllStatus            []func(*deadlock.CUserMessageRequestDllStatus) error
	onCUserMessageRequestUtilAction           []func(*deadlock.CUserMessageRequestUtilAction) error
	onCUserMessageRequestInventory            []func(*deadlock.CUserMessageRequestInventory) error
	onCUserMessageRequestDiagnostic           []func(*deadlock.CUserMessageRequestDiagnostic) error
	onCMsgVDebugGameSessionIDEvent            []func(*deadlock.CMsgVDebugGameSessionIDEvent) error
	onCMsgPlaceDecalEvent                     []func(*deadlock.CMsgPlaceDecalEvent) error
	onCMsgClearWorldDecalsEvent               []func(*deadlock.CMsgClearWorldDecalsEvent) error
	onCMsgClearEntityDecalsEvent              []func(*deadlock.CMsgClearEntityDecalsEvent) error
	onCMsgClearDecalsForSkeletonInstanceEvent []func(*deadlock.CMsgClearDecalsForSkeletonInstanceEvent) error
	onCMsgSource1LegacyGameEventList          []func(*deadlock.CMsgSource1LegacyGameEventList) error
	onCMsgSource1LegacyListenEvents           []func(*deadlock.CMsgSource1LegacyListenEvents) error
	onCMsgSource1LegacyGameEvent              []func(*deadlock.CMsgSource1LegacyGameEvent) error
	onCMsgSosStartSoundEvent                  []func(*deadlock.CMsgSosStartSoundEvent) error
	onCMsgSosStopSoundEvent                   []func(*deadlock.CMsgSosStopSoundEvent) error
	onCMsgSosSetSoundEventParams              []func(*deadlock.CMsgSosSetSoundEventParams) error
	onCMsgSosSetLibraryStackFields            []func(*deadlock.CMsgSosSetLibraryStackFields) error
	onCMsgSosStopSoundEventHash               []func(*deadlock.CMsgSosStopSoundEventHash) error

	pb *proto.Buffer
}

func newCallbacks() *Callbacks {
	return &Callbacks{
		pb: &proto.Buffer{},
	}
}

// OnCDemoStop registers a callback EDemoCommands_DEM_Stop
func (c *Callbacks) OnCDemoStop(fn func(*deadlock.CDemoStop) error) {
	c.onCDemoStop = append(c.onCDemoStop, fn)
}

// OnCDemoFileHeader registers a callback EDemoCommands_DEM_FileHeader
func (c *Callbacks) OnCDemoFileHeader(fn func(*deadlock.CDemoFileHeader) error) {
	c.onCDemoFileHeader = append(c.onCDemoFileHeader, fn)
}

// OnCDemoFileInfo registers a callback EDemoCommands_DEM_FileInfo
func (c *Callbacks) OnCDemoFileInfo(fn func(*deadlock.CDemoFileInfo) error) {
	c.onCDemoFileInfo = append(c.onCDemoFileInfo, fn)
}

// OnCDemoSyncTick registers a callback EDemoCommands_DEM_SyncTick
func (c *Callbacks) OnCDemoSyncTick(fn func(*deadlock.CDemoSyncTick) error) {
	c.onCDemoSyncTick = append(c.onCDemoSyncTick, fn)
}

// OnCDemoSendTables registers a callback EDemoCommands_DEM_SendTables
func (c *Callbacks) OnCDemoSendTables(fn func(*deadlock.CDemoSendTables) error) {
	c.onCDemoSendTables = append(c.onCDemoSendTables, fn)
}

// OnCDemoClassInfo registers a callback EDemoCommands_DEM_ClassInfo
func (c *Callbacks) OnCDemoClassInfo(fn func(*deadlock.CDemoClassInfo) error) {
	c.onCDemoClassInfo = append(c.onCDemoClassInfo, fn)
}

// OnCDemoStringTables registers a callback EDemoCommands_DEM_StringTables
func (c *Callbacks) OnCDemoStringTables(fn func(*deadlock.CDemoStringTables) error) {
	c.onCDemoStringTables = append(c.onCDemoStringTables, fn)
}

// OnCDemoPacket registers a callback EDemoCommands_DEM_Packet
func (c *Callbacks) OnCDemoPacket(fn func(*deadlock.CDemoPacket) error) {
	c.onCDemoPacket = append(c.onCDemoPacket, fn)
}

// OnCDemoSignonPacket registers a callback EDemoCommands_DEM_SignonPacket
func (c *Callbacks) OnCDemoSignonPacket(fn func(*deadlock.CDemoPacket) error) {
	c.onCDemoSignonPacket = append(c.onCDemoSignonPacket, fn)
}

// OnCDemoConsoleCmd registers a callback EDemoCommands_DEM_ConsoleCmd
func (c *Callbacks) OnCDemoConsoleCmd(fn func(*deadlock.CDemoConsoleCmd) error) {
	c.onCDemoConsoleCmd = append(c.onCDemoConsoleCmd, fn)
}

// OnCDemoCustomData registers a callback EDemoCommands_DEM_CustomData
func (c *Callbacks) OnCDemoCustomData(fn func(*deadlock.CDemoCustomData) error) {
	c.onCDemoCustomData = append(c.onCDemoCustomData, fn)
}

// OnCDemoCustomDataCallbacks registers a callback EDemoCommands_DEM_CustomDataCallbacks
func (c *Callbacks) OnCDemoCustomDataCallbacks(fn func(*deadlock.CDemoCustomDataCallbacks) error) {
	c.onCDemoCustomDataCallbacks = append(c.onCDemoCustomDataCallbacks, fn)
}

// OnCDemoUserCmd registers a callback EDemoCommands_DEM_UserCmd
func (c *Callbacks) OnCDemoUserCmd(fn func(*deadlock.CDemoUserCmd) error) {
	c.onCDemoUserCmd = append(c.onCDemoUserCmd, fn)
}

// OnCDemoFullPacket registers a callback EDemoCommands_DEM_FullPacket
func (c *Callbacks) OnCDemoFullPacket(fn func(*deadlock.CDemoFullPacket) error) {
	c.onCDemoFullPacket = append(c.onCDemoFullPacket, fn)
}

// OnCDemoSaveGame registers a callback EDemoCommands_DEM_SaveGame
func (c *Callbacks) OnCDemoSaveGame(fn func(*deadlock.CDemoSaveGame) error) {
	c.onCDemoSaveGame = append(c.onCDemoSaveGame, fn)
}

// OnCDemoSpawnGroups registers a callback EDemoCommands_DEM_SpawnGroups
func (c *Callbacks) OnCDemoSpawnGroups(fn func(*deadlock.CDemoSpawnGroups) error) {
	c.onCDemoSpawnGroups = append(c.onCDemoSpawnGroups, fn)
}

// OnCDemoAnimationData registers a callback EDemoCommands_DEM_AnimationData
func (c *Callbacks) OnCDemoAnimationData(fn func(*deadlock.CDemoAnimationData) error) {
	c.onCDemoAnimationData = append(c.onCDemoAnimationData, fn)
}

// OnCDemoAnimationHeader registers a callback EDemoCommands_DEM_AnimationHeader
func (c *Callbacks) OnCDemoAnimationHeader(fn func(*deadlock.CDemoAnimationHeader) error) {
	c.onCDemoAnimationHeader = append(c.onCDemoAnimationHeader, fn)
}

// OnCNETMsg_NOP registers a callback for NET_Messages_net_NOP
func (c *Callbacks) OnCNETMsg_NOP(fn func(*deadlock.CNETMsg_NOP) error) {
	c.onCNETMsg_NOP = append(c.onCNETMsg_NOP, fn)
}

// OnCNETMsg_SplitScreenUser registers a callback for NET_Messages_net_SplitScreenUser
func (c *Callbacks) OnCNETMsg_SplitScreenUser(fn func(*deadlock.CNETMsg_SplitScreenUser) error) {
	c.onCNETMsg_SplitScreenUser = append(c.onCNETMsg_SplitScreenUser, fn)
}

// OnCNETMsg_Tick registers a callback for NET_Messages_net_Tick
func (c *Callbacks) OnCNETMsg_Tick(fn func(*deadlock.CNETMsg_Tick) error) {
	c.onCNETMsg_Tick = append(c.onCNETMsg_Tick, fn)
}

// OnCNETMsg_StringCmd registers a callback for NET_Messages_net_StringCmd
func (c *Callbacks) OnCNETMsg_StringCmd(fn func(*deadlock.CNETMsg_StringCmd) error) {
	c.onCNETMsg_StringCmd = append(c.onCNETMsg_StringCmd, fn)
}

// OnCNETMsg_SetConVar registers a callback for NET_Messages_net_SetConVar
func (c *Callbacks) OnCNETMsg_SetConVar(fn func(*deadlock.CNETMsg_SetConVar) error) {
	c.onCNETMsg_SetConVar = append(c.onCNETMsg_SetConVar, fn)
}

// OnCNETMsg_SignonState registers a callback for NET_Messages_net_SignonState
func (c *Callbacks) OnCNETMsg_SignonState(fn func(*deadlock.CNETMsg_SignonState) error) {
	c.onCNETMsg_SignonState = append(c.onCNETMsg_SignonState, fn)
}

// OnCNETMsg_SpawnGroup_Load registers a callback for NET_Messages_net_SpawnGroup_Load
func (c *Callbacks) OnCNETMsg_SpawnGroup_Load(fn func(*deadlock.CNETMsg_SpawnGroup_Load) error) {
	c.onCNETMsg_SpawnGroup_Load = append(c.onCNETMsg_SpawnGroup_Load, fn)
}

// OnCNETMsg_SpawnGroup_ManifestUpdate registers a callback for NET_Messages_net_SpawnGroup_ManifestUpdate
func (c *Callbacks) OnCNETMsg_SpawnGroup_ManifestUpdate(fn func(*deadlock.CNETMsg_SpawnGroup_ManifestUpdate) error) {
	c.onCNETMsg_SpawnGroup_ManifestUpdate = append(c.onCNETMsg_SpawnGroup_ManifestUpdate, fn)
}

// OnCNETMsg_SpawnGroup_SetCreationTick registers a callback for NET_Messages_net_SpawnGroup_SetCreationTick
func (c *Callbacks) OnCNETMsg_SpawnGroup_SetCreationTick(fn func(*deadlock.CNETMsg_SpawnGroup_SetCreationTick) error) {
	c.onCNETMsg_SpawnGroup_SetCreationTick = append(c.onCNETMsg_SpawnGroup_SetCreationTick, fn)
}

// OnCNETMsg_SpawnGroup_Unload registers a callback for NET_Messages_net_SpawnGroup_Unload
func (c *Callbacks) OnCNETMsg_SpawnGroup_Unload(fn func(*deadlock.CNETMsg_SpawnGroup_Unload) error) {
	c.onCNETMsg_SpawnGroup_Unload = append(c.onCNETMsg_SpawnGroup_Unload, fn)
}

// OnCNETMsg_SpawnGroup_LoadCompleted registers a callback for NET_Messages_net_SpawnGroup_LoadCompleted
func (c *Callbacks) OnCNETMsg_SpawnGroup_LoadCompleted(fn func(*deadlock.CNETMsg_SpawnGroup_LoadCompleted) error) {
	c.onCNETMsg_SpawnGroup_LoadCompleted = append(c.onCNETMsg_SpawnGroup_LoadCompleted, fn)
}

// OnCNETMsg_DebugOverlay registers a callback for NET_Messages_net_DebugOverlay
func (c *Callbacks) OnCNETMsg_DebugOverlay(fn func(*deadlock.CNETMsg_DebugOverlay) error) {
	c.onCNETMsg_DebugOverlay = append(c.onCNETMsg_DebugOverlay, fn)
}

// OnCSVCMsg_ServerInfo registers a callback for SVC_Messages_svc_ServerInfo
func (c *Callbacks) OnCSVCMsg_ServerInfo(fn func(*deadlock.CSVCMsg_ServerInfo) error) {
	c.onCSVCMsg_ServerInfo = append(c.onCSVCMsg_ServerInfo, fn)
}

// OnCSVCMsg_FlattenedSerializer registers a callback for SVC_Messages_svc_FlattenedSerializer
func (c *Callbacks) OnCSVCMsg_FlattenedSerializer(fn func(*deadlock.CSVCMsg_FlattenedSerializer) error) {
	c.onCSVCMsg_FlattenedSerializer = append(c.onCSVCMsg_FlattenedSerializer, fn)
}

// OnCSVCMsg_ClassInfo registers a callback for SVC_Messages_svc_ClassInfo
func (c *Callbacks) OnCSVCMsg_ClassInfo(fn func(*deadlock.CSVCMsg_ClassInfo) error) {
	c.onCSVCMsg_ClassInfo = append(c.onCSVCMsg_ClassInfo, fn)
}

// OnCSVCMsg_SetPause registers a callback for SVC_Messages_svc_SetPause
func (c *Callbacks) OnCSVCMsg_SetPause(fn func(*deadlock.CSVCMsg_SetPause) error) {
	c.onCSVCMsg_SetPause = append(c.onCSVCMsg_SetPause, fn)
}

// OnCSVCMsg_CreateStringTable registers a callback for SVC_Messages_svc_CreateStringTable
func (c *Callbacks) OnCSVCMsg_CreateStringTable(fn func(*deadlock.CSVCMsg_CreateStringTable) error) {
	c.onCSVCMsg_CreateStringTable = append(c.onCSVCMsg_CreateStringTable, fn)
}

// OnCSVCMsg_UpdateStringTable registers a callback for SVC_Messages_svc_UpdateStringTable
func (c *Callbacks) OnCSVCMsg_UpdateStringTable(fn func(*deadlock.CSVCMsg_UpdateStringTable) error) {
	c.onCSVCMsg_UpdateStringTable = append(c.onCSVCMsg_UpdateStringTable, fn)
}

// OnCSVCMsg_VoiceInit registers a callback for SVC_Messages_svc_VoiceInit
func (c *Callbacks) OnCSVCMsg_VoiceInit(fn func(*deadlock.CSVCMsg_VoiceInit) error) {
	c.onCSVCMsg_VoiceInit = append(c.onCSVCMsg_VoiceInit, fn)
}

// OnCSVCMsg_VoiceData registers a callback for SVC_Messages_svc_VoiceData
func (c *Callbacks) OnCSVCMsg_VoiceData(fn func(*deadlock.CSVCMsg_VoiceData) error) {
	c.onCSVCMsg_VoiceData = append(c.onCSVCMsg_VoiceData, fn)
}

// OnCSVCMsg_Print registers a callback for SVC_Messages_svc_Print
func (c *Callbacks) OnCSVCMsg_Print(fn func(*deadlock.CSVCMsg_Print) error) {
	c.onCSVCMsg_Print = append(c.onCSVCMsg_Print, fn)
}

// OnCSVCMsg_Sounds registers a callback for SVC_Messages_svc_Sounds
func (c *Callbacks) OnCSVCMsg_Sounds(fn func(*deadlock.CSVCMsg_Sounds) error) {
	c.onCSVCMsg_Sounds = append(c.onCSVCMsg_Sounds, fn)
}

// OnCSVCMsg_SetView registers a callback for SVC_Messages_svc_SetView
func (c *Callbacks) OnCSVCMsg_SetView(fn func(*deadlock.CSVCMsg_SetView) error) {
	c.onCSVCMsg_SetView = append(c.onCSVCMsg_SetView, fn)
}

// OnCSVCMsg_ClearAllStringTables registers a callback for SVC_Messages_svc_ClearAllStringTables
func (c *Callbacks) OnCSVCMsg_ClearAllStringTables(fn func(*deadlock.CSVCMsg_ClearAllStringTables) error) {
	c.onCSVCMsg_ClearAllStringTables = append(c.onCSVCMsg_ClearAllStringTables, fn)
}

// OnCSVCMsg_CmdKeyValues registers a callback for SVC_Messages_svc_CmdKeyValues
func (c *Callbacks) OnCSVCMsg_CmdKeyValues(fn func(*deadlock.CSVCMsg_CmdKeyValues) error) {
	c.onCSVCMsg_CmdKeyValues = append(c.onCSVCMsg_CmdKeyValues, fn)
}

// OnCSVCMsg_BSPDecal registers a callback for SVC_Messages_svc_BSPDecal
func (c *Callbacks) OnCSVCMsg_BSPDecal(fn func(*deadlock.CSVCMsg_BSPDecal) error) {
	c.onCSVCMsg_BSPDecal = append(c.onCSVCMsg_BSPDecal, fn)
}

// OnCSVCMsg_SplitScreen registers a callback for SVC_Messages_svc_SplitScreen
func (c *Callbacks) OnCSVCMsg_SplitScreen(fn func(*deadlock.CSVCMsg_SplitScreen) error) {
	c.onCSVCMsg_SplitScreen = append(c.onCSVCMsg_SplitScreen, fn)
}

// OnCSVCMsg_PacketEntities registers a callback for SVC_Messages_svc_PacketEntities
func (c *Callbacks) OnCSVCMsg_PacketEntities(fn func(*deadlock.CSVCMsg_PacketEntities) error) {
	c.onCSVCMsg_PacketEntities = append(c.onCSVCMsg_PacketEntities, fn)
}

// OnCSVCMsg_Prefetch registers a callback for SVC_Messages_svc_Prefetch
func (c *Callbacks) OnCSVCMsg_Prefetch(fn func(*deadlock.CSVCMsg_Prefetch) error) {
	c.onCSVCMsg_Prefetch = append(c.onCSVCMsg_Prefetch, fn)
}

// OnCSVCMsg_Menu registers a callback for SVC_Messages_svc_Menu
func (c *Callbacks) OnCSVCMsg_Menu(fn func(*deadlock.CSVCMsg_Menu) error) {
	c.onCSVCMsg_Menu = append(c.onCSVCMsg_Menu, fn)
}

// OnCSVCMsg_GetCvarValue registers a callback for SVC_Messages_svc_GetCvarValue
func (c *Callbacks) OnCSVCMsg_GetCvarValue(fn func(*deadlock.CSVCMsg_GetCvarValue) error) {
	c.onCSVCMsg_GetCvarValue = append(c.onCSVCMsg_GetCvarValue, fn)
}

// OnCSVCMsg_StopSound registers a callback for SVC_Messages_svc_StopSound
func (c *Callbacks) OnCSVCMsg_StopSound(fn func(*deadlock.CSVCMsg_StopSound) error) {
	c.onCSVCMsg_StopSound = append(c.onCSVCMsg_StopSound, fn)
}

// OnCSVCMsg_PeerList registers a callback for SVC_Messages_svc_PeerList
func (c *Callbacks) OnCSVCMsg_PeerList(fn func(*deadlock.CSVCMsg_PeerList) error) {
	c.onCSVCMsg_PeerList = append(c.onCSVCMsg_PeerList, fn)
}

// OnCSVCMsg_PacketReliable registers a callback for SVC_Messages_svc_PacketReliable
func (c *Callbacks) OnCSVCMsg_PacketReliable(fn func(*deadlock.CSVCMsg_PacketReliable) error) {
	c.onCSVCMsg_PacketReliable = append(c.onCSVCMsg_PacketReliable, fn)
}

// OnCSVCMsg_HLTVStatus registers a callback for SVC_Messages_svc_HLTVStatus
func (c *Callbacks) OnCSVCMsg_HLTVStatus(fn func(*deadlock.CSVCMsg_HLTVStatus) error) {
	c.onCSVCMsg_HLTVStatus = append(c.onCSVCMsg_HLTVStatus, fn)
}

// OnCSVCMsg_ServerSteamID registers a callback for SVC_Messages_svc_ServerSteamID
func (c *Callbacks) OnCSVCMsg_ServerSteamID(fn func(*deadlock.CSVCMsg_ServerSteamID) error) {
	c.onCSVCMsg_ServerSteamID = append(c.onCSVCMsg_ServerSteamID, fn)
}

// OnCSVCMsg_FullFrameSplit registers a callback for SVC_Messages_svc_FullFrameSplit
func (c *Callbacks) OnCSVCMsg_FullFrameSplit(fn func(*deadlock.CSVCMsg_FullFrameSplit) error) {
	c.onCSVCMsg_FullFrameSplit = append(c.onCSVCMsg_FullFrameSplit, fn)
}

// OnCSVCMsg_RconServerDetails registers a callback for SVC_Messages_svc_RconServerDetails
func (c *Callbacks) OnCSVCMsg_RconServerDetails(fn func(*deadlock.CSVCMsg_RconServerDetails) error) {
	c.onCSVCMsg_RconServerDetails = append(c.onCSVCMsg_RconServerDetails, fn)
}

// OnCSVCMsg_UserMessage registers a callback for SVC_Messages_svc_UserMessage
func (c *Callbacks) OnCSVCMsg_UserMessage(fn func(*deadlock.CSVCMsg_UserMessage) error) {
	c.onCSVCMsg_UserMessage = append(c.onCSVCMsg_UserMessage, fn)
}

// OnCSVCMsg_Broadcast_Command registers a callback for SVC_Messages_svc_Broadcast_Command
func (c *Callbacks) OnCSVCMsg_Broadcast_Command(fn func(*deadlock.CSVCMsg_Broadcast_Command) error) {
	c.onCSVCMsg_Broadcast_Command = append(c.onCSVCMsg_Broadcast_Command, fn)
}

// OnCSVCMsg_HltvFixupOperatorStatus registers a callback for SVC_Messages_svc_HltvFixupOperatorStatus
func (c *Callbacks) OnCSVCMsg_HltvFixupOperatorStatus(fn func(*deadlock.CSVCMsg_HltvFixupOperatorStatus) error) {
	c.onCSVCMsg_HltvFixupOperatorStatus = append(c.onCSVCMsg_HltvFixupOperatorStatus, fn)
}

// OnCUserMessageAchievementEvent registers a callback for EBaseUserMessages_UM_AchievementEvent
func (c *Callbacks) OnCUserMessageAchievementEvent(fn func(*deadlock.CUserMessageAchievementEvent) error) {
	c.onCUserMessageAchievementEvent = append(c.onCUserMessageAchievementEvent, fn)
}

// OnCUserMessageCloseCaption registers a callback for EBaseUserMessages_UM_CloseCaption
func (c *Callbacks) OnCUserMessageCloseCaption(fn func(*deadlock.CUserMessageCloseCaption) error) {
	c.onCUserMessageCloseCaption = append(c.onCUserMessageCloseCaption, fn)
}

// OnCUserMessageCloseCaptionDirect registers a callback for EBaseUserMessages_UM_CloseCaptionDirect
func (c *Callbacks) OnCUserMessageCloseCaptionDirect(fn func(*deadlock.CUserMessageCloseCaptionDirect) error) {
	c.onCUserMessageCloseCaptionDirect = append(c.onCUserMessageCloseCaptionDirect, fn)
}

// OnCUserMessageCurrentTimescale registers a callback for EBaseUserMessages_UM_CurrentTimescale
func (c *Callbacks) OnCUserMessageCurrentTimescale(fn func(*deadlock.CUserMessageCurrentTimescale) error) {
	c.onCUserMessageCurrentTimescale = append(c.onCUserMessageCurrentTimescale, fn)
}

// OnCUserMessageDesiredTimescale registers a callback for EBaseUserMessages_UM_DesiredTimescale
func (c *Callbacks) OnCUserMessageDesiredTimescale(fn func(*deadlock.CUserMessageDesiredTimescale) error) {
	c.onCUserMessageDesiredTimescale = append(c.onCUserMessageDesiredTimescale, fn)
}

// OnCUserMessageFade registers a callback for EBaseUserMessages_UM_Fade
func (c *Callbacks) OnCUserMessageFade(fn func(*deadlock.CUserMessageFade) error) {
	c.onCUserMessageFade = append(c.onCUserMessageFade, fn)
}

// OnCUserMessageGameTitle registers a callback for EBaseUserMessages_UM_GameTitle
func (c *Callbacks) OnCUserMessageGameTitle(fn func(*deadlock.CUserMessageGameTitle) error) {
	c.onCUserMessageGameTitle = append(c.onCUserMessageGameTitle, fn)
}

// OnCUserMessageHudMsg registers a callback for EBaseUserMessages_UM_HudMsg
func (c *Callbacks) OnCUserMessageHudMsg(fn func(*deadlock.CUserMessageHudMsg) error) {
	c.onCUserMessageHudMsg = append(c.onCUserMessageHudMsg, fn)
}

// OnCUserMessageHudText registers a callback for EBaseUserMessages_UM_HudText
func (c *Callbacks) OnCUserMessageHudText(fn func(*deadlock.CUserMessageHudText) error) {
	c.onCUserMessageHudText = append(c.onCUserMessageHudText, fn)
}

// OnCUserMessageColoredText registers a callback for EBaseUserMessages_UM_ColoredText
func (c *Callbacks) OnCUserMessageColoredText(fn func(*deadlock.CUserMessageColoredText) error) {
	c.onCUserMessageColoredText = append(c.onCUserMessageColoredText, fn)
}

// OnCUserMessageRequestState registers a callback for EBaseUserMessages_UM_RequestState
func (c *Callbacks) OnCUserMessageRequestState(fn func(*deadlock.CUserMessageRequestState) error) {
	c.onCUserMessageRequestState = append(c.onCUserMessageRequestState, fn)
}

// OnCUserMessageResetHUD registers a callback for EBaseUserMessages_UM_ResetHUD
func (c *Callbacks) OnCUserMessageResetHUD(fn func(*deadlock.CUserMessageResetHUD) error) {
	c.onCUserMessageResetHUD = append(c.onCUserMessageResetHUD, fn)
}

// OnCUserMessageRumble registers a callback for EBaseUserMessages_UM_Rumble
func (c *Callbacks) OnCUserMessageRumble(fn func(*deadlock.CUserMessageRumble) error) {
	c.onCUserMessageRumble = append(c.onCUserMessageRumble, fn)
}

// OnCUserMessageSayText registers a callback for EBaseUserMessages_UM_SayText
func (c *Callbacks) OnCUserMessageSayText(fn func(*deadlock.CUserMessageSayText) error) {
	c.onCUserMessageSayText = append(c.onCUserMessageSayText, fn)
}

// OnCUserMessageSayText2 registers a callback for EBaseUserMessages_UM_SayText2
func (c *Callbacks) OnCUserMessageSayText2(fn func(*deadlock.CUserMessageSayText2) error) {
	c.onCUserMessageSayText2 = append(c.onCUserMessageSayText2, fn)
}

// OnCUserMessageSayTextChannel registers a callback for EBaseUserMessages_UM_SayTextChannel
func (c *Callbacks) OnCUserMessageSayTextChannel(fn func(*deadlock.CUserMessageSayTextChannel) error) {
	c.onCUserMessageSayTextChannel = append(c.onCUserMessageSayTextChannel, fn)
}

// OnCUserMessageShake registers a callback for EBaseUserMessages_UM_Shake
func (c *Callbacks) OnCUserMessageShake(fn func(*deadlock.CUserMessageShake) error) {
	c.onCUserMessageShake = append(c.onCUserMessageShake, fn)
}

// OnCUserMessageShakeDir registers a callback for EBaseUserMessages_UM_ShakeDir
func (c *Callbacks) OnCUserMessageShakeDir(fn func(*deadlock.CUserMessageShakeDir) error) {
	c.onCUserMessageShakeDir = append(c.onCUserMessageShakeDir, fn)
}

// OnCUserMessageWaterShake registers a callback for EBaseUserMessages_UM_WaterShake
func (c *Callbacks) OnCUserMessageWaterShake(fn func(*deadlock.CUserMessageWaterShake) error) {
	c.onCUserMessageWaterShake = append(c.onCUserMessageWaterShake, fn)
}

// OnCUserMessageTextMsg registers a callback for EBaseUserMessages_UM_TextMsg
func (c *Callbacks) OnCUserMessageTextMsg(fn func(*deadlock.CUserMessageTextMsg) error) {
	c.onCUserMessageTextMsg = append(c.onCUserMessageTextMsg, fn)
}

// OnCUserMessageScreenTilt registers a callback for EBaseUserMessages_UM_ScreenTilt
func (c *Callbacks) OnCUserMessageScreenTilt(fn func(*deadlock.CUserMessageScreenTilt) error) {
	c.onCUserMessageScreenTilt = append(c.onCUserMessageScreenTilt, fn)
}

// OnCUserMessageVoiceMask registers a callback for EBaseUserMessages_UM_VoiceMask
func (c *Callbacks) OnCUserMessageVoiceMask(fn func(*deadlock.CUserMessageVoiceMask) error) {
	c.onCUserMessageVoiceMask = append(c.onCUserMessageVoiceMask, fn)
}

// OnCUserMessageSendAudio registers a callback for EBaseUserMessages_UM_SendAudio
func (c *Callbacks) OnCUserMessageSendAudio(fn func(*deadlock.CUserMessageSendAudio) error) {
	c.onCUserMessageSendAudio = append(c.onCUserMessageSendAudio, fn)
}

// OnCUserMessageItemPickup registers a callback for EBaseUserMessages_UM_ItemPickup
func (c *Callbacks) OnCUserMessageItemPickup(fn func(*deadlock.CUserMessageItemPickup) error) {
	c.onCUserMessageItemPickup = append(c.onCUserMessageItemPickup, fn)
}

// OnCUserMessageAmmoDenied registers a callback for EBaseUserMessages_UM_AmmoDenied
func (c *Callbacks) OnCUserMessageAmmoDenied(fn func(*deadlock.CUserMessageAmmoDenied) error) {
	c.onCUserMessageAmmoDenied = append(c.onCUserMessageAmmoDenied, fn)
}

// OnCUserMessageShowMenu registers a callback for EBaseUserMessages_UM_ShowMenu
func (c *Callbacks) OnCUserMessageShowMenu(fn func(*deadlock.CUserMessageShowMenu) error) {
	c.onCUserMessageShowMenu = append(c.onCUserMessageShowMenu, fn)
}

// OnCUserMessageCreditsMsg registers a callback for EBaseUserMessages_UM_CreditsMsg
func (c *Callbacks) OnCUserMessageCreditsMsg(fn func(*deadlock.CUserMessageCreditsMsg) error) {
	c.onCUserMessageCreditsMsg = append(c.onCUserMessageCreditsMsg, fn)
}

// OnCEntityMessagePlayJingle registers a callback for EBaseEntityMessages_EM_PlayJingle
func (c *Callbacks) OnCEntityMessagePlayJingle(fn func(*deadlock.CEntityMessagePlayJingle) error) {
	c.onCEntityMessagePlayJingle = append(c.onCEntityMessagePlayJingle, fn)
}

// OnCEntityMessageScreenOverlay registers a callback for EBaseEntityMessages_EM_ScreenOverlay
func (c *Callbacks) OnCEntityMessageScreenOverlay(fn func(*deadlock.CEntityMessageScreenOverlay) error) {
	c.onCEntityMessageScreenOverlay = append(c.onCEntityMessageScreenOverlay, fn)
}

// OnCEntityMessageRemoveAllDecals registers a callback for EBaseEntityMessages_EM_RemoveAllDecals
func (c *Callbacks) OnCEntityMessageRemoveAllDecals(fn func(*deadlock.CEntityMessageRemoveAllDecals) error) {
	c.onCEntityMessageRemoveAllDecals = append(c.onCEntityMessageRemoveAllDecals, fn)
}

// OnCEntityMessagePropagateForce registers a callback for EBaseEntityMessages_EM_PropagateForce
func (c *Callbacks) OnCEntityMessagePropagateForce(fn func(*deadlock.CEntityMessagePropagateForce) error) {
	c.onCEntityMessagePropagateForce = append(c.onCEntityMessagePropagateForce, fn)
}

// OnCEntityMessageDoSpark registers a callback for EBaseEntityMessages_EM_DoSpark
func (c *Callbacks) OnCEntityMessageDoSpark(fn func(*deadlock.CEntityMessageDoSpark) error) {
	c.onCEntityMessageDoSpark = append(c.onCEntityMessageDoSpark, fn)
}

// OnCEntityMessageFixAngle registers a callback for EBaseEntityMessages_EM_FixAngle
func (c *Callbacks) OnCEntityMessageFixAngle(fn func(*deadlock.CEntityMessageFixAngle) error) {
	c.onCEntityMessageFixAngle = append(c.onCEntityMessageFixAngle, fn)
}

// OnCUserMessageCloseCaptionPlaceholder registers a callback for EBaseUserMessages_UM_CloseCaptionPlaceholder
func (c *Callbacks) OnCUserMessageCloseCaptionPlaceholder(fn func(*deadlock.CUserMessageCloseCaptionPlaceholder) error) {
	c.onCUserMessageCloseCaptionPlaceholder = append(c.onCUserMessageCloseCaptionPlaceholder, fn)
}

// OnCUserMessageCameraTransition registers a callback for EBaseUserMessages_UM_CameraTransition
func (c *Callbacks) OnCUserMessageCameraTransition(fn func(*deadlock.CUserMessageCameraTransition) error) {
	c.onCUserMessageCameraTransition = append(c.onCUserMessageCameraTransition, fn)
}

// OnCUserMessageAudioParameter registers a callback for EBaseUserMessages_UM_AudioParameter
func (c *Callbacks) OnCUserMessageAudioParameter(fn func(*deadlock.CUserMessageAudioParameter) error) {
	c.onCUserMessageAudioParameter = append(c.onCUserMessageAudioParameter, fn)
}

// OnCUserMessageHapticsManagerPulse registers a callback for EBaseUserMessages_UM_HapticsManagerPulse
func (c *Callbacks) OnCUserMessageHapticsManagerPulse(fn func(*deadlock.CUserMessageHapticsManagerPulse) error) {
	c.onCUserMessageHapticsManagerPulse = append(c.onCUserMessageHapticsManagerPulse, fn)
}

// OnCUserMessageHapticsManagerEffect registers a callback for EBaseUserMessages_UM_HapticsManagerEffect
func (c *Callbacks) OnCUserMessageHapticsManagerEffect(fn func(*deadlock.CUserMessageHapticsManagerEffect) error) {
	c.onCUserMessageHapticsManagerEffect = append(c.onCUserMessageHapticsManagerEffect, fn)
}

// OnCUserMessageUpdateCssClasses registers a callback for EBaseUserMessages_UM_UpdateCssClasses
func (c *Callbacks) OnCUserMessageUpdateCssClasses(fn func(*deadlock.CUserMessageUpdateCssClasses) error) {
	c.onCUserMessageUpdateCssClasses = append(c.onCUserMessageUpdateCssClasses, fn)
}

// OnCUserMessageServerFrameTime registers a callback for EBaseUserMessages_UM_ServerFrameTime
func (c *Callbacks) OnCUserMessageServerFrameTime(fn func(*deadlock.CUserMessageServerFrameTime) error) {
	c.onCUserMessageServerFrameTime = append(c.onCUserMessageServerFrameTime, fn)
}

// OnCUserMessageLagCompensationError registers a callback for EBaseUserMessages_UM_LagCompensationError
func (c *Callbacks) OnCUserMessageLagCompensationError(fn func(*deadlock.CUserMessageLagCompensationError) error) {
	c.onCUserMessageLagCompensationError = append(c.onCUserMessageLagCompensationError, fn)
}

// OnCUserMessageRequestDllStatus registers a callback for EBaseUserMessages_UM_RequestDllStatus
func (c *Callbacks) OnCUserMessageRequestDllStatus(fn func(*deadlock.CUserMessageRequestDllStatus) error) {
	c.onCUserMessageRequestDllStatus = append(c.onCUserMessageRequestDllStatus, fn)
}

// OnCUserMessageRequestUtilAction registers a callback for EBaseUserMessages_UM_RequestUtilAction
func (c *Callbacks) OnCUserMessageRequestUtilAction(fn func(*deadlock.CUserMessageRequestUtilAction) error) {
	c.onCUserMessageRequestUtilAction = append(c.onCUserMessageRequestUtilAction, fn)
}

// OnCUserMessageRequestInventory registers a callback for EBaseUserMessages_UM_RequestInventory
func (c *Callbacks) OnCUserMessageRequestInventory(fn func(*deadlock.CUserMessageRequestInventory) error) {
	c.onCUserMessageRequestInventory = append(c.onCUserMessageRequestInventory, fn)
}

// OnCUserMessageRequestDiagnostic registers a callback for EBaseUserMessages_UM_RequestDiagnostic
func (c *Callbacks) OnCUserMessageRequestDiagnostic(fn func(*deadlock.CUserMessageRequestDiagnostic) error) {
	c.onCUserMessageRequestDiagnostic = append(c.onCUserMessageRequestDiagnostic, fn)
}

// OnCMsgVDebugGameSessionIDEvent registers a callback for EBaseGameEvents_GE_VDebugGameSessionIDEvent
func (c *Callbacks) OnCMsgVDebugGameSessionIDEvent(fn func(*deadlock.CMsgVDebugGameSessionIDEvent) error) {
	c.onCMsgVDebugGameSessionIDEvent = append(c.onCMsgVDebugGameSessionIDEvent, fn)
}

// OnCMsgPlaceDecalEvent registers a callback for EBaseGameEvents_GE_PlaceDecalEvent
func (c *Callbacks) OnCMsgPlaceDecalEvent(fn func(*deadlock.CMsgPlaceDecalEvent) error) {
	c.onCMsgPlaceDecalEvent = append(c.onCMsgPlaceDecalEvent, fn)
}

// OnCMsgClearWorldDecalsEvent registers a callback for EBaseGameEvents_GE_ClearWorldDecalsEvent
func (c *Callbacks) OnCMsgClearWorldDecalsEvent(fn func(*deadlock.CMsgClearWorldDecalsEvent) error) {
	c.onCMsgClearWorldDecalsEvent = append(c.onCMsgClearWorldDecalsEvent, fn)
}

// OnCMsgClearEntityDecalsEvent registers a callback for EBaseGameEvents_GE_ClearEntityDecalsEvent
func (c *Callbacks) OnCMsgClearEntityDecalsEvent(fn func(*deadlock.CMsgClearEntityDecalsEvent) error) {
	c.onCMsgClearEntityDecalsEvent = append(c.onCMsgClearEntityDecalsEvent, fn)
}

// OnCMsgClearDecalsForSkeletonInstanceEvent registers a callback for EBaseGameEvents_GE_ClearDecalsForSkeletonInstanceEvent
func (c *Callbacks) OnCMsgClearDecalsForSkeletonInstanceEvent(fn func(*deadlock.CMsgClearDecalsForSkeletonInstanceEvent) error) {
	c.onCMsgClearDecalsForSkeletonInstanceEvent = append(c.onCMsgClearDecalsForSkeletonInstanceEvent, fn)
}

// OnCMsgSource1LegacyGameEventList registers a callback for EBaseGameEvents_GE_Source1LegacyGameEventList
func (c *Callbacks) OnCMsgSource1LegacyGameEventList(fn func(*deadlock.CMsgSource1LegacyGameEventList) error) {
	c.onCMsgSource1LegacyGameEventList = append(c.onCMsgSource1LegacyGameEventList, fn)
}

// OnCMsgSource1LegacyListenEvents registers a callback for EBaseGameEvents_GE_Source1LegacyListenEvents
func (c *Callbacks) OnCMsgSource1LegacyListenEvents(fn func(*deadlock.CMsgSource1LegacyListenEvents) error) {
	c.onCMsgSource1LegacyListenEvents = append(c.onCMsgSource1LegacyListenEvents, fn)
}

// OnCMsgSource1LegacyGameEvent registers a callback for EBaseGameEvents_GE_Source1LegacyGameEvent
func (c *Callbacks) OnCMsgSource1LegacyGameEvent(fn func(*deadlock.CMsgSource1LegacyGameEvent) error) {
	c.onCMsgSource1LegacyGameEvent = append(c.onCMsgSource1LegacyGameEvent, fn)
}

// OnCMsgSosStartSoundEvent registers a callback for EBaseGameEvents_GE_SosStartSoundEvent
func (c *Callbacks) OnCMsgSosStartSoundEvent(fn func(*deadlock.CMsgSosStartSoundEvent) error) {
	c.onCMsgSosStartSoundEvent = append(c.onCMsgSosStartSoundEvent, fn)
}

// OnCMsgSosStopSoundEvent registers a callback for EBaseGameEvents_GE_SosStopSoundEvent
func (c *Callbacks) OnCMsgSosStopSoundEvent(fn func(*deadlock.CMsgSosStopSoundEvent) error) {
	c.onCMsgSosStopSoundEvent = append(c.onCMsgSosStopSoundEvent, fn)
}

// OnCMsgSosSetSoundEventParams registers a callback for EBaseGameEvents_GE_SosSetSoundEventParams
func (c *Callbacks) OnCMsgSosSetSoundEventParams(fn func(*deadlock.CMsgSosSetSoundEventParams) error) {
	c.onCMsgSosSetSoundEventParams = append(c.onCMsgSosSetSoundEventParams, fn)
}

// OnCMsgSosSetLibraryStackFields registers a callback for EBaseGameEvents_GE_SosSetLibraryStackFields
func (c *Callbacks) OnCMsgSosSetLibraryStackFields(fn func(*deadlock.CMsgSosSetLibraryStackFields) error) {
	c.onCMsgSosSetLibraryStackFields = append(c.onCMsgSosSetLibraryStackFields, fn)
}

// OnCMsgSosStopSoundEventHash registers a callback for EBaseGameEvents_GE_SosStopSoundEventHash
func (c *Callbacks) OnCMsgSosStopSoundEventHash(fn func(*deadlock.CMsgSosStopSoundEventHash) error) {
	c.onCMsgSosStopSoundEventHash = append(c.onCMsgSosStopSoundEventHash, fn)
}

func (c *Callbacks) callByDemoType(t int32, buf []byte) error {
	switch t {
	case 0: // deadlock.EDemoCommands_DEM_Stop
		if c.onCDemoStop == nil {
			return nil
		}

		msg := &deadlock.CDemoStop{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoStop {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 1: // deadlock.EDemoCommands_DEM_FileHeader
		if c.onCDemoFileHeader == nil {
			return nil
		}

		msg := &deadlock.CDemoFileHeader{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoFileHeader {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 2: // deadlock.EDemoCommands_DEM_FileInfo
		if c.onCDemoFileInfo == nil {
			return nil
		}

		msg := &deadlock.CDemoFileInfo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoFileInfo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 3: // deadlock.EDemoCommands_DEM_SyncTick
		if c.onCDemoSyncTick == nil {
			return nil
		}

		msg := &deadlock.CDemoSyncTick{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoSyncTick {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 4: // deadlock.EDemoCommands_DEM_SendTables
		if c.onCDemoSendTables == nil {
			return nil
		}

		msg := &deadlock.CDemoSendTables{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoSendTables {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 5: // deadlock.EDemoCommands_DEM_ClassInfo
		if c.onCDemoClassInfo == nil {
			return nil
		}

		msg := &deadlock.CDemoClassInfo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoClassInfo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 6: // deadlock.EDemoCommands_DEM_StringTables
		if c.onCDemoStringTables == nil {
			return nil
		}

		msg := &deadlock.CDemoStringTables{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoStringTables {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 7: // deadlock.EDemoCommands_DEM_Packet
		if c.onCDemoPacket == nil {
			return nil
		}

		msg := &deadlock.CDemoPacket{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoPacket {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 8: // deadlock.EDemoCommands_DEM_SignonPacket
		if c.onCDemoSignonPacket == nil {
			return nil
		}

		msg := &deadlock.CDemoPacket{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoSignonPacket {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 9: // deadlock.EDemoCommands_DEM_ConsoleCmd
		if c.onCDemoConsoleCmd == nil {
			return nil
		}

		msg := &deadlock.CDemoConsoleCmd{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoConsoleCmd {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 10: // deadlock.EDemoCommands_DEM_CustomData
		if c.onCDemoCustomData == nil {
			return nil
		}

		msg := &deadlock.CDemoCustomData{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoCustomData {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 11: // deadlock.EDemoCommands_DEM_CustomDataCallbacks
		if c.onCDemoCustomDataCallbacks == nil {
			return nil
		}

		msg := &deadlock.CDemoCustomDataCallbacks{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoCustomDataCallbacks {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 12: // deadlock.EDemoCommands_DEM_UserCmd
		if c.onCDemoUserCmd == nil {
			return nil
		}

		msg := &deadlock.CDemoUserCmd{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoUserCmd {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 13: // deadlock.EDemoCommands_DEM_FullPacket
		if c.onCDemoFullPacket == nil {
			return nil
		}

		msg := &deadlock.CDemoFullPacket{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoFullPacket {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 14: // deadlock.EDemoCommands_DEM_SaveGame
		if c.onCDemoSaveGame == nil {
			return nil
		}

		msg := &deadlock.CDemoSaveGame{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoSaveGame {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 15: // deadlock.EDemoCommands_DEM_SpawnGroups
		if c.onCDemoSpawnGroups == nil {
			return nil
		}

		msg := &deadlock.CDemoSpawnGroups{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoSpawnGroups {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 16: // deadlock.EDemoCommands_DEM_AnimationData
		if c.onCDemoAnimationData == nil {
			return nil
		}

		msg := &deadlock.CDemoAnimationData{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoAnimationData {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 17: // deadlock.EDemoCommands_DEM_AnimationHeader
		if c.onCDemoAnimationHeader == nil {
			return nil
		}

		msg := &deadlock.CDemoAnimationHeader{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoAnimationHeader {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	}

	if v(1) {
		_debugf("warning: no demo type %d found", t)
	}

	return nil
}

func (c *Callbacks) callByPacketType(t int32, buf []byte) error {
	switch t {
	case 0: // deadlock.NET_Messages_net_NOP
		if c.onCNETMsg_NOP == nil {
			return nil
		}

		msg := &deadlock.CNETMsg_NOP{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_NOP {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 3: // deadlock.NET_Messages_net_SplitScreenUser
		if c.onCNETMsg_SplitScreenUser == nil {
			return nil
		}

		msg := &deadlock.CNETMsg_SplitScreenUser{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SplitScreenUser {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 4: // deadlock.NET_Messages_net_Tick
		if c.onCNETMsg_Tick == nil {
			return nil
		}

		msg := &deadlock.CNETMsg_Tick{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_Tick {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 5: // deadlock.NET_Messages_net_StringCmd
		if c.onCNETMsg_StringCmd == nil {
			return nil
		}

		msg := &deadlock.CNETMsg_StringCmd{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_StringCmd {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 6: // deadlock.NET_Messages_net_SetConVar
		if c.onCNETMsg_SetConVar == nil {
			return nil
		}

		msg := &deadlock.CNETMsg_SetConVar{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SetConVar {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 7: // deadlock.NET_Messages_net_SignonState
		if c.onCNETMsg_SignonState == nil {
			return nil
		}

		msg := &deadlock.CNETMsg_SignonState{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SignonState {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 8: // deadlock.NET_Messages_net_SpawnGroup_Load
		if c.onCNETMsg_SpawnGroup_Load == nil {
			return nil
		}

		msg := &deadlock.CNETMsg_SpawnGroup_Load{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SpawnGroup_Load {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 9: // deadlock.NET_Messages_net_SpawnGroup_ManifestUpdate
		if c.onCNETMsg_SpawnGroup_ManifestUpdate == nil {
			return nil
		}

		msg := &deadlock.CNETMsg_SpawnGroup_ManifestUpdate{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SpawnGroup_ManifestUpdate {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 11: // deadlock.NET_Messages_net_SpawnGroup_SetCreationTick
		if c.onCNETMsg_SpawnGroup_SetCreationTick == nil {
			return nil
		}

		msg := &deadlock.CNETMsg_SpawnGroup_SetCreationTick{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SpawnGroup_SetCreationTick {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 12: // deadlock.NET_Messages_net_SpawnGroup_Unload
		if c.onCNETMsg_SpawnGroup_Unload == nil {
			return nil
		}

		msg := &deadlock.CNETMsg_SpawnGroup_Unload{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SpawnGroup_Unload {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 13: // deadlock.NET_Messages_net_SpawnGroup_LoadCompleted
		if c.onCNETMsg_SpawnGroup_LoadCompleted == nil {
			return nil
		}

		msg := &deadlock.CNETMsg_SpawnGroup_LoadCompleted{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SpawnGroup_LoadCompleted {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 15: // deadlock.NET_Messages_net_DebugOverlay
		if c.onCNETMsg_DebugOverlay == nil {
			return nil
		}

		msg := &deadlock.CNETMsg_DebugOverlay{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_DebugOverlay {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 40: // deadlock.SVC_Messages_svc_ServerInfo
		if c.onCSVCMsg_ServerInfo == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_ServerInfo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_ServerInfo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 41: // deadlock.SVC_Messages_svc_FlattenedSerializer
		if c.onCSVCMsg_FlattenedSerializer == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_FlattenedSerializer{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_FlattenedSerializer {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 42: // deadlock.SVC_Messages_svc_ClassInfo
		if c.onCSVCMsg_ClassInfo == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_ClassInfo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_ClassInfo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 43: // deadlock.SVC_Messages_svc_SetPause
		if c.onCSVCMsg_SetPause == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_SetPause{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_SetPause {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 44: // deadlock.SVC_Messages_svc_CreateStringTable
		if c.onCSVCMsg_CreateStringTable == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_CreateStringTable{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_CreateStringTable {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 45: // deadlock.SVC_Messages_svc_UpdateStringTable
		if c.onCSVCMsg_UpdateStringTable == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_UpdateStringTable{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_UpdateStringTable {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 46: // deadlock.SVC_Messages_svc_VoiceInit
		if c.onCSVCMsg_VoiceInit == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_VoiceInit{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_VoiceInit {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 47: // deadlock.SVC_Messages_svc_VoiceData
		if c.onCSVCMsg_VoiceData == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_VoiceData{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_VoiceData {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 48: // deadlock.SVC_Messages_svc_Print
		if c.onCSVCMsg_Print == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_Print{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_Print {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 49: // deadlock.SVC_Messages_svc_Sounds
		if c.onCSVCMsg_Sounds == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_Sounds{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_Sounds {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 50: // deadlock.SVC_Messages_svc_SetView
		if c.onCSVCMsg_SetView == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_SetView{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_SetView {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 51: // deadlock.SVC_Messages_svc_ClearAllStringTables
		if c.onCSVCMsg_ClearAllStringTables == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_ClearAllStringTables{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_ClearAllStringTables {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 52: // deadlock.SVC_Messages_svc_CmdKeyValues
		if c.onCSVCMsg_CmdKeyValues == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_CmdKeyValues{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_CmdKeyValues {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 53: // deadlock.SVC_Messages_svc_BSPDecal
		if c.onCSVCMsg_BSPDecal == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_BSPDecal{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_BSPDecal {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 54: // deadlock.SVC_Messages_svc_SplitScreen
		if c.onCSVCMsg_SplitScreen == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_SplitScreen{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_SplitScreen {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 55: // deadlock.SVC_Messages_svc_PacketEntities
		if c.onCSVCMsg_PacketEntities == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_PacketEntities{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_PacketEntities {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 56: // deadlock.SVC_Messages_svc_Prefetch
		if c.onCSVCMsg_Prefetch == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_Prefetch{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_Prefetch {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 57: // deadlock.SVC_Messages_svc_Menu
		if c.onCSVCMsg_Menu == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_Menu{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_Menu {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 58: // deadlock.SVC_Messages_svc_GetCvarValue
		if c.onCSVCMsg_GetCvarValue == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_GetCvarValue{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_GetCvarValue {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 59: // deadlock.SVC_Messages_svc_StopSound
		if c.onCSVCMsg_StopSound == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_StopSound{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_StopSound {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 60: // deadlock.SVC_Messages_svc_PeerList
		if c.onCSVCMsg_PeerList == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_PeerList{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_PeerList {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 61: // deadlock.SVC_Messages_svc_PacketReliable
		if c.onCSVCMsg_PacketReliable == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_PacketReliable{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_PacketReliable {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 62: // deadlock.SVC_Messages_svc_HLTVStatus
		if c.onCSVCMsg_HLTVStatus == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_HLTVStatus{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_HLTVStatus {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 63: // deadlock.SVC_Messages_svc_ServerSteamID
		if c.onCSVCMsg_ServerSteamID == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_ServerSteamID{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_ServerSteamID {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 70: // deadlock.SVC_Messages_svc_FullFrameSplit
		if c.onCSVCMsg_FullFrameSplit == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_FullFrameSplit{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_FullFrameSplit {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 71: // deadlock.SVC_Messages_svc_RconServerDetails
		if c.onCSVCMsg_RconServerDetails == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_RconServerDetails{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_RconServerDetails {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 72: // deadlock.SVC_Messages_svc_UserMessage
		if c.onCSVCMsg_UserMessage == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_UserMessage{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_UserMessage {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 74: // deadlock.SVC_Messages_svc_Broadcast_Command
		if c.onCSVCMsg_Broadcast_Command == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_Broadcast_Command{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_Broadcast_Command {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 75: // deadlock.SVC_Messages_svc_HltvFixupOperatorStatus
		if c.onCSVCMsg_HltvFixupOperatorStatus == nil {
			return nil
		}

		msg := &deadlock.CSVCMsg_HltvFixupOperatorStatus{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_HltvFixupOperatorStatus {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 101: // deadlock.EBaseUserMessages_UM_AchievementEvent
		if c.onCUserMessageAchievementEvent == nil {
			return nil
		}

		msg := &deadlock.CUserMessageAchievementEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageAchievementEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 102: // deadlock.EBaseUserMessages_UM_CloseCaption
		if c.onCUserMessageCloseCaption == nil {
			return nil
		}

		msg := &deadlock.CUserMessageCloseCaption{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCloseCaption {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 103: // deadlock.EBaseUserMessages_UM_CloseCaptionDirect
		if c.onCUserMessageCloseCaptionDirect == nil {
			return nil
		}

		msg := &deadlock.CUserMessageCloseCaptionDirect{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCloseCaptionDirect {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 104: // deadlock.EBaseUserMessages_UM_CurrentTimescale
		if c.onCUserMessageCurrentTimescale == nil {
			return nil
		}

		msg := &deadlock.CUserMessageCurrentTimescale{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCurrentTimescale {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 105: // deadlock.EBaseUserMessages_UM_DesiredTimescale
		if c.onCUserMessageDesiredTimescale == nil {
			return nil
		}

		msg := &deadlock.CUserMessageDesiredTimescale{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageDesiredTimescale {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 106: // deadlock.EBaseUserMessages_UM_Fade
		if c.onCUserMessageFade == nil {
			return nil
		}

		msg := &deadlock.CUserMessageFade{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageFade {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 107: // deadlock.EBaseUserMessages_UM_GameTitle
		if c.onCUserMessageGameTitle == nil {
			return nil
		}

		msg := &deadlock.CUserMessageGameTitle{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageGameTitle {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 110: // deadlock.EBaseUserMessages_UM_HudMsg
		if c.onCUserMessageHudMsg == nil {
			return nil
		}

		msg := &deadlock.CUserMessageHudMsg{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageHudMsg {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 111: // deadlock.EBaseUserMessages_UM_HudText
		if c.onCUserMessageHudText == nil {
			return nil
		}

		msg := &deadlock.CUserMessageHudText{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageHudText {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 113: // deadlock.EBaseUserMessages_UM_ColoredText
		if c.onCUserMessageColoredText == nil {
			return nil
		}

		msg := &deadlock.CUserMessageColoredText{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageColoredText {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 114: // deadlock.EBaseUserMessages_UM_RequestState
		if c.onCUserMessageRequestState == nil {
			return nil
		}

		msg := &deadlock.CUserMessageRequestState{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRequestState {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 115: // deadlock.EBaseUserMessages_UM_ResetHUD
		if c.onCUserMessageResetHUD == nil {
			return nil
		}

		msg := &deadlock.CUserMessageResetHUD{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageResetHUD {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 116: // deadlock.EBaseUserMessages_UM_Rumble
		if c.onCUserMessageRumble == nil {
			return nil
		}

		msg := &deadlock.CUserMessageRumble{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRumble {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 117: // deadlock.EBaseUserMessages_UM_SayText
		if c.onCUserMessageSayText == nil {
			return nil
		}

		msg := &deadlock.CUserMessageSayText{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageSayText {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 118: // deadlock.EBaseUserMessages_UM_SayText2
		if c.onCUserMessageSayText2 == nil {
			return nil
		}

		msg := &deadlock.CUserMessageSayText2{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageSayText2 {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 119: // deadlock.EBaseUserMessages_UM_SayTextChannel
		if c.onCUserMessageSayTextChannel == nil {
			return nil
		}

		msg := &deadlock.CUserMessageSayTextChannel{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageSayTextChannel {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 120: // deadlock.EBaseUserMessages_UM_Shake
		if c.onCUserMessageShake == nil {
			return nil
		}

		msg := &deadlock.CUserMessageShake{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageShake {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 121: // deadlock.EBaseUserMessages_UM_ShakeDir
		if c.onCUserMessageShakeDir == nil {
			return nil
		}

		msg := &deadlock.CUserMessageShakeDir{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageShakeDir {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 122: // deadlock.EBaseUserMessages_UM_WaterShake
		if c.onCUserMessageWaterShake == nil {
			return nil
		}

		msg := &deadlock.CUserMessageWaterShake{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageWaterShake {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 124: // deadlock.EBaseUserMessages_UM_TextMsg
		if c.onCUserMessageTextMsg == nil {
			return nil
		}

		msg := &deadlock.CUserMessageTextMsg{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageTextMsg {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 125: // deadlock.EBaseUserMessages_UM_ScreenTilt
		if c.onCUserMessageScreenTilt == nil {
			return nil
		}

		msg := &deadlock.CUserMessageScreenTilt{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageScreenTilt {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 128: // deadlock.EBaseUserMessages_UM_VoiceMask
		if c.onCUserMessageVoiceMask == nil {
			return nil
		}

		msg := &deadlock.CUserMessageVoiceMask{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageVoiceMask {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 130: // deadlock.EBaseUserMessages_UM_SendAudio
		if c.onCUserMessageSendAudio == nil {
			return nil
		}

		msg := &deadlock.CUserMessageSendAudio{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageSendAudio {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 131: // deadlock.EBaseUserMessages_UM_ItemPickup
		if c.onCUserMessageItemPickup == nil {
			return nil
		}

		msg := &deadlock.CUserMessageItemPickup{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageItemPickup {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 132: // deadlock.EBaseUserMessages_UM_AmmoDenied
		if c.onCUserMessageAmmoDenied == nil {
			return nil
		}

		msg := &deadlock.CUserMessageAmmoDenied{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageAmmoDenied {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 134: // deadlock.EBaseUserMessages_UM_ShowMenu
		if c.onCUserMessageShowMenu == nil {
			return nil
		}

		msg := &deadlock.CUserMessageShowMenu{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageShowMenu {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 135: // deadlock.EBaseUserMessages_UM_CreditsMsg
		if c.onCUserMessageCreditsMsg == nil {
			return nil
		}

		msg := &deadlock.CUserMessageCreditsMsg{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCreditsMsg {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 136: // deadlock.EBaseEntityMessages_EM_PlayJingle
		if c.onCEntityMessagePlayJingle == nil {
			return nil
		}

		msg := &deadlock.CEntityMessagePlayJingle{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessagePlayJingle {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 137: // deadlock.EBaseEntityMessages_EM_ScreenOverlay
		if c.onCEntityMessageScreenOverlay == nil {
			return nil
		}

		msg := &deadlock.CEntityMessageScreenOverlay{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessageScreenOverlay {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 138: // deadlock.EBaseEntityMessages_EM_RemoveAllDecals
		if c.onCEntityMessageRemoveAllDecals == nil {
			return nil
		}

		msg := &deadlock.CEntityMessageRemoveAllDecals{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessageRemoveAllDecals {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 139: // deadlock.EBaseEntityMessages_EM_PropagateForce
		if c.onCEntityMessagePropagateForce == nil {
			return nil
		}

		msg := &deadlock.CEntityMessagePropagateForce{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessagePropagateForce {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 140: // deadlock.EBaseEntityMessages_EM_DoSpark
		if c.onCEntityMessageDoSpark == nil {
			return nil
		}

		msg := &deadlock.CEntityMessageDoSpark{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessageDoSpark {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 141: // deadlock.EBaseEntityMessages_EM_FixAngle
		if c.onCEntityMessageFixAngle == nil {
			return nil
		}

		msg := &deadlock.CEntityMessageFixAngle{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessageFixAngle {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 142: // deadlock.EBaseUserMessages_UM_CloseCaptionPlaceholder
		if c.onCUserMessageCloseCaptionPlaceholder == nil {
			return nil
		}

		msg := &deadlock.CUserMessageCloseCaptionPlaceholder{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCloseCaptionPlaceholder {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 143: // deadlock.EBaseUserMessages_UM_CameraTransition
		if c.onCUserMessageCameraTransition == nil {
			return nil
		}

		msg := &deadlock.CUserMessageCameraTransition{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCameraTransition {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 144: // deadlock.EBaseUserMessages_UM_AudioParameter
		if c.onCUserMessageAudioParameter == nil {
			return nil
		}

		msg := &deadlock.CUserMessageAudioParameter{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageAudioParameter {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 150: // deadlock.EBaseUserMessages_UM_HapticsManagerPulse
		if c.onCUserMessageHapticsManagerPulse == nil {
			return nil
		}

		msg := &deadlock.CUserMessageHapticsManagerPulse{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageHapticsManagerPulse {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 151: // deadlock.EBaseUserMessages_UM_HapticsManagerEffect
		if c.onCUserMessageHapticsManagerEffect == nil {
			return nil
		}

		msg := &deadlock.CUserMessageHapticsManagerEffect{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageHapticsManagerEffect {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 153: // deadlock.EBaseUserMessages_UM_UpdateCssClasses
		if c.onCUserMessageUpdateCssClasses == nil {
			return nil
		}

		msg := &deadlock.CUserMessageUpdateCssClasses{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageUpdateCssClasses {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 154: // deadlock.EBaseUserMessages_UM_ServerFrameTime
		if c.onCUserMessageServerFrameTime == nil {
			return nil
		}

		msg := &deadlock.CUserMessageServerFrameTime{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageServerFrameTime {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 155: // deadlock.EBaseUserMessages_UM_LagCompensationError
		if c.onCUserMessageLagCompensationError == nil {
			return nil
		}

		msg := &deadlock.CUserMessageLagCompensationError{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageLagCompensationError {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 156: // deadlock.EBaseUserMessages_UM_RequestDllStatus
		if c.onCUserMessageRequestDllStatus == nil {
			return nil
		}

		msg := &deadlock.CUserMessageRequestDllStatus{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRequestDllStatus {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 157: // deadlock.EBaseUserMessages_UM_RequestUtilAction
		if c.onCUserMessageRequestUtilAction == nil {
			return nil
		}

		msg := &deadlock.CUserMessageRequestUtilAction{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRequestUtilAction {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 160: // deadlock.EBaseUserMessages_UM_RequestInventory
		if c.onCUserMessageRequestInventory == nil {
			return nil
		}

		msg := &deadlock.CUserMessageRequestInventory{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRequestInventory {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 162: // deadlock.EBaseUserMessages_UM_RequestDiagnostic
		if c.onCUserMessageRequestDiagnostic == nil {
			return nil
		}

		msg := &deadlock.CUserMessageRequestDiagnostic{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRequestDiagnostic {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 200: // deadlock.EBaseGameEvents_GE_VDebugGameSessionIDEvent
		if c.onCMsgVDebugGameSessionIDEvent == nil {
			return nil
		}

		msg := &deadlock.CMsgVDebugGameSessionIDEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgVDebugGameSessionIDEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 201: // deadlock.EBaseGameEvents_GE_PlaceDecalEvent
		if c.onCMsgPlaceDecalEvent == nil {
			return nil
		}

		msg := &deadlock.CMsgPlaceDecalEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgPlaceDecalEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 202: // deadlock.EBaseGameEvents_GE_ClearWorldDecalsEvent
		if c.onCMsgClearWorldDecalsEvent == nil {
			return nil
		}

		msg := &deadlock.CMsgClearWorldDecalsEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgClearWorldDecalsEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 203: // deadlock.EBaseGameEvents_GE_ClearEntityDecalsEvent
		if c.onCMsgClearEntityDecalsEvent == nil {
			return nil
		}

		msg := &deadlock.CMsgClearEntityDecalsEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgClearEntityDecalsEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 204: // deadlock.EBaseGameEvents_GE_ClearDecalsForSkeletonInstanceEvent
		if c.onCMsgClearDecalsForSkeletonInstanceEvent == nil {
			return nil
		}

		msg := &deadlock.CMsgClearDecalsForSkeletonInstanceEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgClearDecalsForSkeletonInstanceEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 205: // deadlock.EBaseGameEvents_GE_Source1LegacyGameEventList
		if c.onCMsgSource1LegacyGameEventList == nil {
			return nil
		}

		msg := &deadlock.CMsgSource1LegacyGameEventList{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSource1LegacyGameEventList {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 206: // deadlock.EBaseGameEvents_GE_Source1LegacyListenEvents
		if c.onCMsgSource1LegacyListenEvents == nil {
			return nil
		}

		msg := &deadlock.CMsgSource1LegacyListenEvents{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSource1LegacyListenEvents {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 207: // deadlock.EBaseGameEvents_GE_Source1LegacyGameEvent
		if c.onCMsgSource1LegacyGameEvent == nil {
			return nil
		}

		msg := &deadlock.CMsgSource1LegacyGameEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSource1LegacyGameEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 208: // deadlock.EBaseGameEvents_GE_SosStartSoundEvent
		if c.onCMsgSosStartSoundEvent == nil {
			return nil
		}

		msg := &deadlock.CMsgSosStartSoundEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSosStartSoundEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 209: // deadlock.EBaseGameEvents_GE_SosStopSoundEvent
		if c.onCMsgSosStopSoundEvent == nil {
			return nil
		}

		msg := &deadlock.CMsgSosStopSoundEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSosStopSoundEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 210: // deadlock.EBaseGameEvents_GE_SosSetSoundEventParams
		if c.onCMsgSosSetSoundEventParams == nil {
			return nil
		}

		msg := &deadlock.CMsgSosSetSoundEventParams{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSosSetSoundEventParams {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 211: // deadlock.EBaseGameEvents_GE_SosSetLibraryStackFields
		if c.onCMsgSosSetLibraryStackFields == nil {
			return nil
		}

		msg := &deadlock.CMsgSosSetLibraryStackFields{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSosSetLibraryStackFields {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 212: // deadlock.EBaseGameEvents_GE_SosStopSoundEventHash
		if c.onCMsgSosStopSoundEventHash == nil {
			return nil
		}

		msg := &deadlock.CMsgSosStopSoundEventHash{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSosStopSoundEventHash {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	}

	return nil
}
