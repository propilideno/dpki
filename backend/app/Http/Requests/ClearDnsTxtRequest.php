<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class ClearDnsTxtRequest extends FormRequest
{
    public function authorize(): bool
    {
        return true;
    }
}
