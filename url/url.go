/*
Package url defines url utilities helpers.
*/
package url

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// alphaNumericRegex validates strings containing only letters, numbers, and hyphens
const alphaNumericRegex = "^[a-zA-Z0-9-]+$"

// hostRegex validates hostnames according to RFC 1123
// Rules:
// - Each label must start and end with alphanumeric characters
// - Middle characters can be alphanumeric or hyphens
// - Multiple labels can be joined with dots
const hostRegex = `^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])(\.[a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])*$`

// pathRegex validates URL paths containing letters, numbers, hyphens, underscores  and forward slashes
const pathRegex = "^[a-zA-Z0-9_-]+(?:/[a-zA-Z0-9_-]+)*$"

var alphaNumericRe = regexp.MustCompile(alphaNumericRegex)
var hostRe = regexp.MustCompile(hostRegex)
var pathRe = regexp.MustCompile(pathRegex)

// BuildURL constructs a URL by combining a scheme, host, path, and query parameters.
//
// Parameters:
//   - scheme: The URL scheme (e.g., "http", "https") to use.
//   - host: The host or domain name (e.g., "example.com").
//   - path: The path part of the URL (e.g., "path/to/resource").
//   - query: A map[string]string containing query parameters to append to the URL.
//
// Returns:
//   - string: The constructed URL with the scheme, host, path, and query parameters.
//   - error: An error if the URL could not be parsed or if any issues occurred during URL construction.
//
// Behavior:
//   - The function concatenates the scheme, host, and path into a URL string.
//   - It then attempts to parse the constructed URL and add the query parameters.
//   - If the URL parsing fails, an error is returned.
//   - Query parameters are added one by one using the `url.Values.Add` method to ensure proper encoding.
//
// Example:
//
//	scheme := "https"
//	host := "example.com"
//	path := "search"
//	query := map[string]string{
//	    "q": "golang",
//	    "page": "1",
//	}
//	fullURL, err := BuildURL(scheme, host, path, query)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//	    fmt.Println(fullURL)
//	    // Output: https://example.com/search?q=golang&page=1
//	}
//
// Notes:
//   - If any query parameters are provided, they will be encoded and appended to the URL.
//   - If the path is empty, a trailing slash will be included after the host.
//   - The function ensures the proper encoding of query parameters and safely constructs the final URL.
//
// Usage:
//
//	To construct a URL with query parameters:
//	  queryParams := map[string]string{"key": "value", "anotherKey": "anotherValue"}
//	  url, err := BuildURL("http", "example.com", "search", queryParams)
//	  fmt.Println("Constructed URL:", url)
func BuildURL(scheme, host, path string, query map[string]string) (string, error) {
	var errMessage []string
	if scheme == "" {
		errMessage = append(errMessage, "scheme is required")
	}
	if host != "" {
		if !hostRe.MatchString(host) {
			errMessage = append(errMessage, "the host is not valid")
		}
	}
	if host == "" {
		errMessage = append(errMessage, "host is required")
	}

	if path != "" {
		if !pathRe.MatchString(path) {
			errMessage = append(errMessage, "path is permitted with a-z, 0-9, - and _ characters and multiple path segments")
		}
	}

	if errMessage != nil {
		return "", errors.New(strings.Join(errMessage, "; "))
	}

	parsedUrl := &url.URL{
		Scheme: scheme,
		Host:   host,
	}

	if path == "" {
		parsedUrl.Path = "/"
	} else if !strings.HasPrefix(path, "/") {
		parsedUrl.Path = "/" + path
	} else {
		parsedUrl.Path = path
	}
	queryParams := parsedUrl.Query()
	for key, value := range query {
		queryParams.Add(key, value)
	}
	parsedUrl.RawQuery = queryParams.Encode()

	return parsedUrl.String(), nil
}

// AddQueryParams adds multiple query parameters to a given URL and returns the updated URL.
//
// Parameters:
//   - urlStr: A string representing the base URL to which query parameters should be added.
//   - params: A map[string]string containing key-value pairs of query parameters to add.
//
// Returns:
//   - string: The updated URL with the new query parameters appended.
//   - error: An error if the URL cannot be parsed.
//
// Behavior:
//   - The function parses the provided URL string using `net/url.Parse`.
//   - Iterates through the `params` map, adding each key-value pair as a query parameter to the URL.
//   - Encodes the updated query parameters back into the URL.
//
// Example:
//
//	baseURL := "https://example.com/path"
//	params := map[string]string{
//	    "param1": "value1",
//	    "param2": "value2",
//	}
//	updatedURL, err := AddQueryParams(baseURL, params)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//	    fmt.Println(updatedURL)
//	    // Output: https://example.com/path?param1=value1&param2=value2
//	}
//
// Notes:
//   - If a query parameter key already exists in the URL, `Add` appends the new value instead of overwriting it.
//   - Use this function when you need to dynamically construct URLs with multiple query parameters.
//   - The function ensures proper encoding of query parameters.
//
// Usage:
//
//	To add query parameters to a URL:
//	  params := map[string]string{"key": "value", "foo": "bar"}
//	  updatedURL, err := AddQueryParams("http://example.com", params)
//	  fmt.Println("Result:", updatedURL)
func AddQueryParams(urlStr string, params map[string]string) (string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("URL %s could not be parsed. err: %w", urlStr, err)
	}
	switch parsedURL.Scheme {
	case "http", "https", "ws", "wss", "ftp":
	default:
		return "", fmt.Errorf("URL scheme %s is invalid", parsedURL.Scheme)
	}
	queryParams := parsedURL.Query()

	for key, value := range params {
		err = validateKeyValue(key, value)
		if err != nil {
			return "", err
		}
		queryParams.Add(key, value)
	}
	parsedURL.RawQuery = queryParams.Encode()

	return parsedURL.String(), nil
}

