<?php
/**
 * PublicApi
 * PHP version 5
 *
 * @category Class
 * @package  HydraSDK
 * @author   Swagger Codegen team
 * @link     https://github.com/swagger-api/swagger-codegen
 */

/**
 * ORY Hydra
 *
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * OpenAPI spec version: latest
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 */

/**
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen
 * Do not edit the class manually.
 */

namespace HydraSDK\Api;

use \HydraSDK\ApiClient;
use \HydraSDK\ApiException;
use \HydraSDK\Configuration;
use \HydraSDK\ObjectSerializer;

/**
 * PublicApi Class Doc Comment
 *
 * @category Class
 * @package  HydraSDK
 * @author   Swagger Codegen team
 * @link     https://github.com/swagger-api/swagger-codegen
 */
class PublicApi
{
    /**
     * API Client
     *
     * @var \HydraSDK\ApiClient instance of the ApiClient
     */
    protected $apiClient;

    /**
     * Constructor
     *
     * @param \HydraSDK\ApiClient|null $apiClient The api client to use
     */
    public function __construct(\HydraSDK\ApiClient $apiClient = null)
    {
        if ($apiClient === null) {
            $apiClient = new ApiClient();
        }

        $this->apiClient = $apiClient;
    }

    /**
     * Get API client
     *
     * @return \HydraSDK\ApiClient get the API client
     */
    public function getApiClient()
    {
        return $this->apiClient;
    }

    /**
     * Set the API client
     *
     * @param \HydraSDK\ApiClient $apiClient set the API client
     *
     * @return PublicApi
     */
    public function setApiClient(\HydraSDK\ApiClient $apiClient)
    {
        $this->apiClient = $apiClient;
        return $this;
    }

    /**
     * Operation disconnectUser
     *
     * OpenID Connect Front-Backchannel enabled Logout
     *
     * Client for Hydra
     *
     * @throws \HydraSDK\ApiException on non-2xx response
     * @return void
     */
    public function disconnectUser()
    {
        list($response) = $this->disconnectUserWithHttpInfo();
        return $response;
    }

    /**
     * Operation disconnectUserWithHttpInfo
     *
     * OpenID Connect Front-Backchannel enabled Logout
     *
     * Client for Hydra
     *
     * @throws \HydraSDK\ApiException on non-2xx response
     * @return array of null, HTTP status code, HTTP response headers (array of strings)
     */
    public function disconnectUserWithHttpInfo()
    {
        // parse inputs
        $resourcePath = "/oauth2/sessions/logout";
        $httpBody = '';
        $queryParams = [];
        $headerParams = [];
        $formParams = [];
        $_header_accept = $this->apiClient->selectHeaderAccept(['application/json']);
        if (!is_null($_header_accept)) {
            $headerParams['Accept'] = $_header_accept;
        }
        $headerParams['Content-Type'] = $this->apiClient->selectHeaderContentType(['application/json', 'application/x-www-form-urlencoded']);


        // for model (json/xml)
        if (isset($_tempBody)) {
            $httpBody = $_tempBody; // $_tempBody is the method argument, if present
        } elseif (count($formParams) > 0) {
            $httpBody = $formParams; // for HTTP post (form)
        }
        // make the API Call
        try {
            list($response, $statusCode, $httpHeader) = $this->apiClient->callApi(
                $resourcePath,
                'GET',
                $queryParams,
                $httpBody,
                $headerParams,
                null,
                '/oauth2/sessions/logout'
            );

            return [null, $statusCode, $httpHeader];
        } catch (ApiException $e) {
            switch ($e->getCode()) {
            }

            throw $e;
        }
    }

    /**
     * Operation discoverOpenIDConfiguration
     *
     * OpenID Connect Discovery
     *
     * Client for Hydra
     *
     * @throws \HydraSDK\ApiException on non-2xx response
     * @return \HydraSDK\Model\WellKnown
     */
    public function discoverOpenIDConfiguration()
    {
        list($response) = $this->discoverOpenIDConfigurationWithHttpInfo();
        return $response;
    }

