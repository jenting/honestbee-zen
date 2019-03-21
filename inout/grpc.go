package inout

import (
	"github.com/honestbee/Zen/protobuf"
)

// GRPCCountryCodeMap defines gRPC CountryCode (int32) to internal country code (string) mapping
var GRPCCountryCodeMap = map[protobuf.CountryCode]string{
	protobuf.CountryCode_COUNTRY_CODE_SG: countryCodeSG,
	protobuf.CountryCode_COUNTRY_CODE_HK: countryCodeHK,
	protobuf.CountryCode_COUNTRY_CODE_TW: countryCodeTW,
	protobuf.CountryCode_COUNTRY_CODE_JP: countryCodeJP,
	protobuf.CountryCode_COUNTRY_CODE_TH: countryCodeTH,
	protobuf.CountryCode_COUNTRY_CODE_MY: countryCodeMY,
	protobuf.CountryCode_COUNTRY_CODE_ID: countryCodeID,
	protobuf.CountryCode_COUNTRY_CODE_PH: countryCodePH,
}

// GRPCLocaleMap defines gRPC Locale (int32) to internal locale (string) mapping
var GRPCLocaleMap = map[protobuf.Locale]string{
	protobuf.Locale_LOCALE_EN_US: localeENUS,
	protobuf.Locale_LOCALE_ZH_TW: localeZHTW,
	protobuf.Locale_LOCALE_ZH_CN: localeZHCN,
	protobuf.Locale_LOCALE_JA:    localeJA,
	protobuf.Locale_LOCALE_TH:    localeTH,
	protobuf.Locale_LOCALE_ID:    localeID,
}

// GRPCSortByMap defines gRPC SortBy (int32) to internal sort by (string) mapping
var GRPCSortByMap = map[protobuf.SortBy]string{
	protobuf.SortBy_SORT_BY_POSITION:   sortByPosition,
	protobuf.SortBy_SORT_BY_CREATED_AT: sortByCreatedAt,
	protobuf.SortBy_SORT_BY_UPDATED_AT: sortByUpdatedAt,
}

// GRPCSortOrderMap defines gRPC SortOrder (int32) to internal sort order (string) mapping
var GRPCSortOrderMap = map[protobuf.SortOrder]string{
	protobuf.SortOrder_SORT_ORDER_ASC:  sortOrderAsc,
	protobuf.SortOrder_SORT_ORDER_DESC: sortOrderDesc,
}

// GRPCVoteMap defines gRPC Vote (int32) to internal vote (string) mapping
var GRPCVoteMap = map[protobuf.Vote]string{
	protobuf.Vote_VOTE_UP:   voteUp,
	protobuf.Vote_VOTE_DOWN: voteDown,
}