func validateKeyValue(key, value string) error {
	if key == "" {
		return errors.New("query parameter key cannot be empty")
	}
	if value == "" {
		return fmt.Errorf("query parameter value for key %s cannot be empty", key)
	}
	if !alphaNumericRe.MatchString(key) {
		return fmt.Errorf("query parameter key %s must be alphanumeric", key)
	}
	if !alphaNumericRe.MatchString(value) {
		return fmt.Errorf("query parameter value %s for key %s must be alphanumeric", value, key)
	}

	return nil
}

// IsValidURL checks whether a given URL string is valid and its scheme matches the allowed list.
//
// Parameters:
//   - urlStr: A string representing the URL to validate.
//   - allowedReqSchemes: A slice of strings containing the allowed schemes (e.g., "http", "https").
//
// Returns:
//   - bool: `true` if the URL is valid and its scheme is in the allowed list; otherwise, `false`.
//
// Behavior:
//   - The function attempts to parse the provided URL string using `net/url.Parse`.
//   - If the URL is invalid or its scheme is not in the allowed list, the function returns `false`.
//   - If the URL is valid and the scheme is allowed, the function returns `true`.
//
// Example:
//
//	url := "https://example.com"
//	allowed := []string{"http", "https"}
//	isValid := IsValidURL(url, allowed)
//	if isValid {
//	    fmt.Println("URL is valid and uses an allowed scheme.")
//	} else {
//	    fmt.Println("Invalid URL or scheme.")
//	}
//
// Notes:
//   - The function does not check other parts of the URL (e.g., hostname, path, query parameters).
//   - Use this function when you need to validate both the structure and scheme of a URL.
//
// Usage:
//
//	To validate URLs and restrict their schemes:
//	  valid := IsValidURL("http://example.com", []string{"http", "https"})
//	  fmt.Println("Is valid:", valid)
func IsValidURL(urlStr string, allowedReqSchemes []string) bool {
	if urlStr == "" {
		return false
	}
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return false
	}
	for _, scheme := range allowedReqSchemes {
		if scheme == "" {
			return false
		}
		if parsedURL.Scheme == scheme {
			return true
		}
	}

	return false
}

// ExtractDomain extracts the domain (hostname) from a given URL string.
//
// Parameters:
//   - urlStr: A string representing the URL from which to extract the domain.
//
// Returns:
//   - string: The domain (hostname) extracted from the URL.
//   - error: An error if the URL is invalid or the domain cannot be determined.
//
// Errors:
//   - Returns an error if the provided URL string is invalid or cannot be parsed.
//   - Returns "parameter not found" error if the URL does not contain a hostname.
//
// Example:
//
//	url := "https://example.com/path?query=value"
//	domain, err := ExtractDomain(url)
//	if err != nil {
//	    log.Println("Error:", err)
//	} else {
//	    log.Println("Domain:", domain) // Output: "example.com"
//
// Notes:
//   - This function uses `net/url.ParseRequestURI` to validate and parse the URL.
//   - It extracts the hostname part of the URL and ignores the port, path, query, or fragment.
//
// Usage:
//
//	To extract the domain from a URL:
//	  domain, err := ExtractDomain("http://example.com/some-path")
//	  if err != nil {
//	      fmt.Println("Error:", err)
//	  } else {
//	      fmt.Println("Domain:", domain)
func ExtractDomain(urlStr string) (string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("URL %s could not be parsed. err: %w", urlStr, err)
	}

	host, err := publicsuffix.EffectiveTLDPlusOne(parsedURL.Hostname())
	if err != nil {
		return "", fmt.Errorf("could not extract public suffix from host %s. err: %w", parsedURL.Hostname(), err)
	}

	if host == "" {
		return "", errors.New("public suffix is empty")
	}

	return host, nil
}

// GetQueryParam extracts the value of a specified query parameter from a given URL string.
//
// Parameters:
//   - urlStr: A string representing the URL containing query parameters.
//   - param: The name of the query parameter to retrieve.
//
// Returns:
//   - string: The value of the specified query parameter.
//   - error: An error if the URL is invalid or the parameter is not found.
//
// Errors:
//   - Returns an error if the provided URL string is invalid or cannot be parsed.
//   - Returns "parameter not found" error if the specified parameter does not exist.
//
// Example:
//
//	url := "https://example.com?foo=bar&baz=qux"
//	value, err := GetQueryParam(url, "foo")
//	if err != nil {
//	    log.Println("Error:", err)
//	} else {
//	    log.Println("Value:", value) // Output: "bar"
//
// Notes:
//   - This function uses the `net/url` package for robust URL parsing.
//   - It assumes the URL is properly formatted with query parameters starting after a "?".
//
// Usage:
//
//	To retrieve the value of a query parameter from a URL:
//	  value, err := GetQueryParam("http://example.com?key=value", "key")
//	  if err != nil {
//	      fmt.Println("Error:", err)
//	  } else {
//	      fmt.Println("Value:", value)
func GetQueryParam(urlStr, param string) (string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	queryParams := parsedURL.Query()
	value, exists := queryParams[param]

	if !exists || len(value) == 0 {
		return "", fmt.Errorf("parameter %s not found in URL %s", param, urlStr)
	}

	return value[0], nil
}
