<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class StoreDnsRequest extends FormRequest
{
    public function rules(): array
    {
        return [
            'domain' => ['required', 'string'],
            'txt' => ['required', 'string'],
        ];
    }
}
