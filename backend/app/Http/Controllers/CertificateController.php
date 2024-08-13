<?php

namespace App\Http\Controllers;

use App\Http\Requests\HashRequest;
use App\Http\Requests\SignRequest;
use App\Services\CertificateService;
use Illuminate\Http\JsonResponse;

class CertificateController extends Controller
{
    public function __construct(
        private readonly CertificateService $certificateService,
    ) {
        //
    }

    public function signMessage(SignRequest $request): JsonResponse
    {
        $data = $request->validated();

        $sign = $this->certificateService->signMessage($data['message'], $data['private_key']);

        return response()->json([
            'sign' => $sign,
        ]);
    }

    public function generateCertificate(): JsonResponse
    {
        return response()->json(
            $this->certificateService->generateCertificate()
        );
    }

    public function hash(HashRequest $request): JsonResponse
    {
        $data = $request->validated();

        return response()->json([
            'hash' => base64_encode(hash_hmac('sha256', $data['message'], $data['key'])),
        ]);
    }
}
