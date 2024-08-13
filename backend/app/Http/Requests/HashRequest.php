<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class HashRequest extends FormRequest
{
    public function rules(): array
    {
        return [
            'message' => ['required', 'string'],
            'key' => ['required', 'string'],
        ];
    }
}
