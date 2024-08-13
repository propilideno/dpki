<?php

namespace App\Http\Middleware;

use App\Repositories\Interfaces\DnsRepositoryInterface;
use Closure;
use Exception;
use Illuminate\Http\Request;
use Illuminate\Validation\UnauthorizedException;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\HttpKernel\Exception\UnprocessableEntityHttpException;

class HmacAuthorization
{
    /**
     * @throws Exception
     */
    public function handle(Request $request, Closure $next): Response
    {
        $request->headers->set('X-Requested-With', 'XMLHttpRequest');

        $domain = $request->route('domain') ?? $request->input('domain');

        if (is_null($domain)) {
            throw new Exception('The field domain is required.', 422);
        }

        $dns = app()->make(DnsRepositoryInterface::class)->findBy([
            'domain' => $domain,
        ], false);

        if (is_null($dns)) {
            return $next($request);
        }

        $hashFromHeader = $request->header('hash');
        if (is_null($hashFromHeader)) {
            return $this->errorResponse('There are missing fields in header.', 422);
        }

        $correctHash = base64_encode(hash_hmac('sha256', $domain, $dns->shared_key));

        if ($correctHash !== $hashFromHeader) {
            return $this->errorResponse('Unauthorized.', 401);
        }

        return $next($request);
    }

    private function errorResponse(string $message, int $statusCode): Response
    {
        return response()->json([
            'message' => $message,
        ], $statusCode);
    }
}
