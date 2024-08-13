<?php

use App\Http\Controllers\CertificateController;
use App\Http\Controllers\DnsController;
use Illuminate\Support\Facades\Route;

Route::get('/certificate/new', [CertificateController::class, 'generateCertificate']);
Route::get('/sign', [CertificateController::class, 'signMessage']);
Route::get('/hash', [CertificateController::class, 'hash']);

Route::get('/dns/{domain}', [DnsController::class, 'show']);

Route::middleware('hmac_auth')->group(function ()
{
    Route::post('/dns', [DnsController::class, 'create']);
    Route::patch('/dns/{domain}/clear-txt', [DnsController::class, 'clearDnsTxt']);
});
