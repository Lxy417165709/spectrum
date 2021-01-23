func (s MvpServer) DeleteFavorForGood(ctx context.Context, req *pb.DeleteFavorForGoodReq) (*pb.DeleteFavorForGoodRes, error) {
	logger.Info("DeleteFavorForGood", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.DeleteFavorForGoodRes
}
