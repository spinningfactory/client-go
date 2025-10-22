package polygon

import (
	"context"
	"net/http"

	"github.com/dhiaaayachi/client-go/rest/client"
	"github.com/dhiaaayachi/client-go/rest/iter"
	"github.com/dhiaaayachi/client-go/rest/models"
)

const (
	ListTickersPath               = "/v3/reference/tickers"
	GetTickerDetailsPath          = "/v3/reference/tickers/{ticker}"
	ListTickerNewsPath            = "/v2/reference/news"
	GetTickerRelatedCompaniesPath = "/v1/related-companies/{ticker}"
	GetTickerTypesPath            = "/v3/reference/tickers/types"
	GetMarketHolidaysPath         = "/v1/marketstatus/upcoming"
	GetMarketStatusPath           = "/v1/marketstatus/now"
	ListSplitsPath                = "/v3/reference/splits"
	ListDividendsPath             = "/v3/reference/dividends"
	ListConditionsPath            = "/v3/reference/conditions"
	GetExchangesPath              = "/v3/reference/exchanges"
	GetOptionsContractPath        = "/v3/reference/options/contracts/{ticker}"
	ListOptionsContractsPath      = "/v3/reference/options/contracts"
	ListShortInterestPath         = "/stocks/v1/short-interest"
	ListShortVolumePath           = "/stocks/v1/short-volume"
	ListTreasuryYieldsPath        = "/fed/v1/treasury-yields"
)

// ReferenceClient defines a REST client for the Polygon reference API.
type ReferenceClient struct {
	client.Client
}

