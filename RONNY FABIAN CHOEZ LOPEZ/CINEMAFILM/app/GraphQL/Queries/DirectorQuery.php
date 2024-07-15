<?php

namespace App\GraphQL\Queries;

use App\Models\Director;

class DirectorQuery
{
    public function resolve($root, array $args)
    {
        return Director::all();
    }
}