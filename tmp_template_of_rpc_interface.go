func (s MvpServer) GetAllDeskSets(ctx context.Context, req *pb.GetAllDeskSetsReq) (*pb.GetAllDeskSetsRes, error) {
	logger.Info("GetAllDeskSets", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetAllDeskSetsRes
}
