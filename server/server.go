package server

import (
	"context"
	"fmt"
	"github.com/tonradar/ton-api/config"
	pb "github.com/tonradar/ton-api/proto"
	"github.com/tonradar/tonlib-go"
	"strconv"
	"sync"
)

type TonApiServer struct {
	conf    config.Config
	api     *tonlib.Client
	apiLock sync.Mutex
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

	client, err := tonlib.NewClient(&req, tonlib.Config{}, 10)
	if err != nil {
		return nil, fmt.Errorf("Init client error, error: %v", err)
	}

	return &TonApiServer{
		conf: conf,
		api:  client,
	}, nil
}

func (s *TonApiServer) FetchTransactions(ctx context.Context, in *pb.FetchTransactionsRequest) (*pb.FetchTransactionsResponse, error) {
	s.apiLock.Lock()
	resp, err := s.api.RawGetTransactions(tonlib.NewAccountAddress(in.Address), tonlib.NewInternalTransactionId(in.Hash, tonlib.JSONInt64(in.Lt)))
	if err != nil {
		return nil, err
	}
	s.apiLock.Unlock()

	trxs := make([]*pb.Transaction, 0)

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

		outMsgs := make([]*pb.RawMessage, 0)
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

func (s *TonApiServer) GetAccountState(ctx context.Context, in *pb.GetAccountStateRequest) (*pb.GetAccountStateResponse, error) {
	s.apiLock.Lock()
	resp, err := s.api.RawGetAccountState(tonlib.NewAccountAddress(in.AccountAddress))
	if err != nil {
		return nil, err
	}
	s.apiLock.Unlock()

	transactionId := &pb.InternalTransactionId{
		Hash: resp.LastTransactionId.Hash,
		Lt:   int64(resp.LastTransactionId.Lt),
	}

	return &pb.GetAccountStateResponse{
		Balance:           int64(resp.Balance),
		Code:              resp.Code,
		Data:              resp.Data,
		FrozenHash:        resp.FrozenHash,
		LastTransactionId: transactionId,
		SyncUtime:         resp.SyncUtime,
	}, nil
}

func (s *TonApiServer) GetBetSeed(ctx context.Context, in *pb.GetBetSeedRequest) (*pb.GetBetSeedResponse, error) {
	s.apiLock.Lock()
	address := tonlib.NewAccountAddress(s.conf.TonAPI.DiceAddress)
	smcInfo, err := s.api.SmcLoad(address)
	if err != nil {
		return nil, err
	}
	s.apiLock.Unlock()

	methodName := "get_bet_seed"
	methodID := struct {
		Type  string `json:"@type"`
		Extra string `json:"@extra"`
		Name  string `json:"name"`
	}{
		Type: "smc.methodIdName",
		Name: methodName,
	}

	betId := struct {
		Type   string `json:"@type"`
		Extra  string `json:"@extra"`
		Number string `json:"number"`
	}{
		Type:   "tvm.numberDecimal",
		Number: strconv.Itoa(int(in.BetId)),
	}

	betID := struct {
		Type   string      `json:"@type"`
		Extra  string      `json:"@extra"`
		Number interface{} `json:"number"`
	}{
		Type:   "tvm.stackEntryNumber",
		Number: &betId,
	}

	stack := make([]tonlib.TvmStackEntry, 0)
	stack = append(stack, betID)

	res, err := s.runGetMethod(smcInfo.Id, methodID, stack)
	if err != nil {
		return nil, err
	}
	resNum := res[0].(map[string]interface{})["number"].(map[string]interface{})["number"].(string)

	return &pb.GetBetSeedResponse{
		Seed: resNum,
	}, nil
}

func (s *TonApiServer) GetSeqno(ctx context.Context, in *pb.GetSeqnoRequest) (*pb.GetSeqnoResponse, error) {
	s.apiLock.Lock()
	address := tonlib.NewAccountAddress(s.conf.TonAPI.DiceAddress)
	smcInfo, err := s.api.SmcLoad(address)
	if err != nil {
		return nil, err
	}
	s.apiLock.Unlock()

	methodName := "get_seqno"
	methodID := struct {
		Type  string `json:"@type"`
		Extra string `json:"@extra"`
		Name  string `json:"name"`
	}{
		Type: "smc.methodIdName",
		Name: methodName,
	}

	stack := make([]tonlib.TvmStackEntry, 0)

	res, err := s.runGetMethod(smcInfo.Id, methodID, stack)
	if err != nil {
		return nil, err
	}

	fmt.Println("seqno:", res)

	resNum := res[0].(map[string]interface{})["number"].(map[string]interface{})["number"].(int32)

	return &pb.GetSeqnoResponse{
		Seqno: resNum,
	}, nil
}

func (s *TonApiServer) SendMessage(ctx context.Context, in *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	s.apiLock.Lock()
	resp, err := s.api.RawSendMessage(in.Body)
	if err != nil {
		return nil, err
	}
	s.apiLock.Unlock()

	fmt.Println("send message response:", resp)

	return &pb.SendMessageResponse{
		Ok: resp.Type,
	}, nil
}

func (s *TonApiServer) runGetMethod(id int64, method interface{}, stack []tonlib.TvmStackEntry) ([]tonlib.TvmStackEntry, error) {
	s.apiLock.Lock()
	resp, err := s.api.SmcRunGetMethod(id, method, stack)
	if err != nil {
		return nil, err
	}
	s.apiLock.Unlock()

	return resp.Stack, nil
}
