package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
	"math/big"
	"razor/cmd/mocks"
	"razor/core"
	"razor/core/types"
	"testing"
)

func TestExecuteClaimBounty(t *testing.T) {
	var client *ethclient.Client
	var flagSet *pflag.FlagSet

	type args struct {
		config         types.Configurations
		configErr      error
		password       string
		address        string
		addressErr     error
		bountyId       uint32
		bountyIdErr    error
		claimBountyTxn common.Hash
		claimBountyErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When executeClaimBounty function executes successfully",
			args: args{
				config:         types.Configurations{},
				password:       "test",
				address:        "0x000000000000000000000000000000000000dead",
				bountyId:       2,
				claimBountyTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting config",
			args: args{
				configErr:      errors.New("config error"),
				password:       "test",
				address:        "0x000000000000000000000000000000000000dead",
				bountyId:       2,
				claimBountyTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting address",
			args: args{
				config:         types.Configurations{},
				password:       "test",
				addressErr:     errors.New("address error"),
				bountyId:       2,
				claimBountyTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting bountyId",
			args: args{
				config:         types.Configurations{},
				password:       "test",
				address:        "0x000000000000000000000000000000000000dead",
				bountyIdErr:    errors.New("bountyId error"),
				claimBountyTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error from claimBounty function",
			args: args{
				config:         types.Configurations{},
				password:       "test",
				address:        "0x000000000000000000000000000000000000dead",
				bountyId:       2,
				claimBountyErr: errors.New("claimBounty error"),
			},
			expectedFatal: true,
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			flagSetUtilsMock := new(mocks.FlagSetInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)

			razorUtils = utilsMock
			flagSetUtils = flagSetUtilsMock
			cmdUtils = cmdUtilsMock

			utilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"))
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.password)
			flagSetUtilsMock.On("GetStringAddress", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.address, tt.args.addressErr)
			flagSetUtilsMock.On("GetUint32BountyId", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.bountyId, tt.args.bountyIdErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			cmdUtilsMock.On("ClaimBounty", mock.Anything, mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.claimBountyTxn, tt.args.claimBountyErr)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(1)

			fatal = false
			utils := &UtilsStruct{}
			utils.ExecuteClaimBounty(flagSet)

			if fatal != tt.expectedFatal {
				t.Error("The executeClaimBounty function didn't execute as expected")
			}

		})
	}
}

func TestClaimBounty(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var config types.Configurations
	var client *ethclient.Client
	var bountyInput types.RedeemBountyInput
	var callOpts bind.CallOpts
	var blockTime int64

	type args struct {
		epoch           uint32
		epochErr        error
		bountyLock      types.BountyLock
		bountyLockErr   error
		redeemBountyTxn *Types.Transaction
		redeemBountyErr error
		hash            common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When claimBounty function executes successfully",
			args: args{
				epoch: 70,
				bountyLock: types.BountyLock{
					Amount:      big.NewInt(1000),
					RedeemAfter: 70,
				},
				redeemBountyTxn: &Types.Transaction{},
				hash:            common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When claimBounty function executes successfully after waiting for few epochs",
			args: args{
				epoch: 70,
				bountyLock: types.BountyLock{
					Amount:      big.NewInt(1000),
					RedeemAfter: 80,
				},
				redeemBountyTxn: &Types.Transaction{},
				hash:            common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting epoch",
			args: args{
				epochErr: errors.New("epoch error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("epoch error"),
		},
		{
			name: "Test 4: When there is an error in getting bounty lock",
			args: args{
				epoch:         70,
				bountyLockErr: errors.New("bountyLock error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("bountyLock error"),
		},
		{
			name: "Test 5: When the amount in bounty lock is 0",
			args: args{
				epoch: 70,
				bountyLock: types.BountyLock{
					Amount:      big.NewInt(0),
					RedeemAfter: 70,
				},
			},
			want:    core.NilHash,
			wantErr: errors.New("bounty amount is 0"),
		},
		{
			name: "Test 6: When RedeemBounty transaction fails",
			args: args{
				epoch: 70,
				bountyLock: types.BountyLock{
					Amount:      big.NewInt(1000),
					RedeemAfter: 70,
				},
				redeemBountyErr: errors.New("redeemBounty error"),
			},
			want:    core.NilHash,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stakeManagerMock := new(mocks.StakeManagerInterface)
			utilsMock := new(mocks.UtilsInterface)
			trasactionUtilsMock := new(mocks.TransactionInterface)
			timeMock := new(mocks.TimeInterface)

			razorUtils = utilsMock
			stakeManagerUtils = stakeManagerMock
			transactionUtils = trasactionUtilsMock
			timeUtils = timeMock

			utilsMock.On("GetEpoch", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.epoch, tt.args.epochErr)
			utilsMock.On("GetOptions").Return(callOpts)
			stakeManagerMock.On("GetBountyLock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.CallOpts"), mock.AnythingOfType("uint32")).Return(tt.args.bountyLock, tt.args.bountyLockErr)
			timeMock.On("Sleep", mock.AnythingOfType("time.Duration")).Return()
			utilsMock.On("CalculateBlockTime", mock.AnythingOfType("*ethclient.Client")).Return(blockTime)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			stakeManagerMock.On("RedeemBounty", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.TransactOpts"), mock.AnythingOfType("uint32")).Return(tt.args.redeemBountyTxn, tt.args.redeemBountyErr)
			trasactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStruct{}
			got, err := utils.ClaimBounty(config, client, bountyInput)
			if got != tt.want {
				t.Errorf("Txn hash for claimBounty function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for claimBounty function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for claimBounty function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}
