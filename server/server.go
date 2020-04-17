package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"

	tonlib "github.com/mercuryoio/tonlib-go/v2"
	"github.com/tonradar/ton-api/config"
	pb "github.com/tonradar/ton-api/proto"
)

const (
	TestPassword = "test_password"
)

type TonApiServer struct {
	conf    config.TonAPIConfig
	api     *tonlib.Client
	apiLock sync.Mutex
	key     tonlib.InputKey
}

func NewTonApiServer(conf config.TonAPIConfig) (*TonApiServer, error) {
	options, err := tonlib.ParseConfigFile(conf.TonlibCfgPath)
	if err != nil {
		return nil, fmt.Errorf("Config file not found, error: %v", err)
	}

	req := tonlib.TonInitRequest{
		"init",
		*options,
	}

	client, err := tonlib.NewClient(&req, tonlib.Config{}, 10, true, 9)
	if err != nil {
		return nil, fmt.Errorf("Init client error, error: %v", err)
	}

	loc := tonlib.SecureBytes(TestPassword)
	mem := tonlib.SecureBytes(TestPassword)
	seed := tonlib.SecureBytes("")

	pKey, err := client.CreateNewKey(loc, mem, seed)
	if err != nil {
		return nil, fmt.Errorf("Ton create key for init wallet error", err)
	}

	inputKey := tonlib.InputKey{
		"inputKeyRegular",
		base64.StdEncoding.EncodeToString(loc),
		tonlib.TONPrivateKey{
			pKey.PublicKey,
			pKey.Secret,
		},
	}

	return &TonApiServer{
		conf: conf,
		api:  client,
		key:  inputKey,
	}, nil
}

