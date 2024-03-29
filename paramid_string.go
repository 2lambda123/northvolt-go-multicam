// Code generated by "stringer -type ParamID"; DO NOT EDIT.

package multicam

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AcquisitionModeParam-55640064]
	_ = x[BoardCountParam-1015808]
	_ = x[BoardNameParam-32768]
	_ = x[BoardNameChangeParam-65536]
	_ = x[BoardIdentifierParam-49152]
	_ = x[BoardIndexParam-32358400]
	_ = x[BoardPCIPositionParam-16384]
	_ = x[BoardSerialNumberParam-81920]
	_ = x[BoardTypeParam-98304]
	_ = x[BufferPitchParam-54657024]
	_ = x[BufferSizeParam-54607872]
	_ = x[CamFileParam-180224]
	_ = x[ChannelStateParam-245760]
	_ = x[ClusterParam-196608]
	_ = x[ColorFormatParam-36438016]
	_ = x[ConnectorParam-11173888]
	_ = x[DriverIndexParam-0]
	_ = x[ElapsedPgParam-56606720]
	_ = x[EncoderPitchParam-9699328]
	_ = x[ErrorHandlingParam-802816]
	_ = x[ErrorLogParam-1327104]
	_ = x[ForceTrigParam-819200]
	_ = x[ImageFlipXParam-21954560]
	_ = x[ImageSizeXParam-8568832]
	_ = x[ImageSizeYParam-8585216]
	_ = x[LinePitchParam-9748480]
	_ = x[LineRateModeParam-21757952]
	_ = x[MetadataContentParam-177750016]
	_ = x[MetadataGPPCInputLineParam-183681024]
	_ = x[MetadataGPPCLocationParam-183713792]
	_ = x[MetadataGPPCResetLineParam-183730176]
	_ = x[MetadataInsertionParam-177635328]
	_ = x[MetadataLocationParam-181125120]
	_ = x[MinBufferPitchParam-54640640]
	_ = x[OutputConfigParam-28508160]
	_ = x[OutputStateParam-28540928]
	_ = x[PeriodUsParam-21774336]
	_ = x[SerialNumberParam-81920]
	_ = x[SeqLengthFrParam-55820288]
	_ = x[SeqLengthPgParam-55836672]
	_ = x[SeqLengthLnParam-55721984]
	_ = x[SignalEnableParam-393216]
	_ = x[SurfaceAddrParam-458752]
	_ = x[SurfaceCountParam-1343488]
	_ = x[SurfaceIndexParam-278528]
	_ = x[SurfacePitchParam-475136]
	_ = x[SurfaceSizeParam-442368]
	_ = x[SurfaceStateParam-507904]
	_ = x[AnySignal-0]
	_ = x[StartAcquisitionSignal-11]
	_ = x[EndAcquisitionSignal-10]
	_ = x[AcquisitionFailureSignal-7]
	_ = x[ClusterUnavailableSignal-8]
	_ = x[EndChannelActivitySignal-12]
	_ = x[FrameTriggerViolationSignal-4]
	_ = x[SurfaceProcessingSignal-1]
	_ = x[SurfaceFilledSignal-2]
	_ = x[StartExposureSignal-5]
	_ = x[EndExposureSignal-6]
	_ = x[UnrecoverableOverrunSignal-3]
	_ = x[ReleaseSignal-9]
}

const _ParamID_name = "DriverIndexParamSurfaceProcessingSignalSurfaceFilledSignalUnrecoverableOverrunSignalFrameTriggerViolationSignalStartExposureSignalEndExposureSignalAcquisitionFailureSignalClusterUnavailableSignalReleaseSignalEndAcquisitionSignalStartAcquisitionSignalEndChannelActivitySignalBoardPCIPositionParamBoardNameParamBoardIdentifierParamBoardNameChangeParamBoardSerialNumberParamBoardTypeParamCamFileParamClusterParamChannelStateParamSurfaceIndexParamSignalEnableParamSurfaceSizeParamSurfaceAddrParamSurfacePitchParamSurfaceStateParamErrorHandlingParamForceTrigParamBoardCountParamErrorLogParamSurfaceCountParamImageSizeXParamImageSizeYParamEncoderPitchParamLinePitchParamConnectorParamLineRateModeParamPeriodUsParamImageFlipXParamOutputConfigParamOutputStateParamBoardIndexParamColorFormatParamBufferSizeParamMinBufferPitchParamBufferPitchParamAcquisitionModeParamSeqLengthLnParamSeqLengthFrParamSeqLengthPgParamElapsedPgParamMetadataInsertionParamMetadataContentParamMetadataLocationParamMetadataGPPCInputLineParamMetadataGPPCLocationParamMetadataGPPCResetLineParam"

var _ParamID_map = map[ParamID]string{
	0:         _ParamID_name[0:16],
	1:         _ParamID_name[16:39],
	2:         _ParamID_name[39:58],
	3:         _ParamID_name[58:84],
	4:         _ParamID_name[84:111],
	5:         _ParamID_name[111:130],
	6:         _ParamID_name[130:147],
	7:         _ParamID_name[147:171],
	8:         _ParamID_name[171:195],
	9:         _ParamID_name[195:208],
	10:        _ParamID_name[208:228],
	11:        _ParamID_name[228:250],
	12:        _ParamID_name[250:274],
	16384:     _ParamID_name[274:295],
	32768:     _ParamID_name[295:309],
	49152:     _ParamID_name[309:329],
	65536:     _ParamID_name[329:349],
	81920:     _ParamID_name[349:371],
	98304:     _ParamID_name[371:385],
	180224:    _ParamID_name[385:397],
	196608:    _ParamID_name[397:409],
	245760:    _ParamID_name[409:426],
	278528:    _ParamID_name[426:443],
	393216:    _ParamID_name[443:460],
	442368:    _ParamID_name[460:476],
	458752:    _ParamID_name[476:492],
	475136:    _ParamID_name[492:509],
	507904:    _ParamID_name[509:526],
	802816:    _ParamID_name[526:544],
	819200:    _ParamID_name[544:558],
	1015808:   _ParamID_name[558:573],
	1327104:   _ParamID_name[573:586],
	1343488:   _ParamID_name[586:603],
	8568832:   _ParamID_name[603:618],
	8585216:   _ParamID_name[618:633],
	9699328:   _ParamID_name[633:650],
	9748480:   _ParamID_name[650:664],
	11173888:  _ParamID_name[664:678],
	21757952:  _ParamID_name[678:695],
	21774336:  _ParamID_name[695:708],
	21954560:  _ParamID_name[708:723],
	28508160:  _ParamID_name[723:740],
	28540928:  _ParamID_name[740:756],
	32358400:  _ParamID_name[756:771],
	36438016:  _ParamID_name[771:787],
	54607872:  _ParamID_name[787:802],
	54640640:  _ParamID_name[802:821],
	54657024:  _ParamID_name[821:837],
	55640064:  _ParamID_name[837:857],
	55721984:  _ParamID_name[857:873],
	55820288:  _ParamID_name[873:889],
	55836672:  _ParamID_name[889:905],
	56606720:  _ParamID_name[905:919],
	177635328: _ParamID_name[919:941],
	177750016: _ParamID_name[941:961],
	181125120: _ParamID_name[961:982],
	183681024: _ParamID_name[982:1008],
	183713792: _ParamID_name[1008:1033],
	183730176: _ParamID_name[1033:1059],
}

func (i ParamID) String() string {
	if str, ok := _ParamID_map[i]; ok {
		return str
	}
	return "ParamID(" + strconv.FormatInt(int64(i), 10) + ")"
}
