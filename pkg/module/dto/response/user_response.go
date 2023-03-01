package response

import "github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"

type (
	FindAllActiveUserResponse struct {
		Results ActiveUserResults `json:"result"`

		Errors []errors.ErrorInfo
	}

	FindByIdActiveUserResponse struct {
		Result ActiveUserResult `json:"result"`

		Errors []errors.ErrorInfo
	}

	CreateActiveUserResponse struct {
		Result ActiveUserResult `json:"results"`

		Errors []errors.ErrorInfo
	}

	UpdateActiveUserResponse struct {
		Result ActiveUserResult `json:"results"`

		Errors []errors.ErrorInfo
	}

	DeleteActiveUserResponse struct {
		Errors []errors.ErrorInfo
	}

	LoginUserResponse struct {
		Result LoginUserResult

		Errors []errors.ErrorInfo
	}
)

type (
	ActiveUserResult struct {
		User ActiveUserEntity `json:"user"`
	}
	ActiveUserResults struct {
		Users []ActiveUserEntity `json:"users"`
	}

	HistoryUserResult struct {
		UserHistory *HistoryUserEntity `json:"user_history"`
	}
	HistoryUserResults struct {
		UserHistories []*HistoryUserEntity `json:"user_histories"`
	}

	LoginUserResult struct {
		Status string `json:"status"`
		Token  string `json:"token"`
	}
)

type (
	ActiveUserEntity struct {
	}
	HistoryUserEntity struct {
	}
)
