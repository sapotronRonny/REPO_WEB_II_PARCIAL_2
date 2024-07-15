<?php

namespace App\Http\Controllers;

use App\Models\GeneroDocu;
use Illuminate\Http\Request;

class GeneroDocuController extends Controller
{
    public function index()
    {
        return GeneroDocu::all();
    }

    public function store(Request $request)
    {
        $request->validate([
            'nombre' => 'required|string|max:50',
        ]);

        return GeneroDocu::create($request->all());
    }

    public function show($id_genero)
    {
        return GeneroDocu::findOrFail($id_genero);
    }

    public function update(Request $request, $id_genero)
    {
        $request->validate([
            'nombre' => 'string|max:50',
        ]);

        $genero = GeneroDocu::findOrFail($id_genero);
        $genero->update($request->all());

        return $genero;
    }

    public function destroy($id_genero)
    {
        $genero = GeneroDocu::findOrFail($id_genero);
        $genero->delete();

        return response()->noContent();
    }
}