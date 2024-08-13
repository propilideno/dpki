<?php

namespace App\Http\Controllers;

use App\Http\Requests\StoreDnsRequest;
use App\Repositories\Interfaces\DnsRepositoryInterface;
use App\Services\DnsService;
use Illuminate\Http\JsonResponse;

class DnsController extends Controller
{
    public function __construct(
        private readonly DnsService $dnsService,
        private readonly DnsRepositoryInterface $dnsRepository,
    )
    {
        //
    }

    public function create(StoreDnsRequest $request): JsonResponse
    {
        $attributes = $request->validated();

        return response()->json(
            $this->dnsService->updateOrCreate($attributes)
        );
    }

    public function clearDnsTxt(string $domain): JsonResponse
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
