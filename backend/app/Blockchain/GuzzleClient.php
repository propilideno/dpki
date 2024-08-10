<?php

namespace App\Blockchain;

use Exception;
use GuzzleHttp\Client;
use Psr\Http\Message\ResponseInterface;
use Throwable;

class GuzzleClient extends Client
{
    public function __construct(string $baseUrl)
    {
        $config = [
            'base_uri' => $baseUrl,
            'verify' => false,
        ];

        parent::__construct($config);
    }

    /**
     * @throws Throwable
     */
    public function request(string $method, $uri = '', array $options = [], bool $showApiError = true): ResponseInterface
    {
        try {
            return parent::request($method, $uri, $options);

        } catch (Throwable $e) {
            if ($showApiError) {
                throw $e;
            }

            throw new Exception('Error connecting to Blockchain.');
        }
    }

    /**
     * @throws Throwable
     */
    public function get($uri = '', array $options = [], bool $showApiError = true): ResponseInterface
    {
        return $this->request('GET', $uri, $options, $showApiError);
    }

    /**
     * @throws Throwable
     */
    public function post($uri = '', array $options = [], bool $showApiError = true): ResponseInterface
    {
        return $this->request('POST', $uri, $options, $showApiError);
    }

    public function responseToArray(ResponseInterface $response): array
    {
        return json_decode($response->getBody()->getContents(), true) ?? [];
    }
}