    /**
     * Operation discoverOpenIDConfigurationWithHttpInfo
     *
     * OpenID Connect Discovery
     *
     * Client for Hydra
     *
     * @throws \HydraSDK\ApiException on non-2xx response
     * @return array of \HydraSDK\Model\WellKnown, HTTP status code, HTTP response headers (array of strings)
     */
    public function discoverOpenIDConfigurationWithHttpInfo()
    {
        // parse inputs
        $resourcePath = "/.well-known/openid-configuration";
        $httpBody = '';
        $queryParams = [];
        $headerParams = [];
        $formParams = [];
        $_header_accept = $this->apiClient->selectHeaderAccept(['application/json']);
        if (!is_null($_header_accept)) {
            $headerParams['Accept'] = $_header_accept;
        }
        $headerParams['Content-Type'] = $this->apiClient->selectHeaderContentType(['application/json', 'application/x-www-form-urlencoded']);


        // for model (json/xml)
        if (isset($_tempBody)) {
            $httpBody = $_tempBody; // $_tempBody is the method argument, if present
        } elseif (count($formParams) > 0) {
            $httpBody = $formParams; // for HTTP post (form)
        }
        // make the API Call
        try {
            list($response, $statusCode, $httpHeader) = $this->apiClient->callApi(
                $resourcePath,
                'GET',
                $queryParams,
                $httpBody,
                $headerParams,
                '\HydraSDK\Model\WellKnown',
                '/.well-known/openid-configuration'
            );

            return [$this->apiClient->getSerializer()->deserialize($response, '\HydraSDK\Model\WellKnown', $httpHeader), $statusCode, $httpHeader];
        } catch (ApiException $e) {
            switch ($e->getCode()) {
                case 200:
                    $data = $this->apiClient->getSerializer()->deserialize($e->getResponseBody(), '\HydraSDK\Model\WellKnown', $e->getResponseHeaders());
                    $e->setResponseObject($data);
                    break;
                case 401:
                    $data = $this->apiClient->getSerializer()->deserialize($e->getResponseBody(), '\HydraSDK\Model\GenericError', $e->getResponseHeaders());
                    $e->setResponseObject($data);
                    break;
                case 500:
                    $data = $this->apiClient->getSerializer()->deserialize($e->getResponseBody(), '\HydraSDK\Model\GenericError', $e->getResponseHeaders());
                    $e->setResponseObject($data);
                    break;
            }

            throw $e;
        }
    }

    /**
     * Operation oauth2Token
     *
     * The OAuth 2.0 token endpoint
     *
     * Client for Hydra
     *
     * @param string $grant_type  (required)
     * @param string $code  (optional)
     * @param string $redirect_uri  (optional)
     * @param string $client_id  (optional)
     * @throws \HydraSDK\ApiException on non-2xx response
     * @return \HydraSDK\Model\Oauth2TokenResponse
     */
    public function oauth2Token($grant_type, $code = null, $redirect_uri = null, $client_id = null)
    {
        list($response) = $this->oauth2TokenWithHttpInfo($grant_type, $code, $redirect_uri, $client_id);
        return $response;
    }

