<?php

namespace App\Http\Controllers;

use App\Models\Director;
use Illuminate\Http\Request;

class DirectorController extends Controller
{
    public function index()
    {
        return Director::all();
    }

    public function store(Request $request)
    {
        $request->validate([
            'nombre' => 'required|string|max:50',
        ]);

        return Director::create($request->all());
    }

    public function show($id_director)
    {
        return Director::findOrFail($id_director);
    }

    public function update(Request $request, $id_director)
    {
        $request->validate([
            'nombre' => 'string|max:50',
        ]);

        $genero = Director::findOrFail($id_director);
        $genero->update($request->all());

        return $genero;
    }

    public function destroy($id_director)
    {
        $genero = Director::findOrFail($id_director);
        $genero->delete();

        return response()->noContent();
    }
}