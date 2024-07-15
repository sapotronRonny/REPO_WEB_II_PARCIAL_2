<?php

namespace App\GraphQL\Types;

use App\Models\GeneroDocu;
use GraphQL\Type\Definition\Type;
use Rebing\GraphQL\Support\Type as GraphQLType;

class GeneroDocuType extends GraphQLType
{
    protected $attributes = [
        'name' => 'GeneroDocu',
        'description' => 'A type',
        'model' => GeneroDocu::class,
    ];

    public function fields(): array
    {
        return [
            'id_genero' => [
                'type' => Type::nonNull(Type::int()),
                'description' => 'The ID of the genre',
            ],
            'nombre' => [
                'type' => Type::nonNull(Type::string()),
                'description' => 'The name of the genre',
            ],
        ];
    }
}