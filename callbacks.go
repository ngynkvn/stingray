package stingray

import (
	"github.com/ngynkvn/stingray/deadlock"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"strconv"
)

// Callbacks decodes and routes replay events to callback functions
type Callbacks struct {
	onCDemoStop                                        []func(*deadlock.CDemoStop) error
	onCDemoFileHeader                                  []func(*deadlock.CDemoFileHeader) error
	onCDemoFileInfo                                    []func(*deadlock.CDemoFileInfo) error
	onCDemoSyncTick                                    []func(*deadlock.CDemoSyncTick) error
	onCDemoSendTables                                  []func(*deadlock.CDemoSendTables) error
	onCDemoClassInfo                                   []func(*deadlock.CDemoClassInfo) error
	onCDemoStringTables                                []func(*deadlock.CDemoStringTables) error
	onCDemoPacket                                      []func(*deadlock.CDemoPacket) error
	onCDemoSignonPacket                                []func(*deadlock.CDemoPacket) error
	onCDemoConsoleCmd                                  []func(*deadlock.CDemoConsoleCmd) error
	onCDemoCustomData                                  []func(*deadlock.CDemoCustomData) error
	onCDemoCustomDataCallbacks                         []func(*deadlock.CDemoCustomDataCallbacks) error
	onCDemoUserCmd                                     []func(*deadlock.CDemoUserCmd) error
	onCDemoFullPacket                                  []func(*deadlock.CDemoFullPacket) error
	onCDemoSaveGame                                    []func(*deadlock.CDemoSaveGame) error
	onCDemoSpawnGroups                                 []func(*deadlock.CDemoSpawnGroups) error
	onCDemoAnimationData                               []func(*deadlock.CDemoAnimationData) error
	onCDemoAnimationHeader                             []func(*deadlock.CDemoAnimationHeader) error
	onCNETMsg_NOP                                      []func(*deadlock.CNETMsg_NOP) error
	onCNETMsg_SplitScreenUser                          []func(*deadlock.CNETMsg_SplitScreenUser) error
	onCNETMsg_Tick                                     []func(*deadlock.CNETMsg_Tick) error
	onCNETMsg_StringCmd                                []func(*deadlock.CNETMsg_StringCmd) error
	onCNETMsg_SetConVar                                []func(*deadlock.CNETMsg_SetConVar) error
	onCNETMsg_SignonState                              []func(*deadlock.CNETMsg_SignonState) error
	onCNETMsg_SpawnGroup_Load                          []func(*deadlock.CNETMsg_SpawnGroup_Load) error
	onCNETMsg_SpawnGroup_ManifestUpdate                []func(*deadlock.CNETMsg_SpawnGroup_ManifestUpdate) error
	onCNETMsg_SpawnGroup_SetCreationTick               []func(*deadlock.CNETMsg_SpawnGroup_SetCreationTick) error
	onCNETMsg_SpawnGroup_Unload                        []func(*deadlock.CNETMsg_SpawnGroup_Unload) error
	onCNETMsg_SpawnGroup_LoadCompleted                 []func(*deadlock.CNETMsg_SpawnGroup_LoadCompleted) error
	onCNETMsg_DebugOverlay                             []func(*deadlock.CNETMsg_DebugOverlay) error
	onCSVCMsg_ServerInfo                               []func(*deadlock.CSVCMsg_ServerInfo) error
	onCSVCMsg_FlattenedSerializer                      []func(*deadlock.CSVCMsg_FlattenedSerializer) error
	onCSVCMsg_ClassInfo                                []func(*deadlock.CSVCMsg_ClassInfo) error
	onCSVCMsg_SetPause                                 []func(*deadlock.CSVCMsg_SetPause) error
	onCSVCMsg_CreateStringTable                        []func(*deadlock.CSVCMsg_CreateStringTable) error
	onCSVCMsg_UpdateStringTable                        []func(*deadlock.CSVCMsg_UpdateStringTable) error
	onCSVCMsg_VoiceInit                                []func(*deadlock.CSVCMsg_VoiceInit) error
	onCSVCMsg_VoiceData                                []func(*deadlock.CSVCMsg_VoiceData) error
	onCSVCMsg_Print                                    []func(*deadlock.CSVCMsg_Print) error
	onCSVCMsg_Sounds                                   []func(*deadlock.CSVCMsg_Sounds) error
	onCSVCMsg_SetView                                  []func(*deadlock.CSVCMsg_SetView) error
	onCSVCMsg_ClearAllStringTables                     []func(*deadlock.CSVCMsg_ClearAllStringTables) error
	onCSVCMsg_CmdKeyValues                             []func(*deadlock.CSVCMsg_CmdKeyValues) error
	onCSVCMsg_BSPDecal                                 []func(*deadlock.CSVCMsg_BSPDecal) error
	onCSVCMsg_SplitScreen                              []func(*deadlock.CSVCMsg_SplitScreen) error
	onCSVCMsg_PacketEntities                           []func(*deadlock.CSVCMsg_PacketEntities) error
	onCSVCMsg_Prefetch                                 []func(*deadlock.CSVCMsg_Prefetch) error
	onCSVCMsg_Menu                                     []func(*deadlock.CSVCMsg_Menu) error
	onCSVCMsg_GetCvarValue                             []func(*deadlock.CSVCMsg_GetCvarValue) error
	onCSVCMsg_StopSound                                []func(*deadlock.CSVCMsg_StopSound) error
	onCSVCMsg_PeerList                                 []func(*deadlock.CSVCMsg_PeerList) error
	onCSVCMsg_PacketReliable                           []func(*deadlock.CSVCMsg_PacketReliable) error
	onCSVCMsg_HLTVStatus                               []func(*deadlock.CSVCMsg_HLTVStatus) error
	onCSVCMsg_ServerSteamID                            []func(*deadlock.CSVCMsg_ServerSteamID) error
	onCSVCMsg_FullFrameSplit                           []func(*deadlock.CSVCMsg_FullFrameSplit) error
	onCSVCMsg_RconServerDetails                        []func(*deadlock.CSVCMsg_RconServerDetails) error
	onCSVCMsg_UserMessage                              []func(*deadlock.CSVCMsg_UserMessage) error
	onCSVCMsg_Broadcast_Command                        []func(*deadlock.CSVCMsg_Broadcast_Command) error
	onCSVCMsg_HltvFixupOperatorStatus                  []func(*deadlock.CSVCMsg_HltvFixupOperatorStatus) error
	onCUserMessageAchievementEvent                     []func(*deadlock.CUserMessageAchievementEvent) error
	onCUserMessageCloseCaption                         []func(*deadlock.CUserMessageCloseCaption) error
	onCUserMessageCloseCaptionDirect                   []func(*deadlock.CUserMessageCloseCaptionDirect) error
	onCUserMessageCurrentTimescale                     []func(*deadlock.CUserMessageCurrentTimescale) error
	onCUserMessageDesiredTimescale                     []func(*deadlock.CUserMessageDesiredTimescale) error
	onCUserMessageFade                                 []func(*deadlock.CUserMessageFade) error
	onCUserMessageGameTitle                            []func(*deadlock.CUserMessageGameTitle) error
	onCUserMessageHudMsg                               []func(*deadlock.CUserMessageHudMsg) error
	onCUserMessageHudText                              []func(*deadlock.CUserMessageHudText) error
	onCUserMessageColoredText                          []func(*deadlock.CUserMessageColoredText) error
	onCUserMessageRequestState                         []func(*deadlock.CUserMessageRequestState) error
	onCUserMessageResetHUD                             []func(*deadlock.CUserMessageResetHUD) error
	onCUserMessageRumble                               []func(*deadlock.CUserMessageRumble) error
	onCUserMessageSayText                              []func(*deadlock.CUserMessageSayText) error
	onCUserMessageSayText2                             []func(*deadlock.CUserMessageSayText2) error
	onCUserMessageSayTextChannel                       []func(*deadlock.CUserMessageSayTextChannel) error
	onCUserMessageShake                                []func(*deadlock.CUserMessageShake) error
	onCUserMessageShakeDir                             []func(*deadlock.CUserMessageShakeDir) error
	onCUserMessageWaterShake                           []func(*deadlock.CUserMessageWaterShake) error
	onCUserMessageTextMsg                              []func(*deadlock.CUserMessageTextMsg) error
	onCUserMessageScreenTilt                           []func(*deadlock.CUserMessageScreenTilt) error
	onCUserMessageVoiceMask                            []func(*deadlock.CUserMessageVoiceMask) error
	onCUserMessageSendAudio                            []func(*deadlock.CUserMessageSendAudio) error
	onCUserMessageItemPickup                           []func(*deadlock.CUserMessageItemPickup) error
	onCUserMessageAmmoDenied                           []func(*deadlock.CUserMessageAmmoDenied) error
	onCUserMessageShowMenu                             []func(*deadlock.CUserMessageShowMenu) error
	onCUserMessageCreditsMsg                           []func(*deadlock.CUserMessageCreditsMsg) error
	onCEntityMessagePlayJingle                         []func(*deadlock.CEntityMessagePlayJingle) error
	onCEntityMessageScreenOverlay                      []func(*deadlock.CEntityMessageScreenOverlay) error
	onCEntityMessageRemoveAllDecals                    []func(*deadlock.CEntityMessageRemoveAllDecals) error
	onCEntityMessagePropagateForce                     []func(*deadlock.CEntityMessagePropagateForce) error
	onCEntityMessageDoSpark                            []func(*deadlock.CEntityMessageDoSpark) error
	onCEntityMessageFixAngle                           []func(*deadlock.CEntityMessageFixAngle) error
	onCUserMessageCloseCaptionPlaceholder              []func(*deadlock.CUserMessageCloseCaptionPlaceholder) error
	onCUserMessageCameraTransition                     []func(*deadlock.CUserMessageCameraTransition) error
	onCUserMessageAudioParameter                       []func(*deadlock.CUserMessageAudioParameter) error
	onCUserMessageHapticsManagerPulse                  []func(*deadlock.CUserMessageHapticsManagerPulse) error
	onCUserMessageHapticsManagerEffect                 []func(*deadlock.CUserMessageHapticsManagerEffect) error
	onCUserMessageUpdateCssClasses                     []func(*deadlock.CUserMessageUpdateCssClasses) error
	onCUserMessageServerFrameTime                      []func(*deadlock.CUserMessageServerFrameTime) error
	onCUserMessageLagCompensationError                 []func(*deadlock.CUserMessageLagCompensationError) error
	onCUserMessageRequestDllStatus                     []func(*deadlock.CUserMessageRequestDllStatus) error
	onCUserMessageRequestUtilAction                    []func(*deadlock.CUserMessageRequestUtilAction) error
	onCUserMessageRequestInventory                     []func(*deadlock.CUserMessageRequestInventory) error
	onCUserMessageRequestDiagnostic                    []func(*deadlock.CUserMessageRequestDiagnostic) error
	onCMsgVDebugGameSessionIDEvent                     []func(*deadlock.CMsgVDebugGameSessionIDEvent) error
	onCMsgPlaceDecalEvent                              []func(*deadlock.CMsgPlaceDecalEvent) error
	onCMsgClearWorldDecalsEvent                        []func(*deadlock.CMsgClearWorldDecalsEvent) error
	onCMsgClearEntityDecalsEvent                       []func(*deadlock.CMsgClearEntityDecalsEvent) error
	onCMsgClearDecalsForSkeletonInstanceEvent          []func(*deadlock.CMsgClearDecalsForSkeletonInstanceEvent) error
	onCMsgSource1LegacyGameEventList                   []func(*deadlock.CMsgSource1LegacyGameEventList) error
	onCMsgSource1LegacyListenEvents                    []func(*deadlock.CMsgSource1LegacyListenEvents) error
	onCMsgSource1LegacyGameEvent                       []func(*deadlock.CMsgSource1LegacyGameEvent) error
	onCMsgSosStartSoundEvent                           []func(*deadlock.CMsgSosStartSoundEvent) error
	onCMsgSosStopSoundEvent                            []func(*deadlock.CMsgSosStopSoundEvent) error
	onCMsgSosSetSoundEventParams                       []func(*deadlock.CMsgSosSetSoundEventParams) error
	onCMsgSosSetLibraryStackFields                     []func(*deadlock.CMsgSosSetLibraryStackFields) error
	onCMsgSosStopSoundEventHash                        []func(*deadlock.CMsgSosStopSoundEventHash) error
	onCCitadelUserMessage_Damage                       []func(*deadlock.CCitadelUserMessage_Damage) error
	onCCitadelUserMsg_MapPing                          []func(*deadlock.CCitadelUserMsg_MapPing) error
	onCCitadelUserMsg_TeamRewards                      []func(*deadlock.CCitadelUserMsg_TeamRewards) error
	onCCitadelUserMsg_TriggerDamageFlash               []func(*deadlock.CCitadelUserMsg_TriggerDamageFlash) error
	onCCitadelUserMsg_AbilitiesChanged                 []func(*deadlock.CCitadelUserMsg_AbilitiesChanged) error
	onCCitadelUserMsg_RecentDamageSummary              []func(*deadlock.CCitadelUserMsg_RecentDamageSummary) error
	onCCitadelUserMsg_SpectatorTeamChanged             []func(*deadlock.CCitadelUserMsg_SpectatorTeamChanged) error
	onCCitadelUserMsg_ChatWheel                        []func(*deadlock.CCitadelUserMsg_ChatWheel) error
	onCCitadelUserMsg_GoldHistory                      []func(*deadlock.CCitadelUserMsg_GoldHistory) error
	onCCitadelUserMsg_ChatMsg                          []func(*deadlock.CCitadelUserMsg_ChatMsg) error
	onCCitadelUserMsg_QuickResponse                    []func(*deadlock.CCitadelUserMsg_QuickResponse) error
	onCCitadelUserMsg_PostMatchDetails                 []func(*deadlock.CCitadelUserMsg_PostMatchDetails) error
	onCCitadelUserMsg_ChatEvent                        []func(*deadlock.CCitadelUserMsg_ChatEvent) error
	onCCitadelUserMsg_AbilityInterrupted               []func(*deadlock.CCitadelUserMsg_AbilityInterrupted) error
	onCCitadelUserMsg_HeroKilled                       []func(*deadlock.CCitadelUserMsg_HeroKilled) error
	onCCitadelUserMsg_ReturnIdol                       []func(*deadlock.CCitadelUserMsg_ReturnIdol) error
	onCCitadelUserMsg_SetClientCameraAngles            []func(*deadlock.CCitadelUserMsg_SetClientCameraAngles) error
	onCCitadelUserMsg_MapLine                          []func(*deadlock.CCitadelUserMsg_MapLine) error
	onCCitadelUserMessage_BulletHit                    []func(*deadlock.CCitadelUserMessage_BulletHit) error
	onCCitadelUserMessage_ObjectiveMask                []func(*deadlock.CCitadelUserMessage_ObjectiveMask) error
	onCCitadelUserMessage_ModifierApplied              []func(*deadlock.CCitadelUserMessage_ModifierApplied) error
	onCCitadelUserMsg_CameraController                 []func(*deadlock.CCitadelUserMsg_CameraController) error
	onCCitadelUserMessage_AuraModifierApplied          []func(*deadlock.CCitadelUserMessage_AuraModifierApplied) error
	onCCitadelUserMsg_ObstructedShotFired              []func(*deadlock.CCitadelUserMsg_ObstructedShotFired) error
	onCCitadelUserMsg_AbilityLateFailure               []func(*deadlock.CCitadelUserMsg_AbilityLateFailure) error
	onCCitadelUserMsg_AbilityPing                      []func(*deadlock.CCitadelUserMsg_AbilityPing) error
	onCCitadelUserMsg_PostProcessingAnim               []func(*deadlock.CCitadelUserMsg_PostProcessingAnim) error
	onCCitadelUserMsg_DeathReplayData                  []func(*deadlock.CCitadelUserMsg_DeathReplayData) error
	onCCitadelUserMsg_PlayerLifetimeStatInfo           []func(*deadlock.CCitadelUserMsg_PlayerLifetimeStatInfo) error
	onCCitadelUserMsg_ForceShopClosed                  []func(*deadlock.CCitadelUserMsg_ForceShopClosed) error
	onCCitadelUserMsg_StaminaDrained                   []func(*deadlock.CCitadelUserMsg_StaminaDrained) error
	onCCitadelUserMessage_AbilityNotify                []func(*deadlock.CCitadelUserMessage_AbilityNotify) error
	onCCitadelUserMsg_GetDamageStatsResponse           []func(*deadlock.CCitadelUserMsg_GetDamageStatsResponse) error
	onCCitadelUserMsg_ParticipantSetSoundEventParams   []func(*deadlock.CCitadelUserMsg_ParticipantSetSoundEventParams) error
	onCCitadelUserMsg_ParticipantSetLibraryStackFields []func(*deadlock.CCitadelUserMsg_ParticipantSetLibraryStackFields) error
	onCCitadelUserMessage_CurrencyChanged              []func(*deadlock.CCitadelUserMessage_CurrencyChanged) error
	onCCitadelUserMessage_GameOver                     []func(*deadlock.CCitadelUserMessage_GameOver) error
	onCCitadelUserMsg_BossKilled                       []func(*deadlock.CCitadelUserMsg_BossKilled) error
}

