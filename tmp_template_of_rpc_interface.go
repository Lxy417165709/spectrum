func (s MvpServer) GetAllDeskClasses(ctx context.Context, req *pb.GetAllDeskClassesReq) (*pb.GetAllDeskClassesRes, error) {
	logger.Info("GetAllDeskClasses", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetAllDeskClassesRes
}
