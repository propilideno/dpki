<?php

namespace App\Http\Requests;

use App\Rules\PrivateKeyRule;
use Illuminate\Foundation\Http\FormRequest;

class SignRequest extends FormRequest
{
    public function rules(): array
    {
        return [
            'message' => ['required', 'string'],
            'private_key' => ['required', 'string', new PrivateKeyRule()],
        ];
    }
}
