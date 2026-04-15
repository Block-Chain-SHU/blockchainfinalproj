package chaincode

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// 定义合约结构体
type SmartContract struct {
	contractapi.Contract
}

// 注册用户
func (s *SmartContract) RegisterUser(ctx contractapi.TransactionContextInterface, userID string, userType string, realInfoHash string) error {
	user := User{
		UserID:       userID,
		UserType:     userType,
		RealInfoHash: realInfoHash,
		TraceList:    []*TraceRecord{},
	}
	userAsBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(userID, userAsBytes)
	if err != nil {
		return err
	}
	return nil
}

// 培养信息上链，传入用户ID和分阶段追踪信息
func (s *SmartContract) Uplink(ctx contractapi.TransactionContextInterface, userID string, student_trace_id string, arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 string) (string, error) {
	// 获取用户类型
	userType, err := s.GetUserType(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("failed to get user type: %v", err)
	}

	// 通过学籍追踪码获取历史追踪信息
	traceBytes, err := ctx.GetStub().GetState(student_trace_id)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	// 将历史记录转换为追踪结构体
	var traceRecord TraceRecord
	if traceBytes != nil {
		err = json.Unmarshal(traceBytes, &traceRecord)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal trace record: %v", err)
		}
	}

	//获取时间戳,修正时区
	txtime, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return "", fmt.Errorf("failed to read TxTimestamp: %v", err)
	}
	timeLocation, _ := time.LoadLocation("Asia/Shanghai") // 选择你所在的时区
	time := time.Unix(txtime.Seconds, 0).In(timeLocation).Format("2006-01-02 15:04:05")

	// 获取交易ID
	txid := ctx.GetStub().GetTxID()
	// 设置学籍追踪码
	traceRecord.Student_trace_id = student_trace_id
	// 不同用户类型的上链的参数不一致
	switch userType {
	// 学生
	case "学生":
		traceRecord.Farmer_input.Fa_studentName = arg1
		traceRecord.Farmer_input.Fa_major = arg2
		traceRecord.Farmer_input.Fa_enrollTime = arg3
		traceRecord.Farmer_input.Fa_graduationTime = arg4
		traceRecord.Farmer_input.Fa_mentorName = arg5
		traceRecord.Farmer_input.Fa_img = arg6
		traceRecord.Farmer_input.Fa_Txid = txid
		traceRecord.Farmer_input.Fa_Timestamp = time
	// 课程/实验室
	case "课程/实验室":
		traceRecord.Factory_input.Fac_courseOrLabName = arg1
		traceRecord.Factory_input.Fac_courseBatch = arg2
		traceRecord.Factory_input.Fac_courseTime = arg3
		traceRecord.Factory_input.Fac_platformName = arg4
		traceRecord.Factory_input.Fac_contactNumber = arg5
		traceRecord.Factory_input.Fac_img = arg6
		traceRecord.Factory_input.Fac_Txid = txid
		traceRecord.Factory_input.Fac_Timestamp = time
	// 项目实践/竞赛导师
	case "项目实践/竞赛导师":
		traceRecord.Driver_input.Dr_mentorName = arg1
		traceRecord.Driver_input.Dr_experienceYears = arg2
		traceRecord.Driver_input.Dr_phone = arg3
		traceRecord.Driver_input.Dr_projectCode = arg4
		traceRecord.Driver_input.Dr_contribution = arg5
		traceRecord.Driver_input.Dr_img = arg6
		traceRecord.Driver_input.Dr_Txid = txid
		traceRecord.Driver_input.Dr_Timestamp = time
	// 就业单位/研究生去向
	case "就业单位/研究生去向":
		traceRecord.Shop_input.Sh_offerTime = arg1
		traceRecord.Shop_input.Sh_onboardTime = arg2
		traceRecord.Shop_input.Sh_orgName = arg3
		traceRecord.Shop_input.Sh_orgAddress = arg4
		traceRecord.Shop_input.Sh_contactPhone = arg5
		traceRecord.Shop_input.Sh_img = arg6
		traceRecord.Shop_input.Sh_Txid = txid
		traceRecord.Shop_input.Sh_Timestamp = time

	}

	// 将追踪信息转换为 json
	traceAsBytes, err := json.Marshal(traceRecord)
	if err != nil {
		return "", fmt.Errorf("failed to marshal trace record: %v", err)
	}
	// 将追踪信息存入区块链
	err = ctx.GetStub().PutState(student_trace_id, traceAsBytes)
	if err != nil {
		return "", fmt.Errorf("failed to put trace record: %v", err)
	}
	// 将追踪记录存入用户记录列表
	err = s.AddTraceRecord(ctx, userID, &traceRecord)
	if err != nil {
		return "", fmt.Errorf("failed to add trace record to user: %v", err)

	}

	return txid, nil
}

