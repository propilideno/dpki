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
        $privateKeyPem = CertificateField::PRIVATE_KEY->keyToPemFormat($value);

        if (openssl_pkey_get_private($privateKeyPem) === false) {
            $fail("The field $attribute is not a valid private key.");
        }
    }
}
