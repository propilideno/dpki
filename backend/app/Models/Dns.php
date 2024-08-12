<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Dns extends Model
{
    use HasFactory;

    protected $guarded = [
        'id',
        'created_at',
    ];

    protected $fillable = [
        'domain',
        'txt',
    ];
}
