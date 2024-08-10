<?php

namespace App\Blockchain;

class Blockchain
{
    private static ?self $_instance = null;
    private GuzzleClient $client;

    private function __construct()
    {
        $this->client = new GuzzleClient(self::baseUrl());
    }

    private static function getInstance(): self
    {
        if (is_null(self::$_instance)) {
            self::$_instance = new self();
        }

        return self::$_instance;
    }

    public static function baseUrl(): string
    {
        return rtrim(config('blockchain.base_url'), '/') . '/';
    }

    /**
     * @throws \Throwable
     */
    public static function mine(): array
    {
        $instance = self::getInstance();
        $response = $instance->client->get('mine');

        return $instance->client->responseToArray($response);
    }
}
