<?php

namespace App\Enums;

use Illuminate\Support\Str;

enum CertificateField: string
{
    case CERTIFICATE = 'certificate';

    case PRIVATE_KEY = 'private_key';

    case PUBLIC_KEY = 'public_key';

    public function beginPem(): string
    {
        return "-----BEGIN " . $this->formatValue() . "-----";
    }

    public function endPem(): string
    {
        return "-----END " . $this->formatValue() . "-----";
    }

    private function formatValue(): string
    {
        return strtoupper(str_replace('_', ' ', $this->value));
    }

    public function formatKey(string $key, bool $applyTrim = false): string
    {
        if ($applyTrim) {
            $key = preg_replace('/' . $this->beginPem() . '/', '', $key);
            $key = preg_replace('/' . $this->endPem() . '/', '', $key);
        }

        return str_replace("\n", '', $key);
    }

    public function keyToPemFormat(string $key): string
    {
        $formattedKey = $this->formatKey($key, true);
        $formattedKey = str_replace(' ', '+', $formattedKey);
        $formattedKey = wordwrap($formattedKey, 64, "\n", true);

        return$this->beginPem() . "\n" . $formattedKey . "\n" . $this->endPem();
    }
}
