<?php

namespace App\GraphQL\Types;

use App\Models\Director;
use GraphQL\Type\Definition\Type;
use Rebing\GraphQL\Support\Type as GraphQLType;

class DirectorType extends GraphQLType
{
    protected $attributes = [
        'name' => 'Director',
        'description' => 'A type',
        'model' => Director::class,
    ];

    public function fields(): array
    {
        return [
            'id_director' => [
                'type' => Type::nonNull(Type::int()),
                'description' => 'The ID of the director',
            ],
            'nombre' => [
                'type' => Type::nonNull(Type::string()),
                'description' => 'The name of the director',
            ],
        ];
    }
}