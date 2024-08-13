<?php

namespace App\Providers;

use App\Repositories\DnsRepository;
use App\Repositories\Interfaces\DnsRepositoryInterface;
use Illuminate\Support\ServiceProvider;

class RepositoryServiceProvider extends ServiceProvider
{
    public array $bindings = [
        DnsRepositoryInterface::class => DnsRepository::class,
    ];
}
