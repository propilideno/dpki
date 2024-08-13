<?php

namespace App\Rules;

use App\Enums\CertificateField;
use Closure;
use Illuminate\Contracts\Validation\ValidationRule;
use Illuminate\Support\Str;

class PrivateKeyRule implements ValidationRule
{
    public function validate(string $attribute, mixed $value, Closure $fail): void
    {
        $privateKey = CertificateField::PRIVATE_KEY->removePemHeaderAndFooter($value);
        $privateKey = base64_decode($privateKey);

        if ($privateKey === false) {
            $fail("The field $attribute is not a valid base64 private key.");
        }

        $privateKey = CertificateField::PRIVATE_KEY->addPemHeaderAndFooter($privateKey);

        $privateKeyPem = CertificateField::PRIVATE_KEY->keyToPemFormat($privateKey);

        if (openssl_pkey_get_private($privateKeyPem) === false) {
            $fail("The field $attribute is not a valid private key.");
        }
    }
}
