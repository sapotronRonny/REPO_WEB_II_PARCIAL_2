<?php

namespace App\GraphQL\Types;

use App\Models\Documental;
use GraphQL\Type\Definition\Type;
use Rebing\GraphQL\Support\Type as GraphQLType;


class DocumentalType extends GraphQLType
{
    protected $attributes = [
        'name' => 'Documental',
        'description' => 'A type',
        'model' => Documental::class,
    ];

    public function fields(): array
    {
        return [
            'id_documental' => [
                'type' => Type::nonNull(Type::int()),
                'description' => 'The ID of the documental',
            ],
            'titulo' => [
                'type' => Type::nonNull(Type::string()),
                'description' => 'The title of the documental',
            ],
            'id_genero' => [
                'type' => Type::int(),
                'description' => 'The genre ID of the documental',
            ],
            'id_director' => [
                'type' => Type::int(),
                'description' => 'The director ID of the documental',
            ],
            'duracion' => [
                'type' => Type::int(),
                'description' => 'The duration of the documental',
            ],
            'resumen' => [
                'type' => Type::string(),
                'description' => 'The summary of the documental',
            ],
            'genero' => [
                'type' => \GraphQL::type('GeneroDocu'),
                'description' => 'The genre of the documental',
            ],
            'director' => [
                'type' => \GraphQL::type('Director'),
                'description' => 'The director of the documental',
            ],
        ];
    }
}