// 添加追踪记录到用户列表
func (s *SmartContract) AddTraceRecord(ctx contractapi.TransactionContextInterface, userID string, traceRecord *TraceRecord) error {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return err
	}
	// 检查是否已存在相同追踪码记录
	for _, existingTraceRecord := range user.TraceList {
		if existingTraceRecord.Student_trace_id == traceRecord.Student_trace_id {
			return fmt.Errorf("the trace record with id %s already exists in user %s's trace list", traceRecord.Student_trace_id, userID)
		}
	}
	user.TraceList = append(user.TraceList, traceRecord)
	userAsBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(userID, userAsBytes)
	if err != nil {
		return err
	}
	return nil
}

// 获取用户类型
func (s *SmartContract) GetUserType(ctx contractapi.TransactionContextInterface, userID string) (string, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return "", fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return "", err
	}
	return user.UserType, nil
}

// 获取用户信息
func (s *SmartContract) GetUserInfo(ctx contractapi.TransactionContextInterface, userID string) (*User, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return &User{}, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return &User{}, fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

// 获取单条追踪记录
func (s *SmartContract) GetTraceInfo(ctx contractapi.TransactionContextInterface, student_trace_id string) (*TraceRecord, error) {
	traceBytes, err := ctx.GetStub().GetState(student_trace_id)
	if err != nil {
		return &TraceRecord{}, fmt.Errorf("failed to read from world state: %v", err)
	}

	var traceRecord TraceRecord
	if traceBytes != nil {
		err = json.Unmarshal(traceBytes, &traceRecord)
		if err != nil {
			return &TraceRecord{}, fmt.Errorf("failed to unmarshal trace record: %v", err)
		}
		if traceRecord.Student_trace_id != "" {
			return &traceRecord, nil
		}
	}
	return &TraceRecord{}, fmt.Errorf("the trace record %s does not exist", student_trace_id)
}

// 获取用户追踪记录列表
func (s *SmartContract) GetTraceList(ctx contractapi.TransactionContextInterface, userID string) ([]*TraceRecord, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return nil, fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return nil, err
	}
	return user.TraceList, nil
}

// 获取全部追踪记录
func (s *SmartContract) GetAllTraceInfo(ctx contractapi.TransactionContextInterface) ([]TraceRecord, error) {
	traceListAsBytes, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	defer traceListAsBytes.Close()
	var traceRecords []TraceRecord
	for traceListAsBytes.HasNext() {
		queryResponse, err := traceListAsBytes.Next()
		if err != nil {
			return nil, err
		}
		var traceRecord TraceRecord
		err = json.Unmarshal(queryResponse.Value, &traceRecord)
		if err != nil {
			return nil, err
		}
		if traceRecord.Student_trace_id != "" {
			traceRecords = append(traceRecords, traceRecord)
		}
	}
	return traceRecords, nil
}

// 获取追踪记录历史
func (s *SmartContract) GetTraceHistory(ctx contractapi.TransactionContextInterface, student_trace_id string) ([]HistoryQueryResult, error) {
	log.Printf("GetTraceHistory: ID %v", student_trace_id)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(student_trace_id)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var traceRecord TraceRecord
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &traceRecord)
			if err != nil {
				return nil, err
			}
		} else {
			traceRecord = TraceRecord{
				Student_trace_id: student_trace_id,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, err
		}
		// 指定目标时区
		targetLocation, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			return nil, err
		}

		// 将时间戳转换到目标时区
		timestamp = timestamp.In(targetLocation)
		// 格式化时间戳为指定格式的字符串
		formattedTime := timestamp.Format("2006-01-02 15:04:05")

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: formattedTime,
			Record:    &traceRecord,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}
