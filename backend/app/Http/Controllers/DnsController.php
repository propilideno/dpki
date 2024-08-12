<?php

namespace App\Http\Controllers;

use App\Http\Requests\ClearDnsTxtRequest;
use App\Http\Requests\StoreDnsRequest;
use App\Repositories\Interfaces\DnsRepositoryInterface;
use Illuminate\Http\JsonResponse;

class DnsController extends Controller
{
    public function __construct(private readonly DnsRepositoryInterface $dnsRepository)
    {
        //
    }

    public function create(StoreDnsRequest $request): JsonResponse
    {
        $attributes = $request->validated();

        $matchAttributes = [
            'domain' => $attributes['domain'],
        ];

        return response()->json(
            $this->dnsRepository->updateOrCreate($matchAttributes, $attributes)
        );
    }

    public function clearDnsTXT(ClearDnsTxtRequest $request, string $domain): JsonResponse
    {
        $dns = $this->dnsRepository->findBy([
            'domain' => $domain,
        ]);

        $attributes = [
          'txt' => null,
        ];

        return response()->json(
          $this->dnsRepository->update($attributes, $dns->id)
        );
    }

    public function show(string $domain): JsonResponse
    {
        $dns = $this->dnsRepository->findBy([
            'domain' => $domain,
        ]);

        return response()->json($dns);
    }
}
