package server

import (
	"context"
	"fmt"
	"github.com/mercuryoio/tonlib-go"
	"ton-api/config"
	pb "ton-api/proto"
)

type TonApiServer struct {
	conf config.Config
	api  *tonlib.Client
}

func NewTonApiServer(conf config.Config) (*TonApiServer, error) {
	options, err := tonlib.ParseConfigFile(conf.TonAPI.TonConfig)
	if err != nil {
		return nil, fmt.Errorf("Config file not found, error: %v", err)
	}

	req := tonlib.TonInitRequest{
		"init",
		*options,
	}

	client, err := tonlib.NewClient(&req, tonlib.Config{}, 5)
	if err != nil {
		return nil, fmt.Errorf("Init client error, error: %v", err)
	}

	return &TonApiServer{
		conf: conf,
		api:  client,
	}, nil
}

func (s *TonApiServer) FetchTransactions(ctx context.Context, in *pb.FetchTransactionsRequest) (*pb.FetchTransactionsResponse, error) {
	resp, err := s.api.RawGetTransactions(tonlib.NewAccountAddress(in.Address), tonlib.NewInternalTransactionId(in.Hash, tonlib.JSONInt64(in.Lt)))
	if err != nil {
		return nil, err
	}

	trxs := make([]*pb.Transaction, 10)

	for _, trx := range resp.Transactions {
		inMsg := pb.RawMessage{
			BodyHash:    trx.InMsg.BodyHash,
			CreatedLt:   int64(trx.InMsg.CreatedLt),
			Destination: trx.InMsg.Destination,
			FwdFee:      int64(trx.InMsg.FwdFee),
			IhrFee:      int64(trx.InMsg.IhrFee),
			Message:     trx.InMsg.Message,
			Source:      trx.InMsg.Source,
			Value:       int64(trx.InMsg.Value),
		}

		outMsgs := make([]*pb.RawMessage, 5)
		for _, msg := range trx.OutMsgs {
			tmp := &pb.RawMessage{
				BodyHash:    msg.BodyHash,
				CreatedLt:   int64(msg.CreatedLt),
				Destination: msg.Destination,
				FwdFee:      int64(msg.FwdFee),
				IhrFee:      int64(msg.IhrFee),
				Message:     msg.Message,
				Source:      msg.Source,
				Value:       int64(msg.Value),
			}
			outMsgs = append(outMsgs, tmp)
		}

		transactionId := pb.InternalTransactionId{
			Hash: trx.TransactionId.Hash,
			Lt:   int64(trx.TransactionId.Lt),
		}

		tmp := pb.Transaction{
			Data:          trx.Data,
			Fee:           int64(trx.Fee),
			InMsg:         &inMsg,
			OtherFee:      int64(trx.OtherFee),
			OutMsgs:       outMsgs,
			StorageFee:    int64(trx.StorageFee),
			TransactionId: &transactionId,
		}
		trxs = append(trxs, &tmp)
	}

	return &pb.FetchTransactionsResponse{
		Items: trxs,
	}, nil
}