    /**
     * Operation oauth2TokenWithHttpInfo
     *
     * The OAuth 2.0 token endpoint
     *
     * Client for Hydra
     *
     * @param string $grant_type  (required)
     * @param string $code  (optional)
     * @param string $redirect_uri  (optional)
     * @param string $client_id  (optional)
     * @throws \HydraSDK\ApiException on non-2xx response
     * @return array of \HydraSDK\Model\Oauth2TokenResponse, HTTP status code, HTTP response headers (array of strings)
     */
    public function oauth2TokenWithHttpInfo($grant_type, $code = null, $redirect_uri = null, $client_id = null)
    {
        // verify the required parameter 'grant_type' is set
        if ($grant_type === null) {
            throw new \InvalidArgumentException('Missing the required parameter $grant_type when calling oauth2Token');
        }
        // parse inputs
        $resourcePath = "/oauth2/token";
        $httpBody = '';
        $queryParams = [];
        $headerParams = [];
        $formParams = [];
        $_header_accept = $this->apiClient->selectHeaderAccept(['application/json']);
        if (!is_null($_header_accept)) {
            $headerParams['Accept'] = $_header_accept;
        }
        $headerParams['Content-Type'] = $this->apiClient->selectHeaderContentType(['application/x-www-form-urlencoded']);

        // form params
        if ($grant_type !== null) {
            $formParams['grant_type'] = $this->apiClient->getSerializer()->toFormValue($grant_type);
        }
        // form params
        if ($code !== null) {
            $formParams['code'] = $this->apiClient->getSerializer()->toFormValue($code);
        }
        // form params
        if ($redirect_uri !== null) {
            $formParams['redirect_uri'] = $this->apiClient->getSerializer()->toFormValue($redirect_uri);
        }
        // form params
        if ($client_id !== null) {
            $formParams['client_id'] = $this->apiClient->getSerializer()->toFormValue($client_id);
        }

        // for model (json/xml)
        if (isset($_tempBody)) {
            $httpBody = $_tempBody; // $_tempBody is the method argument, if present
        } elseif (count($formParams) > 0) {
            $httpBody = $formParams; // for HTTP post (form)
        }
        // this endpoint requires HTTP basic authentication
        if (strlen($this->apiClient->getConfig()->getUsername()) !== 0 or strlen($this->apiClient->getConfig()->getPassword()) !== 0) {
            $headerParams['Authorization'] = 'Basic ' . base64_encode($this->apiClient->getConfig()->getUsername() . ":" . $this->apiClient->getConfig()->getPassword());
        }
        // this endpoint requires OAuth (access token)
        if (strlen($this->apiClient->getConfig()->getAccessToken()) !== 0) {
            $headerParams['Authorization'] = 'Bearer ' . $this->apiClient->getConfig()->getAccessToken();
        }
        // make the API Call
        try {
            list($response, $statusCode, $httpHeader) = $this->apiClient->callApi(
                $resourcePath,
                'POST',
                $queryParams,
                $httpBody,
                $headerParams,
                '\HydraSDK\Model\Oauth2TokenResponse',
                '/oauth2/token'
            );

            return [$this->apiClient->getSerializer()->deserialize($response, '\HydraSDK\Model\Oauth2TokenResponse', $httpHeader), $statusCode, $httpHeader];
        } catch (ApiException $e) {
            switch ($e->getCode()) {
                case 200:
                    $data = $this->apiClient->getSerializer()->deserialize($e->getResponseBody(), '\HydraSDK\Model\Oauth2TokenResponse', $e->getResponseHeaders());
                    $e->setResponseObject($data);
                    break;
                case 401:
                    $data = $this->apiClient->getSerializer()->deserialize($e->getResponseBody(), '\HydraSDK\Model\GenericError', $e->getResponseHeaders());
                    $e->setResponseObject($data);
                    break;
                case 500:
                    $data = $this->apiClient->getSerializer()->deserialize($e->getResponseBody(), '\HydraSDK\Model\GenericError', $e->getResponseHeaders());
                    $e->setResponseObject($data);
                    break;
            }

            throw $e;
        }
    }

    /**
     * Operation oauthAuth
     *
     * The OAuth 2.0 authorize endpoint
     *
     * Client for Hydra
     *
     * @throws \HydraSDK\ApiException on non-2xx response
     * @return void
     */
    public function oauthAuth()
    {
        list($response) = $this->oauthAuthWithHttpInfo();
        return $response;
    }

    /**
     * Operation oauthAuthWithHttpInfo
     *
     * The OAuth 2.0 authorize endpoint
     *
     * Client for Hydra
     *
     * @throws \HydraSDK\ApiException on non-2xx response
     * @return array of null, HTTP status code, HTTP response headers (array of strings)
     */
    public function oauthAuthWithHttpInfo()
    {
        // parse inputs
        $resourcePath = "/oauth2/auth";
        $httpBody = '';
        $queryParams = [];
        $headerParams = [];
        $formParams = [];
        $_header_accept = $this->apiClient->selectHeaderAccept(['application/json']);
        if (!is_null($_header_accept)) {
            $headerParams['Accept'] = $_header_accept;
        }
        $headerParams['Content-Type'] = $this->apiClient->selectHeaderContentType(['application/x-www-form-urlencoded']);


        // for model (json/xml)
        if (isset($_tempBody)) {
            $httpBody = $_tempBody; // $_tempBody is the method argument, if present
        } elseif (count($formParams) > 0) {
            $httpBody = $formParams; // for HTTP post (form)
        }
        // make the API Call
        try {
            list($response, $statusCode, $httpHeader) = $this->apiClient->callApi(
                $resourcePath,
                'GET',
                $queryParams,
                $httpBody,
                $headerParams,
                null,
                '/oauth2/auth'
            );

            return [null, $statusCode, $httpHeader];
        } catch (ApiException $e) {
            switch ($e->getCode()) {
                case 401:
                    $data = $this->apiClient->getSerializer()->deserialize($e->getResponseBody(), '\HydraSDK\Model\GenericError', $e->getResponseHeaders());
                    $e->setResponseObject($data);
                    break;
                case 500:
                    $data = $this->apiClient->getSerializer()->deserialize($e->getResponseBody(), '\HydraSDK\Model\GenericError', $e->getResponseHeaders());
                    $e->setResponseObject($data);
                    break;
            }

            throw $e;
        }
    }

