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
        $privateKey = base64_decode($value);

        if ($privateKey === false) {
            $fail("The field $attribute is not a valid base64 private key.");
        }

        $privateKeyPem = CertificateField::PRIVATE_KEY->keyToPemFormat($privateKey);

        if (openssl_pkey_get_private($privateKeyPem) === false) {
            $fail("The field $attribute is not a valid private key.");
        }
    }
}