func (s *TonApiServer) FetchTransactions(ctx context.Context, in *pb.FetchTransactionsRequest) (*pb.FetchTransactionsResponse, error) {
	s.apiLock.Lock()
	resp, err := s.api.RawGetTransactions(*tonlib.NewAccountAddress(in.Address), *tonlib.NewInternalTransactionId(in.Hash, tonlib.JSONInt64(in.Lt)), s.key)
	if err != nil {
		// need to restart container
		//panic(err)
		s.api.UpdateTonConnection()
		return nil, err
	}
	s.apiLock.Unlock()

	trxs := make([]*pb.Transaction, 0)

	for _, trx := range resp.Transactions {
		msgData := trx.InMsg.MsgData.(map[string]interface{})
		msgDataText := ""
		if msgData["@type"] == "msg.dataText" {
			msgDataText = msgData["text"].(string)
		}
		inMsg := pb.RawMessage{
			BodyHash:    trx.InMsg.BodyHash,
			CreatedLt:   int64(trx.InMsg.CreatedLt),
			Destination: trx.InMsg.Destination.AccountAddress,
			FwdFee:      int64(trx.InMsg.FwdFee),
			IhrFee:      int64(trx.InMsg.IhrFee),
			Message:     msgDataText,
			Source:      trx.InMsg.Source.AccountAddress,
			Value:       int64(trx.InMsg.Value),
		}

		outMsgs := make([]*pb.RawMessage, 0)
		for _, msg := range trx.OutMsgs {
			msgData := msg.MsgData.(map[string]interface{})
			msgDataText := ""
			if msgData["@type"] == "msg.dataText" {
				msgDataText = msgData["text"].(string)
			}
			tmp := &pb.RawMessage{
				BodyHash:    msg.BodyHash,
				CreatedLt:   int64(msg.CreatedLt),
				Destination: msg.Destination.AccountAddress,
				FwdFee:      int64(msg.FwdFee),
				IhrFee:      int64(msg.IhrFee),
				Message:     msgDataText,
				Source:      msg.Source.AccountAddress,
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
	resp, err := s.api.RawGetAccountState(*tonlib.NewAccountAddress(in.AccountAddress))
	if err != nil {
		// need to restart container
		//panic(err)
		s.api.UpdateTonConnection()
		return nil, err
	}
	s.apiLock.Unlock()

	if !isRawFullAccountState(resp) {
		return nil, errors.New("Invalid return type for GetAccountState")
	}

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

func (s *TonApiServer) GetActiveBets(ctx context.Context, in *pb.GetActiveBetsRequest) (*pb.GetActiveBetsResponse, error) {
	s.apiLock.Lock()
	address := tonlib.NewAccountAddress(s.conf.ContractAddr)
	smcInfo, err := s.api.SmcLoad(*address)
	if err != nil {
		s.api.UpdateTonConnection()
		return nil, err
	}
	s.apiLock.Unlock()

	methodName := "active_bets"
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

	if len(res) == 0 {
		return nil, fmt.Errorf("empty response")
	}

	var bets CustomTvmStackEntry
	asBytes, err := json.Marshal(res[0])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(asBytes, &bets)
	if err != nil {
		return nil, err
	}

	var activeBets []*pb.ActiveBet
	for _, element := range bets.List.Elements {
		betIdStr := element.Tuple.Elements[0].(map[string]interface{})["number"].(map[string]interface{})["number"].(string)
		betId, err := strconv.Atoi(betIdStr)
		if err != nil {
			return nil, err
		}

		other := element.Tuple.Elements[1]

		var tmp _CustomTvmStackEntryTuple
		asBytes, err := json.Marshal(other)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(asBytes, &tmp)
		if err != nil {
			return nil, err
		}

		params := tmp.Tuple.Elements
		rollUnder, err := strconv.Atoi(params[0].Number.(map[string]interface{})["number"].(string))
		if err != nil {
			return nil, err
		}
		amount, err := strconv.Atoi(params[1].Number.(map[string]interface{})["number"].(string))
		if err != nil {
			return nil, err
		}
		wc1, err := strconv.Atoi(params[2].Number.(map[string]interface{})["number"].(string))
		if err != nil {
			return nil, err
		}
		address1 := params[3].Number.(map[string]interface{})["number"].(string)
		if err != nil {
			return nil, err
		}
		wc2, err := strconv.Atoi(params[4].Number.(map[string]interface{})["number"].(string))
		if err != nil {
			return nil, err
		}
		address2 := params[5].Number.(map[string]interface{})["number"].(string)
		if err != nil {
			return nil, err
		}
		refBonus, err := strconv.Atoi(params[6].Number.(map[string]interface{})["number"].(string))
		if err != nil {
			return nil, err
		}
		seed := params[7].Number.(map[string]interface{})["number"].(string)

		bet := &pb.ActiveBet{
			Id:            int32(betId),
			RollUnder:     int32(rollUnder),
			Amount:        int64(amount),
			PlayerAddress: &pb.TonAddress{Workchain: int32(wc1), Address: address1},
			RefAddress:    &pb.TonAddress{Workchain: int32(wc2), Address: address2},
			RefBonus:      int64(refBonus),
			Seed:          seed,
		}

		activeBets = append(activeBets, bet)
	}

	return &pb.GetActiveBetsResponse{
		Bets: activeBets,
	}, nil
}

func (s *TonApiServer) GetBetSeed(ctx context.Context, in *pb.GetBetSeedRequest) (*pb.GetBetSeedResponse, error) {
	s.apiLock.Lock()
	address := tonlib.NewAccountAddress(s.conf.ContractAddr)
	smcInfo, err := s.api.SmcLoad(*address)
	if err != nil {
		// need to restart container
		//panic(err)
		s.api.UpdateTonConnection()
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

	if len(res) == 0 {
		return nil, fmt.Errorf("Вad result smartcontract get method")
	}

	resNum := res[0].(map[string]interface{})["number"].(map[string]interface{})["number"].(string)

	return &pb.GetBetSeedResponse{
		Seed: resNum,
	}, nil
}

func (s *TonApiServer) GetSeqno(ctx context.Context, in *pb.GetSeqnoRequest) (*pb.GetSeqnoResponse, error) {
	s.apiLock.Lock()
	address := tonlib.NewAccountAddress(s.conf.ContractAddr)
	smcInfo, err := s.api.SmcLoad(*address)
	if err != nil {
		// need to restart container
		//panic(err)
		s.api.UpdateTonConnection()
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

	resNum := res[0].(map[string]interface{})["number"].(map[string]interface{})["number"].(string)

	return &pb.GetSeqnoResponse{
		Seqno: resNum,
	}, nil
}

func (s *TonApiServer) SendMessage(ctx context.Context, in *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	s.apiLock.Lock()
	resp, err := s.api.RawSendMessage(in.Body)
	if err != nil {
		// need to restart container
		//panic(err)
		s.api.UpdateTonConnection()
		return nil, err
	}
	s.apiLock.Unlock()

	return &pb.SendMessageResponse{
		Ok: resp.Type,
	}, nil
}

func (s *TonApiServer) runGetMethod(id int64, method interface{}, stack []tonlib.TvmStackEntry) ([]tonlib.TvmStackEntry, error) {
	s.apiLock.Lock()
	resp, err := s.api.SmcRunGetMethod(id, method, stack)
	if err != nil {
		// need to restart container
		//panic(err)
		s.api.UpdateTonConnection()
		return nil, err
	}
	s.apiLock.Unlock()

	return resp.Stack, nil
}
