<?php

namespace App\GraphQL\Queries;

use App\Models\Documental;

class DocumentalQuery
{
    public function resolve($root, array $args)
    {
        return Documental::all();
    }
}
