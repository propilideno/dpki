<?php

namespace App\Http\Controllers;

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
}
