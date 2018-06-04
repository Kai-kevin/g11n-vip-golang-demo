package constants

const (
	URL_QUERY_TRANSLATION_BY_COMPONET = "http://localhost:18080/i18n/api/v2/translation/products/{productName}/versions/{version}/locales/{locale}/components/{component}"

	URL_QUERY_TRANSLATION_BY_KEY = "http://localhost:18080/i18n/api/v2/translation/products/{productName}/versions/{version}/locales/{locale}/components/{component}/keys/{key}"

	URL_GET_FORMATTING_PATTERN_BY_LOCALE = "http://localhost:18080/i18n/api/v2/formatting/patterns/locales/{locale}"

)