    /**
     * Operation revokeOAuth2Token
     *
     * Revoke OAuth2 tokens
     *
     * Client for Hydra
     *
     * @param string $token  (required)
     * @throws \HydraSDK\ApiException on non-2xx response
     * @return void
     */
    public function revokeOAuth2Token($token)
    {
        list($response) = $this->revokeOAuth2TokenWithHttpInfo($token);
        return $response;
    }

    /**
     * Operation revokeOAuth2TokenWithHttpInfo
     *
     * Revoke OAuth2 tokens
     *
     * Client for Hydra
     *
     * @param string $token  (required)
     * @throws \HydraSDK\ApiException on non-2xx response
     * @return array of null, HTTP status code, HTTP response headers (array of strings)
     */
    public function revokeOAuth2TokenWithHttpInfo($token)
    {
        // verify the required parameter 'token' is set
        if ($token === null) {
            throw new \InvalidArgumentException('Missing the required parameter $token when calling revokeOAuth2Token');
        }
        // parse inputs
        $resourcePath = "/oauth2/revoke";
        $httpBody = '';
        $queryParams = [];
        $headerParams = [];
        $formParams = [];
        $_header_accept = $this->apiClient->selectHeaderAccept(['application/json']);
        if (!is_null($_header_accept)) {
            $headerParams['Accept'] = $_header_accept;
        }
        $headerParams['Content-Type'] = $this->apiClient->selectHeaderContentType(['application/x-www-form-urlencoded']);

        // form params
        if ($token !== null) {
            $formParams['token'] = $this->apiClient->getSerializer()->toFormValue($token);
        }

        // for model (json/xml)
        if (isset($_tempBody)) {
            $httpBody = $_tempBody; // $_tempBody is the method argument, if present
        } elseif (count($formParams) > 0) {
            $httpBody = $formParams; // for HTTP post (form)
        }
        // this endpoint requires HTTP basic authentication
        if (strlen($this->apiClient->getConfig()->getUsername()) !== 0 or strlen($this->apiClient->getConfig()->getPassword()) !== 0) {
            $headerParams['Authorization'] = 'Basic ' . base64_encode($this->apiClient->getConfig()->getUsername() . ":" . $this->apiClient->getConfig()->getPassword());
        }
        // this endpoint requires OAuth (access token)
        if (strlen($this->apiClient->getConfig()->getAccessToken()) !== 0) {
            $headerParams['Authorization'] = 'Bearer ' . $this->apiClient->getConfig()->getAccessToken();
        }
        // make the API Call
        try {
            list($response, $statusCode, $httpHeader) = $this->apiClient->callApi(
                $resourcePath,
                'POST',
                $queryParams,
                $httpBody,
                $headerParams,
                null,
                '/oauth2/revoke'
            );

            return [null, $statusCode, $httpHeader];
        } catch (ApiException $e) {
            switch ($e->getCode()) {
                case 401:
                    $data = $this->apiClient->getSerializer()->deserialize($e->getResponseBody(), '\HydraSDK\Model\GenericError', $e->getResponseHeaders());
                    $e->setResponseObject($data);
                    break;
                case 500:
                    $data = $this->apiClient->getSerializer()->deserialize($e->getResponseBody(), '\HydraSDK\Model\GenericError', $e->getResponseHeaders());
                    $e->setResponseObject($data);
                    break;
            }

            throw $e;
        }
    }

    /**
     * Operation userinfo
     *
     * OpenID Connect Userinfo
     *
     * Client for Hydra
     *
     * @throws \HydraSDK\ApiException on non-2xx response
     * @return \HydraSDK\Model\UserinfoResponse
     */
    public function userinfo()
    {
        list($response) = $this->userinfoWithHttpInfo();
        return $response;
    }

