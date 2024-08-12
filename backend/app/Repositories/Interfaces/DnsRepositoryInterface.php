<?php

namespace App\Repositories\Interfaces;

use App\Models\Dns;

interface DnsRepositoryInterface
{
    public function findOrFail(int $id): Dns;

    public function findBy(array $attributes, bool $fail = true): ?Dns;

    public function create(array $attributes): Dns;

    public function updateOrCreate(array $attributes, array $values): Dns;

    public function update(array $attributes, Dns|int $dns): Dns;
}
