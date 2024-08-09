<?php

use App\Http\Controllers\CertificateController;
use Illuminate\Support\Facades\Route;

Route::get('/certificate/new', [CertificateController::class, 'generateCertificate']);
Route::get('/sign', [CertificateController::class, 'signMessage']);
