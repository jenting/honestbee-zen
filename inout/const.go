package inout

// SupportCountryLocaleMap lists supports country code and locale mapping.
var SupportCountryLocaleMap = map[string][]string{
	countryCodeSG: {localeENUS, localeZHCN},
	countryCodeHK: {localeENUS, localeZHTW},
	countryCodeTW: {localeENUS, localeZHTW},
	countryCodeJP: {localeENUS, localeJA},
	countryCodeTH: {localeENUS, localeTH},
	countryCodeMY: {localeENUS, localeZHCN},
	countryCodeID: {localeENUS, localeID},
	countryCodePH: {localeENUS},
}

const (
	countryCodeSG = "sg"
	countryCodeHK = "hk"
	countryCodeTW = "tw"
	countryCodeJP = "jp"
	countryCodeTH = "th"
	countryCodeMY = "my"
	countryCodeID = "id"
	countryCodePH = "ph"
)

const (
	localeENUS = "en-us"
	localeZHTW = "zh-tw"
	localeZHCN = "zh-cn"
	localeJA   = "ja"
	localeTH   = "th"
	localeID   = "id"
)

const (
	sortByPosition  = "position"
	sortByCreatedAt = "created_at"
	sortByUpdatedAt = "updated_at"
)

const (
	sortOrderAsc  = "asc"
	sortOrderDesc = "desc"
)

const (
	voteUp   = "up"
	voteDown = "down"
)

const (
	maxPerPage         = 100
	minPerPage         = 1
	minPage            = 1
	defaultPage        = 1
	defaultPerPage     = 30
	defaultLocale      = localeENUS
	defaultCountryCode = countryCodeSG
	defaultSortBy      = sortByPosition
	defaultSortOrder   = sortOrderAsc
)

const (
	// SuccessForceSync represents success trigger force sync job
	SuccessForceSync = "success trigger force sync job"
)
