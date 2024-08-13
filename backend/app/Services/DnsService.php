<?php

namespace App\Services;

use App\Models\Dns;
use App\Repositories\Interfaces\DnsRepositoryInterface;

class DnsService
{
    public function __construct(private readonly DnsRepositoryInterface $dnsRepository)
    {
        //
    }

    private function generateUniqueSharedKey(): string
    {
        do {
            $sharedKey = base64_encode(hash_hmac('sha256', str()->random(), str()->random()));

            $dns = $this->dnsRepository->findBy([
                'shared_key' => $sharedKey,
            ], false);
        } while (! is_null($dns));

        return $sharedKey;
    }

    public function updateOrCreate(array $attributes): Dns
    {
        $matchAttributes = [
            'domain' => $attributes['domain'],
        ];

        $dns = $this->dnsRepository->findBy($matchAttributes, false);

        if (is_null($dns)) {
            $attributes['shared_key'] = $this->generateUniqueSharedKey();

            return $this->dnsRepository->create($attributes)
                ->makeVisible('shared_key');
        }

        return $this->dnsRepository->update($attributes, $dns->id);
    }
}