// ListTickers retrieves reference tickers. For more details see
// https://polygon.io/docs/stocks/get_v3_reference_tickers__ticker.
//
// This method returns an iterator that should be used to access the results via this pattern:
//
//	iter := c.ListTickers(context.TODO(), params, opts...)
//	for iter.Next() {
//		log.Print(iter.Item()) // do something with the current value
//	}
//	if iter.Err() != nil {
//		return iter.Err()
//	}
func (c *ReferenceClient) ListTickers(ctx context.Context, params *models.ListTickersParams, options ...models.RequestOption) *iter.Iter[models.Ticker] {
	return iter.NewIter(ctx, ListTickersPath, params, func(uri string) (iter.ListResponse, []models.Ticker, error) {
		res := &models.ListTickersResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// GetTickerDetails retrieves details for a specified ticker. For more details see
// https://polygon.io/docs/stocks/get_v3_reference_tickers__ticker.
func (c *ReferenceClient) GetTickerDetails(ctx context.Context, params *models.GetTickerDetailsParams, options ...models.RequestOption) (*models.GetTickerDetailsResponse, error) {
	res := &models.GetTickerDetailsResponse{}
	err := c.Call(ctx, http.MethodGet, GetTickerDetailsPath, params, res, options...)
	return res, err
}

// ListTickerNews retrieves news articles for a specified ticker. For more details see
// https://polygon.io/docs/stocks/get_v2_reference_news.
//
// This method returns an iterator that should be used to access the results via this pattern:
//
//	iter := c.ListTickerNews(context.TODO(), params, opts...)
//	for iter.Next() {
//		log.Print(iter.Item()) // do something with the current value
//	}
//	if iter.Err() != nil {
//		return iter.Err()
//	}
func (c *ReferenceClient) ListTickerNews(ctx context.Context, params *models.ListTickerNewsParams, options ...models.RequestOption) *iter.Iter[models.TickerNews] {
	return iter.NewIter(ctx, ListTickerNewsPath, params, func(uri string) (iter.ListResponse, []models.TickerNews, error) {
		res := &models.ListTickerNewsResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// GetTickerRelatedCompanies gets a list of related tickers based on news and returns data. For more details see
// https://polygon.io/docs/stocks/get_v1_related-companies__ticker.
func (c *ReferenceClient) GetTickerRelatedCompanies(ctx context.Context, params *models.GetTickerRelatedCompaniesParams, options ...models.RequestOption) (*models.GetTickerRelatedCompaniesResponse, error) {
	res := &models.GetTickerRelatedCompaniesResponse{}
	err := c.Call(ctx, http.MethodGet, GetTickerRelatedCompaniesPath, params, res, options...)
	return res, err
}

// GetTickerTypes retrieves all the possible ticker types that can be queried. For more details see
// https://polygon.io/docs/stocks/get_v3_reference_tickers_types.
func (c *ReferenceClient) GetTickerTypes(ctx context.Context, params *models.GetTickerTypesParams, options ...models.RequestOption) (*models.GetTickerTypesResponse, error) {
	res := &models.GetTickerTypesResponse{}
	err := c.Call(ctx, http.MethodGet, GetTickerTypesPath, params, res, options...)
	return res, err
}

// GetMarketHolidays retrieves upcoming market holidays and their open/close times. For more details see
// https://polygon.io/docs/stocks/get_v1_marketstatus_upcoming.
func (c *ReferenceClient) GetMarketHolidays(ctx context.Context, options ...models.RequestOption) (*models.GetMarketHolidaysResponse, error) {
	res := &models.GetMarketHolidaysResponse{}
	err := c.CallURL(ctx, http.MethodGet, GetMarketHolidaysPath, res, options...)
	return res, err
}

// GetMarketStatus retrieves the current trading status of the exchanges and overall financial markets. For more details
// see https://polygon.io/docs/stocks/get_v1_marketstatus_now.
func (c *ReferenceClient) GetMarketStatus(ctx context.Context, options ...models.RequestOption) (*models.GetMarketStatusResponse, error) {
	res := &models.GetMarketStatusResponse{}
	err := c.CallURL(ctx, http.MethodGet, GetMarketStatusPath, res, options...)
	return res, err
}

// ListSplits retrieves reference splits. For more details see https://polygon.io/docs/stocks/get_v3_reference_splits.
//
// This method returns an iterator that should be used to access the results via this pattern:
//
//	iter := c.ListSplits(context.TODO(), params, opts...)
//	for iter.Next() {
//		log.Print(iter.Item()) // do something with the current value
//	}
//	if iter.Err() != nil {
//		return iter.Err()
//	}
func (c *ReferenceClient) ListSplits(ctx context.Context, params *models.ListSplitsParams, options ...models.RequestOption) *iter.Iter[models.Split] {
	return iter.NewIter(ctx, ListSplitsPath, params, func(uri string) (iter.ListResponse, []models.Split, error) {
		res := &models.ListSplitsResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// ListDividends retrieves reference dividends. For more details see
// https://polygon.io/docs/stocks/get_v3_reference_dividends.
//
// This method returns an iterator that should be used to access the results via this pattern:
//
//	iter := c.ListDividends(context.TODO(), params, opts...)
//	for iter.Next() {
//		log.Print(iter.Item()) // do something with the current value
//	}
//	if iter.Err() != nil {
//		return iter.Err()
//	}
func (c *ReferenceClient) ListDividends(ctx context.Context, params *models.ListDividendsParams, options ...models.RequestOption) *iter.Iter[models.Dividend] {
	return iter.NewIter(ctx, ListDividendsPath, params, func(uri string) (iter.ListResponse, []models.Dividend, error) {
		res := &models.ListDividendsResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// ListConditions retrieves reference conditions. For more details see
// https://polygon.io/docs/stocks/get_v3_reference_conditions.
//
// This method returns an iterator that should be used to access the results via this pattern:
//
//	iter := c.ListConditions(context.TODO(), params, opts...)
//	for iter.Next() {
//		log.Print(iter.Item()) // do something with the current value
//	}
//	if iter.Err() != nil {
//		return iter.Err()
//	}
func (c *ReferenceClient) ListConditions(ctx context.Context, params *models.ListConditionsParams, options ...models.RequestOption) *iter.Iter[models.Condition] {
	return iter.NewIter(ctx, ListConditionsPath, params, func(uri string) (iter.ListResponse, []models.Condition, error) {
		res := &models.ListConditionsResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// GetExchanges lists all exchanges that Polygon knows about. For more details see
// https://polygon.io/docs/stocks/get_v3_reference_exchanges.
func (c *ReferenceClient) GetExchanges(ctx context.Context, params *models.GetExchangesParams, options ...models.RequestOption) (*models.GetExchangesResponse, error) {
	res := &models.GetExchangesResponse{}
	err := c.Call(ctx, http.MethodGet, GetExchangesPath, params, res, options...)
	return res, err
}

// GetOptionsContract retrieves a historical options contract. For more details see
// https://polygon.io/docs/options/get_v3_reference_options_contracts__options_ticker.
func (c *ReferenceClient) GetOptionsContract(ctx context.Context, params *models.GetOptionsContractParams, options ...models.RequestOption) (*models.GetOptionsContractResponse, error) {
	res := &models.GetOptionsContractResponse{}
	err := c.Call(ctx, http.MethodGet, GetOptionsContractPath, params, res, options...)
	return res, err
}

// ListOptionsContracts lists historical options contracts. For more details see
// https://polygon.io/docs/options/get_v3_reference_options_contracts.
//
// This method returns an iterator that should be used to access the results via this pattern:
//
//	iter := c.ListOptionsContracts(context.TODO(), params, opts...)
//	for iter.Next() {
//		log.Print(iter.Item()) // do something with the current value
//	}
//	if iter.Err() != nil {
//		return iter.Err()
//	}
func (c *ReferenceClient) ListOptionsContracts(ctx context.Context, params *models.ListOptionsContractsParams, options ...models.RequestOption) *iter.Iter[models.OptionsContract] {
	return iter.NewIter(ctx, ListOptionsContractsPath, params, func(uri string) (iter.ListResponse, []models.OptionsContract, error) {
		res := &models.ListOptionsContractsResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// ListShortInterest retrieves bi-monthly aggregated short interest data
// reported to FINRA by broker-dealers for a specified stock ticker. Short
// interest represents the total number of shares sold short but not yet
// covered or closed out, serving as an indicator of market sentiment and
// potential price movements. High short interest can signal bearish
// sentiment or highlight opportunities such as potential short squeezes.
// This endpoint provides essential insights for investors monitoring market
// positioning and sentiment.
//
// Use Cases: Market sentiment analysis, short-squeeze prediction, risk
// management, trading strategy refinement.
func (c *ReferenceClient) ListShortInterest(ctx context.Context, params *models.ListShortInterestParams, options ...models.RequestOption) *iter.Iter[models.ShortInterest] {
	return iter.NewIter(ctx, ListShortInterestPath, params, func(uri string) (iter.ListResponse, []models.ShortInterest, error) {
		res := &models.ListShortInterestResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// ListShortVolume retrieves daily aggregated short sale volume data reported
// to FINRA from off-exchange trading venues and alternative trading systems
// (ATS) for a specified stock ticker. Unlike short interest, which measures
// outstanding short positions at specific reporting intervals, short volume
// captures the daily trading activity of short sales. Monitoring short volume
// helps users detect immediate market sentiment shifts, analyze trading
// behavior, and identify trends in short-selling activity that may signal
// upcoming price movements.
//
// Use Cases: Intraday sentiment analysis, short-sale trend identification,
// liquidity analysis, trading strategy optimization.
func (c *ReferenceClient) ListShortVolume(ctx context.Context, params *models.ListShortVolumeParams, options ...models.RequestOption) *iter.Iter[models.ShortVolume] {
	return iter.NewIter(ctx, ListShortVolumePath, params, func(uri string) (iter.ListResponse, []models.ShortVolume, error) {
		res := &models.ListShortVolumeResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// ListTreasuryYields retrieves historical U.S. Treasury yield data for
// standard timeframes ranging from 1-month to 30-years, with daily
// historical records back to 1962. This endpoint lets you query by date or
// date range to see how interest rates have changed over time. Each data
// point reflects the market yield for Treasury securities of a specific
// maturity, helping users understand short- and long-term rate movements.
//
// Use Cases: Charting rate trends, comparing short vs. long-term yields,
// economic research.
func (c *ReferenceClient) ListTreasuryYields(ctx context.Context, params *models.ListTreasuryYieldsParams, options ...models.RequestOption) *iter.Iter[models.TreasuryYield] {
	return iter.NewIter(ctx, ListTreasuryYieldsPath, params, func(uri string) (iter.ListResponse, []models.TreasuryYield, error) {
		res := &models.ListTreasuryYieldsResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}
