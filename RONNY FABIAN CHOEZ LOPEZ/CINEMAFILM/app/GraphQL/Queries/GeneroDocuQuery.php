<?php

namespace App\GraphQL\Queries;

use App\Models\GeneroDocu;

class GeneroDocuQuery
{
    public function resolve($root, array $args)
    {
        return GeneroDocu::all();
    }
}
