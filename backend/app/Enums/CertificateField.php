<?php

namespace App\Enums;

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

    public function formatKey(string $key): string
    {
        $key = $this->removePemHeaderAndFooter($key);
        $key = str_replace("\n", '', $key);
        $key = base64_encode($key);

        return $this->beginPem() . $key . $this->endPem();
    }

    public function keyToPemFormat(string $key): string
    {
        $formattedKey = $this->removePemHeaderAndFooter($key);
        $formattedKey = wordwrap($formattedKey, 64, "\n", true);

        return  $this->addPemHeaderAndFooter($formattedKey, true);
    }

    public function removePemHeaderAndFooter(string $key): string
    {
        $key = preg_replace('/' . $this->beginPem() . '/', '', $key);
        $key = preg_replace('/' . $this->endPem() . '/', '', $key);

        return trim($key);
    }

    public function addPemHeaderAndFooter(string $key, bool $withEol = false): string
    {
        if ($withEol) {
            $key = "\n" . $key . "\n";
        }

        return $this->beginPem() . $key . $this->endPem();
    }
}
