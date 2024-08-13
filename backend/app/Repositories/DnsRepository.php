<?php

namespace App\Repositories;

use App\Models\Dns;
use App\Repositories\Interfaces\DnsRepositoryInterface;
use Illuminate\Support\Arr;

class DnsRepository implements DnsRepositoryInterface
{
    public function __construct(private readonly Dns $model)
    {
        //
    }

    public function findOrFail(int $id): Dns
    {
        /** @var Dns */
        return $this->model->newQuery()
            ->findOrFail($id);
    }

    public function findBy(array $attributes, bool $fail = true): ?Dns
    {
        $query = $this->model->newQuery()
            ->where($attributes);

        if ($fail) {
            /** @var Dns */
            return $query->firstOrFail();
        }

        /** @var ?Dns */
        return $query->first();
    }

    public function create(array $attributes): Dns
    {
        /** @var Dns */
        return $this->model->newQuery()
            ->create($attributes)
            ->fresh();
    }

    public function updateOrCreate(array $attributes, array $values): Dns
    {
        /** @var Dns */
        return $this->model->newQuery()
            ->updateOrCreate($attributes, $values)
            ->fresh();
    }

    public function update(array $attributes, Dns|int $dns): Dns
    {
        $dns = $dns instanceof Dns
            ? $dns
            : $this->model->newQuery()->findOrFail($dns);

        $dns->update($attributes);

        return $dns->fresh();
    }
}