    /**
     * Operation userinfoWithHttpInfo
     *
     * OpenID Connect Userinfo
     *
     * Client for Hydra
     *
     * @throws \HydraSDK\ApiException on non-2xx response
     * @return array of \HydraSDK\Model\UserinfoResponse, HTTP status code, HTTP response headers (array of strings)
     */
    public function userinfoWithHttpInfo()
    {
        // parse inputs
        $resourcePath = "/userinfo";
        $httpBody = '';
        $queryParams = [];
        $headerParams = [];
        $formParams = [];
        $_header_accept = $this->apiClient->selectHeaderAccept(['application/json']);
        if (!is_null($_header_accept)) {
            $headerParams['Accept'] = $_header_accept;
        }
        $headerParams['Content-Type'] = $this->apiClient->selectHeaderContentType(['application/json', 'application/x-www-form-urlencoded']);


        // for model (json/xml)
        if (isset($_tempBody)) {
            $httpBody = $_tempBody; // $_tempBody is the method argument, if present
        } elseif (count($formParams) > 0) {
            $httpBody = $formParams; // for HTTP post (form)
        }
        // this endpoint requires OAuth (access token)
        if (strlen($this->apiClient->getConfig()->getAccessToken()) !== 0) {
            $headerParams['Authorization'] = 'Bearer ' . $this->apiClient->getConfig()->getAccessToken();
        }
        // make the API Call
        try {
            list($response, $statusCode, $httpHeader) = $this->apiClient->callApi(
                $resourcePath,
                'GET',
                $queryParams,
                $httpBody,
                $headerParams,
                '\HydraSDK\Model\UserinfoResponse',
                '/userinfo'
            );

            return [$this->apiClient->getSerializer()->deserialize($response, '\HydraSDK\Model\UserinfoResponse', $httpHeader), $statusCode, $httpHeader];
        } catch (ApiException $e) {
            switch ($e->getCode()) {
                case 200:
                    $data = $this->apiClient->getSerializer()->deserialize($e->getResponseBody(), '\HydraSDK\Model\UserinfoResponse', $e->getResponseHeaders());
                    $e->setResponseObject($data);
                    break;
                case 401:
                    $data = $this->apiClient->getSerializer()->deserialize($e->getResponseBody(), '\HydraSDK\Model\GenericError', $e->getResponseHeaders());
                    $e->setResponseObject($data);
                    break;
                case 500:
                    $data = $this->apiClient->getSerializer()->deserialize($e->getResponseBody(), '\HydraSDK\Model\GenericError', $e->getResponseHeaders());
                    $e->setResponseObject($data);
                    break;
            }

            throw $e;
        }
    }

    /**
     * Operation wellKnown
     *
     * JSON Web Keys Discovery
     *
     * Client for Hydra
     *
     * @throws \HydraSDK\ApiException on non-2xx response
     * @return \HydraSDK\Model\JSONWebKeySet
     */
    public function wellKnown()
    {
        list($response) = $this->wellKnownWithHttpInfo();
        return $response;
    }

    /**
     * Operation wellKnownWithHttpInfo
     *
     * JSON Web Keys Discovery
     *
     * Client for Hydra
     *
     * @throws \HydraSDK\ApiException on non-2xx response
     * @return array of \HydraSDK\Model\JSONWebKeySet, HTTP status code, HTTP response headers (array of strings)
     */
    public function wellKnownWithHttpInfo()
    {
        // parse inputs
        $resourcePath = "/.well-known/jwks.json";
        $httpBody = '';
        $queryParams = [];
        $headerParams = [];
        $formParams = [];
        $_header_accept = $this->apiClient->selectHeaderAccept(['application/json']);
        if (!is_null($_header_accept)) {
            $headerParams['Accept'] = $_header_accept;
        }
        $headerParams['Content-Type'] = $this->apiClient->selectHeaderContentType(['application/json']);


        // for model (json/xml)
        if (isset($_tempBody)) {
            $httpBody = $_tempBody; // $_tempBody is the method argument, if present
        } elseif (count($formParams) > 0) {
            $httpBody = $formParams; // for HTTP post (form)
        }
        // make the API Call
        try {
            list($response, $statusCode, $httpHeader) = $this->apiClient->callApi(
                $resourcePath,
                'GET',
                $queryParams,
                $httpBody,
                $headerParams,
                '\HydraSDK\Model\JSONWebKeySet',
                '/.well-known/jwks.json'
            );

            return [$this->apiClient->getSerializer()->deserialize($response, '\HydraSDK\Model\JSONWebKeySet', $httpHeader), $statusCode, $httpHeader];
        } catch (ApiException $e) {
            switch ($e->getCode()) {
                case 200:
                    $data = $this->apiClient->getSerializer()->deserialize($e->getResponseBody(), '\HydraSDK\Model\JSONWebKeySet', $e->getResponseHeaders());
                    $e->setResponseObject($data);
                    break;
                case 500:
                    $data = $this->apiClient->getSerializer()->deserialize($e->getResponseBody(), '\HydraSDK\Model\GenericError', $e->getResponseHeaders());
                    $e->setResponseObject($data);
                    break;
            }

            throw $e;
        }
    }
}