func newCallbacks() *Callbacks {
	return &Callbacks{}
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

// OnCCitadelUserMessage_Damage registers a callback for CitadelUserMessageIds_k_EUserMsg_Damage
func (c *Callbacks) OnCCitadelUserMessage_Damage(fn func(*deadlock.CCitadelUserMessage_Damage) error) {
	c.onCCitadelUserMessage_Damage = append(c.onCCitadelUserMessage_Damage, fn)
}

// OnCCitadelUserMsg_MapPing registers a callback for CitadelUserMessageIds_k_EUserMsg_MapPing
func (c *Callbacks) OnCCitadelUserMsg_MapPing(fn func(*deadlock.CCitadelUserMsg_MapPing) error) {
	c.onCCitadelUserMsg_MapPing = append(c.onCCitadelUserMsg_MapPing, fn)
}

// OnCCitadelUserMsg_TeamRewards registers a callback for CitadelUserMessageIds_k_EUserMsg_TeamRewards
func (c *Callbacks) OnCCitadelUserMsg_TeamRewards(fn func(*deadlock.CCitadelUserMsg_TeamRewards) error) {
	c.onCCitadelUserMsg_TeamRewards = append(c.onCCitadelUserMsg_TeamRewards, fn)
}

// OnCCitadelUserMsg_TriggerDamageFlash registers a callback for CitadelUserMessageIds_k_EUserMsg_TriggerDamageFlash
func (c *Callbacks) OnCCitadelUserMsg_TriggerDamageFlash(fn func(*deadlock.CCitadelUserMsg_TriggerDamageFlash) error) {
	c.onCCitadelUserMsg_TriggerDamageFlash = append(c.onCCitadelUserMsg_TriggerDamageFlash, fn)
}

// OnCCitadelUserMsg_AbilitiesChanged registers a callback for CitadelUserMessageIds_k_EUserMsg_AbilitiesChanged
func (c *Callbacks) OnCCitadelUserMsg_AbilitiesChanged(fn func(*deadlock.CCitadelUserMsg_AbilitiesChanged) error) {
	c.onCCitadelUserMsg_AbilitiesChanged = append(c.onCCitadelUserMsg_AbilitiesChanged, fn)
}

// OnCCitadelUserMsg_RecentDamageSummary registers a callback for CitadelUserMessageIds_k_EUserMsg_RecentDamageSummary
func (c *Callbacks) OnCCitadelUserMsg_RecentDamageSummary(fn func(*deadlock.CCitadelUserMsg_RecentDamageSummary) error) {
	c.onCCitadelUserMsg_RecentDamageSummary = append(c.onCCitadelUserMsg_RecentDamageSummary, fn)
}

// OnCCitadelUserMsg_SpectatorTeamChanged registers a callback for CitadelUserMessageIds_k_EUserMsg_SpectatorTeamChanged
func (c *Callbacks) OnCCitadelUserMsg_SpectatorTeamChanged(fn func(*deadlock.CCitadelUserMsg_SpectatorTeamChanged) error) {
	c.onCCitadelUserMsg_SpectatorTeamChanged = append(c.onCCitadelUserMsg_SpectatorTeamChanged, fn)
}

// OnCCitadelUserMsg_ChatWheel registers a callback for CitadelUserMessageIds_k_EUserMsg_ChatWheel
func (c *Callbacks) OnCCitadelUserMsg_ChatWheel(fn func(*deadlock.CCitadelUserMsg_ChatWheel) error) {
	c.onCCitadelUserMsg_ChatWheel = append(c.onCCitadelUserMsg_ChatWheel, fn)
}

// OnCCitadelUserMsg_GoldHistory registers a callback for CitadelUserMessageIds_k_EUserMsg_GoldHistory
func (c *Callbacks) OnCCitadelUserMsg_GoldHistory(fn func(*deadlock.CCitadelUserMsg_GoldHistory) error) {
	c.onCCitadelUserMsg_GoldHistory = append(c.onCCitadelUserMsg_GoldHistory, fn)
}

// OnCCitadelUserMsg_ChatMsg registers a callback for CitadelUserMessageIds_k_EUserMsg_ChatMsg
func (c *Callbacks) OnCCitadelUserMsg_ChatMsg(fn func(*deadlock.CCitadelUserMsg_ChatMsg) error) {
	c.onCCitadelUserMsg_ChatMsg = append(c.onCCitadelUserMsg_ChatMsg, fn)
}

// OnCCitadelUserMsg_QuickResponse registers a callback for CitadelUserMessageIds_k_EUserMsg_QuickResponse
func (c *Callbacks) OnCCitadelUserMsg_QuickResponse(fn func(*deadlock.CCitadelUserMsg_QuickResponse) error) {
	c.onCCitadelUserMsg_QuickResponse = append(c.onCCitadelUserMsg_QuickResponse, fn)
}

// OnCCitadelUserMsg_PostMatchDetails registers a callback for CitadelUserMessageIds_k_EUserMsg_PostMatchDetails
func (c *Callbacks) OnCCitadelUserMsg_PostMatchDetails(fn func(*deadlock.CCitadelUserMsg_PostMatchDetails) error) {
	c.onCCitadelUserMsg_PostMatchDetails = append(c.onCCitadelUserMsg_PostMatchDetails, fn)
}

// OnCCitadelUserMsg_ChatEvent registers a callback for CitadelUserMessageIds_k_EUserMsg_ChatEvent
func (c *Callbacks) OnCCitadelUserMsg_ChatEvent(fn func(*deadlock.CCitadelUserMsg_ChatEvent) error) {
	c.onCCitadelUserMsg_ChatEvent = append(c.onCCitadelUserMsg_ChatEvent, fn)
}

// OnCCitadelUserMsg_AbilityInterrupted registers a callback for CitadelUserMessageIds_k_EUserMsg_AbilityInterrupted
func (c *Callbacks) OnCCitadelUserMsg_AbilityInterrupted(fn func(*deadlock.CCitadelUserMsg_AbilityInterrupted) error) {
	c.onCCitadelUserMsg_AbilityInterrupted = append(c.onCCitadelUserMsg_AbilityInterrupted, fn)
}

// OnCCitadelUserMsg_HeroKilled registers a callback for CitadelUserMessageIds_k_EUserMsg_HeroKilled
func (c *Callbacks) OnCCitadelUserMsg_HeroKilled(fn func(*deadlock.CCitadelUserMsg_HeroKilled) error) {
	c.onCCitadelUserMsg_HeroKilled = append(c.onCCitadelUserMsg_HeroKilled, fn)
}

// OnCCitadelUserMsg_ReturnIdol registers a callback for CitadelUserMessageIds_k_EUserMsg_ReturnIdol
func (c *Callbacks) OnCCitadelUserMsg_ReturnIdol(fn func(*deadlock.CCitadelUserMsg_ReturnIdol) error) {
	c.onCCitadelUserMsg_ReturnIdol = append(c.onCCitadelUserMsg_ReturnIdol, fn)
}

// OnCCitadelUserMsg_SetClientCameraAngles registers a callback for CitadelUserMessageIds_k_EUserMsg_SetClientCameraAngles
func (c *Callbacks) OnCCitadelUserMsg_SetClientCameraAngles(fn func(*deadlock.CCitadelUserMsg_SetClientCameraAngles) error) {
	c.onCCitadelUserMsg_SetClientCameraAngles = append(c.onCCitadelUserMsg_SetClientCameraAngles, fn)
}

// OnCCitadelUserMsg_MapLine registers a callback for CitadelUserMessageIds_k_EUserMsg_MapLine
func (c *Callbacks) OnCCitadelUserMsg_MapLine(fn func(*deadlock.CCitadelUserMsg_MapLine) error) {
	c.onCCitadelUserMsg_MapLine = append(c.onCCitadelUserMsg_MapLine, fn)
}

// OnCCitadelUserMessage_BulletHit registers a callback for CitadelUserMessageIds_k_EUserMsg_BulletHit
func (c *Callbacks) OnCCitadelUserMessage_BulletHit(fn func(*deadlock.CCitadelUserMessage_BulletHit) error) {
	c.onCCitadelUserMessage_BulletHit = append(c.onCCitadelUserMessage_BulletHit, fn)
}

// OnCCitadelUserMessage_ObjectiveMask registers a callback for CitadelUserMessageIds_k_EUserMsg_ObjectiveMask
func (c *Callbacks) OnCCitadelUserMessage_ObjectiveMask(fn func(*deadlock.CCitadelUserMessage_ObjectiveMask) error) {
	c.onCCitadelUserMessage_ObjectiveMask = append(c.onCCitadelUserMessage_ObjectiveMask, fn)
}

// OnCCitadelUserMessage_ModifierApplied registers a callback for CitadelUserMessageIds_k_EUserMsg_ModifierApplied
func (c *Callbacks) OnCCitadelUserMessage_ModifierApplied(fn func(*deadlock.CCitadelUserMessage_ModifierApplied) error) {
	c.onCCitadelUserMessage_ModifierApplied = append(c.onCCitadelUserMessage_ModifierApplied, fn)
}

// OnCCitadelUserMsg_CameraController registers a callback for CitadelUserMessageIds_k_EUserMsg_CameraController
func (c *Callbacks) OnCCitadelUserMsg_CameraController(fn func(*deadlock.CCitadelUserMsg_CameraController) error) {
	c.onCCitadelUserMsg_CameraController = append(c.onCCitadelUserMsg_CameraController, fn)
}

// OnCCitadelUserMessage_AuraModifierApplied registers a callback for CitadelUserMessageIds_k_EUserMsg_AuraModifierApplied
func (c *Callbacks) OnCCitadelUserMessage_AuraModifierApplied(fn func(*deadlock.CCitadelUserMessage_AuraModifierApplied) error) {
	c.onCCitadelUserMessage_AuraModifierApplied = append(c.onCCitadelUserMessage_AuraModifierApplied, fn)
}

// OnCCitadelUserMsg_ObstructedShotFired registers a callback for CitadelUserMessageIds_k_EUserMsg_ObstructedShotFired
func (c *Callbacks) OnCCitadelUserMsg_ObstructedShotFired(fn func(*deadlock.CCitadelUserMsg_ObstructedShotFired) error) {
	c.onCCitadelUserMsg_ObstructedShotFired = append(c.onCCitadelUserMsg_ObstructedShotFired, fn)
}

// OnCCitadelUserMsg_AbilityLateFailure registers a callback for CitadelUserMessageIds_k_EUserMsg_AbilityLateFailure
func (c *Callbacks) OnCCitadelUserMsg_AbilityLateFailure(fn func(*deadlock.CCitadelUserMsg_AbilityLateFailure) error) {
	c.onCCitadelUserMsg_AbilityLateFailure = append(c.onCCitadelUserMsg_AbilityLateFailure, fn)
}

// OnCCitadelUserMsg_AbilityPing registers a callback for CitadelUserMessageIds_k_EUserMsg_AbilityPing
func (c *Callbacks) OnCCitadelUserMsg_AbilityPing(fn func(*deadlock.CCitadelUserMsg_AbilityPing) error) {
	c.onCCitadelUserMsg_AbilityPing = append(c.onCCitadelUserMsg_AbilityPing, fn)
}

// OnCCitadelUserMsg_PostProcessingAnim registers a callback for CitadelUserMessageIds_k_EUserMsg_PostProcessingAnim
func (c *Callbacks) OnCCitadelUserMsg_PostProcessingAnim(fn func(*deadlock.CCitadelUserMsg_PostProcessingAnim) error) {
	c.onCCitadelUserMsg_PostProcessingAnim = append(c.onCCitadelUserMsg_PostProcessingAnim, fn)
}

// OnCCitadelUserMsg_DeathReplayData registers a callback for CitadelUserMessageIds_k_EUserMsg_DeathReplayData
func (c *Callbacks) OnCCitadelUserMsg_DeathReplayData(fn func(*deadlock.CCitadelUserMsg_DeathReplayData) error) {
	c.onCCitadelUserMsg_DeathReplayData = append(c.onCCitadelUserMsg_DeathReplayData, fn)
}

// OnCCitadelUserMsg_PlayerLifetimeStatInfo registers a callback for CitadelUserMessageIds_k_EUserMsg_PlayerLifetimeStatInfo
func (c *Callbacks) OnCCitadelUserMsg_PlayerLifetimeStatInfo(fn func(*deadlock.CCitadelUserMsg_PlayerLifetimeStatInfo) error) {
	c.onCCitadelUserMsg_PlayerLifetimeStatInfo = append(c.onCCitadelUserMsg_PlayerLifetimeStatInfo, fn)
}

// OnCCitadelUserMsg_ForceShopClosed registers a callback for CitadelUserMessageIds_k_EUserMsg_ForceShopClosed
func (c *Callbacks) OnCCitadelUserMsg_ForceShopClosed(fn func(*deadlock.CCitadelUserMsg_ForceShopClosed) error) {
	c.onCCitadelUserMsg_ForceShopClosed = append(c.onCCitadelUserMsg_ForceShopClosed, fn)
}

// OnCCitadelUserMsg_StaminaDrained registers a callback for CitadelUserMessageIds_k_EUserMsg_StaminaDrained
func (c *Callbacks) OnCCitadelUserMsg_StaminaDrained(fn func(*deadlock.CCitadelUserMsg_StaminaDrained) error) {
	c.onCCitadelUserMsg_StaminaDrained = append(c.onCCitadelUserMsg_StaminaDrained, fn)
}

// OnCCitadelUserMessage_AbilityNotify registers a callback for CitadelUserMessageIds_k_EUserMsg_AbilityNotify
func (c *Callbacks) OnCCitadelUserMessage_AbilityNotify(fn func(*deadlock.CCitadelUserMessage_AbilityNotify) error) {
	c.onCCitadelUserMessage_AbilityNotify = append(c.onCCitadelUserMessage_AbilityNotify, fn)
}

// OnCCitadelUserMsg_GetDamageStatsResponse registers a callback for CitadelUserMessageIds_k_EUserMsg_GetDamageStatsResponse
func (c *Callbacks) OnCCitadelUserMsg_GetDamageStatsResponse(fn func(*deadlock.CCitadelUserMsg_GetDamageStatsResponse) error) {
	c.onCCitadelUserMsg_GetDamageStatsResponse = append(c.onCCitadelUserMsg_GetDamageStatsResponse, fn)
}

// OnCCitadelUserMsg_ParticipantSetSoundEventParams registers a callback for CitadelUserMessageIds_k_EUserMsg_ParticipantSetSoundEventParams
func (c *Callbacks) OnCCitadelUserMsg_ParticipantSetSoundEventParams(fn func(*deadlock.CCitadelUserMsg_ParticipantSetSoundEventParams) error) {
	c.onCCitadelUserMsg_ParticipantSetSoundEventParams = append(c.onCCitadelUserMsg_ParticipantSetSoundEventParams, fn)
}

// OnCCitadelUserMsg_ParticipantSetLibraryStackFields registers a callback for CitadelUserMessageIds_k_EUserMsg_ParticipantSetLibraryStackFields
func (c *Callbacks) OnCCitadelUserMsg_ParticipantSetLibraryStackFields(fn func(*deadlock.CCitadelUserMsg_ParticipantSetLibraryStackFields) error) {
	c.onCCitadelUserMsg_ParticipantSetLibraryStackFields = append(c.onCCitadelUserMsg_ParticipantSetLibraryStackFields, fn)
}

// OnCCitadelUserMessage_CurrencyChanged registers a callback for CitadelUserMessageIds_k_EUserMsg_CurrencyChanged
func (c *Callbacks) OnCCitadelUserMessage_CurrencyChanged(fn func(*deadlock.CCitadelUserMessage_CurrencyChanged) error) {
	c.onCCitadelUserMessage_CurrencyChanged = append(c.onCCitadelUserMessage_CurrencyChanged, fn)
}

// OnCCitadelUserMessage_GameOver registers a callback for CitadelUserMessageIds_k_EUserMsg_GameOver
func (c *Callbacks) OnCCitadelUserMessage_GameOver(fn func(*deadlock.CCitadelUserMessage_GameOver) error) {
	c.onCCitadelUserMessage_GameOver = append(c.onCCitadelUserMessage_GameOver, fn)
}

// OnCCitadelUserMsg_BossKilled registers a callback for CitadelUserMessageIds_k_EUserMsg_BossKilled
func (c *Callbacks) OnCCitadelUserMsg_BossKilled(fn func(*deadlock.CCitadelUserMsg_BossKilled) error) {
	c.onCCitadelUserMsg_BossKilled = append(c.onCCitadelUserMsg_BossKilled, fn)
}

func (c *Callbacks) callByDemoType(t int32, buf []byte) error {
	switch t {
	case 0: // deadlock.EDemoCommands_DEM_Stop
		if c.onCDemoStop == nil {
			return nil
		}

		msg := &deadlock.CDemoStop{}
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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

// TODO: Cleanup
func (c *Callbacks) getDemoTypeName(t int32) string {
	switch t {
	case 0: //deadlock.EDemoCommands_DEM_Stop
		return "EDemoCommands_DEM_Stop"
	case 1: //deadlock.EDemoCommands_DEM_FileHeader
		return "EDemoCommands_DEM_FileHeader"
	case 2: //deadlock.EDemoCommands_DEM_FileInfo
		return "EDemoCommands_DEM_FileInfo"
	case 3: //deadlock.EDemoCommands_DEM_SyncTick
		return "EDemoCommands_DEM_SyncTick"
	case 4: //deadlock.EDemoCommands_DEM_SendTables
		return "EDemoCommands_DEM_SendTables"
	case 5: //deadlock.EDemoCommands_DEM_ClassInfo
		return "EDemoCommands_DEM_ClassInfo"
	case 6: //deadlock.EDemoCommands_DEM_StringTables
		return "EDemoCommands_DEM_StringTables"
	case 7: //deadlock.EDemoCommands_DEM_Packet
		return "EDemoCommands_DEM_Packet"
	case 8: //deadlock.EDemoCommands_DEM_SignonPacket
		return "EDemoCommands_DEM_SignonPacket"
	case 9: //deadlock.EDemoCommands_DEM_ConsoleCmd
		return "EDemoCommands_DEM_ConsoleCmd"
	case 10: //deadlock.EDemoCommands_DEM_CustomData
		return "EDemoCommands_DEM_CustomData"
	case 11: //deadlock.EDemoCommands_DEM_CustomDataCallbacks
		return "EDemoCommands_DEM_CustomDataCallbacks"
	case 12: //deadlock.EDemoCommands_DEM_UserCmd
		return "EDemoCommands_DEM_UserCmd"
	case 13: //deadlock.EDemoCommands_DEM_FullPacket
		return "EDemoCommands_DEM_FullPacket"
	case 14: //deadlock.EDemoCommands_DEM_SaveGame
		return "EDemoCommands_DEM_SaveGame"
	case 15: //deadlock.EDemoCommands_DEM_SpawnGroups
		return "EDemoCommands_DEM_SpawnGroups"
	case 16: //deadlock.EDemoCommands_DEM_AnimationData
		return "EDemoCommands_DEM_AnimationData"
	case 17: //deadlock.EDemoCommands_DEM_AnimationHeader
		return "EDemoCommands_DEM_AnimationHeader"

	}
	return "UNKNOWN:" + strconv.Itoa(int(t))
}

// TODO: Cleanup
func (c *Callbacks) toDemoString(t int32, buf []byte) string {
	switch t {
	case 0: //deadlock.EDemoCommands_DEM_Stop
		msg := &deadlock.CDemoStop{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 1: //deadlock.EDemoCommands_DEM_FileHeader
		msg := &deadlock.CDemoFileHeader{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 2: //deadlock.EDemoCommands_DEM_FileInfo
		msg := &deadlock.CDemoFileInfo{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 3: //deadlock.EDemoCommands_DEM_SyncTick
		msg := &deadlock.CDemoSyncTick{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 4: //deadlock.EDemoCommands_DEM_SendTables
		msg := &deadlock.CDemoSendTables{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 5: //deadlock.EDemoCommands_DEM_ClassInfo
		msg := &deadlock.CDemoClassInfo{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 6: //deadlock.EDemoCommands_DEM_StringTables
		msg := &deadlock.CDemoStringTables{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 7: //deadlock.EDemoCommands_DEM_Packet
		msg := &deadlock.CDemoPacket{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 8: //deadlock.EDemoCommands_DEM_SignonPacket
		msg := &deadlock.CDemoPacket{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 9: //deadlock.EDemoCommands_DEM_ConsoleCmd
		msg := &deadlock.CDemoConsoleCmd{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 10: //deadlock.EDemoCommands_DEM_CustomData
		msg := &deadlock.CDemoCustomData{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 11: //deadlock.EDemoCommands_DEM_CustomDataCallbacks
		msg := &deadlock.CDemoCustomDataCallbacks{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 12: //deadlock.EDemoCommands_DEM_UserCmd
		msg := &deadlock.CDemoUserCmd{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 13: //deadlock.EDemoCommands_DEM_FullPacket
		msg := &deadlock.CDemoFullPacket{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 14: //deadlock.EDemoCommands_DEM_SaveGame
		msg := &deadlock.CDemoSaveGame{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 15: //deadlock.EDemoCommands_DEM_SpawnGroups
		msg := &deadlock.CDemoSpawnGroups{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 16: //deadlock.EDemoCommands_DEM_AnimationData
		msg := &deadlock.CDemoAnimationData{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 17: //deadlock.EDemoCommands_DEM_AnimationHeader
		msg := &deadlock.CDemoAnimationHeader{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)

	}
	return "UNKNOWN:" + strconv.Itoa(int(t))
}

// TODO: Cleanup
func (c *Callbacks) getPacketTypeName(t int32) string {
	switch t {
	case 0: //deadlock.NET_Messages_net_NOP
		return "NET_Messages_net_NOP"
	case 3: //deadlock.NET_Messages_net_SplitScreenUser
		return "NET_Messages_net_SplitScreenUser"
	case 4: //deadlock.NET_Messages_net_Tick
		return "NET_Messages_net_Tick"
	case 5: //deadlock.NET_Messages_net_StringCmd
		return "NET_Messages_net_StringCmd"
	case 6: //deadlock.NET_Messages_net_SetConVar
		return "NET_Messages_net_SetConVar"
	case 7: //deadlock.NET_Messages_net_SignonState
		return "NET_Messages_net_SignonState"
	case 8: //deadlock.NET_Messages_net_SpawnGroup_Load
		return "NET_Messages_net_SpawnGroup_Load"
	case 9: //deadlock.NET_Messages_net_SpawnGroup_ManifestUpdate
		return "NET_Messages_net_SpawnGroup_ManifestUpdate"
	case 11: //deadlock.NET_Messages_net_SpawnGroup_SetCreationTick
		return "NET_Messages_net_SpawnGroup_SetCreationTick"
	case 12: //deadlock.NET_Messages_net_SpawnGroup_Unload
		return "NET_Messages_net_SpawnGroup_Unload"
	case 13: //deadlock.NET_Messages_net_SpawnGroup_LoadCompleted
		return "NET_Messages_net_SpawnGroup_LoadCompleted"
	case 15: //deadlock.NET_Messages_net_DebugOverlay
		return "NET_Messages_net_DebugOverlay"
	case 40: //deadlock.SVC_Messages_svc_ServerInfo
		return "SVC_Messages_svc_ServerInfo"
	case 41: //deadlock.SVC_Messages_svc_FlattenedSerializer
		return "SVC_Messages_svc_FlattenedSerializer"
	case 42: //deadlock.SVC_Messages_svc_ClassInfo
		return "SVC_Messages_svc_ClassInfo"
	case 43: //deadlock.SVC_Messages_svc_SetPause
		return "SVC_Messages_svc_SetPause"
	case 44: //deadlock.SVC_Messages_svc_CreateStringTable
		return "SVC_Messages_svc_CreateStringTable"
	case 45: //deadlock.SVC_Messages_svc_UpdateStringTable
		return "SVC_Messages_svc_UpdateStringTable"
	case 46: //deadlock.SVC_Messages_svc_VoiceInit
		return "SVC_Messages_svc_VoiceInit"
	case 47: //deadlock.SVC_Messages_svc_VoiceData
		return "SVC_Messages_svc_VoiceData"
	case 48: //deadlock.SVC_Messages_svc_Print
		return "SVC_Messages_svc_Print"
	case 49: //deadlock.SVC_Messages_svc_Sounds
		return "SVC_Messages_svc_Sounds"
	case 50: //deadlock.SVC_Messages_svc_SetView
		return "SVC_Messages_svc_SetView"
	case 51: //deadlock.SVC_Messages_svc_ClearAllStringTables
		return "SVC_Messages_svc_ClearAllStringTables"
	case 52: //deadlock.SVC_Messages_svc_CmdKeyValues
		return "SVC_Messages_svc_CmdKeyValues"
	case 53: //deadlock.SVC_Messages_svc_BSPDecal
		return "SVC_Messages_svc_BSPDecal"
	case 54: //deadlock.SVC_Messages_svc_SplitScreen
		return "SVC_Messages_svc_SplitScreen"
	case 55: //deadlock.SVC_Messages_svc_PacketEntities
		return "SVC_Messages_svc_PacketEntities"
	case 56: //deadlock.SVC_Messages_svc_Prefetch
		return "SVC_Messages_svc_Prefetch"
	case 57: //deadlock.SVC_Messages_svc_Menu
		return "SVC_Messages_svc_Menu"
	case 58: //deadlock.SVC_Messages_svc_GetCvarValue
		return "SVC_Messages_svc_GetCvarValue"
	case 59: //deadlock.SVC_Messages_svc_StopSound
		return "SVC_Messages_svc_StopSound"
	case 60: //deadlock.SVC_Messages_svc_PeerList
		return "SVC_Messages_svc_PeerList"
	case 61: //deadlock.SVC_Messages_svc_PacketReliable
		return "SVC_Messages_svc_PacketReliable"
	case 62: //deadlock.SVC_Messages_svc_HLTVStatus
		return "SVC_Messages_svc_HLTVStatus"
	case 63: //deadlock.SVC_Messages_svc_ServerSteamID
		return "SVC_Messages_svc_ServerSteamID"
	case 70: //deadlock.SVC_Messages_svc_FullFrameSplit
		return "SVC_Messages_svc_FullFrameSplit"
	case 71: //deadlock.SVC_Messages_svc_RconServerDetails
		return "SVC_Messages_svc_RconServerDetails"
	case 72: //deadlock.SVC_Messages_svc_UserMessage
		return "SVC_Messages_svc_UserMessage"
	case 74: //deadlock.SVC_Messages_svc_Broadcast_Command
		return "SVC_Messages_svc_Broadcast_Command"
	case 75: //deadlock.SVC_Messages_svc_HltvFixupOperatorStatus
		return "SVC_Messages_svc_HltvFixupOperatorStatus"
	case 101: //deadlock.EBaseUserMessages_UM_AchievementEvent
		return "EBaseUserMessages_UM_AchievementEvent"
	case 102: //deadlock.EBaseUserMessages_UM_CloseCaption
		return "EBaseUserMessages_UM_CloseCaption"
	case 103: //deadlock.EBaseUserMessages_UM_CloseCaptionDirect
		return "EBaseUserMessages_UM_CloseCaptionDirect"
	case 104: //deadlock.EBaseUserMessages_UM_CurrentTimescale
		return "EBaseUserMessages_UM_CurrentTimescale"
	case 105: //deadlock.EBaseUserMessages_UM_DesiredTimescale
		return "EBaseUserMessages_UM_DesiredTimescale"
	case 106: //deadlock.EBaseUserMessages_UM_Fade
		return "EBaseUserMessages_UM_Fade"
	case 107: //deadlock.EBaseUserMessages_UM_GameTitle
		return "EBaseUserMessages_UM_GameTitle"
	case 110: //deadlock.EBaseUserMessages_UM_HudMsg
		return "EBaseUserMessages_UM_HudMsg"
	case 111: //deadlock.EBaseUserMessages_UM_HudText
		return "EBaseUserMessages_UM_HudText"
	case 113: //deadlock.EBaseUserMessages_UM_ColoredText
		return "EBaseUserMessages_UM_ColoredText"
	case 114: //deadlock.EBaseUserMessages_UM_RequestState
		return "EBaseUserMessages_UM_RequestState"
	case 115: //deadlock.EBaseUserMessages_UM_ResetHUD
		return "EBaseUserMessages_UM_ResetHUD"
	case 116: //deadlock.EBaseUserMessages_UM_Rumble
		return "EBaseUserMessages_UM_Rumble"
	case 117: //deadlock.EBaseUserMessages_UM_SayText
		return "EBaseUserMessages_UM_SayText"
	case 118: //deadlock.EBaseUserMessages_UM_SayText2
		return "EBaseUserMessages_UM_SayText2"
	case 119: //deadlock.EBaseUserMessages_UM_SayTextChannel
		return "EBaseUserMessages_UM_SayTextChannel"
	case 120: //deadlock.EBaseUserMessages_UM_Shake
		return "EBaseUserMessages_UM_Shake"
	case 121: //deadlock.EBaseUserMessages_UM_ShakeDir
		return "EBaseUserMessages_UM_ShakeDir"
	case 122: //deadlock.EBaseUserMessages_UM_WaterShake
		return "EBaseUserMessages_UM_WaterShake"
	case 124: //deadlock.EBaseUserMessages_UM_TextMsg
		return "EBaseUserMessages_UM_TextMsg"
	case 125: //deadlock.EBaseUserMessages_UM_ScreenTilt
		return "EBaseUserMessages_UM_ScreenTilt"
	case 128: //deadlock.EBaseUserMessages_UM_VoiceMask
		return "EBaseUserMessages_UM_VoiceMask"
	case 130: //deadlock.EBaseUserMessages_UM_SendAudio
		return "EBaseUserMessages_UM_SendAudio"
	case 131: //deadlock.EBaseUserMessages_UM_ItemPickup
		return "EBaseUserMessages_UM_ItemPickup"
	case 132: //deadlock.EBaseUserMessages_UM_AmmoDenied
		return "EBaseUserMessages_UM_AmmoDenied"
	case 134: //deadlock.EBaseUserMessages_UM_ShowMenu
		return "EBaseUserMessages_UM_ShowMenu"
	case 135: //deadlock.EBaseUserMessages_UM_CreditsMsg
		return "EBaseUserMessages_UM_CreditsMsg"
	case 136: //deadlock.EBaseEntityMessages_EM_PlayJingle
		return "EBaseEntityMessages_EM_PlayJingle"
	case 137: //deadlock.EBaseEntityMessages_EM_ScreenOverlay
		return "EBaseEntityMessages_EM_ScreenOverlay"
	case 138: //deadlock.EBaseEntityMessages_EM_RemoveAllDecals
		return "EBaseEntityMessages_EM_RemoveAllDecals"
	case 139: //deadlock.EBaseEntityMessages_EM_PropagateForce
		return "EBaseEntityMessages_EM_PropagateForce"
	case 140: //deadlock.EBaseEntityMessages_EM_DoSpark
		return "EBaseEntityMessages_EM_DoSpark"
	case 141: //deadlock.EBaseEntityMessages_EM_FixAngle
		return "EBaseEntityMessages_EM_FixAngle"
	case 142: //deadlock.EBaseUserMessages_UM_CloseCaptionPlaceholder
		return "EBaseUserMessages_UM_CloseCaptionPlaceholder"
	case 143: //deadlock.EBaseUserMessages_UM_CameraTransition
		return "EBaseUserMessages_UM_CameraTransition"
	case 144: //deadlock.EBaseUserMessages_UM_AudioParameter
		return "EBaseUserMessages_UM_AudioParameter"
	case 150: //deadlock.EBaseUserMessages_UM_HapticsManagerPulse
		return "EBaseUserMessages_UM_HapticsManagerPulse"
	case 151: //deadlock.EBaseUserMessages_UM_HapticsManagerEffect
		return "EBaseUserMessages_UM_HapticsManagerEffect"
	case 153: //deadlock.EBaseUserMessages_UM_UpdateCssClasses
		return "EBaseUserMessages_UM_UpdateCssClasses"
	case 154: //deadlock.EBaseUserMessages_UM_ServerFrameTime
		return "EBaseUserMessages_UM_ServerFrameTime"
	case 155: //deadlock.EBaseUserMessages_UM_LagCompensationError
		return "EBaseUserMessages_UM_LagCompensationError"
	case 156: //deadlock.EBaseUserMessages_UM_RequestDllStatus
		return "EBaseUserMessages_UM_RequestDllStatus"
	case 157: //deadlock.EBaseUserMessages_UM_RequestUtilAction
		return "EBaseUserMessages_UM_RequestUtilAction"
	case 160: //deadlock.EBaseUserMessages_UM_RequestInventory
		return "EBaseUserMessages_UM_RequestInventory"
	case 162: //deadlock.EBaseUserMessages_UM_RequestDiagnostic
		return "EBaseUserMessages_UM_RequestDiagnostic"
	case 200: //deadlock.EBaseGameEvents_GE_VDebugGameSessionIDEvent
		return "EBaseGameEvents_GE_VDebugGameSessionIDEvent"
	case 201: //deadlock.EBaseGameEvents_GE_PlaceDecalEvent
		return "EBaseGameEvents_GE_PlaceDecalEvent"
	case 202: //deadlock.EBaseGameEvents_GE_ClearWorldDecalsEvent
		return "EBaseGameEvents_GE_ClearWorldDecalsEvent"
	case 203: //deadlock.EBaseGameEvents_GE_ClearEntityDecalsEvent
		return "EBaseGameEvents_GE_ClearEntityDecalsEvent"
	case 204: //deadlock.EBaseGameEvents_GE_ClearDecalsForSkeletonInstanceEvent
		return "EBaseGameEvents_GE_ClearDecalsForSkeletonInstanceEvent"
	case 205: //deadlock.EBaseGameEvents_GE_Source1LegacyGameEventList
		return "EBaseGameEvents_GE_Source1LegacyGameEventList"
	case 206: //deadlock.EBaseGameEvents_GE_Source1LegacyListenEvents
		return "EBaseGameEvents_GE_Source1LegacyListenEvents"
	case 207: //deadlock.EBaseGameEvents_GE_Source1LegacyGameEvent
		return "EBaseGameEvents_GE_Source1LegacyGameEvent"
	case 208: //deadlock.EBaseGameEvents_GE_SosStartSoundEvent
		return "EBaseGameEvents_GE_SosStartSoundEvent"
	case 209: //deadlock.EBaseGameEvents_GE_SosStopSoundEvent
		return "EBaseGameEvents_GE_SosStopSoundEvent"
	case 210: //deadlock.EBaseGameEvents_GE_SosSetSoundEventParams
		return "EBaseGameEvents_GE_SosSetSoundEventParams"
	case 211: //deadlock.EBaseGameEvents_GE_SosSetLibraryStackFields
		return "EBaseGameEvents_GE_SosSetLibraryStackFields"
	case 212: //deadlock.EBaseGameEvents_GE_SosStopSoundEventHash
		return "EBaseGameEvents_GE_SosStopSoundEventHash"
	case 300: //deadlock.CitadelUserMessageIds_k_EUserMsg_Damage
		return "CitadelUserMessageIds_k_EUserMsg_Damage"
	case 303: //deadlock.CitadelUserMessageIds_k_EUserMsg_MapPing
		return "CitadelUserMessageIds_k_EUserMsg_MapPing"
	case 304: //deadlock.CitadelUserMessageIds_k_EUserMsg_TeamRewards
		return "CitadelUserMessageIds_k_EUserMsg_TeamRewards"
	case 308: //deadlock.CitadelUserMessageIds_k_EUserMsg_TriggerDamageFlash
		return "CitadelUserMessageIds_k_EUserMsg_TriggerDamageFlash"
	case 309: //deadlock.CitadelUserMessageIds_k_EUserMsg_AbilitiesChanged
		return "CitadelUserMessageIds_k_EUserMsg_AbilitiesChanged"
	case 310: //deadlock.CitadelUserMessageIds_k_EUserMsg_RecentDamageSummary
		return "CitadelUserMessageIds_k_EUserMsg_RecentDamageSummary"
	case 311: //deadlock.CitadelUserMessageIds_k_EUserMsg_SpectatorTeamChanged
		return "CitadelUserMessageIds_k_EUserMsg_SpectatorTeamChanged"
	case 312: //deadlock.CitadelUserMessageIds_k_EUserMsg_ChatWheel
		return "CitadelUserMessageIds_k_EUserMsg_ChatWheel"
	case 313: //deadlock.CitadelUserMessageIds_k_EUserMsg_GoldHistory
		return "CitadelUserMessageIds_k_EUserMsg_GoldHistory"
	case 314: //deadlock.CitadelUserMessageIds_k_EUserMsg_ChatMsg
		return "CitadelUserMessageIds_k_EUserMsg_ChatMsg"
	case 315: //deadlock.CitadelUserMessageIds_k_EUserMsg_QuickResponse
		return "CitadelUserMessageIds_k_EUserMsg_QuickResponse"
	case 316: //deadlock.CitadelUserMessageIds_k_EUserMsg_PostMatchDetails
		return "CitadelUserMessageIds_k_EUserMsg_PostMatchDetails"
	case 317: //deadlock.CitadelUserMessageIds_k_EUserMsg_ChatEvent
		return "CitadelUserMessageIds_k_EUserMsg_ChatEvent"
	case 318: //deadlock.CitadelUserMessageIds_k_EUserMsg_AbilityInterrupted
		return "CitadelUserMessageIds_k_EUserMsg_AbilityInterrupted"
	case 319: //deadlock.CitadelUserMessageIds_k_EUserMsg_HeroKilled
		return "CitadelUserMessageIds_k_EUserMsg_HeroKilled"
	case 320: //deadlock.CitadelUserMessageIds_k_EUserMsg_ReturnIdol
		return "CitadelUserMessageIds_k_EUserMsg_ReturnIdol"
	case 321: //deadlock.CitadelUserMessageIds_k_EUserMsg_SetClientCameraAngles
		return "CitadelUserMessageIds_k_EUserMsg_SetClientCameraAngles"
	case 322: //deadlock.CitadelUserMessageIds_k_EUserMsg_MapLine
		return "CitadelUserMessageIds_k_EUserMsg_MapLine"
	case 323: //deadlock.CitadelUserMessageIds_k_EUserMsg_BulletHit
		return "CitadelUserMessageIds_k_EUserMsg_BulletHit"
	case 324: //deadlock.CitadelUserMessageIds_k_EUserMsg_ObjectiveMask
		return "CitadelUserMessageIds_k_EUserMsg_ObjectiveMask"
	case 325: //deadlock.CitadelUserMessageIds_k_EUserMsg_ModifierApplied
		return "CitadelUserMessageIds_k_EUserMsg_ModifierApplied"
	case 326: //deadlock.CitadelUserMessageIds_k_EUserMsg_CameraController
		return "CitadelUserMessageIds_k_EUserMsg_CameraController"
	case 327: //deadlock.CitadelUserMessageIds_k_EUserMsg_AuraModifierApplied
		return "CitadelUserMessageIds_k_EUserMsg_AuraModifierApplied"
	case 329: //deadlock.CitadelUserMessageIds_k_EUserMsg_ObstructedShotFired
		return "CitadelUserMessageIds_k_EUserMsg_ObstructedShotFired"
	case 330: //deadlock.CitadelUserMessageIds_k_EUserMsg_AbilityLateFailure
		return "CitadelUserMessageIds_k_EUserMsg_AbilityLateFailure"
	case 331: //deadlock.CitadelUserMessageIds_k_EUserMsg_AbilityPing
		return "CitadelUserMessageIds_k_EUserMsg_AbilityPing"
	case 332: //deadlock.CitadelUserMessageIds_k_EUserMsg_PostProcessingAnim
		return "CitadelUserMessageIds_k_EUserMsg_PostProcessingAnim"
	case 333: //deadlock.CitadelUserMessageIds_k_EUserMsg_DeathReplayData
		return "CitadelUserMessageIds_k_EUserMsg_DeathReplayData"
	case 334: //deadlock.CitadelUserMessageIds_k_EUserMsg_PlayerLifetimeStatInfo
		return "CitadelUserMessageIds_k_EUserMsg_PlayerLifetimeStatInfo"
	case 336: //deadlock.CitadelUserMessageIds_k_EUserMsg_ForceShopClosed
		return "CitadelUserMessageIds_k_EUserMsg_ForceShopClosed"
	case 337: //deadlock.CitadelUserMessageIds_k_EUserMsg_StaminaDrained
		return "CitadelUserMessageIds_k_EUserMsg_StaminaDrained"
	case 338: //deadlock.CitadelUserMessageIds_k_EUserMsg_AbilityNotify
		return "CitadelUserMessageIds_k_EUserMsg_AbilityNotify"
	case 339: //deadlock.CitadelUserMessageIds_k_EUserMsg_GetDamageStatsResponse
		return "CitadelUserMessageIds_k_EUserMsg_GetDamageStatsResponse"
	case 343: //deadlock.CitadelUserMessageIds_k_EUserMsg_ParticipantSetSoundEventParams
		return "CitadelUserMessageIds_k_EUserMsg_ParticipantSetSoundEventParams"
	case 344: //deadlock.CitadelUserMessageIds_k_EUserMsg_ParticipantSetLibraryStackFields
		return "CitadelUserMessageIds_k_EUserMsg_ParticipantSetLibraryStackFields"
	case 345: //deadlock.CitadelUserMessageIds_k_EUserMsg_CurrencyChanged
		return "CitadelUserMessageIds_k_EUserMsg_CurrencyChanged"
	case 346: //deadlock.CitadelUserMessageIds_k_EUserMsg_GameOver
		return "CitadelUserMessageIds_k_EUserMsg_GameOver"
	case 347: //deadlock.CitadelUserMessageIds_k_EUserMsg_BossKilled
		return "CitadelUserMessageIds_k_EUserMsg_BossKilled"

	}
	return "UNKNOWN:" + strconv.Itoa(int(t))
}

// TODO: Cleanup
func (c *Callbacks) toPacketString(t int32, buf []byte) string {
	switch t {
	case 0: //deadlock.NET_Messages_net_NOP
		msg := &deadlock.CNETMsg_NOP{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 3: //deadlock.NET_Messages_net_SplitScreenUser
		msg := &deadlock.CNETMsg_SplitScreenUser{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 4: //deadlock.NET_Messages_net_Tick
		msg := &deadlock.CNETMsg_Tick{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 5: //deadlock.NET_Messages_net_StringCmd
		msg := &deadlock.CNETMsg_StringCmd{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 6: //deadlock.NET_Messages_net_SetConVar
		msg := &deadlock.CNETMsg_SetConVar{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 7: //deadlock.NET_Messages_net_SignonState
		msg := &deadlock.CNETMsg_SignonState{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 8: //deadlock.NET_Messages_net_SpawnGroup_Load
		msg := &deadlock.CNETMsg_SpawnGroup_Load{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 9: //deadlock.NET_Messages_net_SpawnGroup_ManifestUpdate
		msg := &deadlock.CNETMsg_SpawnGroup_ManifestUpdate{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 11: //deadlock.NET_Messages_net_SpawnGroup_SetCreationTick
		msg := &deadlock.CNETMsg_SpawnGroup_SetCreationTick{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 12: //deadlock.NET_Messages_net_SpawnGroup_Unload
		msg := &deadlock.CNETMsg_SpawnGroup_Unload{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 13: //deadlock.NET_Messages_net_SpawnGroup_LoadCompleted
		msg := &deadlock.CNETMsg_SpawnGroup_LoadCompleted{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 15: //deadlock.NET_Messages_net_DebugOverlay
		msg := &deadlock.CNETMsg_DebugOverlay{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 40: //deadlock.SVC_Messages_svc_ServerInfo
		msg := &deadlock.CSVCMsg_ServerInfo{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 41: //deadlock.SVC_Messages_svc_FlattenedSerializer
		msg := &deadlock.CSVCMsg_FlattenedSerializer{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 42: //deadlock.SVC_Messages_svc_ClassInfo
		msg := &deadlock.CSVCMsg_ClassInfo{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 43: //deadlock.SVC_Messages_svc_SetPause
		msg := &deadlock.CSVCMsg_SetPause{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 44: //deadlock.SVC_Messages_svc_CreateStringTable
		msg := &deadlock.CSVCMsg_CreateStringTable{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 45: //deadlock.SVC_Messages_svc_UpdateStringTable
		msg := &deadlock.CSVCMsg_UpdateStringTable{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 46: //deadlock.SVC_Messages_svc_VoiceInit
		msg := &deadlock.CSVCMsg_VoiceInit{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 47: //deadlock.SVC_Messages_svc_VoiceData
		msg := &deadlock.CSVCMsg_VoiceData{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 48: //deadlock.SVC_Messages_svc_Print
		msg := &deadlock.CSVCMsg_Print{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 49: //deadlock.SVC_Messages_svc_Sounds
		msg := &deadlock.CSVCMsg_Sounds{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 50: //deadlock.SVC_Messages_svc_SetView
		msg := &deadlock.CSVCMsg_SetView{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 51: //deadlock.SVC_Messages_svc_ClearAllStringTables
		msg := &deadlock.CSVCMsg_ClearAllStringTables{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 52: //deadlock.SVC_Messages_svc_CmdKeyValues
		msg := &deadlock.CSVCMsg_CmdKeyValues{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 53: //deadlock.SVC_Messages_svc_BSPDecal
		msg := &deadlock.CSVCMsg_BSPDecal{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 54: //deadlock.SVC_Messages_svc_SplitScreen
		msg := &deadlock.CSVCMsg_SplitScreen{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 55: //deadlock.SVC_Messages_svc_PacketEntities
		msg := &deadlock.CSVCMsg_PacketEntities{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 56: //deadlock.SVC_Messages_svc_Prefetch
		msg := &deadlock.CSVCMsg_Prefetch{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 57: //deadlock.SVC_Messages_svc_Menu
		msg := &deadlock.CSVCMsg_Menu{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 58: //deadlock.SVC_Messages_svc_GetCvarValue
		msg := &deadlock.CSVCMsg_GetCvarValue{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 59: //deadlock.SVC_Messages_svc_StopSound
		msg := &deadlock.CSVCMsg_StopSound{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 60: //deadlock.SVC_Messages_svc_PeerList
		msg := &deadlock.CSVCMsg_PeerList{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 61: //deadlock.SVC_Messages_svc_PacketReliable
		msg := &deadlock.CSVCMsg_PacketReliable{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 62: //deadlock.SVC_Messages_svc_HLTVStatus
		msg := &deadlock.CSVCMsg_HLTVStatus{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 63: //deadlock.SVC_Messages_svc_ServerSteamID
		msg := &deadlock.CSVCMsg_ServerSteamID{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 70: //deadlock.SVC_Messages_svc_FullFrameSplit
		msg := &deadlock.CSVCMsg_FullFrameSplit{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 71: //deadlock.SVC_Messages_svc_RconServerDetails
		msg := &deadlock.CSVCMsg_RconServerDetails{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 72: //deadlock.SVC_Messages_svc_UserMessage
		msg := &deadlock.CSVCMsg_UserMessage{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 74: //deadlock.SVC_Messages_svc_Broadcast_Command
		msg := &deadlock.CSVCMsg_Broadcast_Command{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 75: //deadlock.SVC_Messages_svc_HltvFixupOperatorStatus
		msg := &deadlock.CSVCMsg_HltvFixupOperatorStatus{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 101: //deadlock.EBaseUserMessages_UM_AchievementEvent
		msg := &deadlock.CUserMessageAchievementEvent{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 102: //deadlock.EBaseUserMessages_UM_CloseCaption
		msg := &deadlock.CUserMessageCloseCaption{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 103: //deadlock.EBaseUserMessages_UM_CloseCaptionDirect
		msg := &deadlock.CUserMessageCloseCaptionDirect{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 104: //deadlock.EBaseUserMessages_UM_CurrentTimescale
		msg := &deadlock.CUserMessageCurrentTimescale{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 105: //deadlock.EBaseUserMessages_UM_DesiredTimescale
		msg := &deadlock.CUserMessageDesiredTimescale{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 106: //deadlock.EBaseUserMessages_UM_Fade
		msg := &deadlock.CUserMessageFade{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 107: //deadlock.EBaseUserMessages_UM_GameTitle
		msg := &deadlock.CUserMessageGameTitle{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 110: //deadlock.EBaseUserMessages_UM_HudMsg
		msg := &deadlock.CUserMessageHudMsg{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 111: //deadlock.EBaseUserMessages_UM_HudText
		msg := &deadlock.CUserMessageHudText{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 113: //deadlock.EBaseUserMessages_UM_ColoredText
		msg := &deadlock.CUserMessageColoredText{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 114: //deadlock.EBaseUserMessages_UM_RequestState
		msg := &deadlock.CUserMessageRequestState{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 115: //deadlock.EBaseUserMessages_UM_ResetHUD
		msg := &deadlock.CUserMessageResetHUD{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 116: //deadlock.EBaseUserMessages_UM_Rumble
		msg := &deadlock.CUserMessageRumble{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 117: //deadlock.EBaseUserMessages_UM_SayText
		msg := &deadlock.CUserMessageSayText{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 118: //deadlock.EBaseUserMessages_UM_SayText2
		msg := &deadlock.CUserMessageSayText2{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 119: //deadlock.EBaseUserMessages_UM_SayTextChannel
		msg := &deadlock.CUserMessageSayTextChannel{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 120: //deadlock.EBaseUserMessages_UM_Shake
		msg := &deadlock.CUserMessageShake{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 121: //deadlock.EBaseUserMessages_UM_ShakeDir
		msg := &deadlock.CUserMessageShakeDir{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 122: //deadlock.EBaseUserMessages_UM_WaterShake
		msg := &deadlock.CUserMessageWaterShake{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 124: //deadlock.EBaseUserMessages_UM_TextMsg
		msg := &deadlock.CUserMessageTextMsg{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 125: //deadlock.EBaseUserMessages_UM_ScreenTilt
		msg := &deadlock.CUserMessageScreenTilt{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 128: //deadlock.EBaseUserMessages_UM_VoiceMask
		msg := &deadlock.CUserMessageVoiceMask{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 130: //deadlock.EBaseUserMessages_UM_SendAudio
		msg := &deadlock.CUserMessageSendAudio{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 131: //deadlock.EBaseUserMessages_UM_ItemPickup
		msg := &deadlock.CUserMessageItemPickup{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 132: //deadlock.EBaseUserMessages_UM_AmmoDenied
		msg := &deadlock.CUserMessageAmmoDenied{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 134: //deadlock.EBaseUserMessages_UM_ShowMenu
		msg := &deadlock.CUserMessageShowMenu{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 135: //deadlock.EBaseUserMessages_UM_CreditsMsg
		msg := &deadlock.CUserMessageCreditsMsg{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 136: //deadlock.EBaseEntityMessages_EM_PlayJingle
		msg := &deadlock.CEntityMessagePlayJingle{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 137: //deadlock.EBaseEntityMessages_EM_ScreenOverlay
		msg := &deadlock.CEntityMessageScreenOverlay{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 138: //deadlock.EBaseEntityMessages_EM_RemoveAllDecals
		msg := &deadlock.CEntityMessageRemoveAllDecals{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 139: //deadlock.EBaseEntityMessages_EM_PropagateForce
		msg := &deadlock.CEntityMessagePropagateForce{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 140: //deadlock.EBaseEntityMessages_EM_DoSpark
		msg := &deadlock.CEntityMessageDoSpark{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 141: //deadlock.EBaseEntityMessages_EM_FixAngle
		msg := &deadlock.CEntityMessageFixAngle{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 142: //deadlock.EBaseUserMessages_UM_CloseCaptionPlaceholder
		msg := &deadlock.CUserMessageCloseCaptionPlaceholder{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 143: //deadlock.EBaseUserMessages_UM_CameraTransition
		msg := &deadlock.CUserMessageCameraTransition{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 144: //deadlock.EBaseUserMessages_UM_AudioParameter
		msg := &deadlock.CUserMessageAudioParameter{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 150: //deadlock.EBaseUserMessages_UM_HapticsManagerPulse
		msg := &deadlock.CUserMessageHapticsManagerPulse{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 151: //deadlock.EBaseUserMessages_UM_HapticsManagerEffect
		msg := &deadlock.CUserMessageHapticsManagerEffect{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 153: //deadlock.EBaseUserMessages_UM_UpdateCssClasses
		msg := &deadlock.CUserMessageUpdateCssClasses{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 154: //deadlock.EBaseUserMessages_UM_ServerFrameTime
		msg := &deadlock.CUserMessageServerFrameTime{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 155: //deadlock.EBaseUserMessages_UM_LagCompensationError
		msg := &deadlock.CUserMessageLagCompensationError{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 156: //deadlock.EBaseUserMessages_UM_RequestDllStatus
		msg := &deadlock.CUserMessageRequestDllStatus{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 157: //deadlock.EBaseUserMessages_UM_RequestUtilAction
		msg := &deadlock.CUserMessageRequestUtilAction{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 160: //deadlock.EBaseUserMessages_UM_RequestInventory
		msg := &deadlock.CUserMessageRequestInventory{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 162: //deadlock.EBaseUserMessages_UM_RequestDiagnostic
		msg := &deadlock.CUserMessageRequestDiagnostic{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 200: //deadlock.EBaseGameEvents_GE_VDebugGameSessionIDEvent
		msg := &deadlock.CMsgVDebugGameSessionIDEvent{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 201: //deadlock.EBaseGameEvents_GE_PlaceDecalEvent
		msg := &deadlock.CMsgPlaceDecalEvent{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 202: //deadlock.EBaseGameEvents_GE_ClearWorldDecalsEvent
		msg := &deadlock.CMsgClearWorldDecalsEvent{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 203: //deadlock.EBaseGameEvents_GE_ClearEntityDecalsEvent
		msg := &deadlock.CMsgClearEntityDecalsEvent{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 204: //deadlock.EBaseGameEvents_GE_ClearDecalsForSkeletonInstanceEvent
		msg := &deadlock.CMsgClearDecalsForSkeletonInstanceEvent{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 205: //deadlock.EBaseGameEvents_GE_Source1LegacyGameEventList
		msg := &deadlock.CMsgSource1LegacyGameEventList{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 206: //deadlock.EBaseGameEvents_GE_Source1LegacyListenEvents
		msg := &deadlock.CMsgSource1LegacyListenEvents{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 207: //deadlock.EBaseGameEvents_GE_Source1LegacyGameEvent
		msg := &deadlock.CMsgSource1LegacyGameEvent{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 208: //deadlock.EBaseGameEvents_GE_SosStartSoundEvent
		msg := &deadlock.CMsgSosStartSoundEvent{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 209: //deadlock.EBaseGameEvents_GE_SosStopSoundEvent
		msg := &deadlock.CMsgSosStopSoundEvent{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 210: //deadlock.EBaseGameEvents_GE_SosSetSoundEventParams
		msg := &deadlock.CMsgSosSetSoundEventParams{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 211: //deadlock.EBaseGameEvents_GE_SosSetLibraryStackFields
		msg := &deadlock.CMsgSosSetLibraryStackFields{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 212: //deadlock.EBaseGameEvents_GE_SosStopSoundEventHash
		msg := &deadlock.CMsgSosStopSoundEventHash{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 300: //deadlock.CitadelUserMessageIds_k_EUserMsg_Damage
		msg := &deadlock.CCitadelUserMessage_Damage{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 303: //deadlock.CitadelUserMessageIds_k_EUserMsg_MapPing
		msg := &deadlock.CCitadelUserMsg_MapPing{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 304: //deadlock.CitadelUserMessageIds_k_EUserMsg_TeamRewards
		msg := &deadlock.CCitadelUserMsg_TeamRewards{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 308: //deadlock.CitadelUserMessageIds_k_EUserMsg_TriggerDamageFlash
		msg := &deadlock.CCitadelUserMsg_TriggerDamageFlash{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 309: //deadlock.CitadelUserMessageIds_k_EUserMsg_AbilitiesChanged
		msg := &deadlock.CCitadelUserMsg_AbilitiesChanged{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 310: //deadlock.CitadelUserMessageIds_k_EUserMsg_RecentDamageSummary
		msg := &deadlock.CCitadelUserMsg_RecentDamageSummary{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 311: //deadlock.CitadelUserMessageIds_k_EUserMsg_SpectatorTeamChanged
		msg := &deadlock.CCitadelUserMsg_SpectatorTeamChanged{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 312: //deadlock.CitadelUserMessageIds_k_EUserMsg_ChatWheel
		msg := &deadlock.CCitadelUserMsg_ChatWheel{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 313: //deadlock.CitadelUserMessageIds_k_EUserMsg_GoldHistory
		msg := &deadlock.CCitadelUserMsg_GoldHistory{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 314: //deadlock.CitadelUserMessageIds_k_EUserMsg_ChatMsg
		msg := &deadlock.CCitadelUserMsg_ChatMsg{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 315: //deadlock.CitadelUserMessageIds_k_EUserMsg_QuickResponse
		msg := &deadlock.CCitadelUserMsg_QuickResponse{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 316: //deadlock.CitadelUserMessageIds_k_EUserMsg_PostMatchDetails
		msg := &deadlock.CCitadelUserMsg_PostMatchDetails{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 317: //deadlock.CitadelUserMessageIds_k_EUserMsg_ChatEvent
		msg := &deadlock.CCitadelUserMsg_ChatEvent{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 318: //deadlock.CitadelUserMessageIds_k_EUserMsg_AbilityInterrupted
		msg := &deadlock.CCitadelUserMsg_AbilityInterrupted{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 319: //deadlock.CitadelUserMessageIds_k_EUserMsg_HeroKilled
		msg := &deadlock.CCitadelUserMsg_HeroKilled{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 320: //deadlock.CitadelUserMessageIds_k_EUserMsg_ReturnIdol
		msg := &deadlock.CCitadelUserMsg_ReturnIdol{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 321: //deadlock.CitadelUserMessageIds_k_EUserMsg_SetClientCameraAngles
		msg := &deadlock.CCitadelUserMsg_SetClientCameraAngles{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 322: //deadlock.CitadelUserMessageIds_k_EUserMsg_MapLine
		msg := &deadlock.CCitadelUserMsg_MapLine{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 323: //deadlock.CitadelUserMessageIds_k_EUserMsg_BulletHit
		msg := &deadlock.CCitadelUserMessage_BulletHit{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 324: //deadlock.CitadelUserMessageIds_k_EUserMsg_ObjectiveMask
		msg := &deadlock.CCitadelUserMessage_ObjectiveMask{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 325: //deadlock.CitadelUserMessageIds_k_EUserMsg_ModifierApplied
		msg := &deadlock.CCitadelUserMessage_ModifierApplied{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 326: //deadlock.CitadelUserMessageIds_k_EUserMsg_CameraController
		msg := &deadlock.CCitadelUserMsg_CameraController{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 327: //deadlock.CitadelUserMessageIds_k_EUserMsg_AuraModifierApplied
		msg := &deadlock.CCitadelUserMessage_AuraModifierApplied{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 329: //deadlock.CitadelUserMessageIds_k_EUserMsg_ObstructedShotFired
		msg := &deadlock.CCitadelUserMsg_ObstructedShotFired{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 330: //deadlock.CitadelUserMessageIds_k_EUserMsg_AbilityLateFailure
		msg := &deadlock.CCitadelUserMsg_AbilityLateFailure{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 331: //deadlock.CitadelUserMessageIds_k_EUserMsg_AbilityPing
		msg := &deadlock.CCitadelUserMsg_AbilityPing{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 332: //deadlock.CitadelUserMessageIds_k_EUserMsg_PostProcessingAnim
		msg := &deadlock.CCitadelUserMsg_PostProcessingAnim{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 333: //deadlock.CitadelUserMessageIds_k_EUserMsg_DeathReplayData
		msg := &deadlock.CCitadelUserMsg_DeathReplayData{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 334: //deadlock.CitadelUserMessageIds_k_EUserMsg_PlayerLifetimeStatInfo
		msg := &deadlock.CCitadelUserMsg_PlayerLifetimeStatInfo{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 336: //deadlock.CitadelUserMessageIds_k_EUserMsg_ForceShopClosed
		msg := &deadlock.CCitadelUserMsg_ForceShopClosed{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 337: //deadlock.CitadelUserMessageIds_k_EUserMsg_StaminaDrained
		msg := &deadlock.CCitadelUserMsg_StaminaDrained{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 338: //deadlock.CitadelUserMessageIds_k_EUserMsg_AbilityNotify
		msg := &deadlock.CCitadelUserMessage_AbilityNotify{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 339: //deadlock.CitadelUserMessageIds_k_EUserMsg_GetDamageStatsResponse
		msg := &deadlock.CCitadelUserMsg_GetDamageStatsResponse{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 343: //deadlock.CitadelUserMessageIds_k_EUserMsg_ParticipantSetSoundEventParams
		msg := &deadlock.CCitadelUserMsg_ParticipantSetSoundEventParams{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 344: //deadlock.CitadelUserMessageIds_k_EUserMsg_ParticipantSetLibraryStackFields
		msg := &deadlock.CCitadelUserMsg_ParticipantSetLibraryStackFields{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 345: //deadlock.CitadelUserMessageIds_k_EUserMsg_CurrencyChanged
		msg := &deadlock.CCitadelUserMessage_CurrencyChanged{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 346: //deadlock.CitadelUserMessageIds_k_EUserMsg_GameOver
		msg := &deadlock.CCitadelUserMessage_GameOver{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)
	case 347: //deadlock.CitadelUserMessageIds_k_EUserMsg_BossKilled
		msg := &deadlock.CCitadelUserMsg_BossKilled{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return "ERROR"
		}
		return prototext.Format(msg)

	}
	return "UNKNOWN:" + strconv.Itoa(int(t))
}

func (c *Callbacks) callByPacketType(t int32, buf []byte) error {
	switch t {
	case 0: // deadlock.NET_Messages_net_NOP
		if c.onCNETMsg_NOP == nil {
			return nil
		}

		msg := &deadlock.CNETMsg_NOP{}
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
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
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSosStopSoundEventHash {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 300: // deadlock.CitadelUserMessageIds_k_EUserMsg_Damage
		if c.onCCitadelUserMessage_Damage == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMessage_Damage{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMessage_Damage {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 303: // deadlock.CitadelUserMessageIds_k_EUserMsg_MapPing
		if c.onCCitadelUserMsg_MapPing == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_MapPing{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_MapPing {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 304: // deadlock.CitadelUserMessageIds_k_EUserMsg_TeamRewards
		if c.onCCitadelUserMsg_TeamRewards == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_TeamRewards{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_TeamRewards {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 308: // deadlock.CitadelUserMessageIds_k_EUserMsg_TriggerDamageFlash
		if c.onCCitadelUserMsg_TriggerDamageFlash == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_TriggerDamageFlash{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_TriggerDamageFlash {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 309: // deadlock.CitadelUserMessageIds_k_EUserMsg_AbilitiesChanged
		if c.onCCitadelUserMsg_AbilitiesChanged == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_AbilitiesChanged{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_AbilitiesChanged {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 310: // deadlock.CitadelUserMessageIds_k_EUserMsg_RecentDamageSummary
		if c.onCCitadelUserMsg_RecentDamageSummary == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_RecentDamageSummary{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_RecentDamageSummary {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 311: // deadlock.CitadelUserMessageIds_k_EUserMsg_SpectatorTeamChanged
		if c.onCCitadelUserMsg_SpectatorTeamChanged == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_SpectatorTeamChanged{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_SpectatorTeamChanged {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 312: // deadlock.CitadelUserMessageIds_k_EUserMsg_ChatWheel
		if c.onCCitadelUserMsg_ChatWheel == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_ChatWheel{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ChatWheel {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 313: // deadlock.CitadelUserMessageIds_k_EUserMsg_GoldHistory
		if c.onCCitadelUserMsg_GoldHistory == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_GoldHistory{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_GoldHistory {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 314: // deadlock.CitadelUserMessageIds_k_EUserMsg_ChatMsg
		if c.onCCitadelUserMsg_ChatMsg == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_ChatMsg{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ChatMsg {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 315: // deadlock.CitadelUserMessageIds_k_EUserMsg_QuickResponse
		if c.onCCitadelUserMsg_QuickResponse == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_QuickResponse{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_QuickResponse {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 316: // deadlock.CitadelUserMessageIds_k_EUserMsg_PostMatchDetails
		if c.onCCitadelUserMsg_PostMatchDetails == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_PostMatchDetails{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_PostMatchDetails {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 317: // deadlock.CitadelUserMessageIds_k_EUserMsg_ChatEvent
		if c.onCCitadelUserMsg_ChatEvent == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_ChatEvent{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ChatEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 318: // deadlock.CitadelUserMessageIds_k_EUserMsg_AbilityInterrupted
		if c.onCCitadelUserMsg_AbilityInterrupted == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_AbilityInterrupted{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_AbilityInterrupted {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 319: // deadlock.CitadelUserMessageIds_k_EUserMsg_HeroKilled
		if c.onCCitadelUserMsg_HeroKilled == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_HeroKilled{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_HeroKilled {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 320: // deadlock.CitadelUserMessageIds_k_EUserMsg_ReturnIdol
		if c.onCCitadelUserMsg_ReturnIdol == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_ReturnIdol{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ReturnIdol {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 321: // deadlock.CitadelUserMessageIds_k_EUserMsg_SetClientCameraAngles
		if c.onCCitadelUserMsg_SetClientCameraAngles == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_SetClientCameraAngles{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_SetClientCameraAngles {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 322: // deadlock.CitadelUserMessageIds_k_EUserMsg_MapLine
		if c.onCCitadelUserMsg_MapLine == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_MapLine{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_MapLine {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 323: // deadlock.CitadelUserMessageIds_k_EUserMsg_BulletHit
		if c.onCCitadelUserMessage_BulletHit == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMessage_BulletHit{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMessage_BulletHit {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 324: // deadlock.CitadelUserMessageIds_k_EUserMsg_ObjectiveMask
		if c.onCCitadelUserMessage_ObjectiveMask == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMessage_ObjectiveMask{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMessage_ObjectiveMask {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 325: // deadlock.CitadelUserMessageIds_k_EUserMsg_ModifierApplied
		if c.onCCitadelUserMessage_ModifierApplied == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMessage_ModifierApplied{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMessage_ModifierApplied {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 326: // deadlock.CitadelUserMessageIds_k_EUserMsg_CameraController
		if c.onCCitadelUserMsg_CameraController == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_CameraController{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_CameraController {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 327: // deadlock.CitadelUserMessageIds_k_EUserMsg_AuraModifierApplied
		if c.onCCitadelUserMessage_AuraModifierApplied == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMessage_AuraModifierApplied{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMessage_AuraModifierApplied {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 329: // deadlock.CitadelUserMessageIds_k_EUserMsg_ObstructedShotFired
		if c.onCCitadelUserMsg_ObstructedShotFired == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_ObstructedShotFired{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ObstructedShotFired {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 330: // deadlock.CitadelUserMessageIds_k_EUserMsg_AbilityLateFailure
		if c.onCCitadelUserMsg_AbilityLateFailure == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_AbilityLateFailure{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_AbilityLateFailure {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 331: // deadlock.CitadelUserMessageIds_k_EUserMsg_AbilityPing
		if c.onCCitadelUserMsg_AbilityPing == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_AbilityPing{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_AbilityPing {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 332: // deadlock.CitadelUserMessageIds_k_EUserMsg_PostProcessingAnim
		if c.onCCitadelUserMsg_PostProcessingAnim == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_PostProcessingAnim{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_PostProcessingAnim {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 333: // deadlock.CitadelUserMessageIds_k_EUserMsg_DeathReplayData
		if c.onCCitadelUserMsg_DeathReplayData == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_DeathReplayData{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_DeathReplayData {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 334: // deadlock.CitadelUserMessageIds_k_EUserMsg_PlayerLifetimeStatInfo
		if c.onCCitadelUserMsg_PlayerLifetimeStatInfo == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_PlayerLifetimeStatInfo{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_PlayerLifetimeStatInfo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 336: // deadlock.CitadelUserMessageIds_k_EUserMsg_ForceShopClosed
		if c.onCCitadelUserMsg_ForceShopClosed == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_ForceShopClosed{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ForceShopClosed {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 337: // deadlock.CitadelUserMessageIds_k_EUserMsg_StaminaDrained
		if c.onCCitadelUserMsg_StaminaDrained == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_StaminaDrained{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_StaminaDrained {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 338: // deadlock.CitadelUserMessageIds_k_EUserMsg_AbilityNotify
		if c.onCCitadelUserMessage_AbilityNotify == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMessage_AbilityNotify{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMessage_AbilityNotify {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 339: // deadlock.CitadelUserMessageIds_k_EUserMsg_GetDamageStatsResponse
		if c.onCCitadelUserMsg_GetDamageStatsResponse == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_GetDamageStatsResponse{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_GetDamageStatsResponse {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 343: // deadlock.CitadelUserMessageIds_k_EUserMsg_ParticipantSetSoundEventParams
		if c.onCCitadelUserMsg_ParticipantSetSoundEventParams == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_ParticipantSetSoundEventParams{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ParticipantSetSoundEventParams {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 344: // deadlock.CitadelUserMessageIds_k_EUserMsg_ParticipantSetLibraryStackFields
		if c.onCCitadelUserMsg_ParticipantSetLibraryStackFields == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_ParticipantSetLibraryStackFields{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ParticipantSetLibraryStackFields {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 345: // deadlock.CitadelUserMessageIds_k_EUserMsg_CurrencyChanged
		if c.onCCitadelUserMessage_CurrencyChanged == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMessage_CurrencyChanged{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMessage_CurrencyChanged {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 346: // deadlock.CitadelUserMessageIds_k_EUserMsg_GameOver
		if c.onCCitadelUserMessage_GameOver == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMessage_GameOver{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMessage_GameOver {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 347: // deadlock.CitadelUserMessageIds_k_EUserMsg_BossKilled
		if c.onCCitadelUserMsg_BossKilled == nil {
			return nil
		}

		msg := &deadlock.CCitadelUserMsg_BossKilled{}
		if err := proto.Unmarshal(buf, msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_BossKilled {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	}

	return nil
}
