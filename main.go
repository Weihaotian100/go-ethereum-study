package main

import (
	"go-ethereum-study/account"
	"go-ethereum-study/transaction"
)

func main() {
	//创建客户端
	account.ClientTest()

	//以太坊账户
	account.AddressTest()

	//账户余额
	account.AccountBalanceTest()

	//生成新钱包
	account.WalletGenerateTest()

	//检查地址合法性以及测试是否是合约地址
	account.AddressCheckTest()

	//查询区块
	transaction.BlocksTest()

	//查询交易
	transaction.TransactionsTest()

	//eth转账
	transaction.TransferEthTest()

	//监听新区块
	transaction.BlockSubscribe()
}
