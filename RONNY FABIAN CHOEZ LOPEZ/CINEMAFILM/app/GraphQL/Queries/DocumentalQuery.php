<?php

namespace App\GraphQL\Queries;

use App\Models\Documental;
use GraphQL\Type\Definition\Type;
use Rebing\GraphQL\Support\Facades\GraphQL;
use Rebing\GraphQL\Support\Query;

class DocumentalQuery extends Query
{
    protected $attributes = [
        'name' => 'documentals',
    ];

    public function type(): Type
    {
        return Type::listOf(GraphQL::type('Documental'));
    }

    public function resolve($root, $args)
    {
        return Documental::all();
    }
}