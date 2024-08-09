<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class CertificateRequest extends FormRequest
{
    public function rules(): array
    {
        return [
            'apply_trim' => ['sometimes', 'nullable', 'boolean'],
        ];
    }
}
