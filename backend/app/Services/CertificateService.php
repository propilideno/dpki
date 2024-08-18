<?php

namespace App\Services;

use App\Enums\CertificateField;

class CertificateService
{
    public function generateCertificate(): array
    {
        $config = [
            'digest_alg' => 'sha256',
            'private_key_bits' => 512,
            'private_key_type' => OPENSSL_KEYTYPE_RSA,
            'encrypt_key' => false,
        ];

        $privateKey = openssl_pkey_new($config);
        $csr = openssl_csr_new([], $privateKey);
        $certificate = openssl_csr_sign($csr, null, $privateKey, 365);

        openssl_pkey_export($privateKey, $privateKeyOut);
        $publicKey = openssl_pkey_get_details($privateKey)["key"];
        openssl_x509_export($certificate, $certificateOut);

        return [
            CertificateField::PRIVATE_KEY->value => base64_encode($privateKeyOut),
            CertificateField::PUBLIC_KEY->value => base64_encode($publicKey),
            CertificateField::CERTIFICATE->value => base64_encode($certificateOut),
        ];
    }

    public function signMessage(string $message, string $privateKey): string
    {
        $privateKey = base64_decode($privateKey);

        $privateKeyResource = openssl_pkey_get_private($privateKey);

        openssl_sign($message, $signature, $privateKeyResource);

        return base64_encode($signature);
    }
}
