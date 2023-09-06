package walletrepository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	sharedutils "github.com/autobar-dev/shared-libraries/go/shared-utils"
)

func NewWalletRepository(service_url string, microservice_name string) *WalletRepository {
	client := &http.Client{}

	return &WalletRepository{
		service_url:       service_url,
		http_client:       client,
		microservice_name: microservice_name,
	}
}

func (wr WalletRepository) CreateWallet(
	user_id string,
	currency_code string,
) (*Wallet, error) {
	url := fmt.Sprintf("%s/wallet/create", wr.service_url)

	body := &ServiceCreateWalletRequestBody{
		UserId:       user_id,
		CurrencyCode: currency_code,
	}

	res, err := sharedutils.NewPostRequest(wr.http_client, wr.microservice_name, url, body)
	if err != nil {
		return nil, err
	}

	var scwr ServiceCreateWalletResponse
	if err := json.NewDecoder(res.Body).Decode(&scwr); err != nil {
		return nil, err
	}

	if scwr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing create wallet response: %s", *scwr.Error))
	}

	return scwr.Data, nil
}

func (wr WalletRepository) GetWallet(user_id string) (*Wallet, error) {
	url := fmt.Sprintf("%s/wallet/?user_id=%s", wr.service_url, user_id)

	res, err := sharedutils.NewGetRequest(wr.http_client, wr.microservice_name, url)
	if err != nil {
		return nil, err
	}

	var sgwr ServiceGetWalletResponse
	if err := json.NewDecoder(res.Body).Decode(&sgwr); err != nil {
		return nil, err
	}

	if sgwr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing get wallet response: %s", *sgwr.Error))
	}

	return sgwr.Data, nil
}

func (wr WalletRepository) CreateTransactionDeposit(
	user_id string,
	value int64,
) (*Transaction, error) {
	url := fmt.Sprintf("%s/transaction/create/deposit", wr.service_url)

	body := &ServiceTransactionDepositRequestBody{
		UserId: user_id,
		Value:  value,
	}

	res, err := sharedutils.NewPostRequest(wr.http_client, wr.microservice_name, url, body)
	if err != nil {
		return nil, err
	}

	var sctr ServiceCreateTransactionResponse
	if err := json.NewDecoder(res.Body).Decode(&sctr); err != nil {
		return nil, err
	}

	if sctr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing create transaction deposit response: %s", *sctr.Error))
	}

	return sctr.Data, nil
}

func (wr WalletRepository) CreateTransactionWithdraw(
	user_id string,
	value int64,
) (*Transaction, error) {
	url := fmt.Sprintf("%s/transaction/create/withdraw", wr.service_url)

	body := &ServiceTransactionWithdrawRequestBody{
		UserId: user_id,
		Value:  value,
	}

	res, err := sharedutils.NewPostRequest(wr.http_client, wr.microservice_name, url, body)
	if err != nil {
		return nil, err
	}

	var sctr ServiceCreateTransactionResponse
	if err := json.NewDecoder(res.Body).Decode(&sctr); err != nil {
		return nil, err
	}

	if sctr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing create transaction withdraw response: %s", *sctr.Error))
	}

	return sctr.Data, nil
}

func (wr WalletRepository) CreateTransactionPurchase(
	user_id string,
	value int64,
) (*Transaction, error) {
	url := fmt.Sprintf("%s/transaction/create/purchase", wr.service_url)

	body := &ServiceTransactionPurchaseRequestBody{
		UserId: user_id,
		Value:  value,
	}

	res, err := sharedutils.NewPostRequest(wr.http_client, wr.microservice_name, url, body)
	if err != nil {
		return nil, err
	}

	var sctr ServiceCreateTransactionResponse
	if err := json.NewDecoder(res.Body).Decode(&sctr); err != nil {
		return nil, err
	}

	if sctr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing create transaction purchase response: %s", *sctr.Error))
	}

	return sctr.Data, nil
}

func (wr WalletRepository) CreateTransactionRefund(
	user_id string,
	value int64,
) (*Transaction, error) {
	url := fmt.Sprintf("%s/transaction/create/refund", wr.service_url)

	body := &ServiceTransactionRefundRequestBody{
		UserId: user_id,
		Value:  value,
	}

	res, err := sharedutils.NewPostRequest(wr.http_client, wr.microservice_name, url, body)
	if err != nil {
		return nil, err
	}

	var sctr ServiceCreateTransactionResponse
	if err := json.NewDecoder(res.Body).Decode(&sctr); err != nil {
		return nil, err
	}

	if sctr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing create transaction refund response: %s", *sctr.Error))
	}

	return sctr.Data, nil
}

func (wr WalletRepository) CreateTransactionCurrencyChange(
	user_id string,
	value int64,
) (*Transaction, error) {
	url := fmt.Sprintf("%s/transaction/create/currency_change", wr.service_url)

	body := &ServiceTransactionCurrencyChangeRequestBody{
		UserId: user_id,
		Value:  value,
	}

	res, err := sharedutils.NewPostRequest(wr.http_client, wr.microservice_name, url, body)
	if err != nil {
		return nil, err
	}

	var sctr ServiceCreateTransactionResponse
	if err := json.NewDecoder(res.Body).Decode(&sctr); err != nil {
		return nil, err
	}

	if sctr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing create transaction currency change response: %s", *sctr.Error))
	}

	return sctr.Data, nil
}
